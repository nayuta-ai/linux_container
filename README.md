# linux container
## Environment
- Ubuntu:20.04
```
make build_ubuntu
make run_ubuntu
make exec_ubuntu
```
- Golang:1.19+Alpine
```
make build_alpine
make run_alpine
make exec_alpine
```
- ubuntu:20.04+Golang:1.19
```
make build
make run
make exec
```
## Create Container
1. Expand busybox.tar using the below command.
```
$ cd rootfs
$ tar xvf busybox.tar
```
2. Create a binary file.
```
$ cd ..
$ go build src/myuts.go
```
3. Execute and create the new container
```
$ ./myuts parent /bin/sh
```

## To Do
- [ ] Currently only supports amd64, so make it available for arm64 as well.
  - [ ] Eliminate `netsetgo` and change to code that also works on arm64.
- [ ] Add function for cgroup
  - [ ] Apply cgroupv1
  - [ ] Apply cgroupv2
  - [ ] Set up systemd
