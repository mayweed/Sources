package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//goal: unmarshal json
func main() {
	/* comment else it awaits stdin...
	//first you read the config thing
	//on stdin, not sure it will work
	//should check how it's given...
	var configString string
	fmt.Scan(&configString)

	var config GameConfig
	err := json.Unmarshal([]byte(configString), &config)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(config.UnitsInformation[FILTER])
	/*
		while{
	*/
	//gameTurn.go...
	// $ go run main.go gameTurn.go gameConfig.go gameEvents.go
	//json: cannot unmarshal number into Go value of type []interface {}
	var f Frame
	var testStringF = []byte(`{"p2Units":[[],[],[],[],[],[],[]],"turnInfo":[1,0,0],"p1Stats":[30.0,1.0,5.0,26],"p1Units":[[[8,11,60.0,"2"],[9,11,60.0,"4"],[7,10,60.0,"6"],[7,9,60.0,"8"],[7,8,60.0,"10"],[8,7,60.0,"12"],[9,7,60.0,"14"],[17,11,60.0,"16"],[18,11,60.0,"18"],[18,10,60.0,"20"],[18,9,60.0,"22"],[18,8,60.0,"24"],[17,7,60.0,"26"],[18,7,60.0,"28"],[19,7,60.0,"30"]],[],[[11,7,75.0,"32"],[13,9,75.0,"34"],[15,11,75.0,"36"]],[],[],[],[]],"p2Stats":[-1.0,25.0,5.0,3],"events":{"selfDestruct":[],"breach":[[[2,11],1,3,"55",2],[[2,11],1,3,"56",2],[[2,11],1,3,"57",2],[[2,11],1,3,"58",2],[[2,11],1,3,"59",2],[[2,11],1,3,"60",2],[[2,11],1,3,"61",2],[[2,11],1,3,"62",2]],"damage":[],"shield":[],"move":[],"spawn":[[[8,11],0,"1",1],[[9,11],0,"3",1],[[7,10],0,"5",1],[[7,9],0,"7",1],[[7,8],0,"9",1],[[8,7],0,"11",1],[[9,7],0,"13",1],[[17,11],0,"15",1],[[18,11],0,"17",1],[[18,10],0,"19",1],[[18,9],0,"21",1],[[18,8],0,"23",1],[[17,7],0,"25",1],[[18,7],0,"27",1],[[19,7],0,"29",1],[[11,7],2,"31",1],[[13,9],2,"33",1],[[15,11],2,"35",1]],"death":[[[2,11],3,"55",2,false],[[2,11],3,"56",2,false],[[2,11],3,"57",2,false],[[2,11],3,"58",2,false],[[2,11],3,"59",2,false],[[2,11],3,"60",2,false],[[2,11],3,"61",2,false],[[2,11],3,"62",2,false]],"attack":[],"melee":[]}}`)
	//	var testStringF = []byte(`{"turnInfo":[1,0,0],"p1stats":[30.0,1.0,5.0,26],"p2Stats":[-1.0,25.0,5.0,3],"p1Units":[[[8,11,60.0,"2"],[9,11,60.0,"4"],[7,10,60.0,"6"],[7,9,60.0,"8"],[7,8,60.0,"10"],[8,7,60.0,"12"],[9,7,60.0,"14"],[17,11,60.0,"16"],[18,11,60.0,"18"],[18,10,60.0,"20"],[18,9,60.0,"22"],[18,8,60.0,"24"],[17,7,60.0,"26"],[18,7,60.0,"28"],[19,7,60.0,"30"]],[],[[11,7,75.0,"32"],[13,9,75.0,"34"],[15,11,75.0,"36"]],[],[],[],[]]}`)
	err2 := json.Unmarshal(testStringF, &f)
	if err2 != nil {
		fmt.Println(err2)
	}
	log.Println(json.Valid(testStringF)) //yields true ...
	//fmt.Println(f.P1U[FILTER][0].X)
	//fmt.Println(f.getPlayerUnits(1, "filter"))
	//fmt.Println(f.P1U)
	fmt.Println("SPAWN:", f.Evts.Spawn)
	fmt.Println("BREACH:", f.Evts.Breach)
}
