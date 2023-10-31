package main

import (
	"bufio"
	"fmt"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/util"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	for !exit{
		fmt.Println("===================================")
		initPos := promptInput(scanner, "Initial position: ")
		movements := promptInput(scanner, "Movements: ")
		ok, x, y, face := util.IsInitialPositionValid(initPos)
		if !ok{
			fmt.Println("Invalid position!")
		}else if !util.IsMovesValid(movements){
			fmt.Println("Invalid movement!")
		}else{
			robotCOntroller := entity.NewRobotController(x, y, face)
			newX, newY, newFace := robotCOntroller.MovementsCommand(movements)
			fmt.Printf("Robot last position %d %d %s\n", newX, newY, newFace)
		}
	}
}