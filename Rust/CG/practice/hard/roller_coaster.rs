use std::io;
use std::collections::HashMap;
use std::time::{Duration,Instant};

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let l = parse_input!(inputs[0], i32);
    let c = parse_input!(inputs[1], i32);
    let n = parse_input!(inputs[2], i32);
    eprintln!("l {}, c {}, n {}",l,c,n);
    
    let mut queue:Vec<i64>=Vec::new();
    let mut remaining_places=&l;
    
    //the first i64 == num of ppl in group, the second == dirhams per ride
    let mut cache:HashMap<i64, i64> = HashMap::new();
    
    for i in 0..n as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let pi = parse_input!(input_line, i64);
        queue.push(pi);
    }
    //vec types are indexed on usize
    let mut index:usize=0;
    //can't be i32, does not work (5th test overflow i32!)
    let mut cash_earned:i64=0;
    
    eprintln!("Num of places {}, num of times per day {}, num of groups {}",l,c,n);    
    
    for run in 0..c{
        let mut dirhams_ride:i64=0;
        let mut remaining_places:i64=l as i64;
        let mut num_groups=0;
        let mut num_ppl_groups=0;
        
        loop{
            if index as i32 >= n{index=0};
            //test 4 come on!!
            if l > n && num_groups==n{break};
            if remaining_places-queue[index] < 0{
                break
            }else{
                remaining_places -= queue[index];
                //tu regardes si on a déjà cette valeur et tu l'ajoutes direct
                // match cache.get(remaining_places){
                //Some(remaining_places) => cash_earned + dirhams ride
                //else valeur classique
                dirhams_ride+=queue[index];
                index+=1;
                num_groups+=1;
                num_ppl_groups=remaining_places;
                eprintln!("numGroups {}, pplGroups {}, cash per ride {}, total {}",num_groups,num_ppl_groups,dirhams_ride,cash_earned);
                cache.insert(remaining_places,dirhams_ride);
               
            }
            
        }
        let now= Instant::now();
        cash_earned+=dirhams_ride;
        let new_now= Instant::now();
        //now.elapsed().as_millis() not yet avail on CG
        eprintln!("{:?}",new_now.duration_since(now));
    }
    eprintln!("{:?}",cache);

    println!("{}",cash_earned);
}
