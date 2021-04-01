package stages

// SonarqubeStage stage of sonar
type SonarqubeStage struct {
	Stage        string   `yaml:"stage"`
	Dependencies []string `yaml:"dependencies"`
	Image        string   `yaml:"image"`
	Variables    struct {
		SONARHOSTURL string `yaml:"SONAR_HOST_URL"`
		SONARTOKEN   string `yaml:"SONAR_TOKEN"`
		GITDEPTH     int    `yaml:"GIT_DEPTH"`
	} `yaml:"variables"`
	AllowFailure bool     `yaml:"allow_failure"`
	Script       []string `yaml:"script"`
	Tags         []string `yaml:"tags"`
	Only         []string `yaml:"only"`
}

// NewSonarqubeStage Ioc
func NewSonarqubeStage() *SonarqubeStage {
	return &SonarqubeStage{}
}
