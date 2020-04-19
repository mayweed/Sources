package main

import (
	"encoding/json"
)

// and what if i pass Point as a pointer?? will it unmarshal correclty??
//can i unmarshall my point and the the rest with another unmarshaler?? How??
/*
type SelfDestructEvent struct {
	Coords   Point
	//gosh it's an array of arrays!!
	//you just cant hardcode that...
	Targets  [][]Point
	Damage   float64
	UnitType int
	Id       string
	Owner    int
}

func (s *SelfDestructEvent) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&s.Coords, &s.Targets, &s.Damage, &s.UnitType, &s.Id, &s.Owner}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	return nil
}
*/
type BreachEvent struct {
	Point
	PlayerHealthDamage float64
	BreachedUnitType   int
	BreachedUnitId     string
	Owner              int
}

func (b *BreachEvent) UnmarshalJSON(buf []byte) error {
	var tmp []interface{} //{&b.Coords, &b.PlayerHealthDamage, &b.BreachedUnitType, &b.BreachedUnitId, &b.Owner}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	//should i unmarshal tmp[0] in point struct via unmarshaler interface??
	//was not able to implement a correct umarshaler for Point in array of array...
	//:'(
	//this one is crazy...just failed at writing a custom unmarshaler..
	b.Point.X = int(tmp[0].([]interface{})[0].(float64))
	b.Point.Y = int(tmp[0].([]interface{})[1].(float64))
	b.PlayerHealthDamage = tmp[1].(float64)
	b.BreachedUnitType = int(tmp[2].(float64))
	b.BreachedUnitId = tmp[3].(string)
	b.Owner = int(tmp[4].(float64))

	return nil
}

type DamageEvent struct {
	Point
	DamageSuffered  float64
	DamagedUnitType int
	DamagedUnitId   string
	Owner           int
}

func (d *DamageEvent) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&d.Coords, &d.DamageSuffered, &d.DamagedUnitType, &d.DamagedUnitId, &d.Owner}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	d.Point.X = int(tmp[0].([]interface{})[0].(float64))
	d.Point.Y = int(tmp[0].([]interface{})[1].(float64))
	d.DamageSuffered = tmp[1].(float64)
	d.DamageUnitType = int(tmp[2].(float64))
	d.DamageUnitId = tmp[3].(string)
	d.Owner = int(tmp[4].(float64))

	return nil
}

/*
type ShieldEvent struct {
	//GivingShield   []Point
	//GivenShield    []Point
	ShieldAmount   float64
	ShieldUnitType int
	ShieldUnitId   string
	Owner          int
}

func (s *ShieldEvent) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&s.GivingShield, &s.GivenShield, &s.ShieldAmount, &s.ShieldUnitType, &s.ShieldUnitId, &s.Owner}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	return nil
}

type MoveEvent struct {
	//From          []Point
	//To            []Point
	//Deprecated    []Point
	MovedUnitType int
	MovedUnitId   string
	Owner         int
}

func (m *MoveEvent) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&m.From, &m.To, &m.Deprecated, &m.MovedUnitType, &m.MovedUnitId, &m.Owner}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	return nil
}
*/
type Events struct {
	//SelfDestruct []SelfDestructEvent `json:"selfDestruct"`
	Breach []BreachEvent `json:"breach"`
	Damage []DamageEvent `json:"damage"`
	//Shield       []ShieldEvent       `json:"shield"`
	//Move         []MoveEvent         `json:"move"`
	Spawn []SpawnEvent `json:"spawn"`
	//Death        []DeathEvent        `json:"death"`
	//Attack       []AttackEvent       `json:"attack"`
	//Melee        []MeleeEvent        `json:"melee"`
}

type SpawnEvent struct {
	Point
	SpawnedUnitType int
	SpawnedUnitId   string
	Owner           int
}

func (s *SpawnEvent) UnmarshalJSON(buf []byte) error {
	var tmp []interface{}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	//should i unmarshal tmp[0] in point struct via unmarshaler interface??
	//this one is crazy...might i read reflect doc??
	s.Point.X = int(tmp[0].([]interface{})[0].(float64))
	s.Point.Y = int(tmp[0].([]interface{})[1].(float64))

	s.SpawnedUnitType = int(tmp[1].(float64))
	s.SpawnedUnitId = tmp[2].(string)
	s.Owner = int(tmp[3].(float64))
	return nil
}
