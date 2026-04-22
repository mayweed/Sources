N = tonumber(io.read())

-- point?
function distance (x1,x2,y1,y2)
    local dx = x1-x2
    local dy = y1-y2
    return math.sqrt((dx*dx)+(dy*dy))
end

local cities = {}
for i=1, N do
    next_token = string.gmatch(io.read(), "[^%s]+")
        cities[i] = {
            x = tonumber(next_token()),
            y = tonumber(next_token())
        }
end

local startPoint = cities[1]
local endPoint = cities[1]

local totalDist = 0
local distToOrigin = 0
table.remove(cities,1)

while (#cities ~= 0) do
    --on part de la 1ere ville et on calcule la plus proche
    local minDist = math.huge
    local minIndex = 0

    for j = 1, #cities do
        local d = distance(startPoint.x, cities[j].x,startPoint.y, cities[j].y)
        if d < minDist then
            minDist=d
            minIndex = j
        end
    end
    totalDist = totalDist+ minDist
    startPoint = cities[minIndex]
    if (#cities == 1) then
        distToOrigin = distance(cities[minIndex].x, endPoint.x,cities[minIndex].y,endPoint.y)
    end
    table.remove(cities,minIndex) 
end

function round(x)
    return math.floor(x + 0.5)
end

print(round(totalDist+distToOrigin))
