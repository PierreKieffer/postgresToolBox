#!/bin/bash 

docker run --rm  --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432   postgres
