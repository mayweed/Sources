-- from prog in lua
--
N = 8 -- board size

-- check whether position(n,c) is free from attacks
function isPlaceOK(a,n,c)
    for i=1, n-1, do
        if(a[i] == c) or (a[i]-i == c-n) or (a[i]+i == c+n)
            return false
        end
    end
    return true
end

-- print a board
function printSolution(a)
    for i = 1,N do
        for j = 1,N do
            io.write(a[i] == j and "X" or ".","")
        end
        io.write("\n")
    end
    io.write("\n")
end

--add to board 'a' all queens from 'n' to N
function addQueen(a,n)
    if n>N then
        printSolution(a)
    else
        for c=1, N do
            if isPlaceOK(a,n,c) then
                a[n] = c
                addQueen(a,n+1)
            end
        end
    end
end

--run the program
addQueen({},1)
