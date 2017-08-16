#!/bin/bash

godep_save() {
  missing_package="start"
  while [ "$missing_package" != "" ]
   do
    missing_package=$(godep save $@ 2>&1 | \
      egrep '^godep: Package (.*) not found' | \
      sed 's/.*(\(.*\)).*/\1/');
   [ "$missing_package" != "" ] && {
     echo "Installing missing package: ${missing_package}" ;
     strace -f -e trace=network go get -u "${missing_package}" 2>&1 | pv -i 0.05 > /dev/null 
    }
  done
  godep save $@
}

godep_save $@

