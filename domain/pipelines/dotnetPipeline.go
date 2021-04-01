package pipelines

import (
	stages "fastshop.com.br/create_pipelines/domain/Entities/Stages"
	linq "github.com/ahmetb/go-linq/v3"
)

// DotnetPipeline is a pipeline working in 2020-09-24
type DotnetPipeline struct {
	Variables struct {
		ENVIRONMENT             string `yaml:"ENVIRONMENT"`
		REGION                  string `yaml:"REGION"`
		AWSACCESSKEYIDLOGIC     string `yaml:"AWS_ACCESS_KEY_ID_LOGIC"`
		AWSSECRETACCESSKEYLOGIC string `yaml:"AWS_SECRET_ACCESS_KEY_LOGIC"`
		AWSREGISTRYIMAGELOGIC   string `yaml:"AWS_REGISTRY_IMAGE_LOGIC"`
		GITLABREGISTRYTAG       string `yaml:"GITLAB_REGISTRY_TAG"`
		K8SFILELOGIC            string `yaml:"K8S_FILE_LOGIC"`
	} `yaml:"variables"`
	Stages                       []string `yaml:"stages"`
	stages.BuildTestStage        `yaml:"build-test,omitempty"`
	stages.SonarqubeStage        `yaml:"sonarqube,omitempty"`
	stages.BuildStage            `yaml:"build,omitempty"`
	stages.DockerBuildStage      `yaml:"docker-build,omitempty"`
	stages.SemanticTag           `yaml:"semantic_tag,omitempty"`
	stages.ReleaseStage          `yaml:"release,omitempty"`
	stages.DeployK8SDevelopStage `yaml:"deploy-k8s-develop,omitempty"`
}

// AddStage add a new stage to a pipeline
func (pipe *DotnetPipeline) AddStage(stage string) {
	pipe.Stages = append(pipe.Stages, stage)
}

// DisableSonarQube Disable sonar stage
func (pipe *DotnetPipeline) DisableSonarQube() {
	linq.From(pipe.Stages).Except(linq.From([]string{"sonarqube"})).ToSlice(&pipe.Stages)
	pipe.SonarqubeStage = stages.SonarqubeStage{}
}

// DisableK8s disable kubernetes
func (pipe *DotnetPipeline) DisableK8s() {
	pipe.DeployK8SDevelopStage = stages.DeployK8SDevelopStage{}
}

// NewDotnetPipeline Ioc
func NewDotnetPipeline() *DotnetPipeline {
	return &DotnetPipeline{}
}
