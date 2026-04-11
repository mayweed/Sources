for i=0,4-1 do
    line = io.read()
end

-- func chars qui retourne chars et sa longueur ?
chars = {}
seen = {}

for i=1,#line do
    local c = line:sub(i,i)
    local found = false

    if not seen[c] then
        seen[c] = true
        table.insert(chars,c)
    end
end

-- [[check row and col

--]]

if #chars ~= 4 then print("Invalid")end

io.stderr:write(table.concat(chars, ", ") .. "\n")
