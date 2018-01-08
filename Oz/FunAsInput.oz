%%%%FunnyFunc

declare
fun {FunnyFunc FunL L}
    case L of H|T then {FunL.1 H}|{FunnyFunc FunL.2 T}
    else nil
    end
end

declare
Add=fun{$ X} X+1 end
Sub=fun{$ X} X-1 end

% Argh "You must not write the function FunnyFunc nor the Test procedure, they already exist."
%Not funny!!
declare
proc {Test FunL L SolL}
   SolL={FunnyFunc [Add Sub] [1 2]}
    {Browse {FunnyFunc FunL L} == SolL}
end

%Pass with
{Test [ fun{$ X} X+1 end fun{$ X} X-1 end][1 2][2 1]}

%same length for each list!!
{Browse{FunnyFunc [Add Sub] [1 2]}}

%%%%%%BuildMyFunc%%%%
declare
fun {Build D C}
   %all about X here
    fun{$ D.X} C.X end
end



local F in
    F = {Build [1 2 3] [4 5 6]}
    {Browse {F 1}} % Prints ~1
end

   