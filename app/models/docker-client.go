package models

import (
	"fmt"
	"hash/fnv"
	"regexp"
	"strings"

	"github.com/pottava/docker-webui/app/config"
	"github.com/pottava/docker-webui/app/misc"
)

// DockerClient represents a Docker client
type DockerClient struct {
	ID        string `json:"id"`
	Endpoint  string `json:"endpoint"`
	CertPath  string `json:"certPath"`
	IsActive  bool   `json:"isActive"`
	IsDefault bool   `json:"isDefault"`
}

var dockerClientSavePath string

func init() {
	r, _ := regexp.Compile("[^a-zA-Z0-9_\\.]")
	name := r.ReplaceAllString(strings.ToLower(config.NewConfig().Name), "-")
	dockerClientSavePath = "/tmp/" + name + "-docker-clients.json"
}

// LoadDockerClients returns registered clients
func LoadDockerClients() (clients []*DockerClient, err error) {
	err = misc.ReadFromFile(dockerClientSavePath, &clients)
	return
}

// RemoveDockerClient removes your specified configuration
func RemoveDockerClient(id string) bool {
	prev, _ := LoadDockerClients()
	next := []*DockerClient{}
	for _, client := range prev {
		if client.ID == id {
			if client.IsDefault {
				return false
			}
			continue
		}
		next = append(next, client)
	}
	misc.SaveAsFile(dockerClientSavePath, next)
	return true
}

// RemoveDockerClientByEndpoint removes your specified configuration
func RemoveDockerClientByEndpoint(endpoint string) {
	RemoveDockerClient(fmt.Sprint(hash(endpoint)))
}

// Load select its configuration
func (c *DockerClient) Load() {
	clients, _ := LoadDockerClients()
	for _, client := range clients {
		if client.Endpoint == c.Endpoint {
			c.ID = client.ID
			c.CertPath = client.CertPath
			c.IsActive = client.IsActive
			c.IsDefault = client.IsDefault
			break
		}
	}
}

// Save persists the client configuration
func (c *DockerClient) Save() {
	c.ID = fmt.Sprint(hash(c.Endpoint))

	clients, _ := LoadDockerClients()
	found := false
	for _, client := range clients {
		if client.Endpoint == c.Endpoint {
			client.CertPath = c.CertPath
			client.IsActive = c.IsActive
			found = true
		}
	}
	if !found {
		clients = append(clients, c)
	}
	misc.SaveAsFile(dockerClientSavePath, clients)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
