createdb:
	docker run --name some-postgres \
	-e POSTGRES_USER=myuser \
	-e POSTGRES_PASSWORD=mysecretpassword \
	-e POSTGRES_DB=mydatabase \
	-p 5432:5432 -d postgres

runserver:
	go run cmd/main.go

.PHONY: createdb runserver