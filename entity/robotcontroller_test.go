package entity_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewRobotController(t *testing.T) {
	t.Run("should return not Nil when call", func(t *testing.T) {
		xAxis, yAxis, face := 1,1, "N"

		robotController := entity.NewRobotController(xAxis,yAxis,face)

		assert.NotNil(t, robotController)
	})
}

func TestMovementsCommand(t *testing.T) {
	t.Run("should return 10 9 N when input RARARA", func(t *testing.T) {
		xAxis, yAxis, face := 11,9, "E"
		robotController := entity.NewRobotController(xAxis,yAxis,face)
		expected1, expected2, expected3:= 10, 9, "N"

		newXAxis, newYAxis, newFace := robotController.MovementsCommand("RARARA")

		assert.Equal(t, expected1, newXAxis)
		assert.Equal(t, expected2, newYAxis)
		assert.Equal(t, expected3, newFace)
	})

	t.Run("should return 9 4 W when input RAALAL", func(t *testing.T) {
		xAxis, yAxis, face := 7,3, "N"
		robotController := entity.NewRobotController(xAxis,yAxis,face)
		expected1, expected2, expected3:= 9, 4, "W"

		newXAxis, newYAxis, newFace := robotController.MovementsCommand("RAALAL")

		assert.Equal(t, expected1, newXAxis)
		assert.Equal(t, expected2, newYAxis)
		assert.Equal(t, expected3, newFace)
	})

	t.Run("should return 2 1 E when input LA", func(t *testing.T) {
		xAxis, yAxis, face := 1,1, "S"
		robotController := entity.NewRobotController(xAxis,yAxis,face)
		expected1, expected2, expected3:= 2, 1, "E"

		newXAxis, newYAxis, newFace := robotController.MovementsCommand("LA")

		assert.Equal(t, expected1, newXAxis)
		assert.Equal(t, expected2, newYAxis)
		assert.Equal(t, expected3, newFace)
	})

	t.Run("should return 8 0 S when input RR", func(t *testing.T) {
		xAxis, yAxis, face := 8,0, "N"
		robotController := entity.NewRobotController(xAxis,yAxis,face)
		expected1, expected2, expected3:= 8, 0, "S"

		newXAxis, newYAxis, newFace := robotController.MovementsCommand("RR")

		assert.Equal(t, expected1, newXAxis)
		assert.Equal(t, expected2, newYAxis)
		assert.Equal(t, expected3, newFace)
	})
}