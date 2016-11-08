# How-To's

## Testing with https://github.com/lambci/docker-lambda

1. Clone the repository
2. `cd` into it

### Basic usage

```shell
docker run -v "$PWD":/var/task lambci/lambda index.handler '{"repository": "github.com/chreble/todo"}'
```

### Passing a reference

```shell
docker run -v "$PWD":/var/task lambci/lambda index.handler '{"repository": "github.com/chreble/todo", "reference": "$BRANCH"}'
```

## Makefile

Build the zipfile

`make` or `make build`

Compile `exago-runner` (must be in the `$GOPATH`) and replace the existing binary

`make compile`
