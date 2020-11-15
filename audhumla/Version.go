package audhumla

type Version struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Release string `json:"release"`

	MainClass string `json:"main_class"`
	Downloads struct {
		Client File `json:"client"`
		Server File `json:"server"`
	} `json:"downloads"`

	Assets    []File    `json:"assets"`
	Libraries []Library `json:"libraries"`

	JVMArgs  []Argument `json:"jvm_args"`
	GameArgs []Argument `json:"game_args"`
}
