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
