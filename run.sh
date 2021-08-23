#!/bin/zsh

# This is the bare minimum to run in development. For full list of flags,
# run ./service-monitor -help

go build -o service-monitor cmd/web/*.go && ./service-monitor \
-dbuser='dapperblondie' \
-pusherHost='localhost' \
-pusherKey='abc123' \
-pusherSecret='123abc' \
-pusherApp="1"
-pusherPort="4001"
-pusherSecure=false