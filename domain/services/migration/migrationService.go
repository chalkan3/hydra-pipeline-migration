package migrationservice

import (
	"fmt"
	"time"

	luascript "fastshop.com.br/create_pipelines/infra/lua"

	"fastshop.com.br/create_pipelines/infra/date"

	"fastshop.com.br/create_pipelines/infra/file"
)

// MigrationService service migrtion
type MigrationService struct {
	file *file.File
}

/*
	Private Func:
		-> existMigration
		-> getCurrentFolder
*/
func (m *MigrationService) existMigration() bool {
	date := date.NewDate()
	return len(date.ConvertFileInfoTime(m.file.GetAllDirectory("./scripts/migrations/"))) > 0
}

func (m *MigrationService) getCurrentFolder() (string, time.Time) {
	date := date.NewDate()
	timeFolder := date.GetNowDate(date.ConvertFileInfoTime(m.file.GetAllDirectory("./scripts/migrations/")))
	folder := date.FormatFolderMigration(timeFolder)

	return folder, timeFolder
}

/*
	Public Func:
		-> GetCurrentMigration
		-> CreateTemplate
		-> New
*/

// Migrate migrate new files
func (m *MigrationService) Migrate(args []string) error {

	folder, _ := m.getCurrentFolder()
	luaRunner := luascript.NewLGoLua()
	luaRunner.RunLua(folder)
	return nil
}

// GetCurrentMigration get Current folder
func (m *MigrationService) GetCurrentMigration(args []string) error {
	if m.existMigration() {
		folderToCreate, timeFolder := m.getCurrentFolder()
		fmt.Println("Current migration folder [ "+folderToCreate+" ] date -> [", timeFolder, "]")
	} else {
		fmt.Println("doesn't exist migration. please run [fastshop migrations new] ")
	}

	return nil
}

// CreateTemplate create a new migration template
func (m *MigrationService) CreateTemplate(args []string) error {
	if m.existMigration() {
		fileName := args[0]
		types := args[1]
		projectID := args[2]
		targetBranch := args[3]
		nameBranch := args[4]
		commitMessage := args[5]
		folderToCreate, _ := m.getCurrentFolder()

		m.file.CreateFile(folderToCreate, fileName, projectID, targetBranch, nameBranch, commitMessage)
		m.file.WriteToFile(folderToCreate, fileName, types)
	} else {
		fmt.Println("doesn't exist migration. please run [fastshop migrations new] ")
	}

	return nil
}

// New new migration
func (m *MigrationService) New(args []string) error {
	dt := time.Now()
	folderPath := "./scripts/migrations/" + dt.Format("01-02-2006T15M04M05")
	m.file.CreateFolderMigration(folderPath)
	folder, _ := m.getCurrentFolder()
	m.file.CreateMainFile(folder, "main")

	return nil
}

// NewMigrationService IoC
func NewMigrationService(_file *file.File) *MigrationService {
	return &MigrationService{
		file: _file,
	}
}
