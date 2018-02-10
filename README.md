# consul reader bot in slack

[![Build Status](https://travis-ci.org/matsu-chara/conbot.svg?branch=master)](https://travis-ci.org/matsu-chara/conbot)
[![Go Report Card](https://goreportcard.com/badge/github.com/matsu-chara/conbot)](https://goreportcard.com/report/github.com/matsu-chara/conbot)

## Install

- git clone
- get token bot
- make
- add run.sh and execute

```sh
#!/bin/bash

set -eu

export CONBOT_SLACK_BOT_TOKEN="YOUR_TOKEN"
export CONBOT_CONSUL="localhost:8500"
make build
./conbot
```

## Usage

`@conbot help`

- catalog: get catalog
- note: get notes
- node: get nodes

