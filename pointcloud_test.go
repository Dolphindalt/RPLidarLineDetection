package main

import (
	"math"
	"os"
	"testing"
)

func TestMinMax(t *testing.T) {
	pc := NewPointCloudFromFile("testdata.txt")
	min, max := pc.minMaxPoints()
	t.Logf("Min: %v Max: %v\n", min, max)
	if math.Floor(min.X()) != -1243 && math.Floor(min.Y()) != 344 {
		t.Fatalf("Min point computed incorrectly\n")
	}
	if math.Floor(max.X()) != -768 && math.Floor(max.Y()) != 1843 {
		t.Fatalf("Max point computed incorrectly\n")
	}
}

func TestImage(t *testing.T) {
	pc := NewPointCloudFromFile("testdata.txt")
	pc.toImageSpace()
	min, _ := pc.minMaxPoints()
	if min.X() < 0 || min.Y() < 0 {
		t.Fatalf("Function toImageSpace translated a point cloud into a negative quadrant")
	}
	pc.saveAsImage("test.png", 2, 0.5)
	os.Remove("test.png")
}
