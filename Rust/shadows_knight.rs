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
    x:i32,
    y:i32,
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
    let mut batPos=Point{x:x0,y:y0};
    let mut top=0;
    let mut bottom=h;
    let mut left=0;
    let mut right=w;
    
    // game loop
    loop {
        
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let bomb_dir = input_line.trim().to_string(); // the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
        //here want to use a match...
        if bomb_dir == "DL"{//.chars(){
            top=batPos.y;
            right=batPos.x;
        } else if bomb_dir=="DR"{
            top=batPos.y;
            left=batPos.x;
        }else if bomb_dir=="D"{
            top=batPos.y;
        } else if bomb_dir=="UR"{
            bottom=batPos.y;
            left=batPos.x;
        }else if bomb_dir=="UL"{
            bottom=batPos.y;
            right=batPos.x;
        }else if bomb_dir=="U"{
            bottom=batPos.y;
        }else if bomb_dir=="L"{
            right=batPos.x;
        }else if bomb_dir=="R"{
            left=batPos.x;
        }
        
        // Write an action using println!("message...");
        //print_err!("{}",bomb_dir);

        //Calculate the future coordinates
        let mut x=(left+right)/2;
        let mut y=(top+bottom)/2;
        
        //save the position of Batman!!
        batPos=Point{x,y};
        
        // the location of the next window Batman should jump to.
        println!("{} {}",x,y);
    }
}
