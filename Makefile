include ./makeUtils.mk

compile:
	$(call build.compile_local)

clean:
	$(call build.clean)

test:
	$(call test.run)

check-fmt:
	$(call fmt.check)

fix-fmt:
	$(call fmt.fix)

build: clean check-fmt compile test

build-docker:
	$(call build.compile_linux)
	$(call docker.build_docker)

run-docker: build-docker
	$(call docker.run_docker)

test-docker:
	$(call docker.build_docker_test)
	$(call docker.run_docker_test)
