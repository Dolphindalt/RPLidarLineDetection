package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// PointCloud represents
type PointCloud struct {
	shift  *Vector2f
	points []*Vector2f
}

// toImageSpace translates the point cloud to image space
func (p *PointCloud) toImageSpace() {
	var p1, p2, newshift *Vector2f
	p.minMaxPoints(p1, p2)
	newshift = ScalarQuotient(VectorAdd(p1, p2), 2.0)
	for i := 0; i < len(p.points); i++ {
		p.points[i] = VectorSubtract(p.points[i], newshift)
	}
	p.shift = VectorAdd(p.shift, newshift)
}

// meanValue is also center of gravity
func (p *PointCloud) MeanValue() *Vector2f {
	var v *Vector2f
	for i := 0; i < len(p.points); i++ {
		v = VectorAdd(v, p.points[i])
	}
	if len(p.points) > 0 {
		return ScalarQuotient(v, float64(len(p.points)))
	} else {
		return v
	}
}

func (p *PointCloud) minMaxPoints(minPoint *Vector2f, maxPoint *Vector2f) {
	if len(p.points) > 0 {
		minPoint = p.points[0]
		maxPoint = p.points[1]

		for i := 0; i < len(p.points); i++ {
			cv := p.points[i]
			if minPoint.x > cv.x {
				minPoint.x = cv.x
			}
			if minPoint.y > cv.y {
				minPoint.y = cv.y
			}
			if maxPoint.x < cv.x {
				maxPoint.x = cv.x
			}
			if maxPoint.y < cv.y {
				maxPoint.y = cv.y
			}
		}
	} else {
		minPoint = NewVector2f(0, 0)
		maxPoint = NewVector2f(0, 0)
	}
}

// ReadFromFile reads a pointcloud from a gorplidar scan
func (p *PointCloud) ReadFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to read point cloud from file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		if len(splitLine) != 3 {
			continue // skip other stuff that may be in the file
		} // TODO implement a quality check here
		angle, _ := strconv.ParseFloat(splitLine[1], 64)
		distance, _ := strconv.ParseFloat(splitLine[2], 64)
		p.points = append(p.points, NewVector2f(math.Cos(angle)*distance, math.Sin(angle)*distance))
	}
}

// PointsCloseToLine stores points dx close to line (a, b) in y
func (p *PointCloud) PointsCloseToLine(a *Vector2f, b *Vector2f, dx float64, y *PointCloud) {
	y.points = y.points[:0]
	for i := 0; i < len(p.points); i++ {
		t := ScalarProduct(b, VectorSubtract(p.points[i], a))
		d := VectorSubtract(p.points[i], VectorAdd(a, ScalarMultiplication(b, t)))
		if d.Magnitude() <= dx {
			y.points = append(y.points, p.points[i])
		}
	}
}

// RemovePoints removes points in y from the point cloud
// the order of the points must be the same both clouds
func (p *PointCloud) RemovePoints(y *PointCloud) {
	if len(y.points) == 0 {
		return
	}
	var newPoints []*Vector2f
	i := 0
	j := 0
	// the assumption of order is made here
	for ; i < len(p.points) && j < len(y.points); i++ {
		if p.points[i] == y.points[j] {
			j++
		} else {
			newPoints = append(newPoints, p.points[i])
		}
	}
	// copy other points after y is traversed
	for i = 0; i < len(p.points); i++ {
		newPoints = append(newPoints, p.points[i])
	}

	p.points = newPoints
}
