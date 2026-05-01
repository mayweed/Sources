#include <bits/stdc++.h> //All main STD libraries
#include <x86intrin.h> //SSE Extensions
using namespace std;

#define TESTSIM
//#define DUMP
#define FOR0(i,n) for(int i=0;i<(n);++i)
	
const int DEPTH = 5;

const int ON_AIR = 0;
const int LANDED = 1;
const int LAND_BAD_ANGLE = 2;
const int LAND_BAD_SPEED = 3;
const int CRASH = 4;


const double MAX_LAND_VSPEED = 40.0f;
const double MAX_LAND_HSPEED = 20.0f;
const double GRAVITY = -3.711;
const double T = 1.0;



class Line {
  public:
    double x0,y0,x1,y1,s1_x,s1_y;
    Line(){x0=0.0f;x1=0.0f;y0=0.0f;y1=0.0f;}
    Line(double _x0,double _y0,double _x1,double _y1)
    {
      x0=_x0;y0=_y0;x1=_x1;y1=_y1;
      s1_x = x1 - x0;
      s1_y = y1 - y0;
    }
    bool collides(double landerx0,double landery0,double landerx1,double landery1){
      double s2_x = landerx1 - landerx0;
      double s2_y = landery1 - landery0;
      double coef = 1 / (-s2_x * s1_y + s1_x * s2_y);
      double s = (-s1_y * (x0 - landerx0) + s1_x * (y0 - landery0)) *coef;
      double t = ( s2_x * (y0 - landery0) - s2_y * (x0 - landerx0)) *coef;
      return (s >= 0 && s <= 1 && t >= 0 && t <= 1);      
    }
    bool isLandingZone(){
        return s1_y == 0;
    }
};
int surfaceN;
Line LandingZone;
vector<Line> landLines;
ostream &operator<<(ostream& output, const Line& l){ 
	    output << "Line: ("<<l.x0<<","<<l.y0<<")->("<<l.x1<<","<<l.y1<<")  Dist:"<<l.s1_x<<","<<l.s1_y;
		return output;
	}


class Genoma{
public:
    int rotation[DEPTH];
    double power[DEPTH];
	double Score;
};
ostream &operator<<(ostream& output, const Genoma& g){ 
output << "int[] rotation = {";   
   FOR0(i,DEPTH) output << (i>0?",":"")<<g.rotation[i];
output << "};"<<endl;
output << "int[] power = {";   
   FOR0(i,DEPTH) output << (i>0?",":"")<<g.power[i];
output << "}; //Score:"<<g.Score<<endl;
		return output;
	}
	
	
	
class GameState{
public:
  int turn;
  double prev_x,prev_y,x,y,vx,vy,ax,ay,fuel,power;
  int angle;
  GameState(){turn = -1;ax=0;ay=0;angle=90;}
  
  void ReadConfig(istream& input){
    input >> surfaceN; input.ignore();
    #ifdef DUMP
    cerr << surfaceN<<endl;
    #endif
    double oldX, oldY;
    for (int i = 0; i < surfaceN; i++) {
        double landX;
        double landY;
        input >> landX >> landY; input.ignore();
        #ifdef DUMP
         cerr << landX<<" "<<landY<<endl;
        #endif
        if(i>0){
          Line newLine = Line(oldX,oldY, landX,landY);
          if(newLine.isLandingZone()){
            LandingZone = newLine;
            #ifndef DUMP
            cerr << i<< " LANDZONE "<<newLine<<endl;          
            #endif
          }
            
          
          landLines.push_back(newLine);
        }
        oldX = landX;
        oldY = landY;
    }    
    #ifndef DUMP
     for (int i = 0; i < surfaceN-1; i++)
      cerr <<(i+1)<<" :"<< landLines[i]<<endl;
    #endif      
      
  }
  
