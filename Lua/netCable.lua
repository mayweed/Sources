local houses = {}

N = tonumber(io.read())
for i=0,N-1 do
    next_token = string.gmatch(io.read(), "[^%s]+")
    X = tonumber(next_token())
    Y = tonumber(next_token())
    table.insert(houses, {x=X,y=Y})
end

-- Sort by x, then by y
table.sort(houses, function(a, b)
    if a.x == b.x then
        return a.y < b.y
    else
        return a.x < b.x
    end
end)

function distManhattan (x1,x2,y1,y2)
    return math.abs(x1-x2)+math.abs(y1-y2)
end

-- Write an answer using print()
-- To debug: io.stderr:write("Debug message\n")
local test = distManhattan(houses[1].x, houses[N].x,houses[1].y, houses[N].y)
print(test)
