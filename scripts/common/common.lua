function getObjectToYaml(T)
    -- empty string
    local yaml = {""}
    for k, v in pairs (T) do
        if type(v) ~= "function" then
            table.insert(yaml,  (k .. ":"  .. v))
        end
      end

      return table.concat(yaml, "\n")
end