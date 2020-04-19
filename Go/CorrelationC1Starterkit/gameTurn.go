package main

import (
	"encoding/json"
)

//Units are given in a certain order, always the same
const (
	//firewall types
	FILTER = iota
	ENCRYPTOR
	DESTRUCTOR
	//attack units
	PING
	EMP
	SCRAMBLER
	REMOVE
)

type Point struct {
	X int
	Y int
}

type TurnInfo struct {
	TurnType    int
	TurnNum     int
	ActionPhase int
}

func (t *TurnInfo) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&t.TurnType, &t.TurnNum, &t.ActionPhase}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	return nil
}

type PlayerStats struct {
	Health float64
	Cores  float64
	Bits   float64
	Time   int
}

func (p *PlayerStats) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&p.Health, &p.Cores, &p.Bits, &p.Time}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	return nil
}

//it's a list of list!!Attention!!
type UnitStats struct {
	X          int
	Y          int
	UnitHealth float64
	Id         string
}

func (u *UnitStats) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&u.X, &u.Y, &u.UnitHealth, &u.Id}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	return nil
}

type Frame struct {
	T    TurnInfo      `json:"turnInfo"`
	P1   PlayerStats   `json:"p1stats"`
	P2   PlayerStats   `json:"p2stats"`
	P1U  [][]UnitStats `json:"p1Units"`
	P2U  [][]UnitStats `json:"p2Units"`
	Evts Events        `json:"events"`
	//End Stats??
}

func (f *Frame) getPlayerUnits(id int, t string) []UnitStats {
	switch id {
	case 1:
		switch t {
		case "filter":
			return f.P1U[FILTER]
		case "encryptor":
			return f.P1U[ENCRYPTOR]
		case "ping":
			return f.P1U[PING]
		case "emp":
			return f.P1U[EMP]
		case "scrambler":
			return f.P1U[SCRAMBLER]
		case "remove":
			return f.P1U[REMOVE]

		}
	case 2:
		switch t {
		case "filter":
			return f.P2U[FILTER]
		case "encryptor":
			return f.P2U[ENCRYPTOR]
		case "ping":
			return f.P2U[PING]
		case "emp":
			return f.P2U[EMP]
		case "scrambler":
			return f.P2U[SCRAMBLER]
		case "remove":
			return f.P2U[REMOVE]

		}

	}
	return []UnitStats{}
}
