package main

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

const (
	scaleFactor = 0.5
	pointScale  = 2
)

func main() {
	pc := NewPointCloudFromFile("points.txt")
	pc.toImageSpace()
	pc.saveAsImage("pointcloud.png", pointScale, scaleFactor)
	img := gocv.IMRead("pointcloud.png", gocv.IMReadGrayScale)
	window := gocv.NewWindow("Hough Transform")
	defer window.Close()
	cimg := gocv.NewMat()
	defer cimg.Close()
	gocv.CvtColor(img, &cimg, gocv.ColorGrayToBGR)
	red := color.RGBA{255, 0, 0, 1}
	lines := ExtractEndpoints("pointcloud.png")
	for _, l := range lines {
		gocv.Line(&cimg, image.Pt(int(l.P1.X()), int(l.P1.Y())), image.Pt(int(l.P2.X()), int(l.P2.Y())), red, 3)
	}
	lines = TranslateLines(lines, scaleFactor, pc.shift)
	for _, l := range lines {
		fmt.Printf("(%v,%v) (%v,%v)\n", (l.P1.X()/0.5)+pc.shift.X(), (l.P1.Y()/0.5)+pc.shift.Y(), (l.P2.X()/0.5)+pc.shift.X(), (l.P2.Y()/0.5)+pc.shift.Y())
	}
	gocv.IMWrite("lines.png", cimg)
}
