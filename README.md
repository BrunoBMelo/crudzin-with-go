# crudzin-with-go

sudo docker run -p 3001:3001 --network crudzin-with-go_db -e POSTGRES_DB=PG_DBGO -e POSTGRES_USER=PG_USERGO -e POSTGRES_PASSWORD=PG_PWD2017 -e POSTGRES_PORT=5432 -e POSTGRES_PORT_EXPOSE=5432 -e PG_STRT_CONNECTION="host=golang-crud port=5432 user=PG_USERGO password=PG_PWD2017 dbname=PG_DBGO sslmode=disable" -e PROVIDERNAME=postgres -e APP_PORT=3001 crud-app:1