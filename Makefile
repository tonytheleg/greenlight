db/start:
	podman start postgres

db/stop:
	podman stop postgres


run:
	go run ./cmd/api