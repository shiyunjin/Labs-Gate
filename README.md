# Labs;Gate
[![Build Status](https://travis-ci.com/shiyunjin/Labs-Gate.svg?branch=master)](https://travis-ci.com/shiyunjin/Labs-Gate)
[![Docker Build Status](https://img.shields.io/docker/build/shiyunjin/labs-gate.svg)](https://hub.docker.com/r/shiyunjin/labs-gate/)
[![Coverage Status](https://coveralls.io/repos/github/shiyunjin/Labs-Gate/badge.svg?branch=master)](https://coveralls.io/github/shiyunjin/Labs-Gate?branch=master)
[![Maintainability](https://api.codeclimate.com/v1/badges/f8f91e33ba07913cecb9/maintainability)](https://codeclimate.com/github/shiyunjin/Labs-Gate/maintainability)

University laboratory network management system

# Deploy
[click here](https://github.com/shiyunjin/Labs-Gate-Deploy)

# Development
you need golang node mongodb to create a development environment

use `go get ./...` Install go dependencies

in view to [click here](https://github.com/shiyunjin/Labs-Gate-UI) in `/system/view/`

## config file
must need config file to run

copy `config-dev.json` to `config.json`

`config-dev.json` is development config file with local mongodb

`config.tson` is deploy config file template with `run.sh`
