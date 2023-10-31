package event_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"
	"github.com/stretchr/testify/assert"
)

func TestNewCordinate(t *testing.T) {
	t.Run("should return not nil when call", func(t *testing.T) {
		xAxis, yAxis := 1,1
		
		result := event.NewCordinate(xAxis, yAxis)

		assert.NotNil(t,result)
	})
}

func TestXAxis(t *testing.T) {
	t.Run("should return not nil when call", func(t *testing.T) {
		xAxis, yAxis := 1,1
		cordinate := event.NewCordinate(xAxis, yAxis)
		
		result := cordinate.XAxis()

		assert.NotNil(t,result)
	})
}

func TestYAxis(t *testing.T) {
	t.Run("should return not nil when call", func(t *testing.T) {
		xAxis, yAxis := 1,1
		cordinate := event.NewCordinate(xAxis, yAxis)
		
		result := cordinate.YAxis()

		assert.NotNil(t,result)
	})
}

func TestXCordinatePlus(t *testing.T) {
	t.Run("should return 2 when first cordinate 1,1", func(t *testing.T) {
		xAxis, yAxis := 1,1
		cordinate := event.NewCordinate(xAxis, yAxis)
		expected:= 2
		
		cordinate.XCordinatePlus()
		result := cordinate.XAxis()

		assert.Equal(t, expected, result)
	})
}

func TestYCordinatePlus(t *testing.T) {
	t.Run("should return 2 when first cordinate 1,1", func(t *testing.T) {
		xAxis, yAxis := 1,1
		cordinate := event.NewCordinate(xAxis, yAxis)
		expected:= 2
		
		cordinate.YCordinatePlus()
		result := cordinate.YAxis()

		assert.Equal(t, expected, result)
	})
}

func TestXCordinateMinus(t *testing.T) {
	t.Run("should return 0 when first cordinate 1,1", func(t *testing.T) {
		xAxis, yAxis := 1,1
		cordinate := event.NewCordinate(xAxis, yAxis)
		expected:= 0
		
		cordinate.XCordinateMinus()
		result := cordinate.XAxis()

		assert.Equal(t, expected, result)
	})
}

func TestYCordinateMinus(t *testing.T) {
	t.Run("should return 0 when first cordinate 1,1", func(t *testing.T) {
		xAxis, yAxis := 1,1
		cordinate := event.NewCordinate(xAxis, yAxis)
		expected:= 0
		
		cordinate.YCordinateMinus()
		result := cordinate.YAxis()

		assert.Equal(t, expected, result)
	})
}
