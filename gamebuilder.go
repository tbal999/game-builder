package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type objectStorage struct {
	objectID          []int    `json:"objectID"`
	objectname        []string `json:"objectname"`
	objectdescription []string `json:"objectdescription"`
	objecthealth      []int    `json:"objecthealth"`
	objectattack      []int    `json:"objectattack"`
}

type worldMap struct {
	description []string  `json:"worldMapdescription"`
	zone        [][][]int `json:"worldMapzone"`
	livezone    [][][]int `json:"worldMaplivezone"`
}

func (w *worldMap) interaction(z, y, x int, o *objectStorage) {
	world_map := *w
	object_storage := *o
	checker := 0
	_, _, _, hero_attack := object_storage.grabObject(1)
	object_name, object_description, object_health, object_attack := object_storage.grabObject(world_map.livezone[z][y][x])
	if object_name == "" {
		return
	}
	fmt.Println("You have discovered a " + object_name)
	fmt.Println(object_description)
	chance := randomNumber(1, 10)
	if object_attack == 0 {
		fmt.Println("It looks harmless")
	}
	if object_attack != 0 {
		fmt.Println("It is aggressive, and looks like it wants a fight")
		if hero_attack > object_attack {
			if chance > 6 {
				fmt.Println("It realises you are quite strong and backs away")
				return
			}
			fmt.Println("It leaps at you!")
			object_storage.editObject(0, object_attack, 0)
			fmt.Println("You took " + strconv.Itoa(object_attack) + " damage")
			fmt.Println("It realises that you are quite strong and runs away!")
			world_map.livezone[z][y][x] = 0
			*w = world_map
			*o = object_storage
			return
		}
		for checker == 0 {
			fmt.Println("It leaps at you!")
			object_storage.editObject(0, object_attack, 0)
			fmt.Println("You took " + strconv.Itoa(object_attack) + " damage")
			object_health = object_health - hero_attack
			fmt.Println("The " + object_name + " took " + strconv.Itoa(object_attack) + " damage")
			if object_health <= 0 {
				fmt.Println("The " + object_name + " died")
				world_map.livezone[z][y][x] = 0
				checker = 1
			}
			_, _, hhealth, _ := object_storage.grabObject(1)
			if hhealth <= 0 {
				fmt.Println("You died")
				checker = 1
			}
		}
		*w = world_map
		*o = object_storage
		return
	}

	*w = world_map
	*o = object_storage
}

func (x *objectStorage) editObject(index, dmg, attack int) {
	object_storage := *x
	for i := range object_storage.objectname {
		if i == index {
			object_storage.objecthealth[i] = object_storage.objecthealth[i] - dmg
			object_storage.objectattack[i] = object_storage.objectattack[i] + attack
		}
	}
	*x = object_storage
}

func randomNumber(min, max int) int {
	z := rand.Intn(max)
	if z < min {
		z = min
	}
	return z
}

func (x *objectStorage) grabObject(index int) (string, string, int, int) {
	object_storage := *x
	for i := range object_storage.objectname {
		if i+1 == index {
			return object_storage.objectname[i], object_storage.objectdescription[i], object_storage.objecthealth[i], object_storage.objectattack[i]
		}
	}
	return "", "", 0, 0
}

func (x *objectStorage) allObject() {
	object_storage := *x
	for i := range object_storage.objectname {
		fmt.Println(i, object_storage.objectname[i])
	}
}

func savegame(o objectStorage, w worldMap) { //either JSON or CSV export
	fmt.Println("Not yet implemented")

}

func (world_map *worldMap) buildMap() {
	fmt.Println("")
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
	twice := 0
	yi := 0
	for twice = 0; twice < 2; twice++ {
		slice := [][]int{}
		for yi = 0; yi < yindex; yi++ {
			nest := []int{}
			xi := 0
			for xi = 0; xi < xindex; xi++ {
				nest = append(nest, 0)
			}
			slice = append(slice, nest)
		}
		world_map.buildMap2(slice, mapname, twice)
	}
}

func (x *worldMap) buildMap2(slice [][]int, mapname string, twice int) {
	world_map := *x
	switch twice {
	case 0:
		world_map.zone = append(world_map.zone, slice)
		world_map.description = append(world_map.description, mapname)
	case 1:
		world_map.livezone = append(world_map.livezone, slice)
	}
	*x = world_map
}

func (x worldMap) printMap() {
	fmt.Println("")
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in map index:")
	Scanner.Scan()
	y1 := Scanner.Text()
	y, _ := strconv.Atoi(y1)
	fmt.Println("")
	for i := range x.zone {
		if i == y {
			for o := range x.zone[i] {
				fmt.Println(x.description[i], x.zone[i][o])
			}
		}
	}
	for i := range x.livezone {
		if i == y {
			for o := range x.livezone[i] {
				fmt.Println(x.description[i], x.livezone[i][o])
			}
		}
	}
}

func (x worldMap) printZone(y int) {
	fmt.Println("Your location")
	for i := range x.zone {
		if i == y {
			for o := range x.zone[i] {
				fmt.Println(i, x.zone[i][o])
			}
		}
	}
	fmt.Println("")
	fmt.Println("Objects")
	for i := range x.livezone {
		if i == y {
			for o := range x.livezone[i] {
				fmt.Println(i, x.livezone[i][o])
			}
		}
	}
}

