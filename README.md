# linux container
## Architecture
x86_64
## Environment
- ubuntu:20.04+Golang:1.19
```
# make up
# ssh -p 2222 vagrant@localhost
$ sudo su
```
password: vagrant

## Create Container
0. Set up grub
- Set up GRUB settings
```
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash systemd.unified_cgroup_hierarchy=1"
```
- Update grub and reboot
```
$ update-grub
$ reboot
```
1. Set up
```
# ssh -p 2222 vagrant@localhost
$ sudo su
$ cd linux_container
```
2. Create a binary file.
```
$ PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
$ go build .
```
3. Execute and create the new container
```
$ ./linux_container parent /bin/sh
```
4. Add the binary to path (To Do)
```
$ mv linux_container /usr/local/sbin/linux_container
```

## How to execute integration test for OCI (To Do)
- This container runtime doesn't work well.
- We describe how to execute OCI test when assigning runc in the below.
```
RUNTIME=runc validation/default/default.t
```
