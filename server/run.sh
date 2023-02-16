#!/usr/bin/env sh

set -ex

go run ./... --local ./notes --remote "s3://notescollector/notes.txt" --remote-options "endpoint=s3.fr-par.scw.cloud&profile=scaleway"
