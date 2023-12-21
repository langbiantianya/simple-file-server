goreleaser-snapshot:
	goreleaser release --snapshot --clean

docker-build:goreleaser-snapshot
	docker build -f ./dockerfile -t simplefile-server:latest .