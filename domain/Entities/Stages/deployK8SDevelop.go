package stages

// DeployK8SDevelopStage k8s stage
type DeployK8SDevelopStage struct {
	Image struct {
		Name       string   `yaml:"name"`
		Entrypoint []string `yaml:"entrypoint"`
	} `yaml:"image"`
	Stage     string   `yaml:"stage"`
	Tags      []string `yaml:"tags"`
	Variables struct {
		K8SNAME string `yaml:"K8S_NAME"`
		TAG     string `yaml:"TAG"`
		AKAMAI  string `yaml:"AKAMAI"`
	} `yaml:"variables"`
	Script []string `yaml:"script"`
	Only   []string `yaml:"only"`
}
