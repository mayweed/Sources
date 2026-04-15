-- [[greedy / scheduling approach
-- Always assign the next vault to the earliest finishing slot.
-- time = (10**n) * (5**(c-n))

R = tonumber(io.read())
V = tonumber(io.read())

local vaults ={}


for i=0,V-1 do
    next_token = string.gmatch(io.read(), "[^%s]+")
    C = tonumber(next_token())
    N = tonumber(next_token())
    local time = 10^N * 5^(C-N)
    table.insert(vaults,time)
end

-- initialiser les voleurs
local finish = {}
for i=1, R do
    finish[i]=0
end

for _,t in ipairs(vaults) do
    local minIndex = 1 -- commence à 1 en lua
    -- qui est dispo le plus rapidement?
    for i = 1, R do
        if finish[i] < finish[minIndex] then
            minIndex = i
        end
    end
    finish[minIndex] = finish[minIndex]+t -- calcul le temps maj
end

local result = finish[1]
for i=2,R do
    if finish[i]> result then
        result = finish[i]
    end
end

print(string.format("%d", result))
