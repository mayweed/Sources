%Intuitive fibo
%do not forget capital letter to name proc/func/var
%else "applying non procedure non object error!!"
declare
fun {FibNaive N}
   if N=<1 then N
   else {FibNaive N-1}+{FibNaive N-2}
   end
end

{Browse {FibNaive 8}}

% Accumulators fibo
declare
fun {Fib N}
    local FibAux in
        fun {FibAux N Acc1 Acc2}
	   if N==0 then Acc1
	   elseif N==1 then Acc2
	   else
	      {FibAux N-1 Acc2 Acc1+Acc2}
	   end
	end
	{FibAux N 0 1}
    end
end

{Browse{Fib 42}}