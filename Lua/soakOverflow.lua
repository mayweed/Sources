-- Win the water fight by controlling the most territory, or out-soak your opponent!

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
for i=0,height-1 do
    next_token = string.gmatch(io.read(), "[^%s]+")
    for j=0,width-1 do
        -- x: X coordinate, 0 is left edge
        -- y: Y coordinate, 0 is top edge
        x = tonumber(next_token())
        y = tonumber(next_token())
        tileType = tonumber(next_token())
    end
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

    myAgentCount = tonumber(io.read()) -- Number of alive agents controlled by you
    for i=0,myAgentCount-1 do
        
        -- Write an action using print()
        -- To debug: io.stderr:write("Debug message\n")
        

        -- One line per agent: <agentId>;<action1;action2;...> actions are "MOVE x y | SHOOT id | THROW x y | HUNKER_DOWN | MESSAGE text"
        if i == 0 then
    print(string.format("MOVE 6 3"))
else
    print(string.format("MOVE 6 1"))
end
    end
end
