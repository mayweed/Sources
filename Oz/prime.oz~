declare
X={Float.sqrt {Int.toFloat 37}}

{Browse X}

{Browse {Float.sqrt 37.0}}


declare
fun {Prime N}
   local PrimeAux in
      fun {PrimeAux N M}
	 %stop condition
	 %better use Float.sqrt(N)
	 if M==1 then true
	 elseif N mod M == 0 then false
	 else
	    %this works for very small num
	    {PrimeAux N M-1}
	 end
      end
      if N==1 then false
      else {PrimeAux N N-1}
      end
   end
end

{Browse {Prime 1}}