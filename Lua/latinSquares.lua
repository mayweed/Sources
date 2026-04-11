-- Auto-generated code below aims at helping you parse
-- the standard input according to the problem statement.

for i=0,4-1 do
    line = io.read()
end

chars = {} -- pour mettre les 4 chars trouvés
chars = line[1] -- begin by one in lua

[[for i=2, #line do
 if line[2] not in chars
     then add it
     else continue

end ]]
--[[ check which are the 4 distinct char in the grid
if not => invalid
else read 4 char and see which one is missing
add the missing char to the list
print the liste --]]

-- Write an answer using print()
io.stderr:write("%s\n", line)

print("3 4 3 3")
