# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
    # The most common configuration options are documented and commented below.
    # For a complete reference, please see the online documentation at
    # https://docs.vagrantup.com.
  
    # Every Vagrant development environment requires a box. You can search for
    # boxes at https://vagrantcloud.com/search.
    config.vm.box = "generic/ubuntu2004"
  
    # Disable automatic box update checking. If you disable this, then
    # boxes will only be checked for updates when the user runs
    # `vagrant box outdated`. This is not recommended.
    # config.vm.box_check_update = false
  
    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine. In the example below,
    # accessing "localhost:8080" will access port 80 on the guest machine.
    # NOTE: This will enable public access to the opened port
    # config.vm.network "forwarded_port", guest: 80, host: 8080
  
    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine and only allow access
    # via 127.0.0.1 to disable public access
    # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"
  
    # Create a private network, which allows host-only access to the machine
    # using a specific IP.
    # config.vm.network "private_network", ip: "192.168.33.10"
  
    # Create a public network, which generally matched to bridged network.
    # Bridged networks make the machine appear as another physical device on
    # your network.
    # config.vm.network "public_network"
  
    # Share an additional folder to the guest VM. The first argument is
    # the path on the host to the actual folder. The second argument is
    # the path on the guest to mount the folder. And the optional third
    # argument is a set of non-required options.
    # config.vm.synced_folder "/", "/"
  
    # Provider-specific configuration so you can fine-tune various
    # backing providers for Vagrant. These expose provider-specific options.
    # Example for VirtualBox:
    #
    # config.vm.provider "virtualbox" do |vb|
    #   # Display the VirtualBox GUI when booting the machine
    #   vb.gui = true
    #
    #   # Customize the amount of memory on the VM:
    #   vb.memory = "1024"
    # end
    #
    # View the documentation for the provider you are using for more
    # information on available options.
  
    # Enable provisioning with a shell script. Additional provisioners such as
    # Ansible, Chef, Docker, Puppet and Salt are also available. Please see the
    # documentation for more information about their specific syntax and use.
    config.vm.provision "shell", inline: <<-SHELL
      echo "# vagrant script for every boot" > /etc/profile.d/vagrant.sh
      GO_VERSION=1.19
      GO_PATH=/go
      PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
      wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"
      sudo tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
      sudo rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
      
      sudo echo export GO_PATH=#{ENV["GO_PATH"]} >> /etc/profile.d/vagrant.sh
      sudo echo export PATH=#{ENV["PATH"]} >> /etc/profile.d/vagrant.sh
      sudo mkdir -p "$GOPATH/src" "$GOPATH/bin"
      sudo chmod -R 777 "$GOPATH"
      wget "https://github.com/teddyking/netsetgo/releases/download/0.0.1/netsetgo"
      sudo mv netsetgo /usr/local/bin/
      sudo chown root:root /usr/local/bin/netsetgo
      sudo chmod 4755 /usr/local/bin/netsetgo
      sudo apt-get update
      sudo apt-get install -y init systemd vim git cmake make libseccomp-dev pkg-config\
        gettext-base jq curl tmux g++ lcov libbsd0 libbsd-dev sudo kmod nodejs npm
      sudo npm install -g tap
      cd $HOME
      git clone https://github.com/nayuta-ai/linux_container.git
      git clone https://github.com/opencontainers/runc.git
      git clone https://github.com/opencontainers/runtime-tools.git
      cd $HOME/linux_container/rootfs
      tar xvf busybox.tar
      cd $HOME/runc
      make
      sudo make install
      cd $HOME/runtime-tools
      make runtimetest validation-executables
      rm -r $HOME/runc
    SHELL
  end
  