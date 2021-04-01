package services

import (
	"fastshop.com.br/create_pipelines/domain/pipelines"
	"fastshop.com.br/create_pipelines/infra/file"
	yamlcustom "fastshop.com.br/create_pipelines/infra/yaml"
)

// CreatePipelineService my service to create a pipeline
type CreatePipelineService struct {
	yaml *yamlcustom.YamlCustom
	file *file.File
}

// CreateDotnetPipeline just for clone pipeline
func (service *CreatePipelineService) CreateDotnetPipeline(args []string) error {
	newPipeline := pipelines.NewDotnetPipeline()
	service.yaml.Unmarshal(service.file.GetByteContentFromFile("./domain/vendor/dotnetTemplate.yaml"), newPipeline)
	service.yaml.PrintYAML(service.yaml.Marshal(&newPipeline))
	return nil
}

// CreateNodePipeline is a creating node pipele
func (service *CreatePipelineService) CreateNodePipeline(args []string) error {

	newPipeline := pipelines.NewDotnetPipeline()
	service.yaml.Unmarshal(service.file.GetByteContentFromFile("./domain/vendor/dotnetTemplate.yaml"), newPipeline)
	service.yaml.PrintYAML(service.yaml.Marshal(&newPipeline))

	return nil

}

// NewCreatePipelineService IOC
func NewCreatePipelineService(_yaml *yamlcustom.YamlCustom, _file *file.File) *CreatePipelineService {
	return &CreatePipelineService{
		yaml: _yaml,
		file: _file,
	}
}
