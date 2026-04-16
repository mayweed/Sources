N = tonumber(io.read())

-- point?
function distance (x1,x2,y1,y2)
    local dx = x1-x2
    local dy = y1-y2
    return math.sqrt((dx*dx)+(dy*dy))
end

local cities = {}
for i=0,N-1 do
    next_token = string.gmatch(io.read(), "[^%s]+")
        cities[i] = {
            X = tonumber(next_token()),
            Y = tonumber(next_token())
        }
end

local visited = {}

-- Write an answer using print()
-- To debug: io.stderr:write("Debug message\n")

print("distance")
