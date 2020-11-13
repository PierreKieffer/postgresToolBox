# postgresToolBox 

## Install postgresql-client 
```
sudo apt install -y postgresql-client
```

## deploy 
Deploy a postgresql instance in a container
```bash
./start-db.sh
```

### Connect to instance 
Connect to host : 
```bash 
postgres_host=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgres)

psql -h $postgres_host -U postgres -d postgres
```

## plpsql 
Psql script set.
To execute a script : 

```bash
psql -h $postgres_host -U postgres -f script.sql
```

```bash
psql "user='postgres' password='postgres' host='$postgres_host'" -f plpgsql/create-airflow.sql
```

## listener 
Streaming service listening on a channel for notifications of actions on a table
