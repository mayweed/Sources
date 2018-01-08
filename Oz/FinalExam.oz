declare
local
   fun {MakeMakeFoo X}
     fun {$ Y}
       fun {$ Z} X+Y+Z end
     end
   end
 in
   {Browse {MakeMakeFoo 10}}
end

declare
local A B C in
   A=1 B=2 C=3
   local C D E in
     C=B D=5 E=A
     local A C E in
       A=B C=D E=9
       {Browse A+B+C}
     end
     {Browse A+B+C}
   end
   {Browse A+B+C}
end

declare
L=nil|(a|b)
{Browse L}

declare
 fun {Snake F N X}
    if N==0 then {F X} end
    % Here you simply decrement N til 0 and then apply F to X
    % it's a bit obfuscate..
    else {Snake fun {$ X} {Snake F N-1 X} end N-1 X} end
 end
 {Browse {Snake fun {$ X} X+1 end 2 2}}
 {Browse {Snake fun {$ X} X+1 end 10 5}}