  void ReadTurn(istream& input)
  {
      input >> x >> y >> vx >> vy >> fuel >> angle >> power; input.ignore();
      
      ++turn;
     if (turn == 0)
     {
         prev_x = x;
         prev_y = y;
         #ifdef DUMP
         cerr << x << " "<< y << " " << vx << " " <<vy << " "<< fuel << " "<< angle << " "<< power<<endl;
         #endif
     }
     angle += 90;
  }
  
  bool Equals(const GameState& gs);
  
  int isLanded()
  {
      int result = ON_AIR;
     //Check Landing
     if (x <0) return CRASH;
     if (y <0) return CRASH;
     if (y >=3000) return CRASH;
     if (x >=7000) return CRASH;
     if (LandingZone.collides(prev_x,prev_y,x,y))
     {
         if (angle != 90) 
         {
      //       cerr << "Error: Bad Angle on Land "<<angle<<" != 90"<<endl;
             return LAND_BAD_ANGLE;
         }
         if (abs(vy)>MAX_LAND_VSPEED) 
         {
    //         cerr << "Error: Too Fast on Vspeed "<<vy<<" > "<<MAX_LAND_VSPEED<<endl;
             return LAND_BAD_SPEED;
         }
         if (abs(vx)>MAX_LAND_HSPEED) 
         {
  //           cerr << "Error: Too Fast on Hspeed "<<vy<<" > "<<MAX_LAND_HSPEED<<endl;
             return LAND_BAD_SPEED;
         }
         return LANDED;
     }
     for (int i = 0; i < surfaceN-1; i++)
     {
         if (landLines[i].collides(prev_x,prev_y,x,y))
         {
//             cerr << "Collision with Land"<<(i+1)<<endl;
             return CRASH;
         }
     }
     return result;
  }
  
double toRadians(double angle){
  return angle * (M_PI / 180);
}
 
double sinDeg(double deg){
  return sin(toRadians(deg));
}
 
double cosDeg(double deg){
  return cos(toRadians(deg));
}  
  
  void ApplyGenoma(const Genoma& g,int step)
  {
    if (g.power[step] > power) power = power+1;
    if (g.power[step] < power) power = power-1;
    if (g.rotation[step] > angle) angle += min(15,g.rotation[step]-angle);
    if (g.rotation[step] < angle) angle -= min(15,angle-g.rotation[step]);
    fuel -= power;    
    ax = cosDeg((double)angle)*power;   
    ay = sinDeg((double)angle)*power+GRAVITY;
    prev_x = x;
    prev_y = y;
    x += vx+ ax*0.5;
    y += vy+ ay*0.5;
    vx += ax;
    vy += ay;
  }  
  
  
};
GameState gamestate,working;
ostream &operator<<(ostream& output, const GameState& gs){ 
	    output << "Lander:P:("<<gs.x<<","<<gs.y<<") V:("<<gs.vx<<","<<gs.vy<<") F:"<<gs.fuel<<" A:"<<gs.angle<<" P:"<<gs.power;
		return output;
	}
bool GameState::Equals(const GameState& gs){
      bool result =  (x==round(gs.x) && y == round(gs.y) && vx == round(gs.vx) && vy == round(gs.vy) && fuel == round(gs.fuel));
      //if (!result)
      {
          cerr << "GameState:"<<*this<<endl;
          cerr << "Working:"<<gs<<endl;
          if (!result) abort();
      }
      return result;
  }


void TestSim(istream& input)
{
   while(1)
   {
      gamestate.ReadTurn(input);
      if (gamestate.turn > 0)
      {
         cerr << "Compare Result:"<< gamestate.Equals(working)<<endl;
      }      
      else {
          working = gamestate;
      }
      Genoma g;
      g.rotation[0] = (gamestate.turn*3)%180;
      g.power[0] = min(4,((gamestate.turn+1)%14));      
      working.ApplyGenoma(g,0);                
      cerr << "isLanded="<<working.isLanded()<<endl;
      cout << (g.rotation[0]-90)<<" "<<g.power[0]<<endl;
   }
}


int main(){
   srand(time(0));
   gamestate.ReadConfig(cin);
   TestSim(cin);
}
