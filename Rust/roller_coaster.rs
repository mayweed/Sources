use std::io;
use std::collections::HashMap;

macro_rules! print_err {
    ($($arg:tt)*) => (
        {
            use std::io::Write;
            writeln!(&mut ::std::io::stderr(), $($arg)*).ok();
        }
    )
}

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
    
    let mut queue:Vec<i64>=Vec::new();
    let mut remaining_places=&l;
    let mut cache= HashMap::new();
    //let mut cache:Vec<index_cache>=Vec::new();
    
    for i in 0..n as usize{
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let pi = parse_input!(input_line, i64);
        queue.push(pi);
    }
    
    //vec types are indexed on usize
    let mut index:usize=0;
    //can't be i32, does not work (5th test overflow i32!)
    let mut cash_earned:i64=0;
    
    print_err!("Num of places {}, num of times per day {}, num of groups {}",l,c,n);    
    
    for run in 0..c{
        let mut dirhams_ride:i64=0;
        let mut remaining_places:i64=l as i64;
        //what the hell to cache?
        let mut num_groups=0;
        
        //To test!!
        //Is key in cache? If so add the value to the cash and get
        //to the next iteration
        //oki where should I put that??
        if cache.contains_key(&index){//.is_some(){
               cash_earned+=*cache.get(&index).unwrap();
               continue;
            }

        loop{
            if index as i32 >= n{index=0};
            //test 4 come on!!
            if l > n && num_groups==n{break};
            
            
            
            if remaining_places-queue[index] < 0{
                break
            }else{
                remaining_places -= queue[index];
                dirhams_ride+=queue[index];
                num_groups+=1;
                index+=1;
            }
            
        }
        
        cash_earned+=dirhams_ride;
        cache.insert (index,dirhams_ride);
        //cache.push(index_cache{num_groups,index,dirhams_ride});
        //TEST on case 6...
        if num_groups==190{
        print_err!("index {},cash per ride {}, total {} cache: {:?}",index,dirhams_ride,cash_earned,cache);
        }
    }
    println!("{}",cash_earned);
}
