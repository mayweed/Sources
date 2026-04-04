package main

import (
	"fmt"
	"math"
)

/* ======================================================================
   CONSTANTS
   ====================================================================== */

const (
	MAP_WIDTH               = 23
	MAP_HEIGHT              = 21
	INITIAL_SHIP_HEALTH     = 100
	MAX_SHIP_HEALTH         = 100
	MAX_SHIP_SPEED          = 2
	MINE_VISIBILITY_RANGE   = 5
	LOW_DAMAGE              = 25
	HIGH_DAMAGE             = 50
	MINE_DAMAGE             = 25
	NEAR_MINE_DAMAGE        = 10
	FIRE_DISTANCE_MAX       = 10
	COOLDOWN_CANNON         = 2
	REWARD_RUM_BARREL_VALUE = 30
)

/* ======================================================================
   GEOMETRY : POINT & CUBE COORDS
   ====================================================================== */

type Point struct {
	x, y int
}

func (p Point) neighbour(orientation int) Point {
	var EVEN = [6][2]int{{1, 0}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}}
	var ODD = [6][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, 0}, {0, 1}, {1, 1}}

	if p.y%2 == 0 {
		return Point{
			x: p.x + EVEN[orientation][0],
			y: p.y + EVEN[orientation][1],
		}
	}
	return Point{
		x: p.x + ODD[orientation][0],
		y: p.y + ODD[orientation][1],
	}
}

func (p Point) isInsideMap() bool {
	return p.x >= 0 && p.x < MAP_WIDTH && p.y >= 0 && p.y < MAP_HEIGHT
}

func (p Point) toCubeCoordinate() cubeCoord {
	cx := p.x - (p.y-(p.y&1))/2
	cz := p.y
	cy := -(cx + cz)
	return cubeCoord{cx, cy, cz}
}

func (p Point) distanceTo(dest Point) float64 {
	return p.toCubeCoordinate().distanceTo(dest.toCubeCoordinate())
}

/* ======================================================================
   CUBE COORDINATES
   ====================================================================== */

type cubeCoord struct {
	x, y, z int
}

func (c cubeCoord) distanceTo(d cubeCoord) float64 {
	return (math.Abs(float64(c.x-d.x)) +
		math.Abs(float64(c.y-d.y)) +
		math.Abs(float64(c.z-d.z))) * 0.5
}

/* ======================================================================
   ENTITIES
   ====================================================================== */

type Entity struct {
	id         int
	entityType string
	pos        Point
}

/* ======================================================================
   SHIP STRUCTURE
   ====================================================================== */

type Ship struct {
	Entity

	orientation    int
	speed          int
	health         int
	owner          int
	cannonCooldown int
	mineCooldown   int

	actionType string
	target     Entity

	newOrientation     int
	newPosition        Point
	newBowCoordinate   Point
	newSternCoordinate Point

	initialHealth int
	isDead        bool
}

/* ======================================================================
   SHIP BASIC METHODS
   ====================================================================== */

func (s *Ship) heal(a int) {
	s.health += a
	if s.health > MAX_SHIP_HEALTH {
		s.health = MAX_SHIP_HEALTH
	}
}

func (s *Ship) damage(a int) {
	s.health -= a
	if s.health <= 0 {
		s.health = 0
		s.isDead = true
	}
}

func (s Ship) bow() Point {
	return s.pos.neighbour(s.orientation)
}
func (s Ship) stern() Point {
	return s.pos.neighbour((s.orientation + 3) % 6)
}

func (s Ship) newBow() Point {
	return s.pos.neighbour(s.newOrientation)
}
func (s Ship) newStern() Point {
	return s.pos.neighbour((s.newOrientation + 3) % 6)
}

/* ======================================================================
   COLLISION HELPERS
   ====================================================================== */

func (s Ship) newBowIntersect(other Ship) bool {
	return s.newBowCoordinate == other.newBowCoordinate ||
		s.newBowCoordinate == other.newPosition ||
		s.newBowCoordinate == other.newSternCoordinate
}

func (s Ship) newPositionIntersect(other Ship) bool {
	sternCollision := s.newSternCoordinate == other.newBowCoordinate ||
		s.newSternCoordinate == other.newPosition ||
		s.newSternCoordinate == other.newSternCoordinate

	centerCollision := s.newPosition == other.newBowCoordinate ||
		s.newPosition == other.newPosition ||
		s.newPosition == other.newSternCoordinate

	return s.newBowIntersect(other) || sternCollision || centerCollision
}

