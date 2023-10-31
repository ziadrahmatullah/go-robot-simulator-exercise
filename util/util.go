package util

import (
	"regexp"
	"strconv"
	"strings"
)

func IsMoveCommand(cmd string) bool{
	return cmd == "A"
}

func IsFacingCommand(cmd string) bool{
	return cmd == "R" || cmd == "L"
}

func IsMovesValid(moves string) bool {
	valid := regexp.MustCompile(`^[RLA]+$`)
	return valid.MatchString(moves)
}

func IsInitialPositionValid(initPos string)(eq bool, x, y int, face string){
	init := strings.Split(initPos, " ")
	if len(init) != 3{
		return 
	}
	x, err := strconv.Atoi(init[0])
	y, err2 := strconv.Atoi(init[1])
	if err!= nil || err2 != nil{
		return
	}
	if init[2] == "N" || init[2] == "E" || init[2] == "S" || init[2] == "W"{
		return true, x, y, init[2]
	}
	return
}