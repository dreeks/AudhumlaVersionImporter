package mojang

type File struct {
	URL  string `json:"url"`
	Hash string `json:"sha1"`
	Path string `json:"path"`
	Size int
}
