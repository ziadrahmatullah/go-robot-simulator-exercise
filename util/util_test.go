package util_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/util"
	"github.com/stretchr/testify/assert"
)

func TestIsInitialPositionValid(t *testing.T) {
	t.Run("should return as expected when call", func(t *testing.T) {
		initPos := "1 2 N"
		expected1, expected2, expected3, expected4 := true, 1, 2, "N"

		eq, x, y, face := util.IsInitialPositionValid(initPos)

		assert.Equal(t, expected1, eq)
		assert.Equal(t, expected2, x)
		assert.Equal(t, expected3, y)
		assert.Equal(t, expected4, face)
	})
}

func TestIsMoveValid(t *testing.T) {
	t.Run("should return true when input RRL", func(t *testing.T) {
		input := "RRL"
		expected := true

		result := util.IsMovesValid(input)

		assert.Equal(t, expected, result)
	})
}

func TestIsMoveCommand(t *testing.T) {
	t.Run("should return true when input A", func(t *testing.T) {
		input := "A"
		expected := true

		result := util.IsMoveCommand(input)

		assert.Equal(t, expected, result)
	})
}

func TestIsFacingCommand(t *testing.T) {
	t.Run("should return true when input R", func(t *testing.T) {
		input := "R"
		expected := true

		result := util.IsFacingCommand(input)

		assert.Equal(t, expected, result)
	})
}

