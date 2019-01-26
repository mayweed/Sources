package main

import (
	"fmt"
	"log"
)

//constants
const (
	MAP_HEIGHT   = 1000
	MAP_WIDTH    = 1920
	QUEEN_RADIUS = 30
	QUEEN_SPEED  = 60
)

//MAP-Grid
type (
	Point struct {
		x, y int
	}

	Cell struct {
		pos  Point
		what int //siteID or unitType
	}
)

//SITES
type (
	buildingSites struct {
		siteId int
		pos    Point
		radius int
	}
	Goldmine struct {
		siteId        int
		pos           Point
		radius        int
		structureType int
		owner         int
		incomeRate    int
		param2        int
	}
	Tower struct {
		siteId        int
		pos           Point
		radius        int
		structureType int
		owner         int
		healthPoint   int
		attackRadius  int
	}
	Barrack struct {
		siteId        int
		pos           Point
		radius        int
		structureType int
		owner         int
		turnsInactive int
		creepType     int
	}
)

//UNITS
type Unit struct {
	pos      Point
	owner    int
	unitType int
	health   int
}

//PLAYER
type Player struct {
	//a player sites
	goldmine []Goldmine
	towers   []Tower
	barracks []Barrack
	//units by type
	queen   Unit
	knights []Unit
	archers []Unit
	giants  []Unit
}

//GAME STATE
type State struct {
	players [2]Player
	grid    [MAP_HEIGHT][MAP_WIDTH]Cell
}

