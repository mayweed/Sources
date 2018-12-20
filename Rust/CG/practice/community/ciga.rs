use std::io;
use std::collections::HashMap;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

//yield max num of diff in a list
fn num_diff(cigars:&Vec<i32>) -> Vec<Vec<i32>>{
    let mut idx=0;
    let mut diffs_diffs:Vec<Vec<i32>>=Vec::new();
    for _ in cigars.iter(){
        let mut diffs:Vec<i32>=Vec::new();
        //is there a simpler way??
        let mut start=cigars[idx];
        //should begin from idx
        //drain takes a &mut self!!
        //let mut cigbis=cigars.drain(0..idx);
        for cig in cigars[idx]..cigars[cigars.len()-1]{//cigars.iter(){
            //ternary op in rust?
            let mut diff= if cig > start {cig-start}else {start-cig};
            diffs.push(diff);
            }
        diffs_diffs.push(diffs);
        idx+=1;
        }
    diffs_diffs
}

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let n = parse_input!(input_line, i32);
    let mut cigars:Vec<i32>=Vec::new();
    for i in 0..n as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let lnt = parse_input!(input_line, i32);
        cigars.push(lnt);
    }
    let test=num_diff(&cigars);
    let mut count=HashMap::new();
    for elem in test.iter(){
    //hashmap: k is difference/v is count of diff
        for cnt in elem.iter(){
        //if count.contains_key(cnt){
            let counter_value=count.entry(cnt).or_insert(0);
            *counter_value+=1;
            }
        }
            
    for (k,v) in count{
        eprintln!("{} {}",k,v);
        }
    eprintln!("{:?} {:?}",&cigars,num_diff(&cigars));

    println!("2");
}
