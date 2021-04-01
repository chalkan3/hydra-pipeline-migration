package stages

// AngularDeployment is a angular deployment stage
type AgularDeployment struct {
	Stage string `yaml:"stage"`
	Rules []struct {
		If   string `yaml:"if"`
		When string `yaml:"when"`
	} `yaml:"rules"`
	Image        string   `yaml:"image"`
	BeforeScript []string `yaml:"before_script"`
	Script       []string `yaml:"script"`
	Tags         []string `yaml:"tags"`
	Only         []string `yaml:"only,omitempty"`
}
