package stages

//AngularInspection stage
type AngularInspection struct {
	Stage     string `yaml:"stage"`
	Image     string `yaml:"image"`
	Artifacts struct {
		Name  string   `yaml:"name"`
		When  string   `yaml:"when"`
		Paths []string `yaml:"paths"`
	} `yaml:"artifacts"`
	Cache struct {
		Key   string   `yaml:"key"`
		Paths []string `yaml:"paths"`
	} `yaml:"cache"`
	Services []struct {
		Name string `yaml:"name"`
	} `yaml:"services"`
	BeforeScript []string `yaml:"before_script"`
	Script       []string `yaml:"script"`
	Tags         []string `yaml:"tags"`
	Only         []string `yaml:"only"`
}
