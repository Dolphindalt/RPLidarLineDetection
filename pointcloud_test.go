package main

import (
	"os"
	"testing"
)

func TestMinMax(t *testing.T) {
	pc := NewPointCloudFromFile("testdata.txt")
	min, max := pc.minMaxPoints()
	t.Logf("Min: %v Max: %v\n", min, max)
	if min > max {
		t.Fatalf("Min is greater than max")
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
