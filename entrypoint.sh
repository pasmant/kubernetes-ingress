#!/usr/bin/env bash

# create shared cores if it missing before starting admd
mkdir -p /shared/cores && chmod 777 /shared/cores

/usr/bin/adminstall
/usr/bin/admd -d --standalone > /var/log/adm/admd.log 2>&1
