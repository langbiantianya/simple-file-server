goreleaser-snapshot:
	goreleaser build

docker-build:goreleaser-snapshot
	docker build -f ./dockerfile -t simplefile-server:latest .