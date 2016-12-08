# Copyright (c) 2015 Cisco and/or its affiliates.
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

AUTOMAKE_OPTIONS = foreign subdir-objects


AM_CFLAGS = "-Wall"

#C_INCLUDE_PATH = /home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/vlib-api/lib64:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/vnet/lib64:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/svm/lib64:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/vppinfra/lib64:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/vlib-api/lib64:$(C_INCLUDE_PATH)

lib_LTLIBRARIES = libvpp_cgoclient.la

#LIBRARY_PATH = /home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/vlib-api/lib64:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/vnet/lib64:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/svm/lib64:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/vppinfra/lib64:/home/alagalah/git/work/fdio/vpp/build-root/install-vpp_debug-native/vlib-api/lib64:$(LIBRARY_PATH)

#noinst_PROGRAMS = vpp_client

#vpp_client_SOURCES = vpp_client.c
#vpp_client_LDADD = -lvlibmemoryclient -lvlibapi -lsvm -lvppinfra -lpthread -lm -lrt

libvpp_cgoclient_la_SOURCES = vpp_client.c
libvpp_cgoclient_la_LIBADD = -lvlibmemoryclient -lvlibapi -lsvm -lvppinfra -lpthread -lm -lrt
pkginclude_HEADERS = vpp_client.h