func (x worldMap) fullMap() {
	fmt.Println("")
	fmt.Println("Index, Description.")
	for i := range x.description {
		fmt.Println(i, x.description[i])
	}
}

func (x *objectStorage) createObject() {
	fmt.Println("")
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
	fmt.Println("")
	Input := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in name of object you are looking for:")
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

func (w *worldMap) placeObject(y objectStorage) {
	fmt.Println("")
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
									if objectindex == 0 {
										world_map.zone[i][i2][i3] = objectindex + 1
									}
									if objectindex != 0 {
										world_map.livezone[i][i2][i3] = objectindex + 1
									}
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

func (w *worldMap) moveHero(cmd string, o *objectStorage) {
	world_map := *w
	for i := range world_map.zone {
		for a := range world_map.zone[i] {
			for b := range world_map.zone[i][a] {
				if world_map.zone[i][a][b] == 1 {
					switch cmd {
					case "w":
						if a == 0 {
							world_map.printZone(i)
							return
						}
						if a != 0 {
							world_map.zone[i][a][b] = 0
							world_map.zone[i][a-1][b] = 1 // OBJECT INTERACTION
							world_map.printZone(i)
							w.interaction(i, a-1, b, o)
							*w = world_map
							return
						}
					case "s":
						if a != len(world_map.zone[i])-1 {
							world_map.zone[i][a][b] = 0
							world_map.zone[i][a+1][b] = 1 // OBJECT INTERACTION
							world_map.printZone(i)
							w.interaction(i, a+1, b, o)
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
							world_map.zone[i-1][len(world_map.zone[i-1])-1][len(world_map.zone[i-1][0])-1] = 1 // OBJECT INTERACTION
							world_map.printZone(i - 1)
							w.interaction(i-1, len(world_map.zone[i-1])-1, len(world_map.zone[i-1][0])-1, o)
							*w = world_map
							return
						}
						if b != 0 {
							world_map.zone[i][a][b] = 0
							world_map.zone[i][a][b-1] = 1 // OBJECT INTERACTION
							w.interaction(i, a, b-1, o)
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
							world_map.zone[i+1][len(world_map.zone[i+1])-1][0] = 1 // OBJECT INTERACTION
							world_map.printZone(i + 1)
							w.interaction(i+1, len(world_map.zone[i+1])-1, 0, o)
							*w = world_map
							return
						}
						if b != len(world_map.zone[i][a])-1 {
							world_map.zone[i][a][b] = 0
							world_map.zone[i][a][b+1] = 1 // OBJECT INTERACTION
							world_map.printZone(i)
							w.interaction(i, a, b+1, o)
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
	fmt.Println("===GAME BUILDER===")
	fmt.Println("For help simply type 'help'")
	fmt.Println("Press 'q' to quit")
	fmt.Println("The first object you build must be your hero character")
	for gameover == 0 {
		fmt.Println("===GAME BUILDER===")
		fmt.Println("Input here:")
		Input.Scan()
		result := Input.Text()
		switch result {
		case "help":
			fmt.Println("buildobject: create an object. FIRST OBJECT YOU CREATE IS HERO.")
			fmt.Println("allobject: view all objects by name and index")
			fmt.Println("viewobject: allows you to view object (type in name)")
			fmt.Println("placeobject: place object on the map (type in co-ordinates)")
			fmt.Println("buildmap: allows you to create an X by X map by an index")
			fmt.Println("viewworld: prints out all maps by index.")
			fmt.Println("viewmap: prints out map by index. First map would be 0, second 1 etc.")
			fmt.Println("play: initiates the game")
			fmt.Println("q: exit the game\n")
		case "buildobject":
			object.createObject()
		case "allobject":
			object.allObject()
		case "viewobject":
			object.printObject()
		case "placeobject":
			gamemap.placeObject(object)
		case "buildmap":
			gamemap.buildMap()
		case "viewmap":
			gamemap.printMap()
		case "viewworld":
			gamemap.fullMap()
		case "play":
			playgame := 0
			if len(gamemap.zone) == 0 {
				fmt.Println("No maps! Quitting instance...")
				playgame = 1
				break
			}
			gamemap.printZone(0)
			fmt.Println("Loading instance...")
			fmt.Println("w s a d to move around. p for hero stats. q to quit game")
			for playgame == 0 {
				if object.objecthealth[0] <= 0 {
					playgame = 1
					fmt.Println("Quitting instance...")
					break
				}
				fmt.Println("Input here: [w,s,a,d to move] [p for hero stats]:")
				Input.Scan()
				command := Input.Text()
				switch command {
				case "q":
					fmt.Println("Quitting instance...")
					playgame = 1
				case "w":
					gamemap.moveHero(command, &object)
				case "s":
					gamemap.moveHero(command, &object)
				case "a":
					gamemap.moveHero(command, &object)
				case "d":
					gamemap.moveHero(command, &object)
				case "p":
					fmt.Println("Name: " + object.objectname[0])
					fmt.Println("Description: " + object.objectdescription[0])
					fmt.Println("Attack: " + strconv.Itoa(object.objectattack[0]))
					fmt.Println("Health: " + strconv.Itoa(object.objecthealth[0]))
				}
			}
		case "save":
			savegame(object, gamemap)
		case "q":
			gameover = 1
		}
	}
}
