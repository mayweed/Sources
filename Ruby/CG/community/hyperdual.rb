#http://ruby-doc.com/docs/ProgrammingRuby/
class Horse
  def initialize(v, e)
    @velocity = v
    @elegance = e
  end
  attr_reader :velocity, :elegance
  def to_s
    "Horse: #{@velocity} #{@elegance}"
  end
end

def distance(horses)
  dists = []
  horses.each{|x|
    horses.each{|y|
      if y == x
        next
      else
        dists << (y.velocity-x.velocity).abs+(y.elegance-x.elegance).abs
      end
    }
  }
  return dists
end


n = gets.to_i
horses=[]

n.times do
    v, e = gets.split(" ").collect {|x| x.to_i}
    horses << Horse.new(v,e)
end

d = distance(horses).sort

puts d[0]
