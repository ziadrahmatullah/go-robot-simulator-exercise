package event_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"
	"github.com/stretchr/testify/assert"
)

func TestNewFacing(t *testing.T) {
	t.Run("should return not nil when call", func(t *testing.T) {
		direction := "N"
		
		result := event.NewFacing(direction)

		assert.NotNil(t,result)
	})
}

func TestDirection(t *testing.T) {
	t.Run("should return not nil when call", func(t *testing.T) {
		facing := event.NewFacing("N")
		
		result := facing.Direction()

		assert.NotNil(t,result)
	})
}

func TestPlusDirection(t *testing.T) {
	t.Run("should return 2 when first direction 'N'", func(t *testing.T) {
		facing := event.NewFacing("N")
		facing.PlusDirection()
		expected := 2

		result := facing.Direction()

		assert.Equal(t, expected, result)
	})
}

func TestMinusDirection(t *testing.T) {
	t.Run("should return 1 when first direction 'E'", func(t *testing.T) {
		facing := event.NewFacing("E")
		facing.MinusDirection()
		expected := 1

		result := facing.Direction()

		assert.Equal(t, expected, result)
	})
}

func TestFirstDirection(t *testing.T) {
	t.Run("should return 1 when call function Direction", func(t *testing.T) {
		facing := event.NewFacing("E")
		facing.FirstDirection()
		expected := 1

		result := facing.Direction()

		assert.Equal(t, expected, result)
	})
}

func TestLastDirection(t *testing.T) {
	t.Run("should return 1 when call function Direction", func(t *testing.T) {
		facing := event.NewFacing("E")
		facing.LastDirection()
		expected := 4

		result := facing.Direction()

		assert.Equal(t, expected, result)
	})
}