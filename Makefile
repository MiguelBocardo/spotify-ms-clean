.PHONY: run build docker up loadtest

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

docker:
	docker build -t spotify-ms-clean:latest .

up:
	docker compose up --build

loadtest:
	vegeta attack -duration=10s -rate=200 -targets=loadtest/targets.txt | tee results.bin | vegeta report
	vegeta plot results.bin > loadtest/plot.html
