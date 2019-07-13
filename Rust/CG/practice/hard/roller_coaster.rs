use std::io;

//i think i got the algo: first you take the groups and establish a gain maps of group index ->
//gain
//THEN you apply it numOfTime times
//this way gains are pre-calculated...
//Key is to store two maps / arrays (one of earns per group and one for index of group which goes next) and then just jump here and there and compute total.
macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let places = parse_input!(inputs[0], i32);
    let numOfTime = parse_input!(inputs[1], i32);
    let nbGroups = parse_input!(inputs[2], i32);
    
    let mut groups:Vec<i64>=Vec::new();
    let mut remaining_places=&places;
    
    //fill the seats on demand here
    let mut dirhams_ride:i64=0;
    let mut remaining_places:i64=places as i64;

    for i in 0..nbGroups as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let pi = parse_input!(input_line, i64);
        groups.push(pi);
    }
    
    //vec types are indexed on usize
    let mut index:usize=0;
    //can't be i32, does not work (5th test overflow i32!)
    let mut cash_earned:Vec<i64>;
    
    eprintln!("Num of places {}, num of times per day {}, groups {:?}",places,numOfTime,groups);    
    
    for run in 0..numOfTime{
        let mut num_groups=0;
        let mut num_ppl_groups=0;
        
        loop{
            if index as i32 >= nbGroups{index=0};
            //test 4 come on!!
            if remaining_places > nbGroups.into() && num_groups==nbGroups{break};
            if remaining_places-groups[index] < 0{
                break
            }else{
                remaining_places -= groups[index];
                dirhams_ride+=groups[index];
                index+=1;
                num_groups+=1;
                num_ppl_groups=remaining_places;
                //eprintln!("numGroups {}, pplGroups {}, cash per ride {}, total {}",num_groups,num_ppl_groups,dirhams_ride,cash_earned);
                cash_earned[index]=dirhams_ride;
               
            }
            
        }
        cash_earned+=dirhams_ride;
    }

    println!("{}",cash_earned);
}
