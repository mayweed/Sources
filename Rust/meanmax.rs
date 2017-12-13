use std::io;
use std::f64;
use std::collections::{VecDeque};

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
//idea: looter must stay in watertown or not?
const MAP_RADIUS:f64=6000.0;
const WATERTOWN:f64=3000.0;
const MAX_THRUST:i32 = 300;
const MAX_RAGE:i32 = 300;
const WIN_SCORE:i32 = 50;
const EPSILON:f64=0.00001;
const MIN_IMPULSE:f64 = 30.0;
const IMPULSE_COEFF:f64 = 0.5;

//Tanker
const TANKER_THRUST:i32 = 500;
const TANKER_EMPTY_MASS:f64 = 2.5;
const TANKER_MASS_BY_WATER:f64 = 0.5;
const TANKER_FRICTION :f64= 0.40;
const TANKER_RADIUS_BASE :f64= 400.0;
const TANKER_RADIUS_BY_SIZE :f64= 50.0;
const TANKER_EMPTY_WATER :i32= 1;
const TANKER_MIN_SIZE:f64 = 4.0;
const TANKER_MAX_SIZE :f64= 10.0;
const TANKER_MIN_RADIUS :f64= TANKER_RADIUS_BASE + TANKER_RADIUS_BY_SIZE * TANKER_MIN_SIZE;
const TANKER_MAX_RADIUS :f64= TANKER_RADIUS_BASE + TANKER_RADIUS_BY_SIZE * TANKER_MAX_SIZE;
const TANKER_SPAWN_RADIUS :f64= 8000.0;
const TANKER_START_THRUST :i32= 2000;

//Reaper
const REAPER_MASS :f64= 0.5;
const REAPER_FRICTION :f64= 0.20;
const REAPER_SKILL_DURATION :i32= 3;
const REAPER_SKILL_COST :i32= 30;
const REAPER_SKILL_ORDER :i32= 0;
const REAPER_SKILL_RANGE :f64= 2000.0;
const REAPER_SKILL_RADIUS :f64= 1000.0;
const REAPER_SKILL_MASS_BONUS :f64= 10.0;

//Destroyer
const DESTROYER_MASS :f64= 1.5;
const DESTROYER_FRICTION :f64= 0.30;
const DESTROYER_SKILL_DURATION :i32= 1;
const DESTROYER_SKILL_COST :i32= 60;
const DESTROYER_SKILL_ORDER :i32= 2;
const DESTROYER_SKILL_RANGE :f64= 2000.0;
const DESTROYER_SKILL_RADIUS :f64= 1000.0;
const DESTROYER_NITRO_GRENADE_POWER :i32= 1000;

//Doof
const DOOF_MASS :f64= 1.0;
const DOOF_FRICTION :f64= 0.25;
const DOOF_RAGE_COEF :f64= 1.0 / 100.0;
const DOOF_SKILL_DURATION :i32= 3;
const DOOF_SKILL_COST :i32= 30;
const DOOF_SKILL_ORDER :i32= 1;
const DOOF_SKILL_RANGE :f64= 2000.0;
const DOOF_SKILL_RADIUS :f64= 1000.0;

//Looter common
const LOOTER_RADIUS :f64= 400.0;
const LOOTER_REAPER :i32= 0;
const LOOTER_DESTROYER :i32= 1;
const LOOTER_DOOF:i32 = 2;

//POINT
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
      
    //move the point to an other point for a given distance
    fn moveTo(mut self,p:Point,distance:f64){
        let d:f64=self.distance(&p);
        
        if d > EPSILON{
           let dx:f64=self.x - p.x;
           let dy:f64=self.y - p.y;
           let coef:f64=distance/d;
           
           self.x+= dx*coef;
           self.y+=dy*coef;
           }
    }
    
    //Yields true if a object is in radius range
    fn isInRange (self, p:Point,radius:f64) -> bool{
        if self.distance(&p)<=radius{
            return true
        } else {
            return false
        }
    }
        
}
    
