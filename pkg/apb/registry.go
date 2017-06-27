package apb

import (
	"fmt"

	logging "github.com/op/go-logging"
)

// RegistryConfig - Configuration for the registry
type RegistryConfig struct {
	Name string
	URL  string
	User string
	Pass string
	Org  string // Target org to load playbook bundles from
}

// Registry - Interface that wraps the methods need for a registry
type Registry interface {
	Init(RegistryConfig, *logging.Logger) error
	LoadSpecs() ([]*Spec, int, error)
}

// NewRegistry - Create a new registry from the registry config.
func NewRegistry(config RegistryConfig, log *logging.Logger) (Registry, error) {
	var reg Registry

	log.Info("== REGISTRY CX == ")
	log.Info(fmt.Sprintf("Name: %s", config.Name))
	log.Info(fmt.Sprintf("Url: %s", config.URL))

	switch config.Name {
	case "rhcc":
		reg = &RHCCRegistry{}
	case "dockerhub":
		reg = &DockerHubRegistry{}
	case "mock":
		reg = &MockRegistry{}
	default:
		panic("Unknown registry")
	}

	err := reg.Init(config, log)
	if err != nil {
		return nil, err
	}

	return reg, err
}
