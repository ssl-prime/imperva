package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Travel struct {
	M_X             int    // upper limit of x co-ordinate
	N_Y             int    // upper limit of y co-ordinate
	StartCordinateX int    // initial x co-ordinate
	StartCordinateY int    // initial y co-ordinate
	Direction       string // initial direction
	Instruction     string // set of movement instruction
}

func main() {
	RobotTravel()
}

// RobotTravel ...
func RobotTravel() {
	directionMap := getDirectionMap()
	travelSet := getTravelSet()
	for _, currentSet := range travelSet {
		xCoordinateBoundry := currentSet.M_X
		yCoordinateBoundry := currentSet.N_Y
		startX := currentSet.StartCordinateX
		startY := currentSet.StartCordinateY
		direction := currentSet.Direction
		instruction := strings.Split(currentSet.Instruction, "")

		if direction != "E" && direction != "W" && direction != "N" && direction != "S" {
			fmt.Println("invalid direction given")
			continue
		}
		if startX < 0 || startX > xCoordinateBoundry {
			fmt.Println("invalid start coordinate given")
			continue
		}
		if startY < 0 || startY > yCoordinateBoundry {
			fmt.Println("invalid start coordinate given")
			continue
		}

		msg := ""
		visitedMap := make(map[string]bool)
		visitedMap[getKay(startX, startY)] = true

		for move := 0; move < len(instruction); move++ {
			direction, startX, startY, msg = movement(direction, instruction[move], startX, startY, directionMap,
				visitedMap, xCoordinateBoundry, yCoordinateBoundry, msg)
			if msg != "" {
				break
			}
		}
		fmt.Println(startX, startY, direction)
	}

}

// movement after every instruction
func movement(direction string, instruction string, startX, startY int,
	directionMap map[string]string, visitedMap map[string]bool, xCoordinateBoundry int,
	yCoordinateBoundry int, msg string) (string, int, int, string) {
	switch {
	case instruction == "M" && direction == "W":
		key := getKay(startX-1, startY)
		if msg := checkStopingCondition(xCoordinateBoundry, startX-1, key, visitedMap); msg != "" {
			return direction, startX, startY, msg
		}
		visitedMap[key] = true

		startX--
	case instruction == "M" && direction == "E":
		key := getKay(startX+1, startY)
		if msg := checkStopingCondition(xCoordinateBoundry, startX+1, key, visitedMap); msg != "" {
			return direction, startX, startY, msg
		}
		visitedMap[key] = true

		startX++
	case instruction == "M" && direction == "N":
		key := getKay(startX, startY+1)
		if msg := checkStopingCondition(xCoordinateBoundry, startY+1, key, visitedMap); msg != "" {
			return direction, startX, startY, msg
		}
		visitedMap[key] = true

		startY++
	case instruction == "M" && direction == "S":
		key := getKay(startX, startY-1)
		if msg := checkStopingCondition(xCoordinateBoundry, startY-1, key, visitedMap); msg != "" {
			return direction, startX, startY, msg
		}
		visitedMap[key] = true

		startY--
	case instruction == "L" || instruction == "R":
		direction = directionMap[direction+"-"+instruction]
	}
	return direction, startX, startY, ""
}

func checkStopingCondition(upperLimit int, coordinate int, key string, visitedMap map[string]bool) string {
	if coordinate < 0 {
		return "out of boundry : robot coordinate < 0"
	}
	if coordinate > upperLimit {
		return "out of boundry : robot coordinate > upperlimit"
	}

	if _, isExist := visitedMap[key]; isExist {
		return "already visited"
	}
	return ""
}

func getTravelSet() []Travel {
	TravelSet := make([]Travel, 0)
	TravelSet = append(TravelSet, Travel{
		M_X:             4,
		N_Y:             4,
		StartCordinateX: 0,
		StartCordinateY: 0,
		Direction:       "N",
		Instruction:     "MMMRMMLM",
	},
		Travel{
			M_X:             4,
			N_Y:             4,
			StartCordinateX: 0,
			StartCordinateY: 0,
			Direction:       "E",
			Instruction:     "MMMRMMLMMMMMM",
		},
		Travel{
			M_X:             4,
			N_Y:             4,
			StartCordinateX: 0,
			StartCordinateY: 0,
			Direction:       "S",
			Instruction:     "MMMRMMLMMMMMM",
		},
		Travel{
			M_X:             4,
			N_Y:             4,
			StartCordinateX: 1,
			StartCordinateY: 3,
			Direction:       "W",
			Instruction:     "MMMRMMLMMMMMM",
		},
		Travel{
			M_X:             20,
			N_Y:             20,
			StartCordinateX: 10,
			StartCordinateY: 10,
			Direction:       "W",
			Instruction:     "MMMRMMLMMMMMMMMMRMMLMMMMMM",
		},
		Travel{
			M_X:             20,
			N_Y:             20,
			StartCordinateX: 10,
			StartCordinateY: 10,
			Direction:       "N",
			Instruction:     "MLMLMLML",
		},
		Travel{
			M_X:             20,
			N_Y:             20,
			StartCordinateX: 10,
			StartCordinateY: 10,
			Direction:       "N",
			Instruction:     "MMMMMMMMMMMMMMMMMMMMMMMMMM",
		},
		Travel{
			M_X:             20,
			N_Y:             20,
			StartCordinateX: 10,
			StartCordinateY: 100,
			Direction:       "N",
			Instruction:     "MMMMMMMMMMMMMMMMMMMMMMMMMM",
		},
		Travel{
			M_X:             20,
			N_Y:             20,
			StartCordinateX: 10,
			StartCordinateY: 10,
			Direction:       "X",
			Instruction:     "MMMMMMMMMMMMMMMMMMMMMMMMMM",
		},
	)
	return TravelSet
}

// getDirectionMap set of allowed diection and their final direction after turning left or right
func getDirectionMap() map[string]string {
	directionMap := make(map[string]string)
	directionMap["N-L"] = "W"
	directionMap["S-R"] = "W"
	directionMap["N-R"] = "E"
	directionMap["S-L"] = "E"
	directionMap["E-L"] = "N"
	directionMap["W-R"] = "N"
	directionMap["E-R"] = "S"
	directionMap["W-L"] = "S"

	return directionMap
}

// getKay this will key return for visitedmap tracking
func getKay(startX, startY int) string {
	return "startX-" + strconv.Itoa(startX) + ":" + "startY-" + strconv.Itoa(startY)
}

// output for given set
/*
	2 4 N
	3 0 S
	0 0 S
	0 3 W
	0 12 W
	9 10 E
	10 20 N
	invalid start coordinate given
	invalid direction given
*/
