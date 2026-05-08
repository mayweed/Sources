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
        if tileType ~= 0 then
        --io.stderr:write(string.format("tt is %d\n",tileType))
        end
        tiles[i][j] = {
            x =x,
            y=y,
            tileType=tileType,
        }
    end
end

function manhattan(x1, y1, x2, y2)
    return math.abs(x1 - x2) + math.abs(y1 - y2)
end

-- game loop
while true do
    agentCount = tonumber(io.read()) -- Total number of agents still in the game

    local players = {}

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

    for _,p in ipairs(oppPlayers)do
        if p.wetness < bestWetness then
            bestWetness = p.wetness
            agentToShoot = p.agentId
        end
    end

    --Si le wetness (trempage) d’un agent atteint 100 ou plus, il est retiré de la partie.
    local secondBest
    for _,v in ipairs(oppPlayers) do
        if v.wetness < 100 then
            secondBest = v.agentId
            break
        end
    end

    local emptyTiles = {}

    for _, row in ipairs(tiles) do
        for _, tile in ipairs(row) do
            if tile.tileType == 0 then
                table.insert(emptyTiles, tile)
                --io.stderr:write(string.format("Empty tile at x=%d y=%d\n", tile.x, tile.y))
            end
        end
    end

    myAgentCount = tonumber(io.read()) -- Number of alive agents controlled by you

    for i=1, myAgentCount do 
        -- Write an action using print()
        -- To debug: io.stderr:write("Debug message\n")
        -- One line per agent: <agentId>;<action1;action2;...> actions are "MOVE x y | SHOOT id | THROW x y | HUNKER_DOWN | MESSAGE text"
        print(string.format("%d;SHOOT %d",i,secondBest)) --agentToShoot))
    end
end
