use std::io;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap());
}

/**
 * The while loop represents the game.
 * Each iteration represents a turn of the game
 * where you are given inputs (the heights of the mountains)
 * and where you have to print an output (the index of the mountain to fire on)
 * The inputs you are given are automatically updated according to your last actions.
 **/
 
fn main() {
    loop {
        //put that two inside the loop scope and it works?!
        let mut index=0;
        let mut mountainH=0;
        for i in 0..8 {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let mountain_h = parse_input!(input_line, i32); // represents the height of one mountain.
            //if the height of the current mountain is > to the last ever seen
            //should update my mountainH and index
            if mountain_h>mountainH{
                mountainH=mountain_h;
                index=i;
                }
        }
    println!("{}",index); // The index of the mountain to fire on.
    }
}
