# go-docker

go-docker is a Go client library for accessing the [Docker Remote API](http://docs.docker.io/en/latest/api/docker_remote_api/).

## NOTE

We recommend using this Docker package: https://github.com/fsouza/go-dockerclient

The primary different between the two packages are:

* we've removed the CGO dependency by copying the Docker structs locally into the project
* we've added functionality to create Docker images
* go-dockerclient is better tested and is definitely more actively maintained

We hope that eventually we can just use `go-dockerclient` since there is no reason to
maintain two different implementations.

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