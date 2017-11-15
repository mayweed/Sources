use std::io;

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

struct Point{
    x: i32,
    y: i32,
    }
struct Area {
    //X
    left:i32,
    right:i32,
    //Y
    top:i32,
    bottom:i32,
    }
 
//Output next bat position
fn batPos(zone:Area) -> Point{
    let x=(zone.left+zone.right)/2;
    let y=(zone.top+zone.bottom)/2;
    Point{x,y}
    }
    
    

fn main() {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let w = parse_input!(inputs[0], i32); // width of the building.
    let h = parse_input!(inputs[1], i32); // height of the building.
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let n = parse_input!(input_line, i32); // maximum number of turns before game over.
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let inputs = input_line.split(" ").collect::<Vec<_>>();
    let x0 = parse_input!(inputs[0], i32);
    let y0 = parse_input!(inputs[1], i32);
    let mut batLoc=Point{x:x0,y:y0};
    
    // game loop
    loop {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let bomb_dir = input_line.trim().to_string(); // the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
        //here want to use a match...
        if bomb_dir == "DL" {
            batLoc=batPos(Area{left:0,right:batLoc.x-1,top:batLoc.y+1,bottom:h});
            }
        if bomb_dir =="DR"{
            batLoc=batPos(Area{left:batLoc.x+1,right:w,top:batLoc.y+1,bottom:h});
            }
        if bomb_dir =="UR"{
            batLoc=batPos(Area{left:0,right:batLoc.x-1,top:0,bottom:batLoc.y-1});
            }
        if bomb_dir == "UL"{
            batLoc=batPos(Area{left:batLoc.x+1,right:w,top:0,bottom:batLoc.y-1});
            }
        //those are the easy cases indeed does not pass tower, you have to 
        //recompute the area!!
        if bomb_dir=="D"{
            batLoc=batPos(Area{left:batLoc.x,right:batLoc.x,top:batLoc.y,bottom:h});
            }
        if bomb_dir=="U"{
            batLoc=batPos(Area{left:batLoc.x,right:batLoc.x,top:0,bottom:batLoc.y});
            }
        if bomb_dir=="L"{
            batLoc=batPos(Area{left:0,right:batLoc.x-1,top:batLoc.y,bottom:batLoc.y});
            }
        if bomb_dir=="R"{
            batLoc=batPos(Area{left:batLoc.x+1,right:w,top:batLoc.y,bottom:batLoc.y});
            }
        // Write an action using println!("message...");
        print_err!("W:{}, H:{}, N:{}, bomb_dir:{},x0:{},y0:{}, batLoc.x {} batloc.y {}",w,h,n,bomb_dir,x0,y0,batLoc.x,batLoc.y);


        // the location of the next window Batman should jump to.
        println!("{} {}",batLoc.x,batLoc.y);
    }
}
