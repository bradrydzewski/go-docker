# go-docker

go-docker is a Go client library for accessing the [Docker Remote API](http://docs.docker.io/en/latest/api/docker_remote_api/).

## NOTE

We recommend using this Docker package: https://github.com/fsouza/go-dockerclient

This package adds some missing functionality we needed but was missing in the above package. We
also chose to copy-paste all the Docker structs directy into our project to avoid linking Docker
directly, since it has a CGO dependency.

Once the above issues are resolved we'll probably end up switching to the go-dockerclient package,
since it doesn't make sense to maintain two different package libraries.

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