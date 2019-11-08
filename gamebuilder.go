package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type objectStorage struct {
	objectID          []int
	objectname        []string
	objectdescription []string
	objecthealth      []int
	objectattack      []int
}

type worldMap struct {
	description []string
	zone        [][][]int
}

func (world_map *worldMap) buildMap() {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in name of map:")
	Scanner.Scan()
	mapname := Scanner.Text()
	fmt.Println("Type in width of map:")
	Scanner.Scan()
	yindex1 := Scanner.Text()
	yindex, _ := strconv.Atoi(yindex1)
	fmt.Println("Type in height of map:")
	Scanner.Scan()
	xindex1 := Scanner.Text()
	xindex, _ := strconv.Atoi(xindex1)
	yi := 0
	slice := [][]int{}
	for yi = 0; yi < yindex; yi++ {
		nest := []int{}
		xi := 0
		for xi = 0; xi < xindex; xi++ {
			nest = append(nest, 0)
		}
		slice = append(slice, nest)
	}
	world_map.buildMap2(slice, mapname)
}

func (x *worldMap) buildMap2(slice [][]int, mapname string) {
	world_map := *x
	world_map.zone = append(world_map.zone, slice)
	world_map.description = append(world_map.description, mapname)
	*x = world_map
}

func (x worldMap) printMap() {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in map index:")
	Scanner.Scan()
	y1 := Scanner.Text()
	y, _ := strconv.Atoi(y1)
	for i := range x.zone {
		if i == y {
			for o := range x.zone[i] {
				fmt.Println(x.description[i], x.zone[i][o])
			}
		}
	}
}

func (x worldMap) printZone(y int) {
	for i := range x.zone {
		if i == y {
			for o := range x.zone[i] {
				fmt.Println(x.description[i], x.zone[i][o])
			}
		}
	}
}

func (x worldMap) fullMap() {
	for i := range x.description {
		fmt.Println(x.description[i])
	}
}

func (x *objectStorage) createObject() {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in name of object:")
	Scanner.Scan()
	object_name := Scanner.Text()
	fmt.Println("Type in description:")
	Scanner.Scan()
	object_description := Scanner.Text()
	fmt.Println("Type in health:")
	Scanner.Scan()
	object_health1 := Scanner.Text()
	object_health, _ := strconv.Atoi(object_health1)
	fmt.Println("Type in attack:")
	Scanner.Scan()
	object_attack1 := Scanner.Text()
	object_attack, _ := strconv.Atoi(object_attack1)
	i := *x
	i.objectname = append(i.objectname, object_name)
	i.objectdescription = append(i.objectdescription, object_description)
	i.objecthealth = append(i.objecthealth, object_health)
	i.objectattack = append(i.objectattack, object_attack)
	*x = i
}

func (x objectStorage) printObject() {
	Input := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in object name:")
	Input.Scan()
	result2 := Input.Text()
	for i := range x.objectname {
		if x.objectname[i] == result2 {
			fmt.Printf("name: %s\n", x.objectname[i])
			fmt.Printf("description: %s\n", x.objectdescription[i])
			fmt.Printf("health: %d\n", x.objecthealth[i])
			fmt.Printf("attack: %d\n", x.objectattack[i])
		}
	}
}

func (x objectStorage) grabObject(name string) (string, string, int, int) {
	for i := range x.objectname {
		if x.objectname[i] == name {
			return x.objectname[i], x.objectdescription[i], x.objecthealth[i], x.objectattack[i]
		}
	}
	return "", "", 0, 0
}

func (w *worldMap) placeObject(y objectStorage) {
	world_map := *w
	Input := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in object name:")
	Input.Scan()
	object_name := Input.Text()
	fmt.Println("Type in map index:")
	Input.Scan()
	map_index1 := Input.Text()
	map_index, _ := strconv.Atoi(map_index1)
	fmt.Println("Type in x coord:")
	Input.Scan()
	xcoord1 := Input.Text()
	xcoord, _ := strconv.Atoi(xcoord1)
	fmt.Println("Type in y coord:")
	Input.Scan()
	ycoord1 := Input.Text()
	ycoord, _ := strconv.Atoi(ycoord1)
	for i := range world_map.zone {
		if i == map_index {
			for i2 := range world_map.zone[i] {
				if i2 == ycoord {
					for i3 := range world_map.zone[i][i2] {
						if i3 == xcoord {
							for objectindex := range y.objectname {
								if y.objectname[objectindex] == object_name {
									fmt.Println(world_map.zone[i][i2][i3])
									fmt.Println(objectindex)
									fmt.Println(y.objectname)
									world_map.zone[i][i2][i3] = objectindex + 1
								}
							}
						}
					}
				}
			}
		}

	}
	*w = world_map
}