//UNIT
enum unit_kind{
    //should correlate that with unit_type value
    reaper,
    destroyer,
    doof,
    tanker,
    wreck,
    }
#[derive(Debug)]
struct Unit{
    //kind:unit_kind,
    unitid:i32,
    unitType:i32,
    playerId:i32,
    mass:f64,
    radius:i32,
    point:Point,
    vx:i32,
    vy:i32,
    extra:i32,
    extra2:i32,
}
impl Unit{
    fn new() -> Unit{
        Unit{
            unitid:0,
            unitType:0,
            playerId:0,
            mass:0.0,
            radius:0,
            point:Point::new(),
            vx:0,
            vy:0,
            extra:0,
            extra2:0,
            }
    }
    //should add unit_kind one day...
    fn update(unitid:i32, unitType:i32, playerId:i32, mass:f64, radius:i32, point:Point, vx:i32, vy:i32, extra:i32, extra2:i32) -> Unit{
        Unit{
            //apply field init shorthand
            unitid,
            unitType,
            playerId,
            mass,
            radius,
            point,
            vx,
            vy,
            extra,
            extra2,
            }
    }
    //should I change vx/vy parse types?
    fn unit_move(mut self,t:f64) {
            self.point.x += self.vx as f64 * t;
            self.point.y += self.vy as f64 * t;
        }

    fn unit_speed(&self)->f64{
            return f64::sqrt(self.vx as f64* self.vx as f64+ self.vy as f64* self.vy as f64);
        }
       
     //SHOULD TAKE RAGE INTO ACCOUNT
    //for destroyers
    pub fn moveToTanker (&self,mut tankers:VecDeque<Unit>,rage:i32) -> String{
        //OOPS only if tankers is NOT empty
        //if it's empty you simply output "wait"
        //should modify via unitType to differentiate doof and destroyer
        //so that they act separately/independently!!
        if tankers.len() as i32 !=0{
            //persevere on the first?
            let mut tanker=tankers.pop_front().unwrap();
            if tanker.point.distance(&self.point)<REAPER_SKILL_RANGE && rage>DESTROYER_SKILL_COST{
                format!("SKILL {} {}",&tanker.point.x,&tanker.point.y)
                }else{
                format!("{} {} 300",&tanker.point.x,&tanker.point.y)
                }
        } else {
            format!("WAIT")
            }
        }
        
    //for reapers    
    //should evaluate the wrecks to target. Idea: a very simple wreck-scoring algorithm, (water content) * FACTOR - (distance to reaper). I forget what the factor was, but it was high like 12000 or something
    pub fn moveToWreck(&self,mut wrecks:VecDeque<Unit>) -> String{
        //OOPS only if there is wrecks!!
        //if it's empty you simply output "wait"
        if wrecks.len() as i32 !=0{
            //persevere on the first?
            let mut target=wrecks.pop_front().unwrap();
            if target.point.isInRange(self.point,LOOTER_RADIUS){
                format!("{} {} 200",&target.point.x,&target.point.y)
                }else{
                format!("{} {} 300",&target.point.x,&target.point.y)
                }
        } else {
            format!("WAIT")
            }
        }
    
    //for doofs 
    //TODO: modify to take into accounts ALL reapers from enemy + target
    // precisely which one to attack (score based? distance based?...)
    pub fn chaseTheReaper(&self,mut reaper:Unit,rage:i32) -> String{
        if reaper.point.distance(&self.point)<DOOF_SKILL_RANGE&&rage >DOOF_SKILL_COST {
            format!("SKILL {} {}",&reaper.point.x,&reaper.point.y)
        }else{
            format!("{} {} 300",&reaper.point.x,&reaper.point.y)
        }
        //    } else{
        //        format!("WAIT")
        //    }
        }
}

//PLAYER
#[derive(Debug)]
struct Player{
    id:i32, //0 for me, 1 and 2 for others
    score:i32,
    rage:i32,
    reaper:Unit,
    destroyer:Unit,
    doof:Unit,
    }
    