/* OLD contest code
func dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}
func distUnitToSite(q Unit, s Site) float64 {
	return dist(q.pos.x, q.pos.y, s.pos.x, s.pos.y)
}
func distUnitToUnit(q Unit, r Unit) float64 {
	return dist(q.pos.x, q.pos.y, r.pos.x, r.pos.y)
}
func inQueenRange(dist float64) bool {
	return dist <= QUEEN_RADIUS
}
func In(sl []int, s int) bool {
	for _, num := range sl {
		if num == s {
			return true
		}
	}
	return false
}

//commands
func mv(from, to int) string {
	s := fmt.Sprintf("MOVE %d %d", from, to)
	return s
}
func formatTrainComm(sl []int) string {
	var s = []string{"TRAIN"}
	if len(sl) > 0 {
		for _, site := range sl {
			s = append(s, strconv.Itoa(site))
		}
		return strings.Join(s, " ")
	} else {
		return "TRAIN"
	}
}

//no pointer, map not adressable
func (s *Site) updateSite(s2 Site) {
	s.structureType = s2.structureType
	s.owner = s2.owner
	s.param1 = s2.param1
	s.param2 = s2.param2
}

func (p Player) pickNextSite(sl map[int]*Site) int {
	var minDist = MAP_WIDTH
	var siteMinId int
	for _, site := range sl {
		//avoid enemy towers to spare health?
		if site.structureType != 1 && site.owner != 1 {
			dd := distUnitToSite(p.queen, *site)
			if dd < minDist {
				minDist = dd
				siteMinId = site.siteId
			}
		}
	}
	return siteMinId
}

//if knights in range, build tower!!
func (p Player) knightsInRange(k []Unit) bool {
	for _, knight := range k {
		dd := distUnitToUnit(p.queen, knight)
		log.Println(dd)
		//Queen radius + 30
		if dd < 60 {
			return true
		}
	}
	return false
}

//Will be helpful with pickNextSite() for ex
func (p PlayingField) getSiteCoord(id int) (float64, float64) {
	for _, site := range p.sites {
		if site.siteId == id {
			return site.x, site.y
		}
	}
	//if no site corresponds to id?
	return -1, -1
}

//establish potential target list
func (p PlayingField) filterListSites(s map[int]*Site) map[int]*Site {
	var newTargets = make(map[int]*Site)
	for id, site := range s {
		//add enemy barracks,avoid towers..
		if site.structureType == -1 || site.owner == 1 && site.structureType == 2 {
			newTargets[id] = site
		}
	}
	return newTargets
}

//BE PRAGMATIC: to go silver i must be more healthy, first avoid enemy towers
//and get to cover behind my towers...
//gives the x point where there is most towers?
//IDEA: add that tier as a member in struct site and init it at the beginning of the game??
func (p Player) coverYourAss() (int, int, int) {
	//0-640 (tier0)/641-1280(tier1)/1281-1920(tier2)
	var tier0 int
	var tier1 int
	var tier2 int
	for _, site := range p.towers {
		if site.x <= 640 {
			tier0 += 1
		} else if site.x > 640 && site.x < 1280 {
			tier1 += 1
		} else {
			tier2 += 1
		}
	}
	return tier0, tier1, tier2
}
*/
//MAIN
func main() {

	board := State{}
	me := board.players[0]
	opp := board.players[1]

	//should init my grid
	//then setSites, units etc...

	var numSites int
	fmt.Scan(&numSites)

	for i := 0; i < numSites; i++ {
		var x, y int
		var siteId, radius int
		fmt.Scan(&siteId, &x, &y, &radius)
		board.sites[siteId] = &Site{siteId: siteId, x: x, y: y, radius: radius}
		grid[y][x] = Cell{}
	}

	for {
		// touchedSite: -1 if none
		var gold, touchedSite int
		fmt.Scan(&gold, &touchedSite)
		me.queen.gold = gold
		if touchedSite != -1 {
			if !In(me.queen.hasTouched, touchedSite) {
				me.queen.hasTouched = append(me.queen.hasTouched, touchedSite)
			}
		}

		for i := 0; i < numSites; i++ {
			// ignore1: used in future leagues
			// ignore2: used in future leagues
			// structureType: -1 = No structure, 0=Goldmine, 1= Tower, 2 = Barracks
			// owner: -1 = No structure, 0 = Friendly, 1 = Enemy
			var siteId, ignore1, ignore2, structureType, owner, param1, param2 int
			fmt.Scan(&siteId, &ignore1, &ignore2, &structureType, &owner, &param1, &param2)
			if owner == 0 {
				switch structureType {
				case 0:
					me.goldmine = append(me.goldmine, Site{siteId: siteId, structureType: structureType, owner: owner, param1: param1, param2: param2})
				case 1:
					me.towers = append(me.towers, Site{siteId: siteId, structureType: structureType, owner: owner, param1: param1, param2: param2})
				case 2:
					me.barracks = append(me.barracks, Site{siteId: siteId, structureType: structureType, owner: owner, param1: param1, param2: param2})
				}
				board.sites[siteId].updateSite(Site{structureType: structureType, owner: owner, param1: param1, param2: param2})
			} else if owner == 1 {
				switch structureType {
				case 0:
					opp.goldmine = append(opp.goldmine, Site{siteId: siteId, structureType: structureType, owner: owner, param1: param1, param2: param2})
				case 1:
					opp.towers = append(opp.towers, Site{siteId: siteId, structureType: structureType, owner: owner, param1: param1, param2: param2})
				case 2:
					opp.barracks = append(opp.barracks, Site{siteId: siteId, structureType: structureType, owner: owner, param1: param1, param2: param2})
				}
				board.sites[siteId].updateSite(Site{structureType: structureType, owner: owner, param1: param1, param2: param2})
			} else {
				//no structure, no owner...
				board.sites[siteId].updateSite(Site{structureType: structureType, owner: owner, param1: param1, param2: param2})
			}
		}
		var numUnits int
		fmt.Scan(&numUnits)

		for i := 0; i < numUnits; i++ {
			// unitType: -1 = QUEEN, 0 = KNIGHT, 1 = ARCHER, 2=GIANT
			var x, y float64
			var owner, unitType, health int
			fmt.Scan(&x, &y, &owner, &unitType, &health)
			if unitType == -1 {
				if owner == 0 {
					me.queen = Unit{x: x, y: y, owner: owner, unitType: unitType, health: health}
				} else {
					opp.queen = Unit{x: x, y: y, owner: owner, unitType: unitType, health: health}
				}
			} else if unitType == 0 {
				if owner == 0 {
					me.knights = append(me.knights, Unit{x: x, y: y, owner: owner, unitType: unitType, health: health})
				} else {
					opp.knights = append(opp.knights, Unit{x: x, y: y, owner: owner, unitType: unitType, health: health})
				}
			} else if unitType == 1 {
				if owner == 0 {
					me.archers = append(me.archers, Unit{x: x, y: y, owner: owner, unitType: unitType, health: health})
				} else {
					opp.archers = append(opp.archers, Unit{x: x, y: y, owner: owner, unitType: unitType, health: health})
				}
			} else if unitType == 2 {
				if owner == 0 {
					me.giants = append(me.giants, Unit{x: x, y: y, owner: owner, unitType: unitType, health: health})
				} else {
					opp.giants = append(opp.giants, Unit{x: x, y: y, owner: owner, unitType: unitType, health: health})
				}
			}
		}
		//LOGS
		//log.Println(board.filterListSites(board.sites))
		//log.Println(me.knightsInRange(opp.knights))
		log.Println(me.coverYourAss())
		log.Println(me.queenInTowerAttackRadius(opp.towers))

		// First line: A valid queen action
		//FLAW: I dont take into account opp moves in knighted and towered!!
		//filter the target list, could refine that I think (taking into account the opp for ex?)
		//IDEA: check queen range from radius to see if there is units/sites int it
		//in that case tower!!
		//IDEA: use WAIT to hide queen behind towers and spare health
		//IF enemy tower is in range of queen (radiius) just keep it at distance
		//IDEA: use giant in sites center of the map!!
		//IF creeps in range queen, build tower!!
		var target = me.pickNextSite(board.filterListSites(board.sites))
		//there is a structure so move on!!

		if len(me.towers) > len(opp.towers) && (!me.knightsInRange(opp.knights)) { //&& board.sites[touchedSite].structureType==1{
			fmt.Println("WAIT")
		} else if touchedSite == -1 || board.sites[touchedSite].structureType != -1 {
			s := mv(int(board.sites[target].x), int(board.sites[target].y))
			fmt.Println(s)
		} else {
			if gold < 30 {
				sc := fmt.Sprintf("BUILD %d MINE", touchedSite)
				fmt.Println(sc)
			} else if len(opp.towers) >= 2 || len(me.knights) >= 3 || me.knightsInRange(opp.knights) {
				sb := fmt.Sprintf("BUILD %d TOWER", touchedSite)
				fmt.Println(sb)
				//idea: you got a defensive wall of tower, stay behind
			} else {
				sd := fmt.Sprintf("BUILD %d BARRACKS-KNIGHT", touchedSite)
				me.knighted = append(me.knighted, touchedSite)
				fmt.Println(sd)
			}
		}

		// Second line: A set of training instructions
		//bugged func here: Gold is not accounted for!!
		//my func is a BLOB mixing formatting and calculs so...
		sum := 0
		var sl []int
		for _, id := range me.knighted {
			if sum += 80; sum < gold {
				sl = append(sl, id)
			} else {
				break
			}
		}
		t := formatTrainComm(sl)
		fmt.Println(t)

		//clear at the end of each turn to avoid false num of sites etc..
		me.goldmine = []Site{}
		opp.goldmine = []Site{}
		me.towers = []Site{}
		opp.towers = []Site{}
		me.barracks = []Site{}
		opp.barracks = []Site{}
		me.knights = []Unit{}
		me.archers = []Unit{}
		me.giants = []Unit{}
		opp.knights = []Unit{}
		opp.archers = []Unit{}
		opp.giants = []Unit{}
	}
}
