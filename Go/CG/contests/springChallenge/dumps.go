package main

import (
	"bytes"
	"strconv"
)

//of no need in main now

//no pointer: Grid has a string meth, not *Grid !!
func (g *Grid) String() string {
	var buf bytes.Buffer
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			if g.c[x][y].what == -1 {
				buf.WriteString("#")
			} else if g.c[x][y].what != 0 {
				buf.WriteString(strconv.Itoa(g.c[x][y].what))

			} else {
				buf.WriteString(" ")
			}
		}
		buf.WriteString("\n")
	}
	//fmt.Println()
	return buf.String()
}
