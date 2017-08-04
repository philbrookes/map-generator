compile_to_linux_docker: build_linux_api docker_build

docker_build:
	docker build -t generator .

build_linux_api:
	cd cmd/rest-api && env GOOS=linux go build .

build_linux_generator:
	cd cmd/generator && env GOOS=linux go build .

build_mac_api:
	cd cmd/rest-api && env GOOS=darwin go build .

build_mac_generator:
	cd cmd/generator && env GOOS=darwin go build .

build_linux: build_linux_api build_linux_generator
build_mac: build_mac_api build_mac_generator