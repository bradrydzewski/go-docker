package docker

import (
	"fmt"
	"io"
)

type ContainerService struct {
	*Client
}

// List only running containers.
func (c *ContainerService) List() ([]*Containers, error) {
	containers := []*Containers{}
	err := c.do("GET", "/containers/json?all=0", nil, &containers)
	return containers, err
}

// List all containers
func (c *ContainerService) ListAll() ([]*Containers, error) {
	containers := []*Containers{}
	err := c.do("GET", "/containers/json?all=1", nil, &containers)
	return containers, err
}

// List all containers
func (c *ContainerService) Create(conf *Config) (*Run, error) {
	run := Run{}
	err := c.do("POST", "/containers/create", conf, &run)
	return &run, err
}

// Start the container id
func (c *ContainerService) Start(id string) error {
	return c.do("POST", fmt.Sprintf("/containers/%s/start", id), &HostConfig{}, nil)
}

// Stop the container id
func (c *ContainerService) Stop(id string, timeout int) error {
	return c.do("POST", fmt.Sprintf("/containers/%s/stop?t=%v", id, timeout), nil, nil)
}

// Remove the container id from the filesystem.
func (c *ContainerService) Remove(id string) error {
	return c.do("DELETE", fmt.Sprintf("/containers/%s", id), nil, nil)
}

// Block until container id stops, then returns the exit code
func (c *ContainerService) Wait(id string) (*Wait, error) {
	wait := Wait{}
	err := c.do("POST", fmt.Sprintf("/containers/%s/wait", id), nil, &wait)
	return &wait, err
}

// Attach to the container to stream the stdout and stderr
func (c *ContainerService) Attach(id string, out io.Writer) error {
	path := fmt.Sprintf("/containers/%s/attach?&stream=1&stdout=1&stderr=1", id)
	return c.hijack("POST", path, false, out)
}

// Stop the container id
func (c *ContainerService) Inspect(id string) (*Container, error) {
	container := Container{}
	err := c.do("GET", fmt.Sprintf("/containers/%s/json", id), nil, &container)
	return &container, err
}

// Run the container
func (c *ContainerService) Run(conf *Config, out io.Writer) (*Wait, error) {
	run, err := c.Create(conf)
	if err != nil {
		return nil, err
	}

	// attach to the container
	go func() {
		c.Attach(run.ID, out)
	}()

	// start the container
	if err := c.Start(run.ID); err != nil {
		return nil, err
	}

	// wait for the container to stop
	wait, err := c.Wait(run.ID)
	if err != nil {
		return nil, err
	}

	return wait, nil
}

// Run the container as a Daemon
func (c *ContainerService) RunDaemon(conf *Config) (*Run, error) {
	run, err := c.Create(conf)
	if err != nil {
		return nil, err
	}

	// start the container
	err = c.Start(run.ID)
	return run, err
}
