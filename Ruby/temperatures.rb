n = gets.to_i 
p gets.split.map(&:to_i).sort!{|x,y|y<=>x}.sort_by{|t|t.abs}[0]||0
