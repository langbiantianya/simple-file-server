goreleaser-release:
	goreleaser release --skip=publish --clean

goreleaser-snapshot:
	goreleaser --snapshot --skip=publish --clean --skip-validate

docker-debian-bookworm-build:goreleaser-snapshot
	docker build -f ./deploy/debian/Dockerfile -t simplefile-server:debian-bookworm .

docker-opensuseTumbleweed-build:goreleaser-snapshot
	docker build -f ./deploy/opensuse/Dockerfile -t simplefile-server:opensuse-tumbleweed .

docker-openEulr-20.03-build:goreleaser-snapshot
	docker build -f ./deploy/openEulr/Dockerfile -t simplefile-server:openeulr-20.03-lts .

docker-build:docker-openEulr-20.03-build
	docker build -f ./deploy/openEulr/Dockerfile -t simplefile-server:latest .