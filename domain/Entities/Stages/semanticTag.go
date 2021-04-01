package stages

// SemanticTag semantic stage
type SemanticTag struct {
	Stage        string   `yaml:"stage"`
	Dependencies []string `yaml:"dependencies"`
	Tags         []string `yaml:"tags"`
	Image        string   `yaml:"image"`
	Script       []string `yaml:"script"`
	Only         []string `yaml:"only"`
}
