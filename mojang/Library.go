package mojang

import "encoding/json"

type Library struct {
	Name      string
	Downloads struct {
		Artifact    File
		Classifiers map[string]File
	} `json:"downloads"`
	Natives map[string]string
	Rules   []json.RawMessage
}
