use std::io;
use std::collections::VecDeque;

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}
struct Card{
    value:i32,
    suit:char,
    }

//use enum
impl card{
    fn parse_card(c:str) -> Card{
        let mut value;
        let mut chars=c.chars();
        if let y=chars.next(){
        //here a match to get JQKA?
            if y.is_digit() {
                value=y;
            }
        //idea: iterate on str and fulfill the struct
        //if Some(chars.next() char::is_numeric) => value
        //else => suit
        }
    }
fn main() {
    let mut my_deck:VecDeque<Card>=VecDeque::new();
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let n = parse_input!(input_line, i32); // the number of cards for player 1
    for i in 0..n as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let cardp_1 = input_line.trim().to_string(); // the n cards of player 1
        //my_deck.push_back(card{
    }
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let m = parse_input!(input_line, i32); // the number of cards for player 2
    for i in 0..m as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let cardp_2 = input_line.trim().to_string(); // the m cards of player 2
    }

    // Write an action using println!("message...");
    // To debug: eprintln!("Debug message...");

    println!("PAT");
}
