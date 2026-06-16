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

	--on the cash he currently has
	local betAmount = math.ceil(CASH / 4) --Round up

	if str == "ODD" then
		if n1 % 2 ~= 0 then
			CASH = CASH + betAmount -- pas retiré betAmount de cash
		else
			CASH = CASH - betAmount --he lost
		end
	elseif str == "EVEN" then
		if n1 ~= 0 and n1 % 2 == 0 then
			CASH = CASH + betAmount
		else
			io.stderr:write("round ", i, " n1 ", n1, " cash ", CASH, "\n")
			CASH = CASH - betAmount --he lost
		end
	elseif str == "PLAIN" then
		if n1 == n2 then
			CASH = CASH + (betAmount * 35)
		else
			CASH = CASH - betAmount
		end
	end
end
print(CASH)
