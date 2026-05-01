n = tonumber(io.read()) -- the number of relationships of influence

local graph = {}
local result = 0

for i=0,n-1 do
    -- x: a relationship of influence between two people (x influences y)
    next_token = string.gmatch(io.read(), "[^%s]+")
    x = tonumber(next_token())
    y = tonumber(next_token())
    graph[x]= graph[x] or {}
    table.insert(graph[x],y)
end

function DFS (graph, startNode)
    local maxLen = 1
    if graph[startNode] then
        for _,nodes in pairs(graph[startNode]) do   
            -- pour chaque noeud X : longest(x) = 1 + max(longest(y)) pour tous les voisins y
            local len = 1+DFS(graph,nodes)
            if len > maxLen then
                maxLen=len
            end      
        end
    end
    return maxLen
end

for node,_ in pairs(graph) do
    len = DFS(graph,node)
    if len > result then
        result=len
    end
end

-- The number of people involved in the longest succession of influences
print(result)
