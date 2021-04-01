dofile("./scripts/common/common.lua")
local golangFunctions = require("golangFunctions")


-- Meta Class
SpaPromotionManagement = { type="", stages = {}, 
						   variables = { enviroment = "", awsAcessKeyId = "", awsAcessKeyLogic = "", awsS3EnviromentLogic = "", awsDistributionLogic = "", region = "" },
						   projectId = 0,
						   targetBranch = "",
}

-- Class
function SpaPromotionManagement:new (type)
	local obj = {}
	setmetatable(obj, self)
	self.__index = self
	self.type = type or ""
	self.stages = { "build", "sonarqube", "tag", "deployment" }
	self.variables = { 
					   enviroment = "$CI_COMMIT_REF_NAME", 
					   awsAcessKeyId = "(if [ \"$${CI_COMMIT_REF_NAME}\" == \"master\" ]; then echo $${AWS_ACCESS_KEY_ID_MASTER}; else echo $${AWS_ACCESS_KEY_ID_DEV}; fi);",
					   awsAcessKeyLogic = "(if [ \"$${CI_COMMIT_REF_NAME}\" == \"master\" ]; then echo $${AWS_SECRET_ACCESS_KEY_MASTER}; else echo $${AWS_SECRET_ACCESS_KEY_DEV}; fi);",
					   awsS3EnviromentLogic = "(if [ \"$${CI_COMMIT_REF_NAME}\" != \"master\" ]; then echo \"$S3_BUCKET-$CI_COMMIT_REF_NAME\"; else S3_BUCKET; fi);",
					   awsDistributionLogic = "(echo ${MAP_DISTRIBUTION_ID[$CI_COMMIT_REF_NAME]});",
					   region = "$AWS_DEFAULT_REGION"
					 }
	self.projectId = 21473976
	self.targetBranch = "master"
	return obj
end

function SpaPromotionManagement:migrate ()
golangFunctions.migrate(self.type, self.stages, self.variables, self.projectId, self.targetBranch)
end