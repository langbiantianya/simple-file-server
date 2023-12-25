goreleaser-snapshot:
	goreleaser release --skip=publish --clean

docker-build:goreleaser-snapshot
	docker build -f ./dockerfile -t simplefile-server:latest .