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
goreleaser-release:
	goreleaser release --skip=publish --clean

goreleaser-snapshot:
	goreleaser --snapshot --skip=publish --clean --skip-validate

docker-debian-bookworm-build:docker-goreleaser-release
	docker build -f ./deploy/debian/Dockerfile -t simplefile-server:debian-bookworm .

docker-opensuseTumbleweed-build:docker-goreleaser-release
	docker build -f ./deploy/opensuse/Dockerfile -t simplefile-server:opensuse-tumbleweed .

docker-openEulr-20.03-build:docker-goreleaser-release
	docker build -f ./deploy/openEulr/Dockerfile -t simplefile-server:openeulr-20.03-lts .

docker-build:docker-debian-bookworm-build