impl Player{
    //idea: get the max score player's reapers targeted by doofs!! 
    //what about equal scores?
    //fn best_score(self,p1:Player,p2:Player)-> Player{
        //empty struct?
        //let mut max_score=self;
        //match??
        //if self.score > p1.score && self.score > p2.score{
           // self
        //}
        //if p1.score > self.score && p1.score > p2.score{
        //    p1
        //    }
        //if p2.score > self.score && p2.score > p1.score{
       //     p2
       // }
        //max_score
    //}
}

//MAIN
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
        
        //init_entities gamestate?
        //should write a new new() func for this struct?
        let mut me=Player{id:0,score:my_score,rage:my_rage,reaper:Unit::new(),destroyer:Unit::new(),doof:Unit::new()};
        let mut enemy1=Player{id:1,score:enemy_score_1,rage:enemy_rage_1,reaper:Unit::new(),destroyer:Unit::new(),doof:Unit::new()};
        let mut enemy2=Player{id:2,score:enemy_score_2,rage:enemy_rage_2,reaper:Unit::new(),destroyer:Unit::new(),doof:Unit::new()};
        let mut tankers=VecDeque::new();
        let mut wrecks = VecDeque::new();
        
        for i in 0..unit_count as usize {
            let mut input_line = String::new();
            io::stdin().read_line(&mut input_line).unwrap();
            let inputs = input_line.split(" ").collect::<Vec<_>>();
            let unit_id = parse_input!(inputs[0], i32);
            let unit_type = parse_input!(inputs[1], i32);
            let player = parse_input!(inputs[2], i32);
            let mass = parse_input!(inputs[3], f64);
            let radius = parse_input!(inputs[4], i32);
            let x = parse_input!(inputs[5], f64);
            let y = parse_input!(inputs[6], f64);
            let point=Point{x,y};
            let vx = parse_input!(inputs[7], i32);
            let vy = parse_input!(inputs[8], i32);
            let extra = parse_input!(inputs[9], i32);
            let extra_2 = parse_input!(inputs[10], i32);
            
            if player == 0 {
                if unit_type == 0{
                    me.reaper=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                } else if unit_type ==1{
                    me.destroyer=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                } else if unit_type==2{
                    me.doof=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                }
            }else if player ==1{
                if unit_type == 0{
                    enemy1.reaper=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                } else if unit_type ==1{
                    enemy1.destroyer=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                } else if unit_type==2{
                    enemy1.doof=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                }
            }else if player==2{
                if unit_type == 0{
                    enemy2.reaper=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                } else if unit_type ==1{
                    enemy2.destroyer=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                } else if unit_type==2{
                    enemy2.doof=Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2);
                }
            } else if unit_type == 3{
                tankers.push_back(Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
            } else if unit_type == 4{    
                //should sort by distance:the center of the reaper in the wreck and reaper radius ==400 
                if me.reaper.point.distance(&Point{x,y}) > LOOTER_RADIUS{
                wrecks.push_back(Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                }else{
                wrecks.push_front(Unit::update(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                }
            }
        }
        
        //opt for the nearest tank with no enemy reapers on it?
        //let reaperGuillaume=me.reapers.pop_front().unwrap();
        let str1=me.reaper.moveToWreck(wrecks);
        println!("{}",str1);
        
        //let destroyerGuillaume=me.destroyers.pop_front().unwrap();
        let str2=me.destroyer.moveToTanker(tankers,me.rage);
        println!("{}",str2);
        
        //let doofGuillaume=me.doof.pop_front().unwrap();
        //should take into account all of the reapers!! and pick the
        //best one to hinder!!should target the enemy with bestScore!!
        //modify and use a value outside its scope in rust...
        if enemy1.score > enemy2.score{
            let str3=me.doof.chaseTheReaper(enemy1.reaper,me.rage);
            println!("{}",str3);
            }else {
            let str3=me.doof.chaseTheReaper(enemy2.reaper,me.rage);
            println!("{}",str3);
            }
           
        //THREE input lines!!
        //first line reaper
        //second destroyer
        //third doof
    }
}
