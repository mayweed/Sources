myId = tonumber(io.read()) -- Your player id (0 or 1)
agentDataCount = tonumber(io.read()) -- Total number of agents in the game

local myAgents = {}
local oppAgents = {}

for i = 0, agentDataCount - 1 do
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
for i = 0, height - 1 do
	next_token = string.gmatch(io.read(), "[^%s]+")
	tiles[i] = {}
	for j = 0, width - 1 do
		-- x: X coordinate, 0 is left edge
		-- y: Y coordinate, 0 is top edge
		x = tonumber(next_token())
		y = tonumber(next_token())
		tileType = tonumber(next_token())
		tiles[i][j] = {
			x = x,
			y = y,
			tileType = tileType,
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

function hasNoCover(x, y)
	local dirs = {
		{ 0, -1 },
		{ 0, 1 },
		{ -1, 0 },
		{ 1, 0 },
	}

	for _, d in ipairs(dirs) do
		local nx = x + d[1]
		local ny = y + d[2]

		if nx >= 0 and nx < width and ny >= 0 and ny < height then
			if tiles[ny][nx].tileType == 0 then
				return true
			end
		end
	end

	return false
end

function hasLowCover(x, y)
	local dirs = {
		{ 0, -1 },
		{ 0, 1 },
		{ -1, 0 },
		{ 1, 0 },
	}

	for _, d in ipairs(dirs) do
		local nx = x + d[1]
		local ny = y + d[2]

		if nx >= 0 and nx < width and ny >= 0 and ny < height then
			if tiles[ny][nx].tileType == 1 then
				return true
			end
		end
	end

	return false
end

function getBestCoverPosition(coverTile, enemies)
	local dirs = {
		{ 0, -1 },
		{ 0, 1 },
		{ -1, 0 },
		{ 1, 0 },
	}

	local best = nil

	for _, d in ipairs(dirs) do
		local nx = coverTile.x + d[1]
		local ny = coverTile.y + d[2]

		if nx >= 0 and nx < width and ny >= 0 and ny < height then
			if tiles[ny][nx].tileType == 0 then
				local pos = { x = nx, y = ny }

				if isGoodSide(pos, coverTile, enemies) then
					return pos -- jackpot direct
				end

				if not best then
					best = pos -- fallback
				end
			end
		end
	end

	return best
end

function isGoodSide(pos, cover, enemies)
	for _, e in ipairs(enemies) do
		-- si la cover est entre toi et l'ennemi
		if
			(pos.x < cover.x and e.x >= cover.x)
			or (pos.x > cover.x and e.x <= cover.x)
			or (pos.y < cover.y and e.y >= cover.y)
			or (pos.y > cover.y and e.y <= cover.y)
		then
			return true
		end
	end
	return false
end

--[[ game loop
TODO 
- Chercher l’agent ennemi le moins bien protégé et tirer.
]]

while true do
	agentCount = tonumber(io.read()) -- Total number of agents still in the game

	local players = {}
	local commands = {} --store the commands needed

	for i = 0, agentCount - 1 do
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
			wetness = wetness,
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
			table.insert(mePlayers, p)
		else
			table.insert(oppPlayers, p)
		end
	end

	table.sort(players, function(a, b)
		return a.agentId < b.agentId
	end)

	for _, p in ipairs(mePlayers) do
		local bestDist = math.huge
		local bestTile = nil
		local bestIndex = nil
		local bestMove = nil

		for i, t in ipairs(highCover) do
			--local pos = getBestCoverPosition(t,oppPlayers)
			--if pos then
			local dist = manhattan(p, t)

			if dist < bestDist then
				bestDist = dist
				bestTile = t
				--bestMove = pos
				bestIndex = i
			end
			--end
		end

		if bestTile then
			--if bestMove then
			table.insert(commands, {
				player = p.agentId,
				action = "MOVE",
				x = bestTile.x,
				y = bestTile.y,
			})
			--end

			table.remove(highCover, bestIndex)
		end
	end
	-- on part de la case où le jour sera…
	for _, p in ipairs(oppPlayers) do
		if hasNoCover(p.x, p.y) then
			pToShoot = p
			--- elseif pToShoot is empty then check the one with lowest cover
		end
		--if not pToShoot then
		if hasLowCover(p.x, p.y) then
			pToShoot = p
			--
			if pToShoot then
				table.insert(commands, {
					player = p.agentId, --non!! c’est l’id de l’agent qui doit tirer!!
					action = "SHOOT",
					x = pToShoot.x,
					y = pToShoot.y,
				})
			end
		end
	end

	for _, c in ipairs(commands) do
		io.stderr:write(c.player .. "," .. c.action .. "," .. c.x .. "," .. c.y .. "\n")
	end

	local out = {}

	for _, p in ipairs(mePlayers) do
		local c = commands[p.agentId]

		if c then
			local parts = {}

			if c.move then
				table.insert(parts, string.format("MOVE %d %d", c.move.x, c.move.y))
			end

			if c.shoot then
				table.insert(parts, string.format("SHOOT %d", c.shoot))
			end

			-- concat des actions avec ;
			print(string.format("%d;%s", p.agentId, table.concat(parts, ";")))
		end
	end
end
