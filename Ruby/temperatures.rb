n = gets.to_i 
inputs = gets.split(" ")
t = []
for i in 0..(n-1)
  t << inputs[i].to_i
end
if t.empty?
  puts 0
else
  t.sort!{|x,y| y<=>x}
  puts t.sort_by{|t| t.abs}.first
end
