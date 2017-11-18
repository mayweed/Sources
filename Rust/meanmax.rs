use std::io;

//macro_rules! print_err {
//    ($($arg:tt)*) => (
//        {
//            use std::io::Write;
//            writeln!(&mut ::std::io::stderr(), $($arg)*).ok();
//        }
//    )
//}

macro_rules! parse_input {
    ($x:expr, $t:ident) => ($x.trim().parse::<$t>().unwrap())
}

//POINT
struct Point{
    x:i32,
    y:i32,
}
    
//UNIT
struct Unit{
    unitid:i32,
    unitType:i32,
    playerId:i32,
    mass:f64,
    radius:i32,
    x:i32,
    y:i32,
    vx:i32,
    vy:i32,
    extra:i32,
    extra2:i32,
}
impl Unit{
    fn new(unitid:i32, unitType:i32, playerId:i32, mass:f64, radius:i32, x:i32, y:i32, vx:i32, vy:i32, extra:i32, extra2:i32) -> Unit{
        Unit{
            unitid:unitid,
            unitType:unitType,
            playerId:playerId,
            mass:mass,
            radius:radius,
            x:x,
            y:y,
            vx:vx,
            vy:vy,
            extra:extra,
            extra2:extra2,
            }
    }
    pub fn getTanks(&self) -> bool{
        if self.unitType == 4{
            true
        }else{
            false
            }
    }
    pub fn getReaper(&self) -> bool{
        if self.unitType == 0{
            true
        }else{
            false
            }
    }
                
    //take a reaper move it to nearby tanks
    //pub fn moveToTanks{
}

//GAMESTATE
//??

fn main() { 
    // game loop
    loop {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let my_score = parse_input!(input_line, i32);
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let enemy_score_1 = parse_input!(input_line, i32);
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let enemy_score_2 = parse_input!(input_line, i32);
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let my_rage = parse_input!(input_line, i32);
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let enemy_rage_1 = parse_input!(input_line, i32);
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let enemy_rage_2 = parse_input!(input_line, i32);
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let unit_count = parse_input!(input_line, i32);
        
        let mut units=vec![];
        
        for i in 0..unit_count as usize {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let inputs = input_line.split(" ").collect::<Vec<_>>();
            let unit_id = parse_input!(inputs[0], i32);
            let unit_type = parse_input!(inputs[1], i32);
            let player = parse_input!(inputs[2], i32);
            let mass = parse_input!(inputs[3], f64);
            let radius = parse_input!(inputs[4], i32);
            let x = parse_input!(inputs[5], i32);
            let y = parse_input!(inputs[6], i32);
            let vx = parse_input!(inputs[7], i32);
            let vy = parse_input!(inputs[8], i32);
            let extra = parse_input!(inputs[9], i32);
            let extra_2 = parse_input!(inputs[10], i32);
            units.push(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
        }

        // Write an action using println!("message...");
        // To debug: print_err!("Debug message...");

        println!("WAIT");
        println!("WAIT");
        println!("WAIT");
    }
}
