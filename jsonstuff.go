package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type worldMap struct {
	Description []string  `json:"worldMapDescription"`
	Zone        [][][]int `json:"worldMapZone"`
	LiveZone    [][][]int `json:"worldMapLiveZone"`
}

func (w worldMap) saveMap() {
	fmt.Println(w)
	convert := &w
	output, err := json.Marshal(convert)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, _ := json.Marshal(output)
	_ = ioutil.WriteFile("worldmapsavedata.json", file, 0755)
}

func (w worldMap) loadMap() {
	outputmap := worldMap{}
	jsonFile, _ := ioutil.ReadFile("worldmapsavedata.json") ///////////// BIT THAT DOESN'T WORK /////////////
	_ = json.Unmarshal([]byte(jsonFile), &outputmap)
	fmt.Println(outputmap)
}

func main() {
	testmap := worldMap{}
	xlength := []int{0, 0, 0}
	ylength := [][]int{}
	var x = 0
	for x = 0; x < 4; x++ {
		ylength = append(ylength, xlength)
	}
	testmap.Description = append(testmap.Description, "It's a test map")
	testmap.Zone = append(testmap.Zone, ylength)
	testmap.LiveZone = append(testmap.LiveZone, ylength)
	testmap.saveMap()
	testmap.loadMap()
}
