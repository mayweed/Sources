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

//POINT
struct Point{
    x:i32,
    y:i32,
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
    x:i32,
    y:i32,
    vx:i32,
    vy:i32,
    extra:i32,
    extra2:i32,
}
impl Unit{
    //why not update instead of new??
    fn new(unitid:i32, unitType:i32, playerId:i32, mass:f64, radius:i32, x:i32, y:i32, vx:i32, vy:i32, extra:i32, extra2:i32) -> Unit{
        Unit{
            //apply field init shorthand
            unitid,
            unitType,
            playerId,
            mass,
            radius,
            x,
            y,
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
    //tankers list should be pass as pointer you modify it??
    pub fn moveToTanker (&self,mut tankers:VecDeque<Unit>) -> String{
        //OOPS only if tankers is NOT empty
        //if it's empty you simply output "wait"
        //should modify via unitType to differentiate doof and destroyer
        //so that they act separately/independently!!
        if tankers.len() as i32 !=0{
            //persevere on the first?
            let mut tanker=tankers.pop_front().unwrap();
            format!("{} {} 300",&tanker.x,&tanker.y)
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
            format!("{} {} 300",&target.x,&target.y)
        } else {
            format!("WAIT")
            }
        }
        
}

//idea: get the max score player's reapers targeted by doofs!! 
#[derive(Debug)]
struct Player{
    id:i32, //0 for me, 1 and 2 for others
    score:i32,
    //rage:i32,??
    reapers:VecDeque<Unit>,
    destroyers:VecDeque<Unit>,
    doofs:VecDeque<Unit>,
    }

//GAMESTATE
//??
#[derive(Debug)]
struct GameState{
    players:Player,
    }
    


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
            let x = parse_input!(inputs[5], i32);
            let y = parse_input!(inputs[6], i32);
            let vx = parse_input!(inputs[7], i32);
            let vy = parse_input!(inputs[8], i32);
            let extra = parse_input!(inputs[9], i32);
            let extra_2 = parse_input!(inputs[10], i32);
            
            //match => non exhaustive pattern in unit_type '_'
            //UGLY of the UGLIEST!! should factor that one day!!
            if player == 0 {
                if unit_type == 0{
                    myReapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    myDestroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    myDoofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                }
            }else if player ==1{
                if unit_type == 0{
                    enemy1Reapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    enemy1Destroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    enemy1Doofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                }
            }else if player==2{
                if unit_type == 0{
                    enemy2Reapers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                } else if unit_type ==1{
                    enemy2Destroyers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                } else if unit_type==2{
                    enemy2Doofs.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
                }
            } else if unit_type == 3{
                tankers.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
            } else if unit_type == 4{    
                wreckTanks.push_back(Unit::new(unit_id,unit_type,player,mass,radius,x,y,vx,vy,extra,extra_2));
            }
        }
        //print_err!("{:#?}",myDestroyers);
        //opt for the nearest tank with no enemy reapers on it?
        let mut reaperGuillaume=myReapers.pop_front().unwrap();
        let mut str1=reaperGuillaume.moveToWreck(wreckTanks);
        
        let mut destroyerGuillaume=myDestroyers.pop_front().unwrap();
        let mut str2=destroyerGuillaume.moveToTanker(tankers);
        
        //just test not good!!
        let mut doofGuillaume=myDoofs.pop_front().unwrap();
        let mut str3=doofGuillaume.moveToTanker(tankers);
        
            
        //THREE input lines!!
        //first line reaper
        //second destroyer
        //third doof
        println!("{}",str1);
        println!("{}",str2);
        println!("WAIT");
        
        print_err!("{}",str2);
    }
}
