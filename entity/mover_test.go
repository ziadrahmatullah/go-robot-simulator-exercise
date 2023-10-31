package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"
)

func TestMove(t *testing.T) {
	t.Run("should return different robot when call Move", func(t *testing.T) {
		xAxis, yAxis, face := 1, 1, "N"
		cordinate := event.NewCordinate(xAxis, yAxis)
		facing := event.NewFacing(face)
		robot := entity.NewRobot(cordinate, facing)
		mover := entity.MoveAdvance{}

		result := mover.Move(robot)

		assert.NotEqual(t, robot, result)
	})

}
