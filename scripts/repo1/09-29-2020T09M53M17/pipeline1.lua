dofile("./scripts/common/common.lua")
local golangFunctions = require("golangFunctions")


-- Meta Class
SpaPaymentTerm = { type="", stages = {}, variables = { enviroment = "", awsAcessKeyId = "", awsAcessKeyLogic = ""} , projectId = 0, targetBranch = "", 
		nameBranch= "", commitMenssage="", angularInspection={ nodeVersion = 0, beforeScripts = {}, script = {} }, 
		angularBuild = { nodeVersion = 0, beforeScripts = {}, script = {} } }

-- Class
function SpaPaymentTerm:new (type)
	local obj = {}
	setmetatable(obj, self)
	self.__index = self
	self.type = type or ""
	self.stages = { "build", "sonarqube", "docker-build", "tag", "release", "deployment"}
	self.projectId = 10
	self.targetBranch = "develop"
	self.nameBranch = "feature/pipeline-devops"
	self.commitMessage = "fix(pipeline): colocando develop"
	self.angularInspection = {
		nodeVersion = 10, beforeScripts = { "npm install --silent", "npm i -g @angular-devkit/build-angular@0.800.6", "npm i -g @angular/cli@8.0.6", "ng version"}, 
		script = { "ng build --configuration=local --aot --build-optimizer" } 
	}
	self.angularBuild = { 
		nodeVersion = 10, beforeScripts = { "npm install --silent", "npm i -g @angular-devkit/build-angular@0.800.6", "npm i -g @angular/cli@8.0.6", "ng version "}, 
		script = {  "echo $CI_COMMIT_REF_NAME" ,"echo $ENVIRONMENT" ,"if [[ $ENVIRONMENT =~ ^qa.* ]]; then ENVIRONMENT=qa; fi;" ,"ng build --configuration=$ENVIRONMENT --aot --build-optimizer"} 
	}


	return obj
end

function SpaPaymentTerm:migrate ()
	golangFunctions.migrate(self.type, self.projectId, self.targetBranch, self.nameBranch, self.commitMessage, self.stages, self.variables, self.angularInspection, self.angularBuild)
end