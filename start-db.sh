#!/bin/bash 

docker run --rm  --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432   postgres

sleep 2
postgres_host=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgres)

psql "user='postgres' password='postgres' host='$postgres_host'" -f create-airflow.sql



