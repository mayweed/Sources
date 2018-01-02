//the 1st in rust got 2296 of fuel left!!
//objective here: simulate mars landing and use monte carlo or gen algo
//https://pastebin.com/3SUUjgSF
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
//CONSTANTS
const GRAVITY:f64=3.711;

//FLAT_GROUND
#[derive(Debug,Clone,Copy)]
struct flat_ground{
    x0:f64,
    y0:f64,
    x1:f64,
    y1:f64,
    s1_x:f64,
    s1_y:f64,
    }
    
impl flat_ground{
    fn new(x0:f64,y0:f64,x1:f64,y1:f64) -> flat_ground{
        flat_ground{
            x0,
            y0,
            x1,
            y1,
            //if y0_y1 between two points == 0 it's flat
            s1_x:x0-x1,
            s1_y:y0-y1,
            }
        }
        
    //return a bool
    fn collide(self,landerx0:f64,landery0:f64,landerx1:f64,landery1:f64) -> bool{
        let mut s2_x = landerx1 - landerx0;
        let mut s2_y = landery1 - landery0;
        let mut coef = 1.0 / (-s2_x * self.s1_y + self.s1_x * s2_y);
        let mut s = (-self.s1_y * (self.x0 - landerx0) + self.s1_x * (self.y0 - landery0)) *coef;
        let mut t = ( s2_x * (self.y0 - landery0) - s2_y * (self.x0 - landerx0)) *coef;
        s >= 0.0 && s <= 1.0 && t >= 0.0 && t <= 1.0    
    }
    fn is_landing_zone(self) -> bool{
        self.s1_y ==0.0
    }
}

fn main() {
    //should factor that under flat_ground as method "find_landing_zone"
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).unwrap();
    let surface_n = parse_input!(input_line, i32); // the number of points used to draw the surface of Mars.
    let mut old_x=0.0;
    let mut old_y=0.0;
    let mut landlines:Vec<flat_ground>=Vec::new();
    //am not satisfied with that!! it works but well...
    //i do want my exp above to work...
    let mut landing_zone=flat_ground::new(old_x,old_y,0.0,0.0);
    
    for i in 0..surface_n as usize {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let land_x = parse_input!(inputs[0], f64); // X coordinate of a surface point. (0 to 6999)
        let land_y = parse_input!(inputs[1], f64); // Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
        let line=flat_ground::new(old_x,old_y,land_x,land_y);
        landlines.push(line);
        if line.is_landing_zone() { landing_zone=line;};
        old_x=land_x;
        old_y=land_y;
    }
    print_err!("{:?}",landing_zone);
    //cant get that type of thing to work in rust!!
    //mismatched types: expected (); found reference
    //let landing_zone= for line in landlines.iter(){
    //    if line.is_landing_zone(){
    //        line
    //        }
    //     };
         
    // game loop
    loop {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let x = parse_input!(inputs[0], i32);
        let y = parse_input!(inputs[1], i32);
        let h_speed = parse_input!(inputs[2], i32); // the horizontal speed (in m/s), can be negative.
        let v_speed = parse_input!(inputs[3], i32); // the vertical speed (in m/s), can be negative.
        let fuel = parse_input!(inputs[4], i32); // the quantity of remaining fuel in liters.
        let rotate = parse_input!(inputs[5], i32); // the rotation angle in degrees (-90 to 90).
        let power = parse_input!(inputs[6], i32); // the thrust power (0 to 4).
        
        print_err!("hspeed:{}, vspeed:{}, fuel:{}, rotate:{}, power:{}",h_speed,v_speed,fuel,rotate,power);
        
        // rotate power. rotate is the desired rotation angle. power is the desired thrust power.
        println!("0 2");
    }
}
