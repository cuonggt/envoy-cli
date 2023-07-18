# Envoy CLI

<a name="introduction"></a>
## Introduction

Envoy CLI is a tool for executing common tasks you run on your remote servers. Using YAML syntax, you can easily setup tasks for deployment, and more.

<a name="installation"></a>
## Installation

Install Envoy CLI from a [prebuilt binary], package manager, or package repository. Please see the installation instructions for your operating system:

- [macOS]
- [Linux]
- [Windows]
- [DragonFly BSD, FreeBSD, NetBSD, and OpenBSD]

<a name="build-from-source"></a>
## Build from source

```shell
go install github.com/cuonggt/envoy-cli@latest
```

<a name="quickstart"></a>
## Quickstart

Once Envoy has been installed, the Envoy binary will be available. Inside your app directory, run `envoy init`. Now edit the new file `Envoyfile`. It could look as simple as this:

```
servers:
  web:
    - user@192.168.1.1

tasks:
  deploy:
    name: Deploy
    script: |
      cd /path/to/site
```

Now you’re ready to run the task `deploy` to deploy to the servers:

```shell
envoy run deploy
```

<a name="writing-tasks"></a>
## Writing Tasks

<a name="defining-tasks"></a>
### Defining Tasks

Tasks are the basic building block of Envoy. Tasks define the shell commands that should execute on your remote servers when the task is invoked. For example, you might define a task that restarts the `supervisor` service on all of your application's queue worker servers.

All of your Envoy tasks should be defined in an `Envoyfile` file at the root of your application. Here's an example to get you started:

```
servers:
  web:
    - user@192.168.1.1
  workers:
    - user@192.168.1.2

tasks:
  restart-supervisor:
    on: workers
    script: |
      sudo supervisorctl restart
```

<a name="local-tasks"></a>
#### Local Tasks

You can force a script to run on your local computer by specifying the server's IP address as `127.0.0.1`:

```
servers:
  localhost: 127.0.0.1
```

<a name="multiple-servers"></a>
### Multiple Servers

Envoy allows you to easily run a task across multiple servers. First, add additional servers to your `servers` declaration. Each server should be assigned a unique name. Once you have defined your additional servers you may list each of the servers in the task's `on` array:

```
servers:
  web-1: 192.168.1.1
  web-2: 192.168.1.2

tasks:
  deploy:
    on:
      - web-1
      - web-2
    script: |
      cd /home/user/example.com
      git pull origin master
```

<a name="parallel-execution"></a>
#### Parallel Execution

By default, tasks will be executed on each server serially. In other words, a task will finish running on the first server before proceeding to execute on the second server. If you would like to run a task across multiple servers in parallel, add the `parallel` option to your task declaration:

```
servers:
  web-1: 192.168.1.1
  web-2: 192.168.1.2

tasks:
  deploy:
    on:
      - web-1
      - web-2
    parallel: true
    script: |
      cd /home/user/example.com
      git pull origin master
```

<a name="stories"></a>
### Stories

Stories group a set of tasks under a single, convenient name. For instance, a `deploy` story may run the `update-code` and `build-code` tasks by listing the task names within its definition:

```
servers:
  web:
    - user@192.168.1.1

tasks:
  update-code:
    cd /home/user/example.com
    git pull origin master
  build-code:
    cd /home/user/example.com
    npm install
    npm run build

stories:
  deploy:
    - update-code
    - build-code
```

Once the story has been written, you may invoke it in the same way you would invoke a task:

```shell
envoy run deploy
```

<a name="running-tasks"></a>
## Running Tasks

To run a task or story that is defined in your application's `Envoyfile` file, execute Envoy's `run` command, passing the name of the task or story you would like to execute. Envoy will execute the task and display the output from your remote servers as the task is running:

```shell
envoy run deploy
```

<a name="notifications"></a>
## Notifications

<a name="slack"></a>
### Slack

<a name="discord"></a>
### Discord

<a name="telegram"></a>
### Telegram

<a name="microsoft-teams"></a>
### Microsoft Teams