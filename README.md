# madock
Local development environment based on Docker for Magento

## Description
`madock` is a local Docker-based environment that allows you to run Magento2 projects.
This project is written on Golang and it is distributed under a MIT License.

## Key Features
* Automatic project setup
* Two or more projects can work simultaneously
* Cron support
* Flexible configuration for each project
* Database import and export in two clicks
* Simple viewing of logs with one command
* Debug support
* Synchronization of the local database and media files with the dev site
* Additional services: phpmyadmin, redis, rabbitMQ, elasticsearch, Kibana, ioncube, xdebug, cron
* LiveReload with [Google Chrome plugin](https://chrome.google.com/webstore/detail/livereload-for-madock/cmablbpbnbbgmakinefjgmgpolfahdbo)
* MailHog (email testing tool for developers)
* Magento Cloud
* Composer patches

## Tested on
* Linux (Ubuntu 20.04)
* macOS (Monterey)

## Video

[![madock - install the two Magento 2 projects](https://i9.ytimg.com/vi/_9NvZak_kt8/mq1.jpg?sqp=CPTN95cG&rs=AOn4CLCdHqilfuAftZYHtejLn8v52qWP3g)](https://www.youtube.com/watch?v=_9NvZak_kt8)

## Installation

You need 5 things on your local machine: `git`, `docker`, `docker-compose`, `golang` and `madock`

Follow the installation steps for your system.
<details>
<summary>Mac</summary>

1. Install [Docker](https://docs.docker.com/docker-for-mac/install/)
2. Install [Golang](https://go.dev/doc/install)
3. Clone this repo and follow into folder "madock"
```
git clone git@github.com:faradey/madock.git
```
If you got error "git@github.com: Permission denied (publickey)." see [solution](https://docs.github.com/en/authentication/troubleshooting-ssh/error-permission-denied-publickey#verify-the-public-key-is-attached-to-your-account)

4. Go to the cloned directory
```shell
cd madock
```
5. Compile
```
Run command below for Apple M1

GOARCH=arm64 go build -o madock
```
```
Run command below for Apple Intel

go build -o madock
```
6. Add `madock` bin into your `$PATH`
```shell
Run command below for Apple M1

ln -s absolute_path_to_your_madock_dir/madock /opt/homebrew/bin/
```
```shell
Run command below for Apple Intel

ln -s absolute_path_to_your_madock_dir/madock /usr/local/bin/
```
7. Open a new terminal tab/window and check that `madock` works
```
which madock
madock
```
8. Optionally you can also apply these performance tweaks
    * [http://markshust.com/2018/01/30/performance-tuning-docker-mac](http://markshust.com/2018/01/30/performance-tuning-docker-mac)
</details>

<details>
<summary>Linux</summary>

1. Install docker
   * Install Docker on [Debian](https://docs.docker.com/engine/installation/linux/docker-ce/debian/)
   * Install Docker on [Ubuntu](https://docs.docker.com/engine/installation/linux/docker-ce/ubuntu/)
   * Install Docker on [CentOS](https://docs.docker.com/engine/installation/linux/docker-ce/centos/)
2. Configure permissions
   * [Manage Docker as a non-root user](https://docs.docker.com/install/linux/linux-postinstall/)
3. Install [Docker-compose](https://docs.docker.com/compose/install/)
4. Install [Golang](https://go.dev/doc/install)
5. Clone this repo and follow into folder "madock"
```
git clone git@github.com:faradey/madock.git
```
If you got error "git@github.com: Permission denied (publickey)." see [solution](https://docs.github.com/en/authentication/troubleshooting-ssh/error-permission-denied-publickey#verify-the-public-key-is-attached-to-your-account)

6. Compile
```
go build -o madock
```
7. Add `madock` bin into your `$PATH`
```
ln -s absolute_path_to_your_madock_dir/madock /usr/local/bin/
```
8. Open a new terminal tab/window and check that `madock` works
```
which madock
madock
```
</details>

## Project Setup
```shell
cd <your_project>
madock setup --download # for a new project
madock setup --install # for an existing project
```

## Usage
### Start Application
```
madock start
madock composer install
sudo vim /etc/hosts
// Add -> 127.0.0.1 <your-domain>
```
### Workflow
See detailed documentation about development workflow with madock
* [Development Workflow](docs/workflow.md)

## More Documentation

* [PHPStorm + Xdebug Setup](docs/xdebug_phpstorm.md)
* [Docker images list](docs/docker_images.md)
* [Customizations](docs/customizations.md)
* [Database import, export, synchronization, phpmyadmin](docs/database.md)
* [Media synchronization](docs/media.md)
* [Cron](docs/cron.md)
* Kibana. URL http://{you_domain_name}/kibana
* Mailhog. Default URL http://localhost:8025

## Donations
If you find it useful and want to invite us for a beer, just click on the donation button. Thanks!

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/faradey)

## Resources
This project has been possible thanks to the following resources:

* [docker-magento](https://github.com/markoshust/docker-magento)
* [dockergento](https://github.com/ModestCoders/magento2-dockergento)

## License

* [The MIT License](https://opensource.org/licenses/MIT)

## Copyright
(c) faradey
