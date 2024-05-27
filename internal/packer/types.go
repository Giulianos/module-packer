package packer

import (
	"gopkg.in/yaml.v3"
	"os"
)

type PackingSpec struct {
	KernelPath string       `yaml:"kernel_path"`
	TargetPath string       `yaml:"target_path"`
	Modules    []ModuleSpec `yaml:"modules"`
}

type ModuleSpec struct {
	Path       string            `yaml:"path"`
	Attributes map[string]string `yaml:"attributes"`
}

func LoadPackingSpecFromFile(path string) (*PackingSpec, error) {
	packingSpecFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer packingSpecFile.Close()

	var packingSpec PackingSpec
	err = yaml.NewDecoder(packingSpecFile).Decode(&packingSpec)
	if err != nil {
		return nil, err
	}

	return &packingSpec, nil
}
