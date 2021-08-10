package util

import (
	"encoding/json"
)

// ParseAppDescriptor parse the application descriptor
func ParseAppDescriptor(appJson string) (*AIflowAppDescriptor, error) {
	descriptor := &AIflowAppDescriptor{}

	err := json.Unmarshal([]byte(appJson), descriptor)

	if err != nil {
		return nil, err
	}

	return descriptor, nil
}

// AIflowAppDescriptor is the descriptor for a AIflow application
type AIflowAppDescriptor struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	AppModel    string   `json:"appModel,omitempty"`
	Imports     []string `json:"imports"`

	Triggers []*AIflowTriggerConfig `json:"triggers"`
}

type AIflowTriggerConfig struct {
	Id   string `json:"id"`
	Ref  string `json:"ref"`
	Type string `json:"type"`
}
