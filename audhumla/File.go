package audhumla

type File struct {
	URL      string `json:"url"`
	Path     string `json:"path"`
	Hash     string `json:"hash"`
	Filesize int64  `json:"filesize"`
}
