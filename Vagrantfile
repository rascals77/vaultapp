$script = <<-SCRIPT
git clone https://github.com/rascals77/vaultapp.git /root/vaultapp
chmod +x /root/vaultapp/run.sh
/root/vaultapp/run.sh
SCRIPT

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/bionic64"
  config.vm.box_version = "20200129.1.0"
  config.vm.synced_folder ".", "/vagrant", disabled: true
  config.vm.network "forwarded_port", guest: 8080, host: 8080, auto_correct: true
  config.vm.provision "shell", inline: $script

  config.vm.provider :virtualbox do |vb|
    vb.memory = "2048"
    vb.cpus = "2"
  end

  config.vm.define "vaultapp" do |vaultapp|
    vaultapp.vm.hostname = "vaultapp"
  end
end
