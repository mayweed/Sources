# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

n = gets.to_i # the number of temperatures to analyse
inputs = gets.split(" ")
t = []
for i in 0..(n-1)
    # t: a temperature expressed as an integer ranging from -273 to 5526
  t << inputs[i].to_i
end

# Write an answer using puts
# To debug: STDERR.puts "Debug messages..."
if t.empty?
  puts 0
else
  puts t.sort_by{|t| t.abs}
end
