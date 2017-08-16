#!/bin/bash

export CONTIV_NODES=2
export CONTIV_NODE_OS=ubuntu
export CONTIV_DOCKER_VERSION=1.13.0
export VAGRANT_DEFAULT_PROVIDER=virtualbox
VAGRANT_HTTP_PROXY=http://proxy-wsa.esl.cisco.com:80 VAGRANT_HTTPS_PROXY=http://proxy-wsa.esl.cisco.com:80 vagrant up
