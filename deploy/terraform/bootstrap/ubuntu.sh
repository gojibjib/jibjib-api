#!/usr/bin/env bash

# Update system
sudo apt-get update
sudo DEBIAN_FRONTEND=noninteractive apt-get -y \
    -o DPkg::options::="--force-confdef" \
    -o DPkg::options::="--force-confold" \
    upgrade

# Install Docker
curl -fsSL get.docker.com -o get-docker.sh
sh get-docker.sh

# Add user to docker group
sudo usermod -aG docker ubuntu

# Install Docker compose
sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose