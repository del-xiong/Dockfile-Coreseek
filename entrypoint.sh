#!/bin/bash
/usr/local/bin/indexer --all
echo "start coreseek"
/usr/local/bin/searchd --nodetach
