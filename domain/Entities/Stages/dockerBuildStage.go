package stages

// DockerBuildStage docker build stages
type DockerBuildStage struct {
	Stage    string `yaml:"stage"`
	Services []struct {
		Name string `yaml:"name"`
	} `yaml:"services"`
	Dependencies []string `yaml:"dependencies"`
	Cache        struct {
		Policy string `yaml:"policy"`
		Key    string `yaml:"key"`
	} `yaml:"cache"`
	BeforeScript []string `yaml:"before_script"`
	Script       []string `yaml:"script"`
	Tags         []string `yaml:"tags"`
	Only         []string `yaml:"only"`
}
