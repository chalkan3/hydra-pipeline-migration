package menus

import (
	migrationservice "fastshop.com.br/create_pipelines/domain/services/migration"
	helpersmenu "fastshop.com.br/create_pipelines/infra/provider/helpers/menu"
)

// MigrationMenu is a menu command pipeline [separar o commum em uma base]
type MigrationMenu struct {
	migrationService *migrationservice.MigrationService
	new              *helpersmenu.Menu
	mainManu         *helpersmenu.Menu
	mainCommand      string
}

// RunMenu is startupmenu
func (m *MigrationMenu) RunMenu(args []string) {
	m.new.Call(args)
}

// MountMenu integrate all menus
func (m *MigrationMenu) MountMenu() *MigrationMenu {
	m.mainManu.CreateItem(m.mainCommand, nil).CreateSubMenu(m.new)
	return m
}

// Create create menu
func (m *MigrationMenu) Create() *MigrationMenu {
	m.new.CreateItem("new",
		m.migrationService.New).CreateItem("create",
		m.migrationService.CreateTemplate).CreateItem("current",
		m.migrationService.GetCurrentMigration).CreateItem("migrate", m.migrationService.Migrate)

	return m
}

// NewMigrationMenu Ioc
func NewMigrationMenu(_new *helpersmenu.Menu, _migrationService *migrationservice.MigrationService,
	_main *helpersmenu.Menu) *MigrationMenu {

	return &MigrationMenu{
		migrationService: _migrationService,
		new:              _new,
		mainManu:         _main,
		mainCommand:      "migration",
	}
}
