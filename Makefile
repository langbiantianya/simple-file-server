docker-goreleaser-snapshot:
	docker run \
		--rm \
		-e CGO_ENABLED=1\
		-e GO111MODULE=on\
		-e GOPROXY=https://goproxy.cn \
		-v `pwd`:/go/src/simple-file-server \
		-w /go/src/simple-file-server \
		goreleaser/goreleaser-cross:latest \
		--snapshot --skip=publish --clean --skip-validate
docker-goreleaser-release:
	docker run \
		--rm \
		-e CGO_ENABLED=1\
		-e GO111MODULE=on\
		-e GOPROXY=https://goproxy.cn \
		-v `pwd`:/go/src/simple-file-server \
		-w /go/src/simple-file-server \
		goreleaser/goreleaser-cross:latest \
		release --skip=publish --clean

# goreleaser-release:
# 	goreleaser release --skip=publish --clean

# goreleaser-snapshot:
# 	goreleaser --snapshot --skip=publish --clean --skip-validate

docker-debian-bookworm-snapshot:docker-goreleaser-snapshot
	docker build -f ./deploy/debian/Dockerfile -t simplefile-server:debian-bookworm-snapshot .

docker-debian-bookworm-release:docker-goreleaser-release
	docker build -f ./deploy/debian/Dockerfile -t simplefile-server:debian-bookworm-release .

docker-opensuseTumbleweed-snapshot:docker-goreleaser-snapshot
	docker build -f ./deploy/opensuse/Dockerfile -t simplefile-server:opensuse-tumbleweed-snapshot .

docker-opensuseTumbleweed-release:docker-goreleaser-release
	docker build -f ./deploy/opensuse/Dockerfile -t simplefile-server:opensuse-tumbleweed-release .

docker-openEulr-snapshot:docker-goreleaser-snapshot
	docker build -f ./deploy/openEulr/Dockerfile -t simplefile-server:openeulr-snapshot .

docker-openEulr-release:docker-goreleaser-release
	docker build -f ./deploy/openEulr/Dockerfile -t simplefile-server:openeulr-release .

docker-build-snapshot:docker-openEulr-snapshot

docker-build-release:docker-openEulr-release