func (s Ship) newBowsIntersect(ships []Ship) bool {
	for _, other := range ships {
		if other.id != s.id && !other.isDead {
			if s.newBowIntersect(other) {
				return true
			}
		}
	}
	return false
}

func (s Ship) newPositionsIntersect(ships []Ship) bool {
	for _, o := range ships {
		if o.id != s.id && s.newPositionIntersect(o) {
			return true
		}
	}
	return false
}

/* ======================================================================
   SET ACTIONS
   ====================================================================== */

func (s *Ship) wait()      { s.actionType = "WAIT" }
func (s *Ship) faster()    { s.actionType = "FASTER" }
func (s *Ship) slower()    { s.actionType = "SLOWER" }
func (s *Ship) port()      { s.actionType = "PORT" }
func (s *Ship) starboard() { s.actionType = "STARBOARD" }

func (s *Ship) move(p Point) {
	s.actionType = "MOVE"
	s.target.pos = p
}

func (s *Ship) fire(p Point) {
	s.actionType = "FIRE"
	s.target.pos = p
}

/* ======================================================================
   ENTITIES
   ====================================================================== */

type Barrel struct {
	Entity
	health int
}

type Mine struct {
	Entity
	isTargeted bool
}

type cannonball struct {
	target         Point
	fromShip       int
	remainingTurns int
}

/* ======================================================================
   GAME STATE
   ====================================================================== */

type State struct {
	entityCount           int
	myShipCount           int
	turn                  int
	enemyShips            []Ship
	allyShips             []Ship
	ships                 []Ship
	barrels               []Barrel
	mines                 []Mine
	cannonballs           []cannonball
	cannonballsExplosions []Point
}

/* ======================================================================
   COPY (deep copy with synchronization!)
   ====================================================================== */

func (s *State) Copy() State {
	cp := *s

	// Basic copies
	cp.ships = append([]Ship(nil), s.ships...)
	cp.allyShips = append([]Ship(nil), s.allyShips...)
	cp.enemyShips = append([]Ship(nil), s.enemyShips...)
	cp.barrels = append([]Barrel(nil), s.barrels...)
	cp.mines = append([]Mine(nil), s.mines...)
	cp.cannonballs = append([]cannonball(nil), s.cannonballs...)
	cp.cannonballsExplosions = append([]Point(nil), s.cannonballsExplosions...)

	// IMPORTANT : synchronize allyShips with ships by ID
	for i := range cp.allyShips {
		id := cp.allyShips[i].id
		for j := range cp.ships {
			if cp.ships[j].id == id {
				cp.allyShips[i] = cp.ships[j]
				break
			}
		}
	}

	for i := range cp.enemyShips {
		id := cp.enemyShips[i].id
		for j := range cp.ships {
			if cp.ships[j].id == id {
				cp.enemyShips[i] = cp.ships[j]
				break
			}
		}
	}

	return cp
}

/* ======================================================================
   SIMULATION STEP: decrement rum, cannonballs
   ====================================================================== */

func (s *State) updateInitialRum() {
	for i := range s.ships {
		s.ships[i].initialHealth = s.ships[i].health
	}
}

func (s *State) decrementRum() {
	for i := range s.ships {
		s.ships[i].damage(1)
	}
}

func (s *State) movecannonballs() {
	for i := len(s.cannonballs) - 1; i >= 0; i-- {
		if s.cannonballs[i].remainingTurns == 0 {
			s.cannonballsExplosions = append(s.cannonballsExplosions, s.cannonballs[i].target)
			s.cannonballs = append(s.cannonballs[:i], s.cannonballs[i+1:]...)
		} else {
			s.cannonballs[i].remainingTurns--
		}
	}
}

/* ======================================================================
   APPLY ACTIONS
   ====================================================================== */

func (s *State) applyActions() {
	for i := range s.ships {
		ship := &s.ships[i]

		ship.newOrientation = ship.orientation

		switch ship.actionType {
		case "FASTER":
			if ship.speed < MAX_SHIP_SPEED {
				ship.speed++
			}
		case "SLOWER":
			if ship.speed > 0 {
				ship.speed--
			}
		case "PORT":
			ship.newOrientation = (ship.orientation + 1) % 6
		case "STARBOARD":
			ship.newOrientation = (ship.orientation + 5) % 6

		case "MOVE":
			// no immediate effect; used during movement phase

		case "FIRE":
			if ship.cannonCooldown == 0 {
				d := ship.bow().distanceTo(ship.target.pos)
				if d <= FIRE_DISTANCE_MAX {
					flight := int(1 + math.Round(d/3.0))
					s.cannonballs = append(s.cannonballs,
						cannonball{target: ship.target.pos, fromShip: ship.id, remainingTurns: flight})
					ship.cannonCooldown = COOLDOWN_CANNON
				}
			}
		}
	}
}

