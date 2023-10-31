package entity_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"
	"github.com/stretchr/testify/assert"
)

func TestFaceRight(t *testing.T) {
	t.Run("should diferent facing when inputed to the FaceRight", func(t *testing.T) {
		facing := &event.Facing{}
		facer := entity.FaceRight{}

		result := facer.Face(facing)

		assert.NotEqual(t, facing, result)
	})
}

func TestLeft(t *testing.T) {
	t.Run("should diferent facing when inputed to the FaceLeft", func(t *testing.T) {
		facing := &event.Facing{}
		facer := entity.FaceLeft{}

		result := facer.Face(facing)

		assert.NotEqual(t, facing, result)
	})
}