package docker

import (
	"fmt"
	"io"
	"time"
)

type Images struct {
	ID          string   `json:"Id"`
	RepoTags    []string `json:",omitempty"`
	Created     int64
	Size        int64
	VirtualSize int64
	ParentId    string `json:",omitempty"`

	// DEPRECATED
	Repository string `json:",omitempty"`
	Tag        string `json:",omitempty"`
}

type Image struct {
	ID              string    `json:"id"`
	Parent          string    `json:"parent,omitempty"`
	Comment         string    `json:"comment,omitempty"`
	Created         time.Time `json:"created"`
	Container       string    `json:"container,omitempty"`
	ContainerConfig Config    `json:"container_config,omitempty"`
	DockerVersion   string    `json:"docker_version,omitempty"`
	Author          string    `json:"author,omitempty"`
	Config          *Config   `json:"config,omitempty"`
	Architecture    string    `json:"architecture,omitempty"`
	Size            int64
}

type Delete struct {
	Deleted  string `json:",omitempty"`
	Untagged string `json:",omitempty"`
}

type ImageService struct {
	*Client
}

// List Images
func (c *ImageService) List() ([]*Images, error) {
	images := []*Images{}
	err := c.do("GET", "/images/json?all=0", nil, &images)
	return images, err
}

// Create an image, either by pull it from the registry or by importing it.
func (c *ImageService) Create(image string) error {
	return c.do("POST", fmt.Sprintf("/images/create?fromImage=%s"), nil, nil)
}

func (c *ImageService) Pull(name, tag string, in io.Reader, out io.Writer) error {
	path := fmt.Sprintf("/images/create?fromImage=%s&tag=%s", name, tag)
	//path := fmt.Sprintf("/images/create?fromImage=%s", name)
	return c.stream("POST", path, in, out)
}

// Remove the image name from the filesystem
func (c *ImageService) Remove(image string) ([]*Delete, error) {
	resp := []*Delete{}
	err := c.do("DELETE", fmt.Sprintf("/images/%s", image), nil, &resp)
	return resp, err
}

// Inspect the image
func (c *ImageService) Inspect(name string) (*Image, error) {
	image := Image{}
	err := c.do("GET", fmt.Sprintf("/images/%s/json", name), nil, &image)
	return &image, err
}
