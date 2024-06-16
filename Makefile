db/init:
	podman run -d -P -p 127.0.0.1:5432:5432 -e POSTGRES_PASSWORD="letsgo" --name postgres postgres

db/start:
	podman start postgres

db/stop:
	podman stop postgres

db/login:
	psql postgresql://postgres:letsgo@localhost:5432/postgres

db/backup:
	pg_dumpall -U postgres -h localhost -p 5432 > backup.sql


db/restore:
	psql -f backup.sql postgresql://postgres:letsgo@localhost:5432/postgres

run:
	go run ./cmd/api
