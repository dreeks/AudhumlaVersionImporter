package audhumla

import "encoding/json"

type Argument struct {
	Values []string `json:"values"`
	Rules  []json.RawMessage
}
