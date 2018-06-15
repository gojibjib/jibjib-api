# Deploy the [JibJib](https://github.com/gojibjib) stack remotely

Leverage [Terraform](https://www.terraform.io/) and/or [SaltStack](https://saltstack.com/) to deploy the following services to AWS:

- The [jibjib-api](https://github.com/gojibjib/jibjib-api/) service
- A [MongoDB](https://www.mongodb.com/) instance for the API
- The [jibjib-query](https://github.com/gojibjib/jibjib-query) service, containg a [Flask](http://flask.pocoo.org/) API and a [TensorFlow Serving](https://www.tensorflow.org/serving/) container to serve the [jibjib-model](https://gojibjib/jibjib-model).

## Local prerequisites
Clone the repo:

```
git clone https://github.com/gojibjib/jibjib-api
cd jibjib-api/deploy
```

Install [Docker](https://docs.docker.com/install/#server) and [Terraform](https://www.terraform.io/intro/getting-started/install.html) (optional)

## Terraform
The plan provided in the `terraform/` directory is for an [AWS](https://aws.amazon.com/free/) installation. If you want to deploy it on a different platform you will have to create your own plan or skip this part. By default, the setup should not exceed free tier limits if you set your `instance_type` and `query_instance_type` variables to `t2.micro`. However, if you're using the `jibjib_model` at least 2GB of RAM is recommended for the [query service](https://github.com/gojibjib/jibjib-query) instance. 

The plan will provision the following:

- AWS key pair (needs to be generated first)
- 3x AWSinstances (+ provision it with `bootstrap/ubuntu.sh`)
- 3x EIP
- Security group:
	- Ingress:
		- From all to 22
		- From all to 8080
      - from all to all inside own SG
	- Egress:
		- From all to all

### Preparation

Generate a key pair:

```
ssh-keygen -b 4096 -t rsa -f terraform/keys/example_key -q -N "" -C ""
```

Initiliaze Terraform:

```
cd terraform
terraform init
```

Create and populate your `terraform.tfvars`:

```
touch terraform.tfvars
```

For your keys, assuming you have the following layout:

```
keys/
├── example_key.pem
└── example_key.pub
```

Make sure to set the relative paths to the files in your `terraform.tfvars`:

```
key_name = "example_key"
public_key = "keys/example_key.pub"
private_key = "keys/example_key.pem"
```

### Deploy

Plan your deployment:

```
terraform plan
```

Then apply it:

```
terraform apply
```

Make sure to save the following outputs, substitute those variables down below with your IPs:

- `jibjib_api_eip_public_ip`
- `jibjib_db_eip_public_ip`
- `jibjib_db_eip_private_ip`
- `jibjib_query_eip_public_ip`
- `jibjib_query_eip_private_ip`


## SaltStack

The in `saltstack/` provided Formula is intended to be used with agentless [Salt SSH](https://docs.saltstack.com/en/latest/topics/ssh/). You can either use your own environment to send commands & apply states, use a [Docker container](https://github.com/obitech/docker-salt) or use [Vagrant](https://www.vagrantup.com/).

### Preparation

Copy your SSH private key into the `saltstack/keys` directory (in this example the key generated for the Terraform plan will be used):

```
cp terraform/keys/*.pem saltstack/keys
```

Create a [roster](https://docs.saltstack.com/en/latest/topics/ssh/roster.html):

```
cat << EOF > saltstack/salt/roster
jibjib_api:
  host: jibjib_api_eip_public_ip
  port: 22
  user: ubuntu
  priv: /root/keys/example_key.pem
  sudo: True

jibjib_db:
  host: jibjib_db_eip_public_ip
  port: 22
  user: ubuntu
  priv: /root/keys/example_key.pem
  sudo: True

jibjib_query:
  host: jibjib_query_eip_public_ip
  port: 22
  user: ubuntu
  priv: /root/keys/example_key.pem
  sudo: True
EOF
```

Create a [Saltfile](https://docs.saltstack.com/en/latest/topics/ssh/index.html#define-cli-options-with-saltfile):

```
cat << EOF > saltstack/salt/Saltfile
salt-ssh:
  roster_file: /root/salt/roster
  config_dir: /root/salt
  priv: /root/keys/example_key.pem
EOF
```

Create your Pillar data:

```
# Create your MongoDB root user
cat << EOF > saltstastack/salt/pillar/db_root_auth.sls
root_user:
  user: root
  pw: changeMe
EOF
```

```
# Create a user with read-only rights on the bird database
cat << EOF > saltstack/salt/pillar/db_api_auth.sls
db_user:
  user: prod-r
  pw: changeMe
EOF
```

```
# The IP of your DB instance
cat << EOF > saltstack/salt/pillar/db_ip.sls
db_ip: jibjib_db_eip_private_ip
EOF
```

### Highstate
Applying Highstate will provision the following, simultaneously:
#### All hosts

1. Install Docker & docker-compose

#### Database host

1. Create necessary folders and copy files
2. Grab DB data as JSON from [github](https://github.com/gojibjib/voice-grabber/tree/master/meta)
3. Starts an initial MongoDB container to:
  - Create root user
  - Create read-only user on production database
  - Imports JSON into production database
4. Starts the actual MongoDB container (with `--auth` flag and persisting, bind-mounted data folder)

#### API host

1. Download & run `docker-compose.yml`

#### Query host

1. Download the [model in protobuf format](https://s3-eu-west-1.amazonaws.com/jibjib/model/jibjib_model_serving.tgz)
2. Download [ID mappings in pickle format](https://s3-eu-west-1.amazonaws.com/jibjib/pickle/mapping_pickles.tgz)
3. Download & run `docker-compose.yml`

### Vagrant
To use this `Vagrantfile` out of the box, you need to have [VirtualBox](https://www.virtualbox.org/wiki/Downloads) installed.

Create your VM:

```
cd saltstack
vagrant up
```

After Vagrant has provisioned the VM, ssh into it:

```
vagrant ssh
```

Folders are mounted into `/root` and commands require `sudo`. First test your connection (might take a while):

```
vagrant@jibjib-api:~$ sudo su
root@jibjib-api:~# cd ~/salt/
root@jibjib-api:~# salt-ssh -i "*" test.ping
jibjib_api:
    True
jibjib_db:
    True
jibjib_query:
    True
```

Apply Highstate:

```
root@jibjib-api:~/salt# salt-ssh "*" state.highstate
```

### Docker
Alternatively, you can send commands from a Docker container with Salt installed.

Start the container from within the `saltstack/` directory:

```
docker container run --rm -it -v $(pwd)/keys:/root/keys -v $(pwd)/salt:/root/salt obitech/salt bash
```

Test the connection (might take a while):

```
root@20af504a6e01:/# cd /root/salt
root@20af504a6e01:~/salt# salt-ssh -i "*" test.ping
jibjib_api:
    True
jibjib_db:
    True
jibjib_query:
    True
```

Apply Highstate:

```
root@20af504a6e01:~/salt# salt-ssh "*" state.highstate
```
