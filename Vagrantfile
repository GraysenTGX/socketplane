# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

$ubuntu = <<SCRIPT
echo ====> Updating Packages
export DEBIAN_FRONTEND=noninteractive
# -qq is pointless, it doesn't work :S
apt-get update > /dev/null
echo ====> Installing Packages
apt-get install -qq -y --no-install-recommends docker.io openvswitch-switch
ln -s /vagrant/socketplane.sh /usr/bin/socketplane
echo ====> Installing SocketPlane
socketplane install
SCRIPT

$redhat = <<SCRIPT
echo ====> Updating Packages
yum -qy update
echo ====> Installing Packages
yum -qy remove docker
yum -qy install docker-io openvswitch
ln -s /vagrant/socketplane.sh /usr/bin/socketplane
echo ====> Installing SocketPlane
socketplane install
SCRIPT

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  # Socketplane demo boxes
  num_nodes = (ENV['SOCKETPLANE_NODES'] || 3).to_i
  base_ip = "10.254.101."
  socketplane_ips = num_nodes.times.collect { |n| base_ip + "#{n+21}" }

  num_nodes.times do |n|
    config.vm.define "socketplane-#{n+1}" do |socketplane|
      socketplane.vm.box = "chef/ubuntu-14.04"
      socketplane_ip = socketplane_ips[n]
      socketplane_index = n+1
      socketplane.vm.hostname = "socketplane-#{socketplane_index}"
      socketplane.vm.network :private_network, ip: "#{socketplane_ip}", virtualbox__intnet: true
      socketplane.vm.provider :virtualbox do |vb|
        vb.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
      end
      socketplane.vm.provision :shell, inline: $ubuntu
    end
  end

  # Boxes for testing installer
  config.vm.define "ubuntu", autostart: false do |ubuntu|
    ubuntu.vm.box = "chef/ubuntu-14.04"
    ubuntu.vm.hostname = "ubuntu"
    ubuntu.vm.network :private_network, ip: "10.254.101.10", virtualbox__intnet: true
    ubuntu.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
    end
    #ubuntu.vm.provision :shell, inline: $ubuntu
  end
  config.vm.define "centos", autostart: false do |centos|
    centos.vm.box = "chef/centos-7"
    centos.vm.hostname = "centos"
    centos.vm.network :private_network, ip: "10.254.101.11", virtualbox__intnet: true
    centos.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
    end
    #centos.vm.provision :shell, inline: $redhat
  end
  config.vm.define "fedora", autostart: false do |fedora|
    fedora.vm.box = "chef/fedora-20"
    fedora.vm.hostname = "fedora"
    fedora.vm.network :private_network, ip: "10.254.101.12", virtualbox__intnet: true
    fedora.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--nicpromisc2", "allow-all"]
    end
    #fedora.vm.provision :shell, inline: $redhat
  end
end