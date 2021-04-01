dofile("./scripts/common/common.lua")
-- local yaml = require("yaml")
T = Teste:new(nil, "abc", "dfg")

-- criar funcao de conersaõ para key
local testeYaml = getObjectToYaml(T)


print(testeYaml)


-- local tb = yaml.parse(T.bla)
