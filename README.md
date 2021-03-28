# go-shopping-cart

### Commands

1. Docker Mode - Requires docker installed on your local

- To build the docker image
```
    make build-docker
```

- To run the docker container
```
    make run-docker
```

- To run the tests in docker
```
    make test-docker
```

2. Local Mode - Requires Go 1.15 installed on your local

- To compile the source and build the executable
```
    make compile
```

- To clean the build dir
```
    make clean
```

- To run tests
```
    make test
```

- To check for formatting errors
```
    make check-fmt
```

- To fix formatting errors
```
    make fix-fmt
```
