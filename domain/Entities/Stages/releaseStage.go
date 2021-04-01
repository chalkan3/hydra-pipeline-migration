package stages

// ReleaseStage a realase stage
type ReleaseStage struct {
	Stage        string   `yaml:"stage"`
	BeforeScript []string `yaml:"before_script"`
	Script       []string `yaml:"script"`
	Tags         []string `yaml:"tags"`
	Only         []string `yaml:"only"`
}
