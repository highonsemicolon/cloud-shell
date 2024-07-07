#!bin/bash
id=$1
docker stop $id || true && docker rm $id
