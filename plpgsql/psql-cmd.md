### docker start 
docker run --network host -e POSTGRES_PASSWORD=postgres -d postgres

### Connect to instance 
```
psql -h localhost -U postgres -d postgres
```

### Execute plpsql file 
```
psql -h 127.0.0.1 -U postgres -f create-table.sql
```
