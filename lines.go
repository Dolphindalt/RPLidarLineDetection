package main

// cv stands for computer vision

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
	"gocv.io/x/gocv"
)

const (
	ksize      = 5
	rho        = 1
	theta      = math.Pi / 180
	threshold  = 8
	minLineLen = 30
	maxLineGap = 50
)

// Line stores the end points of a line
type Line struct {
	P1 mgl64.Vec2
	P2 mgl64.Vec2
}

// ExtractEndpoints takes and image and detects lines in it
func ExtractEndpoints(imageName string) []Line {
	lines := []Line{}
	img := gocv.IMRead(imageName, gocv.IMReadGrayScale)
	gocv.MedianBlur(img, &img, ksize)
	houghLines := gocv.NewMat()
	gocv.HoughLinesPWithParams(img, &houghLines, rho, theta, threshold, minLineLen, maxLineGap)
	for i := 0; i < houghLines.Rows(); i += 4 {
		x0 := float64(houghLines.GetIntAt(0, i))
		y0 := float64(houghLines.GetIntAt(0, i+1))
		x1 := float64(houghLines.GetIntAt(0, i+2))
		y1 := float64(houghLines.GetIntAt(0, i+3))
		lines = append(lines, Line{mgl64.Vec2{x0, y0}, mgl64.Vec2{x1, y1}})
	}
	return lines
}

// TranslateLines takes the line array and disregards image scaling and shifting
// This is done to make the end points relative to the lidar at proper scale
func TranslateLines(lines []Line, scaleFactor float64, shift mgl64.Vec2) []Line {
	for _, l := range lines {
		l.P1 = mgl64.Vec2{(l.P1.X() / scaleFactor) + shift.X(), (l.P1.Y() / scaleFactor) + shift.Y()}
		l.P2 = mgl64.Vec2{(l.P2.X() / scaleFactor) + shift.X(), (l.P2.Y() / scaleFactor) + shift.Y()}
	}
	return lines
}
