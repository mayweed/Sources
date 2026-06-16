ROUNDS = tonumber(io.read())
CASH = tonumber(io.read())

for i = 0, ROUNDS - 1 do
	PLAY = io.read()
	local n1, str, n2 = PLAY:match("(%d+)%s+(PLAIN)%s+(%d+)")

	if not n1 then
		n1, str = PLAY:match("(%d+)%s+(%S+)")
	end
	n1 = tonumber(n1)
	n2 = tonumber(n2)

	--io.stderr:write("n1 " .. tostring(n1) .. " call " .. tostring(str) .. " n2 " .. tostring(n2) .. "\n")

	--on the cash he currently has
	local betAmount = math.ceil(CASH / 4) --Round up
	--    io.stderr:write("Bet : "..betAmount.."\n")

	if str == "ODD" then
		if n1 % 2 ~= 0 then
			CASH = CASH + betAmount -- pas retiré betAmount de cash
		--            io.stderr:write("Cash : "..CASH.."\n")
		else
			CASH = CASH - betAmount --he lost
		end
	elseif str == "EVEN" then
		if n1 ~= 0 and n1 % 2 == 0 then
			CASH = CASH + betAmount
			--io.stderr:write("round ", i, " n1 ", n1, " cash ", CASH, "\n")
		else
			io.stderr:write("round ", i, " n1 ", n1, " cash ", CASH, "\n")
			CASH = CASH - betAmount --he lost
		end
	elseif str == "PLAIN" then
		if n1 == n2 then
			CASH = CASH + (betAmount * 35)
		else
			CASH = CASH - betAmount
			--           io.stderr:write("Round : ",i," cash : ", CASH, " bet : "..betAmount,"\n")
		end
	end
end
print(CASH)
