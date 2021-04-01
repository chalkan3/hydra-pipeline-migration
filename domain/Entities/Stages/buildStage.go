package stages

// BuildStage build stage
type BuildStage struct {
	Stage     string   `yaml:"stage"`
	Image     string   `yaml:"image"`
	Tags      []string `yaml:"tags"`
	Artifacts struct {
		Name  string   `yaml:"name"`
		When  string   `yaml:"when"`
		Paths []string `yaml:"paths"`
	} `yaml:"artifacts"`
	Cache struct {
		Policy string `yaml:"policy"`
		Key    string `yaml:"key"`
	} `yaml:"cache"`
	Script []string `yaml:"script"`
	Only   []string `yaml:"only"`
}
