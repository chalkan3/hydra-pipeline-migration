package menus

import (
	services "fastshop.com.br/create_pipelines/domain/services/pipelines"
	helpersmenu "fastshop.com.br/create_pipelines/infra/provider/helpers/menu"
)

// PipelinesMenu is a menu command pipeline [separar o commum em uma base]
type PipelinesMenu struct {
	createPipelineService *services.CreatePipelineService
	create                *helpersmenu.Menu
	mainManu              *helpersmenu.Menu
	mainCommand           string
}

// RunMenu is startupmenu
func (m *PipelinesMenu) RunMenu(args []string) {
	m.create.Call(args)
}

// MountMenu integrate all menus
func (m *PipelinesMenu) MountMenu() *PipelinesMenu {
	m.mainManu.CreateItem(m.mainCommand, nil).CreateSubMenu(m.create)
	return m
}

// Create create menu
func (m *PipelinesMenu) Create() *PipelinesMenu {
	m.create.CreateItem("dotnet", m.createPipelineService.CreateDotnetPipeline).CreateItem("node", m.createPipelineService.CreateNodePipeline)
	return m
}

// NewPipelinesMenu Ioc
func NewPipelinesMenu(_create *helpersmenu.Menu, _createPipelineService *services.CreatePipelineService, _main *helpersmenu.Menu) *PipelinesMenu {

	return &PipelinesMenu{
		createPipelineService: _createPipelineService,
		create:                _create,
		mainManu:              _main,
		mainCommand:           "create",
	}
}
