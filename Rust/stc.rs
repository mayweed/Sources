extern crate rand;
use rand::{thread_rng,Rng};

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

//usize easier for indexing Vec<i32> which does not implement trait index<i32> blablabla...
const HEIGHT:usize=12;
const WIDTH:usize=6;
const SKULL:i32=0;
const FREE:i32=-1;

#[derive(Debug)]
struct Stone{
    color_a:i32,
    color_b:i32,
}

//wont accept Copy??
#[derive(Clone,Debug)]
struct Board{
    row:Vec<i32>,
    columns:Vec<i32>,
    grid:Vec<Vec<i32>>,
    }

impl Board{
fn new() -> Board{
    Board{
        row:Vec::new(),
        columns:Vec::new(),
        grid:vec![vec![FREE;WIDTH];HEIGHT],
        }
    }
//should i use a closure here?
fn get_column(&self,index:usize) -> Vec<i32>{
    let mut cols:Vec<i32>=Vec::new();
    for column in self.grid.iter(){
            cols.push(column[index])
        }
    cols
    }
    
fn get_row(&self,index:usize) -> &Vec<i32>{
    &self.grid[index]
    }
    
//closure on |color| ??
fn is_empty(colrow:&Vec<i32>) -> bool{
    for color in colrow.iter(){
        if *color != FREE{
            return false
        }
    }
    true
}
// I DONT NEED THAT: should mem last color and last column in gamestate wtf!!            
fn check_color(self,color_stone:i32) -> i32{
    let mut col:i32=0;
    for num in 0..6{
        let mut c=&self.get_column(num);
        for color_col in c.iter(){
            if *color_col==color_stone{ //et case précédente libre{
                col=num as i32;
                break;
                }
            }
        }
       col 
    }
        
    
//then you place stones during 99ms (timeout==100ms)
//and you choose the best one?? DEPTH??
//fn place_stones(self,s:Vec<Stone>) -> i32{}
  
}
//very simple idea:
//if current col==last col and col[1] FREE output last column

struct game_state{
    boardstate:Board,
    last_color:i32,
    last_column:i32,
    }
//need to build a simu of the grid
//with fmt::Display to watch it

//take a string output a line of i32 in a vec
//should be a board func?
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

fn main() {
    
    let mut turn:i32=0;
    
    //to store future pieces
    let mut pieces:VecDeque<Stone>=VecDeque::new();
    
    let mut my_board:Board=Board::new();
    
    let mut new_board=vec![vec![FREE;WIDTH as usize];HEIGHT as usize];
    print_err!("{:?}",new_board);
    
    let mut last_col:u32=0;
    
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
            my_board.grid.push(parse_row(row));
        }
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let score_2 = parse_input!(input_line, i32);
        for i in 0..12 as usize {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let row = input_line.trim().to_string();
            
        }
        let mut rng = thread_rng();
        let n: u32 = rng.gen_range(0,6);
        let color=pieces[0].color_a;
        
        //not twice the same :)
        if n==last_col{
            let mut m = rng.gen_range(0,6);
            println!("{}",m);
        }else{
            println!("{}", n);
        }
            
        //TESTS
        print_err!("{:?}",&my_board.get_column(0));
        //should work on test!!
        print_err!("{:?}",Board::is_empty(&my_board.get_row(11)));
        //it works?? no need of indexMut???
        print_err!("{:?}",&my_board.grid[11-1][1]);//(11));
        print_err!("{} {}",n,last_col);
        
        my_board.grid.clear();
        pieces.pop_front();
        turn+=1;
        last_col=n;
    }
}