func (w *worldMap) moveHero(cmd string) {
	world_map := *w
	for i := range world_map.zone {
		for a := range world_map.zone[i] {
			for b := range world_map.zone[i][a] {
				if world_map.zone[i][a][b] == 1 {
					fmt.Println("Hero spotted!")
					switch cmd {
					case "w":
						if a == 0 {
							world_map.printZone(i)
							return
						}
						if a != 0 {
							world_map.zone[i][a][b] = 0
							world_map.zone[i][a-1][b] = 1
							world_map.printZone(i)
							*w = world_map
							return
						}
					case "s":
						if a != len(world_map.zone[i])-1 {
							world_map.zone[i][a][b] = 0
							world_map.zone[i][a+1][b] = 1
							world_map.printZone(i)
							*w = world_map
							return
						}
						if a == len(world_map.zone[i])-1 {
							world_map.printZone(i)
							return
						}
					case "a":
						if b == 0 {
							if i == 0 {
								world_map.printZone(i)
								return
							}
							world_map.zone[i][a][b] = 0
							world_map.zone[i-1][len(world_map.zone[i-1])-1][len(world_map.zone[i-1][0])-1] = 1
							world_map.printZone(i - 1)
							*w = world_map
							return
						}
						if b != 0 {
							world_map.zone[i][a][b] = 0
							world_map.zone[i][a][b-1] = 1
							world_map.printZone(i)
							*w = world_map
							return
						}
					case "d":
						if b == len(world_map.zone[i][a])-1 {
							if i == len(world_map.zone)-1 {
								world_map.printZone(i)
								return
							}
							world_map.zone[i][a][b] = 0
							world_map.zone[i+1][len(world_map.zone[i+1])-1][0] = 1
							world_map.printZone(i + 1)
							*w = world_map
							return
						}
						if b != len(world_map.zone[i][a])-1 {
							world_map.zone[i][a][b] = 0
							world_map.zone[i][a][b+1] = 1
							world_map.printZone(i)
							*w = world_map
							return
						}
					}

				}

			}
		}

	}
}

func main() {
	//initial things necessary for game to work
	object := objectStorage{}
	gamemap := worldMap{}
	Input := bufio.NewScanner(os.Stdin)
	//game running
	gameover := 0
	for gameover == 0 {
		fmt.Println("===GAME BUILDER===")
		fmt.Println("For help/commands simply enter 'help'")
		fmt.Println("q to quit")
		fmt.Println("Input here:")
		Input.Scan()
		result := Input.Text()
		switch result {
		case "help":
			fmt.Println("buildobject: create an object (name, description, attack, health)")
			fmt.Println("viewobject: allows you to view object (type in name)")
			fmt.Println("placeobject: place object on the map (type in co-ordinates)")
			fmt.Println("buildmap: allows you to create an X by X map by an index")
			fmt.Println("fullmap: prints out all maps by index.")
			fmt.Println("printmap: prints out map by index. First map would be 0, second 1 etc.")
			fmt.Println("play: initiates the game")
			fmt.Println("q: exit the game\n")
		case "buildobject":
			object.createObject()
		case "viewobject":
			object.printObject()
		case "placeobject":
			gamemap.placeObject(object)
		case "buildmap":
			gamemap.buildMap()
		case "printmap":
			gamemap.printMap()
		case "fullmap":
			gamemap.fullMap()
		case "play":
			fmt.Println("Playing Game!")
			fmt.Println("w s a d to move around. q to quit game")
			playgame := 0
			for playgame == 0 {
				fmt.Println("Input here:")
				Input.Scan()
				result2 := Input.Text()
				switch result2 {
				case "q":
					fmt.Println("Quitting game...")
					playgame = 1
				case "w":
					gamemap.moveHero(result2)
				case "s":
					gamemap.moveHero(result2)
				case "a":
					gamemap.moveHero(result2)
				case "d":
					gamemap.moveHero(result2)
				}
			}
		case "save":

		case "q":
			gameover = 1
		}
	}
}
