# amqp-purge

A simple tool to purge all queues from a rabbitmq cluster using `rabbitmqadmin` which, for example, can be useful for cleaning staging environments.

## Setup

You need to have [rabbitmqadmin](https://www.rabbitmq.com/management-cli.html) in your path before using amqp-purge.

Download the [latest release](https://github.com/selency/amqp-purge/releases) binary and save it to `/usr/local/bin` or any executable path.

## Usage

Purge all queues

```SHELL
amqp-purge --username=foo --password=bar --host=localhost --port=15672
```
