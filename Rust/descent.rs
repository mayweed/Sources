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
 //SHOULD YIELD THE *INDEX* not the number in itself but it works my first working rust func ;)
 fn largest (high:&[i32]) -> i32{
    let mut max=high[0];
    for &item in high.iter() {
        if item > max{
            max = item;
        }
    }
    max 
}
fn main() {

    let mut high = vec![];
    loop {
        for i in 0..8 {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let mountain_h = parse_input!(input_line, i32); // represents the height of one mountain.
            high.push(mountain_h);
        }
    let maxi=largest(&high); 
    println!("{}",maxi); // The index of the mountain to fire on.
    }
}
