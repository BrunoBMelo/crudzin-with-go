# crudzin-with-go

```docker docker build -t crudapp:a .```

```docker docker run -p 3001:3001 --network crudzin-with-go_db -e PG_STRT_CONNECTION="host=golang-crud port=5432 user=PG_USERGO password=PG_PWD2017 dbname=PG_DBGO sslmode=disable" -e PROVIDERNAME=postgres -e APP_PORT=3001 crudapp:1```
