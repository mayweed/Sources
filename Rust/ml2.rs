//the 1st in rust got 2296 of fuel left!!
//objective here: simulate mars landing and use monte carlo or gen algo
//https://pastebin.com/3SUUjgSF
use std::io;
use std::f64::consts::PI;

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
const GRAVITY:f64=-3.711;
const MAX_LAND_VSPEED:f64 = 40.0;
const MAX_LAND_HSPEED:f64 = 20.0;
const T:f64 = 1.0;

const ON_AIR:i32 = 0;
const LANDED:i32 = 1;
const LAND_BAD_ANGLE:i32 = 2;
const LAND_BAD_SPEED:i32 = 3;
const CRASH:i32 = 4;

//TEST
const DEPTH:i32 = 5;

//POINT should I use <T> here??
#[derive(Debug,Copy,Clone)]
struct Point{
    x:f64,
    y:f64,

}

impl Point{
    fn new() -> Point{
        Point{
            x:0.0,
            y:0.0,
            }
    }
    fn distance (self, p:&Point) -> f64{
        let x=self.x;
        let y=self.y; 
        let other_x=p.x;
        let other_y = p.y;
        f64::sqrt((x-other_x)*(x-other_x) + (y-other_y)*(y-other_y)) 
       }
}

//FLAT_GROUND
#[derive(Debug,Clone,Copy)]
struct flat_ground{
    point0:Point,
    point1:Point,
    s1_x:f64,
    s1_y:f64,
    }
    
impl flat_ground{
    fn new(x0:f64,y0:f64,x1:f64,y1:f64) -> flat_ground{
        flat_ground{
            point0:Point{x:x0,y:y0},
            point1:Point{x:x1,y:y1},
            //if y0_y1 between two points == 0 it's flat
            s1_x:x0-x1,
            s1_y:y0-y1,
            }
    }
    
    //return a bool:will next pos of lander collides with surface?
    fn collide(self,landerx0:f64,landery0:f64,landerx1:f64,landery1:f64) -> bool{
        let mut s2_x = landerx1 - landerx0;
        let mut s2_y = landery1 - landery0;
        let mut coef = 1.0 / (-s2_x * self.s1_y + self.s1_x * s2_y);
        let mut s = (-self.s1_y * (self.point0.x - landerx0) + self.s1_x * (self.point0.y - landery0)) *coef;
        let mut t = ( s2_x * (self.point0.y - landery0) - s2_y * (self.point0.x - landerx0)) *coef;
        s >= 0.0 && s <= 1.0 && t >= 0.0 && t <= 1.0    
    }
    fn is_landing_zone(self) -> bool{
        self.s1_y ==0.0
    }
    
    fn find_landing_zone() -> flat_ground{
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let surface_n = parse_input!(input_line, i32); // the number of points used to draw the surface of Mars.
        let mut old_x=0.0;
        let mut old_y=0.0;
        let mut landlines:Vec<flat_ground>=Vec::new();
        //am not satisfied with that!! it works but well...
        //i do want my exp below to work...
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
        //diff iter() into_iter() ter(), which iterates over &T.
        //iter_mut(), which iterates over &mut T.
        //into_iter(), which iterates over T.
        //let for line in landlines.into_iter(){
        //    if line.is_landing_zone(){
        //        line
        //        }
        //     } = landing_zone;
        landing_zone
    }
    //where to put that with variables??
    //fn is_Landed(&self)->i32 {
        //result = ON_AIR;
        //Check Landing
        //if x <0{return CRASH}
        //if y <0{return CRASH}
        //if y >=3000{return CRASH}
        //if x >=7000{return CRASH}
     
        //if self.collide(prev_x,prev_y,x,y){
            //if (angle != 90)
                //{
                //    return LAND_BAD_ANGLE;
                //}
            //if (abs(vy)>MAX_LAND_VSPEED)
              //  {
            //        return LAND_BAD_SPEED;
            //    }
            //if (abs(vx)>MAX_LAND_HSPEED)
            //    {
            //    return LAND_BAD_SPEED;
            //    }
            //return LANDED;
        //}
        
        //for (int i = 0; i < surfaceN-1; i++){
         //if (landLines[i].collides(prev_x,prev_y,x,y))
         //{
         //    return CRASH;
         //}
     //}
     //return result;
  //}
}

//TRIGO => macros??
fn toRadians (angle:f64) -> f64{
  angle * (PI / 180.0)
}
 
fn sinDeg(deg:f64) -> f64{
  f64::sin(toRadians(deg))
}
 
fn cosDeg (deg:f64) -> f64{
  f64::cos(toRadians(deg))
}

//GAMESTATE
struct game_state{
    turn:i32,
    last_position:Point, //for oldx/oldy
    curr_position:Point,
    vx:f64,
    vy:f64,
    ax:f64,
    ay:f64,
    fuel:f64,
    power:f64,
    }
//impl game_state -> i32 {

    
struct Genoma{
    rotation:Vec<f64>,
    power:Vec<f64>,
    score:f64,
}
fn main() {
    //first find a landing spot
    let mut landing_ground=flat_ground::find_landing_zone();
    let mut turn=1;
    let mut prev_x=0.0;
    let  mut prev_y=0.0;
    
    // game loop
    loop {
        let mut input_line = String::new();
        io::stdin().read_line(&mut input_line).unwrap();
        let inputs = input_line.split(" ").collect::<Vec<_>>();
        let x = parse_input!(inputs[0], f64);
        let y = parse_input!(inputs[1], f64);
        let h_speed = parse_input!(inputs[2], f64); // the horizontal speed (in m/s), can be negative.
        let v_speed = parse_input!(inputs[3], f64); // the vertical speed (in m/s), can be negative.
        let fuel = parse_input!(inputs[4], f64); // the quantity of remaining fuel in liters.
        let rotate = parse_input!(inputs[5], i32); // the rotation angle in degrees (-90 to 90).
        let power = parse_input!(inputs[6], i32); // the thrust power (0 to 4).
        
        let mut position:Point=Point{x:x, y:y };
        turn+=1;
        print_err!("hspeed:{}, vspeed:{}, fuel:{}, rotate:{}, power:{},turn:{},x:{},prev_x:{}",h_speed,v_speed,fuel,rotate,power,turn,x,prev_x);
        prev_x=x;
        prev_y=y;
        // rotate power. rotate is the desired rotation angle. power is the desired thrust power.
       
        if position.distance(&landing_ground.point0) <=1500.0{
            println!("0 4");
        }else{
            //if v_speed > -50.0{
                println!("-20 3");
               // }else{
               // println!("-20 4");
               // }
            }
        print_err!("{}",position.distance(&landing_ground.point0));
    }
    
}
