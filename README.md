# go-docker

go-docker is a Go client library for accessing the [Docker Remote API](http://docs.docker.io/en/latest/api/docker_remote_api_v1.6/).

## Usage

```go
import "github.com/bradrydzewski/go-docker/docker"
```

Construct a new Docker client, then use the various services on the client
to access different parts of the API. For example, to retrieve a list of
containers:

```go
client := docker.New()
containers, err := client.Containers.List()
```
