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

//Looter
const LOOTER_RADIUS :f64= 400.0;
const LOOTER_REAPER :i32= 0;
const LOOTER_DESTROYER :i32= 1;
const LOOTER_DOOF:i32 = 2;

//POINT
#[derive(Debug)]
struct Point{
    x:f64,
    y:f64,

}

impl Point{
    fn distance (&self, p:&Point) -> f64{
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
    fn isInRange (&self, p:Point,radius:f64) -> bool{
        if self.distance(&p)<=radius{
            true
        } else {
            false
        }
    }
        
}
    
//UNIT
enum unit_kind{
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
    //should add unit_kind one day...
    fn new(unitid:i32, unitType:i32, playerId:i32, mass:f64, radius:i32, point:Point, vx:i32, vy:i32, extra:i32, extra2:i32) -> Unit{
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
                
    //idea: output a string with X/Y/THROTTLE
    //IDEA SKILL=when adverse reaper is in range of my reaper, Skill instead of ...
    //tankers list should be pass as pointer you modify it??
    pub fn moveToTanker (&self,mut tankers:VecDeque<Unit>) -> String{
        //OOPS only if tankers is NOT empty
        //if it's empty you simply output "wait"
        //should modify via unitType to differentiate doof and destroyer
        //so that they act separately/independently!!
        if tankers.len() as i32 !=0{
            //persevere on the first?
            let mut tanker=tankers.pop_front().unwrap();
            format!("{} {} 300",&tanker.point.x,&tanker.point.y)
        } else {
            format!("WAIT")
            }
        }
        
    //for reapers    
    pub fn moveToWreck(&self,mut wrecks:VecDeque<Unit>) -> String{
        //OOPS only if there is wrecks!!
        //if it's empty you simply output "wait"
        if wrecks.len() as i32 !=0{
            //persevere on the first?
            let mut target=wrecks.pop_front().unwrap();
            format!("{} {} 300",&target.point.x,&target.point.y)
        } else {
            format!("WAIT")
            }
        }
    
    //for doofs very minimalist (no dist etc...)
    //mut reapers: cannot borrow immutable argument 'reapers' as mutable
    pub fn chaseTheReaper(&self,mut reapers:VecDeque<Unit>) -> String{
        if reapers.len() as i32 !=0{
            let mut target=reapers.pop_front().unwrap();
            format!("{} {} 200",&target.point.x,&target.point.y)
            } else{
                format!("WAIT")
            }
        }
        
}

//PLAYER
#[derive(Debug)]
struct Player{
    id:i32, //0 for me, 1 and 2 for others
    score:i32,
    rage:i32,
    reapers:VecDeque<Unit>,
    destroyers:VecDeque<Unit>,
    doofs:VecDeque<Unit>,
    }
    
impl Player{
    //idea: get the max score player's reapers targeted by doofs!! 
    //what about equal scores?
    //yield the id of the player?
    fn best_score(self,p1:Player,p2:Player)->i32{
        let mut max_score=0;
        if self.score > p1.score && self.score > p2.score{
            max_score=self.id;
        }else if p1.score > self.score && p1.score > p2.score{
            max_score=p1.id;
        }else if p2.score > self.score && p2.score > p1.score{
            max_score=p2.id;
        }
        max_score
    }
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
        let mut me=Player{id:0,score:my_score,rage:my_rage,reapers:VecDeque::new(),destroyers:VecDeque::new(),doofs:VecDeque::new()};
        let mut enemy1=Player{id:1,score:enemy_score_1,rage:enemy_rage_1,reapers:VecDeque::new(),destroyers:VecDeque::new(),doofs:VecDeque::new()};
        let mut enemy2=Player{id:2,score:enemy_score_2,rage:enemy_rage_2,reapers:VecDeque::new(),destroyers:VecDeque::new(),doofs:VecDeque::new()};
        let mut tankers=VecDeque::new();
        let mut wreckTanks = VecDeque::new();
        
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
                    me.reapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    me.destroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    me.doofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                }
            }else if player ==1{
                if unit_type == 0{
                    enemy1.reapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    enemy1.destroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    enemy1.doofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                }
            }else if player==2{
                if unit_type == 0{
                    enemy2.reapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    enemy2.destroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    enemy2.doofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                }
            } else if unit_type == 3{
                tankers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
            } else if unit_type == 4{    
                wreckTanks.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
            }
        }
        
        //cant make it work with all that borrowing thing!!
        //let me= Player{id:0,score:my_score,rage:my_rage,reapers:&myReapers,destroyers:&myDestroyers,doofs:&myDoofs};
        //print_err!("{:#?}",myDestroyers);
        //opt for the nearest tank with no enemy reapers on it?
        let reaperGuillaume=me.reapers.pop_front().unwrap();
        let str1=reaperGuillaume.moveToWreck(wreckTanks);
        
        let destroyerGuillaume=me.destroyers.pop_front().unwrap();
        let str2=destroyerGuillaume.moveToTanker(tankers);
        
        //let doofGuillaume=myDoofs.pop_front().unwrap();
        //let str3={
           
        //};
        //if score1 > score2 tu prends les reapers 1 comme cible
        //si score2 > score 1 tu prends les reapers 2 comme cible
        
        
        //THREE input lines!!
        //first line reaper
        //second destroyer
        //third doof
        println!("{}",str1);
        println!("{}",str2);
        
         //if enemy_score_1 > enemy_score_2{
               // println!("{}",doofGuillaume.chaseTheReaper(enemy1Reapers));
                //try to pass the ref to the string outside the scope
                //&s;
                //}
           // if enemy_score_1 < enemy_score_2{
           //     println!("{}",doofGuillaume.chaseTheReaper(enemy2Reapers));
                //&s;
           // }  
           // else {
           //     println!("WAIT");
                //&s;
            //    }
        println!("WAIT");
        //println!("{}",&str3);
        
        //print_err!("{}",str2);
    }
}
