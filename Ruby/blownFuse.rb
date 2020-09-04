#http://ruby-doc.com/docs/ProgrammingRuby/
n, m, c = gets.split(" ").collect {|x| x.to_i}
inputs = gets.split(" ")
powerConsumption = {}
for i in 0..(n-1)
    nx = inputs[i].to_i
	powerConsumption[i]= nx
end

n = 0
max = 0
alreadySeen = {}
inputs = gets.split(" ")
for i in 0..(m-1)
    mx = inputs[i].to_i
		#appliance was off
	if !alreadySeen[mx] 
	    n+=powerConsumption[mx-1]
        if n > max
		  max=n
	    end
	    alreadySeen[mx] = true
	else
		n-=powerConsumption[mx-1]
        alreadySeen[mx]=false
	end
end

STDERR.puts max 

# Write an answer using puts
#STDERR.puts powerConsumption,alreadySeen

puts "Fuse was not blown."
puts "Maximal consumed current was XX A."