/* ======================================================================
   MOVE SHIPS (with collision resolution)
   ====================================================================== */

func (s *State) moveShips() {
	for step := 1; step <= MAX_SHIP_SPEED; step++ {

		// Compute candidate new positions
		for i := range s.ships {
			ship := &s.ships[i]

			if ship.isDead || step > ship.speed {
				continue
			}

			next := ship.pos.neighbour(ship.orientation)
			if next.isInsideMap() {
				ship.newPosition = next
				ship.newBowCoordinate = next.neighbour(ship.orientation)
				ship.newSternCoordinate = next.neighbour((ship.orientation + 3) % 6)
			} else {
				ship.speed = 0
				ship.newPosition = ship.pos
				ship.newBowCoordinate = ship.bow()
				ship.newSternCoordinate = ship.stern()
			}
		}

		// Resolve bow collisions
		for {
			collision := false

			for i := range s.ships {
				ship := &s.ships[i]
				if ship.isDead {
					continue
				}

				if ship.newBowsIntersect(s.ships) {
					ship.newPosition = ship.pos
					ship.newBowCoordinate = ship.bow()
					ship.newSternCoordinate = ship.stern()
					ship.speed = 0
					collision = true
				}
			}

			if !collision {
				break
			}
		}

		// Apply
		for i := range s.ships {
			ship := &s.ships[i]
			if !ship.isDead {
				ship.pos = ship.newPosition
			}
		}

		s.checkCollisions()
	}
}

/* ======================================================================
   ROTATION PHASE
   ====================================================================== */

func (s *State) rotateShips() {
	for i := range s.ships {
		ship := &s.ships[i]
		ship.newPosition = ship.pos
		ship.newBowCoordinate = ship.newBow()
		ship.newSternCoordinate = ship.newStern()
	}

	for {
		block := false

		for i := range s.ships {
			ship := &s.ships[i]

			if ship.newPositionsIntersect(s.ships) {
				ship.newOrientation = ship.orientation
				ship.newBowCoordinate = ship.bow()
				ship.newSternCoordinate = ship.stern()
				ship.speed = 0
				block = true
			}
		}

		if !block {
			break
		}
	}

	for i := range s.ships {
		s.ships[i].orientation = s.ships[i].newOrientation
	}

	s.checkCollisions()
}

/* ======================================================================
   COLLISIONS (barrels, mines)
   ====================================================================== */

func (s *State) checkCollisions() {
	for si := range s.ships {
		ship := &s.ships[si]
		if ship.isDead {
			continue
		}

		bow := ship.bow()
		stern := ship.stern()
		center := ship.pos

		// Barrels
		for i := len(s.barrels) - 1; i >= 0; i-- {
			if s.barrels[i].pos == bow ||
				s.barrels[i].pos == stern ||
				s.barrels[i].pos == center {

				ship.heal(s.barrels[i].health)
				s.barrels = append(s.barrels[:i], s.barrels[i+1:]...)
			}
		}

		// Mines
		for i := len(s.mines) - 1; i >= 0; i-- {
			if s.mines[i].pos == bow ||
				s.mines[i].pos == stern ||
				s.mines[i].pos == center {

				ship.damage(MINE_DAMAGE)

				// near-mine damage
				for j := range s.ships {
					other := &s.ships[j]
					if other.id == ship.id || other.isDead {
						continue
					}

					if s.mines[i].pos.distanceTo(other.pos) == 1 ||
						s.mines[i].pos.distanceTo(other.bow()) == 1 ||
						s.mines[i].pos.distanceTo(other.stern()) == 1 {
						other.damage(NEAR_MINE_DAMAGE)
					}
				}

				s.mines = append(s.mines[:i], s.mines[i+1:]...)
			}
		}
	}
}

/* ======================================================================
   EXPLOSIONS
   ====================================================================== */

func (s *State) explodeShips() {
	for i := len(s.cannonballsExplosions) - 1; i >= 0; i-- {
		ex := s.cannonballsExplosions[i]

		for si := range s.ships {
			ship := &s.ships[si]
			if ship.isDead {
				continue
			}

			if ex == ship.bow() || ex == ship.stern() {
				ship.damage(LOW_DAMAGE)
				s.cannonballsExplosions = append(s.cannonballsExplosions[:i], s.cannonballsExplosions[i+1:]...)
				break
			}
			if ex == ship.pos {
				ship.damage(HIGH_DAMAGE)
				s.cannonballsExplosions = append(s.cannonballsExplosions[:i], s.cannonballsExplosions[i+1:]...)
				break
			}
		}
	}
}

