#!/bin/bash
  
echo "start coreseek"
/usr/local/bin/indexer --all
/usr/local/bin/searchd --nodetach