name := go-shopping-cart
build_out := out/build
bin := $(build_out)/$(name)
linux_bin := $(build_out)/$(name)-linux
main_dir := ./
main_dir_files := ./...
test_docker_image_name := go-shopping-cart-test
test_docker_file := ./Dockerfile.test

define build.compile_local
	go build -o $(bin) $(main_dir)
endef

define build.compile_linux
	GOOS=linux go build -o $(linux_bin) $(main_dir)
endef

define build.clean
	go clean -cache -testcache $(main_dir)
endef

define test.run
	go test -coverpkg=$(main_dir_files) -coverprofile c.out $(main_dir_files)
	go tool cover -func c.out
endef

define fmt.check
	gofmt -d -l -e $(main_dir)
endef

define fmt.fix
	gofmt -d -l -w $(main_dir)
endef

define docker.build_docker
	docker build . -t $(name)
endef

define docker.run_docker
	docker run $(name)
endef

define docker.build_docker_test
	docker build -t $(test_docker_image_name) -f ${test_docker_file} .
 endef

define docker.run_docker_test
	docker run $(test_docker_image_name)
endef