func (s *State) explodeMines() {
	for i := len(s.cannonballsExplosions) - 1; i >= 0; i-- {
		ex := s.cannonballsExplosions[i]

		for j := len(s.mines) - 1; j >= 0; j-- {
			if s.mines[j].pos == ex {
				s.mines = append(s.mines[:j], s.mines[j+1:]...)
				s.cannonballsExplosions = append(s.cannonballsExplosions[:i], s.cannonballsExplosions[i+1:]...)
				break
			}
		}
	}
}

func (s *State) explodeBarrels() {
	for i := len(s.cannonballsExplosions) - 1; i >= 0; i-- {
		ex := s.cannonballsExplosions[i]

		for j := len(s.barrels) - 1; j >= 0; j-- {
			if s.barrels[j].pos == ex {
				s.barrels = append(s.barrels[:j], s.barrels[j+1:]...)
				s.cannonballsExplosions = append(s.cannonballsExplosions[:i], s.cannonballsExplosions[i+1:]...)
				break
			}
		}
	}
}

/* ======================================================================
   SHIP DEATH -> BARRELS
   ====================================================================== */

func (s *State) sinkShipMakeRum() {
	for si := range s.ships {
		ship := &s.ships[si]

		if ship.health <= 0 {
			reward := int(math.Min(float64(ship.initialHealth), float64(REWARD_RUM_BARREL_VALUE)))

			if reward > 0 {
				s.barrels = append(s.barrels, Barrel{
					Entity: Entity{pos: ship.pos},
					health: reward,
				})
			}
		}
	}
}

/* ======================================================================
   FULL SIMULATION
   ====================================================================== */

func simulateTurn(s *State) {
	s.updateInitialRum()
	s.movecannonballs()
	s.decrementRum()
	s.applyActions()
	s.moveShips()
	s.rotateShips()
	s.explodeShips()
	s.explodeMines()
	s.explodeBarrels()
	s.sinkShipMakeRum()
	s.turn++
}

/* ======================================================================
   AI UTILITIES
   ====================================================================== */

func (s *State) getClosestEnemyShip(my Ship) Ship {
	best := Ship{}
	bestDist := 999.0

	for _, e := range s.enemyShips {
		d := my.pos.distanceTo(e.pos)
		if d < bestDist {
			bestDist = d
			best = e
		}
	}
	return best
}

func (s *State) getClosestBarrel(my Ship) Barrel {
	best := Barrel{}
	bestDist := 999.0

	for _, b := range s.barrels {
		d := my.pos.distanceTo(b.pos)
		if d < bestDist {
			bestDist = d
			best = b
		}
	}
	return best
}

func (s Ship) nextPosShip(t int) Point {
	p := s.pos
	for i := 0; i < t; i++ {
		p = p.neighbour(s.orientation)
	}
	return p
}

/* ======================================================================
   EVALUATION FUNCTION
   ====================================================================== */

func (s *State) evaluateState(myId int) int {
	score := 0

	// Find my ship
	var me Ship
	for _, sh := range s.ships {
		if sh.id == myId {
			me = sh
		}
	}

	if me.isDead {
		return -999999
	}

	score += me.health * 5

	// Enemy health → we want it low
	for _, e := range s.enemyShips {
		score += (100 - e.health) * 4
		if e.health <= 0 {
			score += 2000
		}
	}

	enemy := s.getClosestEnemyShip(me)
	dist := me.pos.distanceTo(enemy.pos)
	score += int((15 - dist) * 20)

	// Barrels proximity
	for _, b := range s.barrels {
		score += int((20 - me.pos.distanceTo(b.pos)) * 3)
	}

	return score
}

/* ======================================================================
   GENERATE CANDIDATE ACTIONS
   ====================================================================== */

func generateActions(s *State, ship Ship) []string {
	actions := []string{
		"WAIT",
		"FASTER",
		"SLOWER",
		"PORT",
		"STARBOARD",
	}

	// Move to closest barrel
	if len(s.barrels) > 0 {
		b := s.getClosestBarrel(ship)
		actions = append(actions, fmt.Sprintf("MOVE %d %d", b.pos.x, b.pos.y))
	}

	// Firing options (predictive)
	enemy := s.getClosestEnemyShip(ship)

	for t := 1; t <= 3; t++ {
		pred := enemy.nextPosShip(t)
		if ship.bow().distanceTo(pred) <= FIRE_DISTANCE_MAX {
			actions = append(actions,
				fmt.Sprintf("FIRE %d %d", pred.x, pred.y))
		}
	}

	return actions
}

