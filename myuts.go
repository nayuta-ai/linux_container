package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var cgroupRootPath = "/sys/fs/cgroup"
var cgroupChildPath = filepath.Join(cgroupRootPath, "child")

func createChildCgroup() {
	os.Mkdir(cgroupChildPath, os.ModePerm)
}

func enableCgroup() {
	must(ioutil.WriteFile(filepath.Join(cgroupChildPath, "memory.max"), []byte("2M"), 0700))
	must(ioutil.WriteFile(filepath.Join(cgroupChildPath, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func waitForNetwork() error {
	maxWait := time.Second * 3
	checkInterval := time.Second
	timeStarted := time.Now()
	for {
		interfaces, err := net.Interfaces()
		if err != nil {
			return err
		}
		// pretty basic check ...
		// > 1 as a lo device will already exist
		if len(interfaces) > 1 {
			return nil
		}
		if time.Since(timeStarted) > maxWait {
			return fmt.Errorf("Timeout after %s waiting for network", maxWait)
		}
		time.Sleep(checkInterval)
	}
}

func pivotRoot(newroot string) error {
	putold := filepath.Join(newroot, "/.pivot_root")
	// bind mount newroot to itself - this is a slight hack needed to satisfy the
	// pivot_root requirement that newroot and putold must not be on the same filesystem
	// as the current root
	if err := syscall.Mount(newroot, newroot, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}
	// create putold directory
	if err := os.MkdirAll(putold, 07000); err != nil {
		return err
	}

	// call pivot_root
	if err := syscall.PivotRoot(newroot, putold); err != nil {
		return err
	}

	// ensure current working directory is set to new root
	if err := os.Chdir("/"); err != nil {
		return err
	}

	// umount putold, which now lives at /.pivot_root
	putold = "/.pivot_root"
	if err := syscall.Unmount(putold, syscall.MNT_DETACH); err != nil {
		return err
	}

	// remove puttold
	if err := os.RemoveAll(putold); err != nil {
		return err
	}
	return nil
}

// the parent function invoked from the main program which sets up the needed namespaces
func parent() {
	createChildCgroup()
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"name=shashank"}
	// command below creates the UTS, PID, IPC, NETWORK and USERNAMESPACES
	// and adds the user and group mappings
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
	must(cmd.Start())

	pid := fmt.Sprintf("%d", cmd.Process.Pid)
	// Code velow does the following
	// Create the bridge on the host
	// Create the veth pair
	// Attaches one end of veth to bridge
	// Attaches the other end to the network namespace. This is interresting
	// as we now have access to the host side and the network side until
	// we block.
	netsetgoCmd := exec.Command("/usr/local/bin/netsetgo", "-pid", pid)
	if err := netsetgoCmd.Run(); err != nil {
		fmt.Printf("Error running netsetgo - %s\n", err)
		os.Exit(1)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error waiting for reexec.Command - %s\n", err)
		os.Exit(1)
	}
}

// this is the child process which is a copy of the parent program itself.
func child() {
	// enable the cgroup functionality
	enableCgroup()
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error fetching current working directory - %s\n", err)
		os.Exit(1)
	}
	// make a call to mountProc function which would mount the proc filesystem
	// to the already created mount namespace
	must(mountProc(filepath.Join(dir, "rootfs")))
	// the command below sets the hostname to myhost. Idea here is to showcase
	// the use of UTS namespace
	must(syscall.Sethostname([]byte("myhost")))
	if err := pivotRoot(filepath.Join(dir, "rootfs")); err != nil {
		fmt.Printf("Error running pivot_root - %s\n", err)
		os.Exit(1)
	}
	if err := waitForNetwork(); err != nil {
		fmt.Printf("Error waiting for network - %s\n", err)
		os.Exit(1)
	}
	// this command executes the shell which is passed as a program argumnent
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error starting the reexec.Command - %s\n", err)
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		fmt.Printf("Error - %s\n", err)
	}
}

// this function mounts the proc filesystem which the
// new mount namespace
func mountProc(newroot string) error {
	source := "proc"
	target := filepath.Join(newroot, "/proc")
	fstype := "proc"
	flags := 0
	data := ""
	// make a Mount system call to mount the proc filesystem within the mount namespace
	os.MkdirAll(target, 0755)
	if err := syscall.Mount(
		source,
		target,
		fstype,
		uintptr(flags),
		data,
	); err != nil {
		return err
	}
	return nil
}

func parse() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error fetching current working directory - %s\n", err)
		os.Exit(1)
	}
	configPath := filepath.Join(dir, "config/config.json")
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("Error opening config file - %s\n", err)
		os.Exit(1)
	}
	dec := json.NewDecoder(f)
	var config Config

	if err := dec.Decode(&config); err != nil {
		fmt.Printf("Error decoding config file - %s\n", err)
		os.Exit(1)
	}
	f.Close()
	fmt.Println(config.OciVersion)
}
