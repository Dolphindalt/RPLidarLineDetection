package main

import (
	"math"
)

const (
	maxTheta         int     = 180
	thetaStep        float64 = math.Pi / float64(maxTheta)
	neighborhoodSize int     = 16
)

type Hough struct {
	width        int
	height       int
	accumulator  [][]int
	doubleHeight int
	houghHeight  int
	centerX      float64
	centerY      float64
}

type HoughLine struct {
	theta            float64
	accumulatorTheta int
	p                int
	x                int
	y                int
}

type Line struct {
	x1 int16
	x2 int16
	y1 int16
	y2 int16
}

//=======================================
/*
Point Stuff.
Run the toImageSpace functions to convert to useable values.
Run the toPolarSpace to put it back into Lidar info.
*/
type PolarPoint struct {
	r     float64
	theta float64
}

func CartesianPointFrom(pp PolarPoint) Point {
	return Point{int(pp.r * math.Cos(pp.theta)), int(pp.r * math.Cos(pp.theta))}
}

type Point struct {
	x int
	y int
}

func (p *Point) toImageSpace(offset int) {
	p.x += offset
	p.y += offset
}

func (p *Point) toDiscreteSpace(offset int) {
	p.x -= offset
	p.y -= offset
}

func arrayToImageSpace(points []Point, offset int) {
	for i := 0; i < len(points); i++ {
		points[i].toImageSpace(offset)
	}
}

func arrayToDiscreteSpace(points []Point, offset int) {
	for i := 0; i < len(points); i++ {
		points[i].toDiscreteSpace(offset)
	}
}

//========================================

var sinCache = make([]float64, maxTheta)
var cosCache = make([]float64, maxTheta)

func HoughInit() {
	for i := 0; i < maxTheta; i++ {
		thetaReal := float64(i) * thetaStep
		sinCache[i] = math.Sin(thetaReal)
		cosCache[i] = math.Cos(thetaReal)
	}
}

func NewHoughTransform(width int, height int, points []Point) *Hough {
	centerX := float64(width) / 2
	centerY := float64(height) / 2
	houghHeight := (int)(math.Sqrt2 * math.Max(float64(height), float64(width)))
	doubleHeight := houghHeight * 2

	accumulator := make([][]int, maxTheta)
	for i := 0; i < maxTheta; i++ {
		accumulator[i] = make([]int, doubleHeight)
	}

	// Fills the accumulator array of the Hough.
	for _, p := range points {
		for j := 0; j < maxTheta; j++ {
			h := int(((float64(p.x) - centerX) * cosCache[j]) + ((float64(p.y) - centerY) * sinCache[j]))
			h += houghHeight
			if h < 0 || h >= houghHeight {
				continue
			}
			accumulator[j][h]++
		}
	}

	return &Hough{width, height, accumulator, doubleHeight, houghHeight, centerX, centerY}
}

func (h *Hough) InfiniteLines(threshold int) []HoughLine {
	var lines []Hough

	for r := 0; r < maxTheta; r++ {
		for p := neighborhoodSize; p < h.doubleHeight; p++ {
		OuterLoop:
			if h.accumulator[r][p] > threshold {
				peak := h.accumulator[r][p]
				for dx := -neighborhoodSize; dx <= neighborhoodSize; dx++ {
					for dy := -neighborhoodSize; dy <= neighborhoodSize; dy++ {
						dr := r + dx
						dp := p + dy
						if dr < 0 {
							dr += maxTheta
						} else if dt >= maxTheta {
							dr -= maxTheta
						}
						if h.accumulator[dr][dp] > peak {
							p++
							break OuterLoop
						}
					}
				}
				/*
										double theta = ((double)r) * theta_step;
					                	if(list->max_length == list->current_length)
					                {
					                    list->max_length += 20;
					                    list->lines = realloc(list->lines, sizeof(line_t) * list->max_length);
					                }
					                	line_t line;
					               		line.accumulator_theta = r;
					                	line.theta = theta;
					                	line.r = p;
										list->lines[list->current_length++] = line;
				*/
			}

		}
	}
	return list
}

/*
func (h *Hough) LineSegments(inifinteLines []HoughLine) []Line {

}
*/
