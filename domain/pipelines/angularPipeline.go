package pipelines

import (
	stages "fastshop.com.br/create_pipelines/domain/Entities/Stages"
	"fastshop.com.br/create_pipelines/infra/file"
	"fastshop.com.br/create_pipelines/infra/gitlab"
	yamlcustom "fastshop.com.br/create_pipelines/infra/yaml"
	gl "github.com/xanzy/go-gitlab"
)

// AngularPipeline pipeline angular
type AngularPipeline struct {
	Variables struct {
		ENVIRONMENT             string `yaml:"ENVIRONMENT"`
		AWSACCESSKEYIDLOGIC     string `yaml:"AWS_ACCESS_KEY_ID_LOGIC"`
		AWSSECRETACCESSKEYLOGIC string `yaml:"AWS_SECRET_ACCESS_KEY_LOGIC"`
		S3ENVIRONMENTLOGIC      string `yaml:"S3_ENVIRONMENT_LOGIC"`
		DISTRIBUTIONIDLOGIC     string `yaml:"DISTRIBUTION_ID_LOGIC"`
		REGION                  string `yaml:"REGION"`
	} `yaml:"variables"`
	Stages                   []string `yaml:"stages"`
	stages.AngularInspection `yaml:"angular-inspection,omitempty"`
	stages.AngularBuild      `yaml:"angular-build,omitempty"`
	stages.SonarqubeStage    `yaml:"sonarqube,omitempty"`
	stages.SemanticTag       `yaml:"semantic_tag,omitempty"`
	stages.AgularDeployment  `yaml:"deployment,omitempty"`
}

// LoadTemplate load template
func (pipeline *AngularPipeline) LoadTemplate() *AngularPipeline {
	yaml := yamlcustom.NewYamlCustom()
	file := file.NewFile()

	yaml.Unmarshal(file.GetByteContentFromFile("./domain/vendor/defaultTemplate/angularTemplate.yaml"), pipeline)
	return pipeline
}

// PrintYAML print yaml
func (pipeline *AngularPipeline) PrintYAML() {
	yaml := yamlcustom.NewYamlCustom()
	yaml.PrintYAML(yaml.Marshal(&pipeline))

}

// GetYaml print yaml
func (pipeline *AngularPipeline) GetYaml() string {
	yaml := yamlcustom.NewYamlCustom()
	return string(yaml.Marshal(&pipeline))
}

//CreateMergeRequest for this pipeline
func (pipeline *AngularPipeline) CreateMergeRequest(idProject int, targetBranch string, nameBranch string, commitMessage string) {

	gitlabAPI := gitlab.NewGitlab().StartClient(idProject)
	gitlabAPI.CreateBranch(nameBranch, targetBranch)
	branch := gitlabAPI.Commit(".gitlab-ci.yml", gl.FileUpdate, pipeline.GetYaml(), commitMessage, nameBranch)
	gitlabAPI.CreateMergeRequest(branch.Name, targetBranch, branch.Commit.Title)

}

// NewAngularPipeline IoC
func NewAngularPipeline() *AngularPipeline {
	return &AngularPipeline{}
}
