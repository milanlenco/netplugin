#!/usr/bin/env bash

# Copyright (c) 2016 Cisco and/or its affiliates.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at:
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

function pause(){
   read -p "$*"
}

# Restart VPP so there is no config
echo "Restarting vpp so there is no config..."
echo "$ sudo service vpp restart"
sudo service vpp restart

# Restart host
pause 'Press [Enter] key to restart host services'
echo "Restarting host services..."
echo "$ make host-restart"
make host-restart

# Show vpp initial interface state
pause 'Press [Enter] key to show vpp interfaces'
echo "$ sudo vppctl sh int"
sudo vppctl sh int

# Show vpp initial bridge domain state
pause 'Press [Enter] key to show vpp bridge domains'
echo "sudo vppctl sh br"
sudo vppctl sh br

# Create a network
pause 'Press [Enter] key to create a contiv net'
echo "$ netctl net create contiv-net --subnet=20.1.1.0/24"
netctl net create contiv-net --subnet=20.1.1.0/24

# Check if vpp created the bridge domain
pause 'Press [Enter] key to see the created nework in contiv and vpp'
echo "netctl net ls"
netctl net ls
echo "sudo vppctl sh br"
sudo vppctl sh br

# Create two docker containers attached to contiv-net
pause 'Press [Enter] key to create two docker containers attached to contiv-net'
echo "$ docker run -itd --name=web --net=contiv-net alpine /bin/sh"
docker run -itd --name=web --net=contiv-net alpine /bin/sh &>/dev/null
echo "$ docker run -itd --name=db --net=contiv-net alpine /bin/sh"
docker run -itd --name=db --net=contiv-net alpine /bin/sh &>/dev/null

etcdctl ls /contiv.io/oper/eps


# Show vpp initial interface state
pause 'Press [Enter] key to show vpp interfaces'
echo "$ sudo vppctl sh int"
sudo vppctl sh int

# Show vpp initial bridge domain state with details
pause 'Press [Enter] key to show vpp bridge domain 1 details'
echo "$ sudo vppctl sh br 1 detail"
sudo vppctl sh br 1 detail