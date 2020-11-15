package audhumla

import "encoding/json"

type Library struct {
	Name     string            `json:"name"`
	Artifact File              `json:"artifact"`
	Natives  map[string]File   `json:"natives"`
	Rules    []json.RawMessage `json:"rules"`
}
