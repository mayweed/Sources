N = tonumber(io.read())

local words = {}
for i=0,N-1 do
    s = io.read()
    local w = string.lower(s)
    table.insert(words,w)
end

--table.sort(words)

function checkLetters(w1, w2)
    local freq = {}
    for i=1, #w1 do
        local c = w1:sub(i,i)
        freq[c]= (freq[c] or 0) + 1
    end

    local count = 0
    for j=1,#w2 do
        local c = w2:sub(j,j)
        if freq[c] and freq[c] > 0 then
            freq[c]= freq[c]-1
            count = count +1
        end
    end
    return count
end

function isValid(word,subset) 
    for _,w in ipairs(subset) do
        if checkletters(word,w) >= 3 then
            return false
        end
    end
    return true
end

local subset = {}
for _,word in ipairs(words) do
    local ok = true
    for _,j in ipairs(subset) do
        if checkLetters(word,j) >= 3 then
            ok = false
            break
        end
    end
    if ok then
        table.insert(subset,word)
    end
end

print(string.format("%d",#subset))

-- Write an answer using print()
--local test = checkLetters("Xanax", "Viagra")
--io.stderr:write(string.format (test))
