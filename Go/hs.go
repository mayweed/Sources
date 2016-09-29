package main

import (
	"fmt"
	"log"
	"strings"
)

type position struct {
	x, y int
}

func NewPosition(x, y int) position {
	return position{
		x: x,
		y: y,
	}
}

type commande struct {
	command string
	pos     position
}

//should be intialized differently for a bomb or a player in NewEntity func
type entity struct {
	entityType, owner, x, y, param1, param2 int
}

func NewEntity(entityType, owner, x, y, param1, param2 int) entity {
	return entity{
		entityType: entityType,
		owner:      owner,
		x:          x,
		y:          y,
		param1:     param1,
		param2:     param2,
	}
}

type myBot struct {
	//my position should i use anon struct position here?
	x int
	y int
	//my num of bombs left
	myBombs int
}

func main() {
	var width, height, myId int
	fmt.Scan(&width, &height, &myId)

	for {
		var grid = make([][]string, height)
		for i := 0; i < height; i++ {
			var row string
			fmt.Scan(&row)
			inputs := strings.Split(row, "")
			grid[i] = make([]string, width)
			for j := 0; j < width; j++ {
				grid[i][j] = inputs[j]
			}
		}
		//log.Println(grid)

		var entities int
		fmt.Scan(&entities)
		//log.Println(entities)

		var players []entity
		var bombs []entity
		for i := 0; i < entities; i++ {
			var entityType, owner, x, y, param1, param2 int
			fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)
			if entityType == 0 {
				players = append(players, NewEntity(entityType, owner, x, y, param1, param2))
			} else if entityType == 1 {
				bombs = append(bombs, NewEntity(entityType, owner, x, y, param1, param2))
			}

		}
		log.Println(players)
		log.Println(myId)

		var me myBot
		for _, p := range players {
			if p.owner == myId {
				me = myBot{x: p.x, y: p.y, myBombs: p.param1}
			}
		}
		log.Println(me.x)

		//MOVE test
		var xx = me.x + 1
		var yy = me.y + 1
		if grid[xx][me.y] == "." {
			fmt.Println("MOVE", xx, me.y)
		} else if grid[xx][me.y] == "1" {
			if grid[xx][yy] == "." {
				fmt.Println("MOVE", xx, yy)
			}
		}
		//fmt.Fprintln(os.Stderr,entities)
		//fmt.Println("MOVE 11 13")// Write action to stdout
	}
}
