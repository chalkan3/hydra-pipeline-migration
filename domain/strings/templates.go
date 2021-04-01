package strings

import (
	"strings"
)

// Stringstemplate is a strings templates
type Stringstemplate struct {
}

// MakeLuaTemplate return lua template
func (s *Stringstemplate) MakeLuaTemplate(className string, projectID string, targetBranch string, nameBranch string, commitMessage string) string {
	return `dofile("./scripts/common/common.lua")
local golangFunctions = require("golangFunctions")


-- Meta Class
` + strings.Title(className) + ` = { type="", stages = {}, variables = { enviroment = "", awsAcessKeyId = "", awsAcessKeyLogic = ""} , projectId = 0, targetBranch = "", 
		nameBranch= "", commitMenssage="", angularInspection={ nodeVersion = 0, beforeScripts = {}, script = {} }, 
		angularBuild = { nodeVersion = 0, beforeScripts = {}, script = {} }, 
			
}

-- Class
function ` + strings.Title(className) + ":new (type)" + `
	local obj = {}
	setmetatable(obj, self)
	self.__index = self
	self.type = type or ""
	self.stages = { "build", "sonarqube", "docker-build", "tag", "release", "deployment"}
	self.projectId = ` + projectID + `
	self.targetBranch = "` + targetBranch + `"
	self.nameBranch = "` + nameBranch + `"
	self.commitMessage = "` + commitMessage + `"
	return obj
end

function ` + strings.Title(className) + ":migrate ()" + `
	golangFunctions.migrate(self.type, self.projectId, self.targetBranch, self.nameBranch, self.commitMessage, self.stages, self.variables)
end`
}

// WriteOnMainLua add new migrations
func (s *Stringstemplate) WriteOnMainLua(folder string, className string, types string) string {
	return `
-- AUTO GEN BY CLI	
dofile("./scripts/migrations/` + folder + `/` + className + `.lua")` + `
local ` + strings.ToLower(className) + ` = ` + strings.Title(className) + `:new("` + strings.ToLower(types) + `")
` + strings.ToLower(className) + `:migrate()`

}

// NewStringTemplate IoC
func NewStringTemplate() *Stringstemplate {
	return &Stringstemplate{}
}
