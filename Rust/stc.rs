use std::io;
use std::collections::{VecDeque};

//First let's keep things simple:
// store the future pieces
// check the current board state
// if any pieces got the same color and there is space left in col
// put the stones here!!

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

const HEIGHT:i32=12;
const WIDTH:i32=6;
const FREE:i32=-1;

#[derive(Debug)]
struct Stone{
    color_a:i32,
    color_b:i32,
}


//probably the most important struct !!
//#[derive(Debug)]
type Board=Vec<Vec<i32>>;
//impl Board{
    //yield an empty *free* board
fn new() -> Board{
    let mut new_board:Board=vec![vec![FREE;WIDTH as usize];HEIGHT as usize];
    //w/o return it does not work...
    return new_board
  }

//take a string output a line of i32 in a vec
//fill the board quoi
//take the board, take the pieces yield the int num of a col
//fn place_stones(self,s:Vec<Stone>) -> i32{
//this is not a board func?
fn parse_row(r:String) -> Vec<i32>{
    let mut dum:Vec<i32>=Vec::new();
    for elem in r.chars(){
        if elem == '.'{
            dum.push(FREE);
        }else{
            //quite complicated but as i32 alone wont make it
            dum.push(elem.to_digit(10).unwrap() as i32);
            }
        }
    dum
    }
//check board: is there group of same colors?


fn main() {
    
    let mut turn:i32=0;
    
    //to store future pieces
    let mut pieces:VecDeque<Stone>=VecDeque::new();
    
    let mut board:Vec<Vec<i32>>=Vec::new();
    
    // game loop
    loop {
    
        for i in 0..8 as usize {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let inputs = input_line.split(" ").collect::<Vec<_>>();
            let color_a = parse_input!(inputs[0], i32); // color of the first block
            let color_b = parse_input!(inputs[1], i32); // color of the attached block
            pieces.push_back(Stone{color_a,color_b});
        }
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let score_1 = parse_input!(input_line, i32);
        for i in 0..12 as usize {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let row = input_line.trim().to_string(); // One line of the map ('.' = empty, '0' = skull block, '1' to '5' = colored block)
            board.push(parse_row(row));
        }
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let score_2 = parse_input!(input_line, i32);
        for i in 0..12 as usize {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let row = input_line.trim().to_string();
            
        }

        for num in 0..6{
            println!("{}", num); // "x": the column in which to drop your blocks
            }
            
        turn+=1;
        
        print_err!("{:?}",board[11][0]);
        board.clear();
    }
}
