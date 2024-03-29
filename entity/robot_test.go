package entity_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"
	"github.com/stretchr/testify/assert"
)

func TestNewRobot(t *testing.T) {
	t.Run("should return not Nil wen call", func(t *testing.T) {
		xAxis, yAxis := 1,1
		direction := "N"
		cordinate := event.NewCordinate(xAxis, yAxis)
		facing := event.NewFacing(direction)

		robot := entity.NewRobot(cordinate, facing)

		assert.NotNil(t, robot)
	})
}