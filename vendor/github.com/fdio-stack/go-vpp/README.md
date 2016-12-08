# GO vppapi server and testclient

## DISCLAIMER AND TODOS

This code is the kind of script kiddy douche baggery that gives programming a bad name.

There's a very simple pattern that needs to be done to clean this up:

1. VPP APIs and their callbacks follow a simple regex pattern
2. A simple library of a dictionary structs with parms to, values returned for each API may suffice (side effect)
3. A mutex or similar to register in a dictionary the calls and callbacks with said struct passed as parm
4. Hand that struct back to GO calling function ie

Instead of the sync.Wait wanton fuckery repeated everywhere simple have a func that has stuff like:
err, _ = knit_me_an_af_packet(*af_packet_struct)
    where the struct has the sw_if_index (side effect) and the name of interface

This simple idea can be extended to all CRUD including dumps ... but proper programming takes time
and attention span, of which I am in rapidly decreasing supply.

This doesn't need to be a full channel implemented War and Peace science project, but I'm super uncomfortable
releasing this into the wild as is... its the GO equivalent of a shitty BASH script....

## Check VPP libraries installed

```
 alagalah@thing1:~ $ ldconfig -p | grep vpp
     libvppinfra.so.0 (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libvppinfra.so.0
     libvppinfra.so (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libvppinfra.so
     libjvpp_registry.so.0 (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libjvpp_registry.so.0
     libjvpp_registry.so (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libjvpp_registry.so
     libjvpp_core.so.0 (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libjvpp_core.so.0
     libjvpp_core.so (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libjvpp_core.so
     libjvpp_common.so.0 (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libjvpp_common.so.0
     libjvpp_common.so (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libjvpp_common.so
     libgmodvpp.so (libc6,x86-64) => /usr/lib/x86_64-linux-gnu/libgmodvpp.so
```

## Do once - assumes you are in project root.

```
  cd c
  autoreconf -i -f
  ./configure
  make
  sudo make install
  sudo ldconfig -v
```

## From emacs or CLI - assumes you are in project root
```
  make
```
## To test 
### (assumes you have docker installed with user added to docker group)
```
  ./dockerenv.sh
  docker ps
  make testclient
  sudo service vpp restart
  sudo ./testclient
```
