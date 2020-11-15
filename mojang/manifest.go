package mojang

import "encoding/json"

type Manifest struct {
	Arguments struct {
		Game []json.RawMessage `json:"game"`
		JVM  []json.RawMessage `json:"jvm"`
	} `json:"arguments"`
	AssetIndex struct {
		Id        string
		Hash      string `json:"sha1"`
		Size      int
		TotalSize int `json:"totalSize"`
	} `json:"assetIndex"`
	Assets          string
	ComplianceLevel int `json:"complianceLevel"`
	Downloads       Downloads
	ID              string `json:"id"`
	Libraries       []Library
	MainClass       string `json:"mainClass"`
	Time            string `json:"releaseTime"`
	Type            string `json:"type"`
}
