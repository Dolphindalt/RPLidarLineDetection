package main

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Dolphindalt/gorplidar"
	"github.com/fogleman/gg"
	"github.com/go-gl/mathgl/mgl64"
)

const (
	scanCycles int = 3
)

// PointCloud stores a translation of, and all the of, the points
type PointCloud struct {
	shift  mgl64.Vec2
	points []mgl64.Vec2
}

// NewPointCloudFromLidar constructs a PointCloud by performing a lidar scan
// PRECONDITION: Lidar connected and motor spinning
func NewPointCloudFromLidar(lidar *gorplidar.RPLidar) (*PointCloud, error) {
	if lidar.Connected == false || lidar.MotorActive == false {
		return nil, errors.New("Tried to construct PointCloud, but lidar not connected or spinning")
	}
	scanResults, err := lidar.StartScan(scanCycles)
	if err != nil {
		return nil, err
	}
	cloud := &PointCloud{mgl64.Vec2{0, 0}, []mgl64.Vec2{}}
	for _, p := range scanResults {
		cloud.points = append(cloud.points, mgl64.Vec2{float64(p.X), float64(p.Y)})
	}
	return cloud, nil
}

// NewPointCloudFromFile constructs a PointCloud from a given file
// This function is designed for testing and debugging
func NewPointCloudFromFile(fileName string) *PointCloud {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to read PointCloud from file: %v\n", err)
	}
	defer file.Close()
	cloud := &PointCloud{mgl64.Vec2{0, 0}, []mgl64.Vec2{}}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		if len(splitLine) != 3 {
			continue // to skip junk that may exist in the file
		}
		angle, _ := strconv.ParseFloat(splitLine[1], 64)
		distance, _ := strconv.ParseFloat(splitLine[2], 64)
		cloud.points = append(cloud.points, mgl64.Vec2{math.Cos(angle*(math.Pi/180.0)) * distance, math.Sin(angle*(math.Pi/180.0)) * distance})
	}
	return cloud
}

// toImageSpace translates the point cloud so that the center of the
// image is relative to the center of the first quardrant.
func (p *PointCloud) toImageSpace() {
	min, _ := p.minMaxPoints()
	newShift := min
	for i := 0; i < len(p.points); i++ {
		p.points[i] = p.points[i].Sub(min)
	}
	p.shift = p.shift.Add(newShift)
}

// minMaxPoints returns two Vec2 representing the max and min values of
// each axis of the point cloud.
func (p *PointCloud) minMaxPoints() (mgl64.Vec2, mgl64.Vec2) {
	if len(p.points) > 0 {
		min := p.points[0]
		max := p.points[1]
		for i := 0; i < len(p.points); i++ {
			cmp := p.points[i]
			if min[0] > cmp[0] {
				min[0] = cmp[0]
			}
			if min[1] > cmp[1] {
				min[1] = cmp[1]
			}
			if max[0] < cmp[0] {
				max[0] = cmp[0]
			}
			if max[1] < cmp[1] {
				max[1] = cmp[1]
			}
		}
		return min, max
	}
	return mgl64.Vec2{0, 0}, mgl64.Vec2{0, 0}
}

func (p *PointCloud) saveAsImage(nameExt string, pointScale float64, scaleFactor float64) {
	_, max := p.minMaxPoints()
	dc := gg.NewContext(int(max.X()*scaleFactor), int(max.Y()*scaleFactor))
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.Fill()
	dc.SetRGB(1, 1, 1)
	for i := 0; i < len(p.points); i++ {
		l := p.points[i]
		dc.DrawPoint(l.X()*scaleFactor, l.Y()*scaleFactor, 1*pointScale)
	}
	dc.Fill()
	dc.SavePNG(nameExt)
}
