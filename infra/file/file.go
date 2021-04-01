package file

import (
	"fmt"
	"io/ioutil"
	"os"

	"fastshop.com.br/create_pipelines/domain/strings"
)

// File is a file manipulator
type File struct {
}

// GetByteContentFromFile get file from ./domain/vendor/pipeline.yaml
func (file *File) GetByteContentFromFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return content
}

// GetAllDirectory return all files names
func (file *File) GetAllDirectory(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
	}

	return files
}

// WriteToFile to main lua
func (file *File) WriteToFile(folder string, fileName string, types string) {
	f, err := os.OpenFile("./scripts/migrations/"+folder+"/main.lua", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	stringTemplate := strings.NewStringTemplate()
	if err != nil {
		fmt.Println(err)
	}

	if _, errWrite := f.Write([]byte(stringTemplate.WriteOnMainLua(folder, fileName, types))); err != nil {
		fmt.Println(errWrite)
	}

	if errClose := f.Close(); err != nil {
		fmt.Println(errClose)
	}
}

// CreateFile is a create file
func (file *File) CreateFile(folder string, fileName string, projectID string, targetBranch string, nameBranch string, commitMessage string) {
	stringTemplate := strings.NewStringTemplate()
	err := ioutil.WriteFile("./scripts/migrations/"+folder+"/"+fileName+".lua", []byte(stringTemplate.MakeLuaTemplate(fileName, projectID, targetBranch, nameBranch, commitMessage)), 0755)

	if err != nil {
		fmt.Println(err)
	}

}

// CreateMainFile Create EntryPoint Lua
func (file *File) CreateMainFile(folder string, fileName string) {
	err := ioutil.WriteFile("./scripts/migrations/"+folder+"/"+fileName+".lua", []byte(""), 0755)

	if err != nil {
		fmt.Println(err)
	}

}

// CreateFolderMigration is a createfolder
func (file *File) CreateFolderMigration(path string) *File {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	return file
}

// NewFile Ioc
func NewFile() *File {
	return &File{}
}
