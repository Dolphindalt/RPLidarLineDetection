package main

import (
	"image"
	"image/color"
	"os"
	"testing"

	"gocv.io/x/gocv"
)

func TestExtraction(t *testing.T) {
	pc := NewPointCloudFromFile("testdata.txt")
	pc.toImageSpace()
	pc.saveAsImage("test.png", 2, 0.5)
	lines := ExtractEndpoints("test.png")
	img := gocv.IMRead("test.png", gocv.IMReadGrayScale)
	cimg := gocv.NewMat()
	defer cimg.Close()
	defer img.Close()
	gocv.CvtColor(img, &cimg, gocv.ColorGrayToBGR)
	red := color.RGBA{255, 0, 0, 1}
	for _, l := range lines {
		gocv.Line(&cimg, image.Pt(int(l.P1.X()), int(l.P1.Y())), image.Pt(int(l.P2.X()), int(l.P2.Y())), red, 2)
	}
	lines = TranslateLines(lines, 0.5, pc.shift)
	gocv.IMWrite("lines_test.png", cimg)
	os.Remove("test.png")
}
