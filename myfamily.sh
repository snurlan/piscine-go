#! /bin/bash

curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -r " .[] | select( .id == $HERO_ID ) | .connections.relatives" | tr -d \" | sed ':a;N;$!ba;s/\n/\\n/g'
