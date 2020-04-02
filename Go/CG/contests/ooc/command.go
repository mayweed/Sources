package main

import (
	"fmt"
	"strings"
)

//COMMANDS
func (m *Me) isTorpCharge() bool {
	var c bool
	if m.torpedoCooldown > 0 {
		c = true
		m.canFireTorpedo = false
	} else {
		c = false
		m.canFireTorpedo = true
	}
	return c
}
func (m *Me) isSonarCharge() bool {
	var c bool
	if m.sonarCooldown > 0 {
		c = true
	} else {
		c = false
		m.canUseSonar = true
	}
	return c
}
func (m *Me) isSilenceCharge() bool {
	var c bool
	if m.silenceCooldown > 0 {
		c = true
	} else {
		c = false
		m.canUseSilence = true
	}
	return c
}
func (m *Me) move(dir string) {
	var command string
	switch dir {
	case "N":
		command = fmt.Sprintf("MOVE %s", dir)
	case "S":
		command = fmt.Sprintf("MOVE %s", dir)
	case "W":
		command = fmt.Sprintf("MOVE %s", dir)
	case "E":
		command = fmt.Sprintf("MOVE %s", dir)
	}
	//having torp charge is important but sonar?
	//no dynamic, be charged that's all
	if m.isTorpCharge() {
		command = command + " TORPEDO"
	}
	if !m.isTorpCharge() && m.isSonarCharge() {
		command = command + " SONAR"
	}
	if !m.isTorpCharge() && !m.isSonarCharge() && m.isSilenceCharge() {
		command = command + " SILENCE"
	}
	m.commands = append(m.commands, command)
}
func (m *Me) surface() {
	m.commands = append(m.commands, "SURFACE")
}
func (m *Me) sonar(sector int) {
	command := fmt.Sprintf("SONAR %d", sector)
	m.commands = append(m.commands, command)
}
func (m *Me) torpedo(tile Tile) {
	command := fmt.Sprintf("TORPEDO %d %d", tile.pos.x, tile.pos.y)
	m.commands = append(m.commands, command)
}
func (m *Me) silence(direction string, distance int) {
	command := fmt.Sprintf("SILENCE %s %d", direction, distance)
	m.commands = append(m.commands, command)
}
func (m *Me) msg(s string) {
	command := fmt.Sprintf("MSG %s", s)
	m.commands = append(m.commands, command)
}
func (m *Me) sendTurn() {
	if len(m.commands) == 1 {
		fmt.Print(m.commands[0])
	} else {
		fmt.Print(strings.Join(m.commands, "|"))
	}
	fmt.Println()
}
