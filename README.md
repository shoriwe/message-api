# message-api

[![Build](https://github.com/shoriwe/message-api/actions/workflows/build.yaml/badge.svg)](https://github.com/shoriwe/message-api/actions/workflows/build.yaml)
[![codecov](https://codecov.io/gh/shoriwe/message-api/branch/main/graph/badge.svg?token=RU4KKCQPUV)](https://codecov.io/gh/shoriwe/message-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/shoriwe/message-api)](https://goreportcard.com/report/github.com/shoriwe/message-api)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hidromatologia-v2/stations)

REST API for [message-app](https://github.com/shoriwe/message-app)

## Documentation

| File                                             | Description                                |
| ------------------------------------------------ | ------------------------------------------ |
| [docs/spec.openapi.yaml](docs/spec.openapi.yaml) | OpenAPI documentation for the REST service |
| [docs/CLI.md](docs/CLI.md)                       | CLI documentation for the main binary      |

## Running

### Docker

```shell
docker pull ghcr.io/shoriwe/message-api:latest
```

### Docker compose

Make sure to customize the [example.env](example.env) file. Then:

```shell
docker compose up -d
```

### Binary

You can find pre-compiled binaries at the [releases](https://github.com/shoriwe/message-api/releases) section.

```shell
export SECRET = "MY_SECRET"
export FIREBASE_PROJECT_ID = "MY_FIREBASE_PROJECT"
export FIREBASE_CONFIGURATION_FILE = "firebase-adminsdk.json"
export DATABASE_FILE = "database.db"

message-api :5000
```

#### Building from source

```shell
go install github.com/shoriwe/message-api@latest
```