/* ======================================================================
   APPLY STRING ACTION TO SHIP (in simulated state)
   ====================================================================== */

func applyActionFromString(ship *Ship, action string) {
	if action == "WAIT" {
		ship.wait()
		return
	}
	if action == "FASTER" {
		ship.faster()
		return
	}
	if action == "SLOWER" {
		ship.slower()
		return
	}
	if action == "PORT" {
		ship.port()
		return
	}
	if action == "STARBOARD" {
		ship.starboard()
		return
	}

	if len(action) > 4 && action[:4] == "MOVE" {
		var x, y int
		fmt.Sscanf(action, "MOVE %d %d", &x, &y)
		ship.move(Point{x, y})
		return
	}

	if len(action) > 4 && action[:4] == "FIRE" {
		var x, y int
		fmt.Sscanf(action, "FIRE %d %d", &x, &y)
		ship.fire(Point{x, y})
		return
	}
}

/* ======================================================================
   FIND SHIP POINTER BY ID IN SIMULATED FUTURE
   ====================================================================== */

func findShipById(s *State, id int) *Ship {
	for i := range s.ships {
		if s.ships[i].id == id {
			return &s.ships[i]
		}
	}
	return nil
}

/* ======================================================================
   AI DEPTH 2 (SIMULATION-BASED)
   ====================================================================== */

func (s *State) think() {

	for i := range s.allyShips {

		myShip := s.allyShips[i]
		myId := myShip.id

		// Generate actions for this ship
		A0 := generateActions(s, myShip)

		bestScore := -99999999
		bestAction := "WAIT"

		// ------------------------------
		// DEPTH = 2
		// ------------------------------
		for _, a0 := range A0 {

			// FUTURE AFTER ACTION 0
			f1 := s.Copy()

			// reset actions
			for si := range f1.ships {
				f1.ships[si].actionType = ""
			}

			fs1 := findShipById(&f1, myId)
			if fs1 == nil {
				continue
			}

			applyActionFromString(fs1, a0)
			simulateTurn(&f1)

			// Now second layer of actions
			A1 := generateActions(&f1, *fs1)

			for _, a1 := range A1 {

				// FUTURE AFTER ACTION 1
				f2 := f1.Copy()

				// reset actions before second simulation
				for si := range f2.ships {
					f2.ships[si].actionType = ""
				}

				fs2 := findShipById(&f2, myId)
				if fs2 == nil {
					continue
				}

				applyActionFromString(fs2, a1)
				simulateTurn(&f2)

				score := f2.evaluateState(myId)

				if score > bestScore {
					bestScore = score
					bestAction = a0 // Only first action is executed in real world
				}
			}
		}

		// PRINT DECISION FOR THIS SHIP
		fmt.Println(bestAction)
	}
}

/* ======================================================================
   PARSE INPUT
   ====================================================================== */

func (s *State) readEntities() {
	s.enemyShips = []Ship{}
	s.allyShips = []Ship{}
	s.ships = []Ship{}
	s.barrels = []Barrel{}
	s.mines = []Mine{}
	s.cannonballs = []cannonball{}
	s.cannonballsExplosions = []Point{}

	fmt.Scan(&s.myShipCount)
	fmt.Scan(&s.entityCount)

	for i := 0; i < s.entityCount; i++ {
		var id, x, y, a1, a2, a3, a4 int
		var t string
		fmt.Scan(&id, &t, &x, &y, &a1, &a2, &a3, &a4)

		switch t {

		case "SHIP":
			ship := Ship{
				Entity:      Entity{id, t, Point{x, y}},
				orientation: a1,
				speed:       a2,
				health:      a3,
				owner:       a4,
			}

			if a4 == 1 {
				s.allyShips = append(s.allyShips, ship)
				s.ships = append(s.ships, ship)
			} else {
				s.enemyShips = append(s.enemyShips, ship)
				s.ships = append(s.ships, ship)
			}

		case "BARREL":
			s.barrels = append(s.barrels, Barrel{
				Entity: Entity{id, t, Point{x, y}},
				health: a1,
			})

		case "MINE":
			s.mines = append(s.mines, Mine{
				Entity: Entity{id, t, Point{x, y}},
			})

		case "CANNONBALL":
			s.cannonballs = append(s.cannonballs, cannonball{
				target:         Point{x, y},
				fromShip:       a1,
				remainingTurns: a2,
			})
		}
	}
}

/* ======================================================================
   MAIN LOOP
   ====================================================================== */

func main() {
	state := State{}

	for {
		state.readEntities()

		// Compute actions with depth‑2 search
		state.think()
	}
}
