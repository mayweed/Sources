myId = tonumber(io.read()) -- Your player id (0 or 1)
agentDataCount = tonumber(io.read()) -- Total number of agents in the game

local myAgents = {}
local oppAgents = {}

for i=0,agentDataCount-1 do
    -- agentId: Unique identifier for this agent
    -- player: Player id of this agent
    -- shootCooldown: Number of turns between each of this agent's shots
    -- optimalRange: Maximum manhattan distance for greatest damage output
    -- soakingPower: Damage output within optimal conditions
    -- splashBombs: Number of splash bombs this can throw this game
    next_token = string.gmatch(io.read(), "[^%s]+")
    agentId = tonumber(next_token())
    player = tonumber(next_token())
    shootCooldown = tonumber(next_token())
    optimalRange = tonumber(next_token())
    soakingPower = tonumber(next_token())
    splashBombs = tonumber(next_token())

    local agent = {
        agentId = agentId,
        player = player,
        shootCooldown = shootCooldown,
        optimalRange = optimalRange,
        soakingPower = soakingPower,
        splashBombs = splashBombs,
    }

    if player == myId then
        table.insert(myAgents, agent)
    else
        table.insert(oppAgents, agent)
    end
end

-- width: Width of the game map
-- height: Height of the game map
next_token = string.gmatch(io.read(), "[^%s]+")
width = tonumber(next_token())
height = tonumber(next_token())
local tiles = {}
for i=0,height-1 do
    next_token = string.gmatch(io.read(), "[^%s]+")
    tiles[i] = {}
    for j=0,width-1 do
        -- x: X coordinate, 0 is left edge
        -- y: Y coordinate, 0 is top edge
        x = tonumber(next_token())
        y = tonumber(next_token())
        tileType = tonumber(next_token())
        tiles[i][j] = {
            x =x,
            y=y,
            tileType=tileType,
        }
    end
end

-- sort by tileType
local emptyTiles = {}
local lowCover = {}
local highCover = {}

for _, row in ipairs(tiles) do
    for _, tile in ipairs(row) do
        if tile.tileType == 0 then
            table.insert(emptyTiles, tile)
        end
        if tile.tileType == 1 then
            table.insert(lowCover, tile)
        end
        if tile.tileType == 2 then
            table.insert(highCover, tile)
        end
    end
end

function manhattan(a, b)
    return math.abs(a.x - b.x) + math.abs(a.y - b.y)
end

--[[ game loop
TODO 
- Chercher l’agent ennemi le moins bien protégé et tirer.
]]

while true do
    agentCount = tonumber(io.read()) -- Total number of agents still in the game

    local players = {}
    local commands = {} --store the commands needed

    for i=0,agentCount-1 do
        -- cooldown: Number of turns before this agent can shoot
        -- wetness: Damage (0-100) this agent has taken
        next_token = string.gmatch(io.read(), "[^%s]+")
        agentId = tonumber(next_token())
        x = tonumber(next_token())
        y = tonumber(next_token())
        cooldown = tonumber(next_token())
        splashBombs = tonumber(next_token())
        wetness = tonumber(next_token())

        players[i] = {
        agentId = agentId,
        x = x,
        y = y,
        cooldown = cooldown,
        splashBombs = splashBombs,
        wetness = wetness
        }
    end

    local oppPlayers = {}
    local mePlayers = {}
    local bestWetness = 101
    local agentToShoot = -1

    for _, p in pairs(players) do
        local isEnemy = false
        for _, enemy in ipairs(oppAgents) do
            if p.agentId == enemy.agentId then
                isEnemy = true
                break
            end
        end

        if not isEnemy then 
            table.insert(mePlayers,p) 
        else 
            table.insert(oppPlayers,p)
        end
    end

    table.sort(players, function(a,b)
        return a.agentId < b.agentId end )

    for _,p in ipairs(mePlayers) do
        local bestDist = math.huge
        local bestTile = nil

        for i,t in ipairs(highCover)do          
            local dist = manhattan(p,t)
            
            if  dist < bestDist then
                bestDist = dist
                bestTile = t
                bestIndex = i
            end
        end
                
        if bestTile then
            table.insert(commands, { 
                player = p.agentId,
                action = "MOVE",
                x = bestTile.x,
                y = bestTile.y
                })
        table.remove(highCover,bestIndex)
        end      
    end 

    for _,c in ipairs(commands)do
        io.stderr:write(c.player..","..c.action..","..c.x..","..c.y.."\n")
    end
end
