#!/bin/bash

source /etc/profile.d/envvar.sh
cd /opt/gopath/src/github.com/contiv/netplugin/scripts/python
PYTHONIOENCODING=utf-8 ./startPlugin.py -nodes ${CLUSTER_NODE_IPS}
