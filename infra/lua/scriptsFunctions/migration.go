package scriptsfunctions

import (
	"fastshop.com.br/create_pipelines/domain/pipelines"
	"fastshop.com.br/create_pipelines/infra/file"
	yamlcustom "fastshop.com.br/create_pipelines/infra/yaml"
	lua "github.com/yuin/gopher-lua"
)

// MigrationFunction Lua function
type MigrationFunction struct {
	yaml *yamlcustom.YamlCustom
	file *file.File
}

// Migrate is a migrationFunction
func (f *MigrationFunction) Migrate(L *lua.LState) int {
	types := L.ToString(1)
	idProject := L.ToInt(2)
	targetBranch := L.ToString(3)
	nameBranch := L.ToString(4)
	commitMessage := L.ToString(5)
	stages := L.ToTable(6)
	variables := L.ToTable(7)

	switch types {
	case "dotnet":
		f.migrateDotnetFunction(stages)
	case "angular":
		angularinspection := L.ToTable(8)
		angularBuild := L.ToTable(9)
		f.migrateAngularFunction(stages, variables, idProject, targetBranch, nameBranch, commitMessage, angularinspection, angularBuild)
	}

	return 1
}

// MigrateDotnetFunction Function migration dotnet
func (f *MigrationFunction) migrateDotnetFunction(stages *lua.LTable) {
	var teste []string
	stages.ForEach(func(key lua.LValue, value lua.LValue) {
		teste = append(teste, value.String())
	})

	newPipeline := pipelines.NewDotnetPipeline()
	newPipeline.Stages = teste
	f.yaml.Unmarshal(f.file.GetByteContentFromFile("./domain/vendor/defaultTemplate/dotnetTemplate.yaml"), newPipeline)
	f.yaml.PrintYAML(f.yaml.Marshal(&newPipeline))
}

func (f *MigrationFunction) migrateAngularFunction(stages *lua.LTable, variables *lua.LTable,
	idProject int, targetBranch string, nameBranch string, commitMessage string, angularInspection *lua.LTable, angularBuild *lua.LTable) {
	angularPipeline := pipelines.NewAngularPipeline().LoadTemplate()

	// Angular inspections
	angularInspection.ForEach(func(key lua.LValue, value lua.LValue) {
		switch key.String() {
		case "nodeVersion":
			angularPipeline.AngularInspection.Image = "node:" + value.String()
		case "beforeScripts":
			var beforeScripts []string
			value.(*lua.LTable).ForEach(func(key lua.LValue, value lua.LValue) {
				beforeScripts = append(beforeScripts, value.String())
			})
			angularPipeline.AngularInspection.BeforeScript = beforeScripts
		case "script":
			var scripts []string
			value.(*lua.LTable).ForEach(func(key lua.LValue, value lua.LValue) {
				scripts = append(scripts, value.String())
			})
			angularPipeline.AngularInspection.Script = scripts

		}
	})

	angularBuild.ForEach(func(key lua.LValue, value lua.LValue) {
		switch key.String() {
		case "nodeVersion":
			angularPipeline.AngularBuild.Image = "node:" + value.String()
		case "beforeScripts":
			var beforeScripts []string
			value.(*lua.LTable).ForEach(func(key lua.LValue, value lua.LValue) {
				beforeScripts = append(beforeScripts, value.String())
			})
			angularPipeline.AngularBuild.BeforeScript = beforeScripts
		case "script":
			var scripts []string
			value.(*lua.LTable).ForEach(func(key lua.LValue, value lua.LValue) {
				scripts = append(scripts, value.String())

			})
			angularPipeline.AngularBuild.Script = scripts

		}
	})
	// angularPipeline.PrintYAML()
	angularPipeline.CreateMergeRequest(idProject, targetBranch, nameBranch, commitMessage)

}

// NewMigrationFunction IoC
func NewMigrationFunction() *MigrationFunction {
	return &MigrationFunction{
		yaml: yamlcustom.NewYamlCustom(),
		file: file.NewFile(),
	}
}
