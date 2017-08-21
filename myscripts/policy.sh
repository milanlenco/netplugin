#!/bin/bash

netctl net create contiv-net2 --subnet=10.1.2.0/24 
netctl policy create prodWeb2
#netctl policy rule-add prodWeb 0 -d in -j deny
netctl policy rule-add prodWeb2 1 -direction=in -protocol=tcp -port=88 -action=deny -priority=1
#netctl policy rule-add prodWeb2 2 -i 10.1.2.2  -j deny -l tcp -p 88
netctl policy rule-ls prodWeb2
netctl group create -p prodWeb2 contiv-net web3
