
-- AUTO GEN BY CLI	
dofile("./scripts/migrations/09-29-2020T09M53M17/spaPaymentTerm.lua")
dofile("./scripts/migrations/09-29-2020T09M53M17/spaPaymentTermQa.lua")
dofile("./scripts/migrations/09-29-2020T09M53M17/spaPaymentTermProd.lua")
local a = SpaPaymentTerm:new("angular")
a:migrate()
local b = SpaPaymentTermProd:new("angular")
b:migrate()

local c = SpaPaymentTermQa:new("angular")
c:migrate()