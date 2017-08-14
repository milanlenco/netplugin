#!/bin/bash

sudo pkill netmaster
sudo pkill netplugin

source /etc/profile.d/envvar.sh
cd /opt/gopath/src/github.com/contiv/netplugin
make run-build install-shell-completion
