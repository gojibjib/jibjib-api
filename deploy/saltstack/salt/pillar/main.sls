system:
  lookup:
    pkg:
      required:
        - curl
        - vim
        - wget
        - apt-transport-https
        - ca-certificates
        - software-properties-common
        - python-apt

docker:
  lookup:
    pkg:
      name: docker-ce
      key_url: https://download.docker.com/linux/ubuntu/gpg
    users:
      - ubuntu

docker-compose:
  lookup:
    version: 1.21.2

jibjib:
  lookup:
    api:
      compose: https://raw.githubusercontent.com/gojibjib/jibjib-api/master/docker-compose.yml