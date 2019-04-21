#!/bin/bash
/usr/local/bin/indexer --all
echo "start coreseek"
nohup /usr/local/etc/sphinx/coreseek_autoindex > /dev/null 2>&1 &
/usr/local/bin/searchd --nodetach
