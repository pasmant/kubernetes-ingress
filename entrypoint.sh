#!/usr/bin/env bash
/usr/bin/adminstall
/usr/bin/admd -d --standalone > /var/log/adm/admd.log 2>&1
