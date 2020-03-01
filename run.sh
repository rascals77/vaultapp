#!/bin/bash

set -e

echo -e "\x1b[1;32mStep 1: Install packages\x1b[0m"
apt-get update
apt-get -y install docker.io python3-pip python3-venv

echo -e "\x1b[1;32mStep 2: Enable and start docker\x1b[0m"
systemctl enable docker
systemctl start docker

echo -e "\x1b[1;32mStep 3: Initialize pip env\x1b[0m"
python3 -m venv --system-site-packages /root/PY

echo -e "\x1b[1;32mStep 4: Install pip requirements\x1b[0m"
source /root/PY/bin/activate
pip install -r /root/vaultapp/ansible/requirements.txt

echo -e "\x1b[1;32mStep 5: Run Vault\x1b[0m"
ansible-playbook -i localhost, /root/vaultapp/ansible/vault.yml

echo -e "\x1b[1;32mStep 6: Build webapp\x1b[0m"
docker build -t local/webapp:0.1 /root/vaultapp/webapp
docker image prune --filter label=webapp=builder --force

echo -e "\x1b[1;32mStep 7: Run webapp\x1b[0m"
DOCKER_IP=$(ip addr list docker0 | grep "inet " | awk '{print $2}' | cut -d/ -f1)
docker run -d --rm -p 8080:8080 --name webapp local/webapp:0.1 /web-app -vaultip ${DOCKER_IP}
