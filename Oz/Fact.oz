declare
fun {FactL N}
   local Fact in
      fun {Fact N A F}
	 %2 accumulators: one to match N the other
	 %to match Fact result (multiply A by past results in F!)
	 if A>N then nil
	 else A*F|{Fact N A+1 A*F}
	 end
      end
      {Fact N 1 1}
   end
end

      
{Browse {FactL 4}}

%Lubien's code do like it totally tail-rec
local Fact in
   fun {Fact N}
      local Aux in
         fun {Aux End Curr Acc}
            if Curr > End then Acc %{Reverse Acc}
            else
               {Aux End (Curr + 1) ((Curr * Acc.1)|Acc)}
            end
         end
         {Aux N 2 [1]}
      end
   end
   {Browse {Fact 4}}
end

