use std::io;
use std::collections::VecDeque;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}
#[derive(Debug,Clone)]
struct Robot{
    target:String,
    eta:i32,
    storage_a:i32,
    storage_b:i32,
    storage_c:i32,
    storage_d:i32,
    storage_e:i32,
    expertise_a:i32,
    expertise_b:i32,
    expertise_c:i32,
    expertise_d:i32,
    expertise_e:i32,
    }
//should I write a module trait to intervene in each??
#[derive(Debug,Clone)]
struct Sample{
    sample_id:i32,
    carried_by:i32,
    rank:i32,
    expertise_gain:String,
    health:i32,
    cost_a:i32,
    cost_b:i32,
    cost_c:i32,
    cost_d:i32,
    cost_e:i32,
    }
    
//Implement bot commands
//trait or enum for modules??
//https://rustbyexample.com/custom_types/enum.html
pub enum Command {
    //where string is either DIAGNOSIS etc...enum??
    //should be a string or whatever modules will be counted for??
    Goto(i32),
    Connect(i32),
}
impl Command {
    pub fn encode(&self) -> String {
        match *self {
            Command::Goto(s) => format!("GOTO {}", s),
            Command::Connect(s) => format!("CONNECT {}", s),
        }
    }
}

/**
 * Bring data on patient samples from the diagnosis machine to the laboratory with enough molecules to produce medicine!
 **/
fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let project_count = parse_input!(input_line, i32);
    for i in 0..project_count as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let a = parse_input!(inputs[0], i32);
        let b = parse_input!(inputs[1], i32);
        let c = parse_input!(inputs[2], i32);
        let d = parse_input!(inputs[3], i32);
        let e = parse_input!(inputs[4], i32);
    }

    // game loop
    loop {
        for i in 0..2 as usize {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let inputs = input_line.split(" ").collect::<Vec<_>>();
            let target = inputs[0].trim().to_string();
            let eta = parse_input!(inputs[1], i32);
            let score = parse_input!(inputs[2], i32);
            let storage_a = parse_input!(inputs[3], i32);
            let storage_b = parse_input!(inputs[4], i32);
            let storage_c = parse_input!(inputs[5], i32);
            let storage_d = parse_input!(inputs[6], i32);
            let storage_e = parse_input!(inputs[7], i32);
            let expertise_a = parse_input!(inputs[8], i32);
            let expertise_b = parse_input!(inputs[9], i32);
            let expertise_c = parse_input!(inputs[10], i32);
            let expertise_d = parse_input!(inputs[11], i32);
            let expertise_e = parse_input!(inputs[12], i32);
        }
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let available_a = parse_input!(inputs[0], i32);
        let available_b = parse_input!(inputs[1], i32);
        let available_c = parse_input!(inputs[2], i32);
        let available_d = parse_input!(inputs[3], i32);
        let available_e = parse_input!(inputs[4], i32);
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let sample_count = parse_input!(input_line, i32);
        
        let mut sample_queue:VecDeque<Sample>=VecDeque::new();
        
        for i in 0..sample_count as usize {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let inputs = input_line.split(" ").collect::<Vec<_>>();
            let sample_id = parse_input!(inputs[0], i32);
            let carried_by = parse_input!(inputs[1], i32);
            let rank = parse_input!(inputs[2], i32);
            let expertise_gain = inputs[3].trim().to_string();
            let health = parse_input!(inputs[4], i32);
            let cost_a = parse_input!(inputs[5], i32);
            let cost_b = parse_input!(inputs[6], i32);
            let cost_c = parse_input!(inputs[7], i32);
            let cost_d = parse_input!(inputs[8], i32);
            let cost_e = parse_input!(inputs[9], i32);
            sample_queue.push_back(Sample{sample_id,carried_by,rank,expertise_gain,health,cost_a,cost_b,cost_c,cost_d,cost_e});
        }

        // Write an action using println!("message...");
        eprintln!("{:?}",sample_queue);

        println!("GOTO DIAGNOSIS");
    }
}
