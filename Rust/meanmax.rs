use std::io;
use std::f64;
//use std::fmt; no need with format
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

const MAX_THRUST:i32=300;
const LOOTER_RADIUS:f64=400.0;
const EPSILON:f64=0.00001;

//POINT
#[derive(Debug)]
struct Point{
    x:f64,
    y:f64,

}

impl Point{
    fn distance (&self, p:Point) -> f64{
        //cast here...
        let x=self.x as f64;
        let y=self.y as f64;
        let other_x=p.x as f64;
        let other_y = p.y as f64;
        f64::sqrt((x-other_x)*(x-other_x) + (y-other_y)*(y-other_y)) 
       }
      
    //move the point to an other point for a given distance
    fn moveTo(self,p:Point,distance:f64){
        let d:f64=self.distance(p);
        
        if d > EPSILON{
           let dx:f64=self.x - p.x;
           let dy:f64=self.y - p.y;
           let coef:f64=distance/d;
           
           //cast no good here??
           //change as f64 parse input x/y
           self.x+= dx*coef;
           self.y+=dy*coef;
           }
    }
    
      //r cf looter radius
    fn isInRange (&self, p:Point,r:f64) -> bool{
        if self.distance(p)<=r{
            true
        } else {
            false
        }
    }
        
}
    
//UNIT
#[derive(Debug)]
struct Unit{
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
    //why not update instead of new??
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

//idea: get the max score player's reapers targeted by doofs!! 
#[derive(Debug)]
struct Player<'a>{
    id:i32, //0 for me, 1 and 2 for others
    score:i32,
    //rage:i32,??
    //just need ref wont make any change here!!
    reapers:&'a VecDeque<Unit>,
    destroyers:&'a VecDeque<Unit>,
    doofs:&'a VecDeque<Unit>,
    }
    
impl <'a> Player<'a>{
    //pub fn new(id:i32,score:i32,reapers:&VecDeque<Unit>,destroyers:&VecDeque<Unit>,doofs:&VecDeque<Unit>) -> Player{
    //    Player{
    //        id,
    //        score,
    //        reapers:&reapers,
    //        destroyers:&destroyers,
    //        doofs:&doofs,
    //    }
    //}
    
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
    //fn getReapers(reapers)
    //for x in reapers if player id==1 
}

//GAMESTATE
//??
//struct GameState{
// expected liftime parameters for player??
//    players:Player,
//    }
    


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
        let mut myReapers=VecDeque::new();
        let mut enemy1Reapers=VecDeque::new();
        let mut enemy2Reapers=VecDeque::new();
        
        let mut myDestroyers=VecDeque::new();
        let mut enemy1Destroyers=VecDeque::new();
        let mut enemy2Destroyers=VecDeque::new();
        
        let mut myDoofs=VecDeque::new(); 
        let mut enemy1Doofs=VecDeque::new(); 
        let mut enemy2Doofs=VecDeque::new(); 
        
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
            
            //match => non exhaustive pattern in unit_type '_'
            //UGLY of the UGLIEST!! should factor that one day!!
            if player == 0 {
                if unit_type == 0{
                    myReapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    myDestroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    myDoofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                }
            }else if player ==1{
                if unit_type == 0{
                    enemy1Reapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    enemy1Destroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    enemy1Doofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                }
            }else if player==2{
                if unit_type == 0{
                    enemy2Reapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    enemy2Destroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    enemy2Doofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
                }
            } else if unit_type == 3{
                tankers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
            } else if unit_type == 4{    
                wreckTanks.push_back(Unit::new(unit_id,unit_type,player,mass,radius,point,vx,vy,extra,extra_2));
            }
        }
        
        //cant make it work with all that borrowing thing!!
        //let me= Player{id:0,score:my_score,reapers:&myReapers,destroyers:&myDestroyers,doofs:&myDoofs};
        //print_err!("{:#?}",myDestroyers);
        //opt for the nearest tank with no enemy reapers on it?
        let reaperGuillaume=myReapers.pop_front().unwrap();
        let str1=reaperGuillaume.moveToWreck(wreckTanks);
        
        let destroyerGuillaume=myDestroyers.pop_front().unwrap();
        let str2=destroyerGuillaume.moveToTanker(tankers);
        
        //if score1 > score2 tu prends les reapers 1 comme cible
        //si score2 > score 1 tu prends les reapers 2 comme cible
        let doofGuillaume=myDoofs.pop_front().unwrap();
        let str3=if enemy_score_1 > enemy_score_2{
            doofGuillaume.chaseTheReaper(enemy1Reapers); 
        }else if enemy_score_1 < enemy_score_2{
            doofGuillaume.chaseTheReaper(enemy2Reapers); 
        };
        
        //THREE input lines!!
        //first line reaper
        //second destroyer
        //third doof
        println!("{}",str1);
        println!("{}",str2);
        println!("{}",str3);
        
        print_err!("{:?}",str3);
    }
}
