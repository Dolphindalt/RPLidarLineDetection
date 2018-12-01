package main

import (
	"math"
)

// Hough the hough transform representation
type Hough struct {
	votingSpace []int
	circle      *Circle
	numB        int
	dx          float64
	maxX        float64
	numX        int
}

// NewHough asks for a description of the parameter space and allocates the accumulator
func NewHough(minP *Vector2f, maxP *Vector2f, vardx float64, circleGranularity int) *Hough {
	var circle *Circle
	circle = &Circle{[]*Vector2f{}, []int{}}
	circle.FromIcosahedron(circleGranularity)
	numB := len(circle.vertices)

	maxX := math.Max(maxP.Magnitude(), minP.Magnitude())
	rangeX := 2 * maxX
	dx := vardx
	if dx == 0.0 {
		dx = rangeX / 64.0
	}
	numX := int(roundNearest(rangeX / dx))

	votingSpace := make([]int, numX*numX*numB)

	return &Hough{votingSpace, circle, numB, dx, maxX, numX}
}

// getLine returns a line that was voted on the most
// a is the point, b is direction
func (h *Hough) getLine(a *Vector2f, b *Vector2f) (int, *Vector2f, *Vector2f) {
	votes := 0
	index := 0
	for i := 0; i < len(h.votingSpace); i++ {
		if h.votingSpace[i] > votes {
			votes = h.votingSpace[i]
			index = i
		}
	}

	x := int(index / (h.numX * h.numB))
	index -= int(x * h.numX * h.numB)
	xr := float64(x)*h.dx - h.maxX

	y := int(index / h.numB)
	index -= int(y * h.numB)
	yr := float64(y)*h.dx - h.maxX

	b = h.circle.vertices[index]
	// equation 3 follows, with the 3rd dimension ignored
	a.x = xr*(1-(b.x*b.x)) - yr*(b.x*b.y)
	a.y = xr*(-(b.x * b.y)) + yr*(1-(b.y*b.y))

	return votes, a, b
}

// add adds all points from the cloud into voting space
func (h *Hough) add(p *PointCloud) {
	for j := 0; j < len(p.points); j++ {
		h.pointVote(p.points[j], true)
	}
}

// subtract subtracts all points from the cloud from the voting space
func (h *Hough) subtract(p *PointCloud) {
	for j := 0; j < len(p.points); j++ {
		h.pointVote(p.points[j], false)
	}
}

// pointVote adds or subtracts one point from the voting space
func (h *Hough) pointVote(point *Vector2f, add bool) {
	for j := 0; j < len(h.circle.vertices); j++ {
		b := h.circle.vertices[j]
		// the denominator in equation 2 is just zero due to no z plane
		// x' left side of equation 2
		xNew := ((1 - (b.x * b.x)) * point.x) - ((b.x * b.y) * point.y) - (b.x)
		// y' left side of equation 2
		yNew := ((b.x * b.y) * point.x) + ((1 - (b.y * b.y)) * point.y) - (b.y)

		xi := int(roundNearest((xNew + h.maxX) / h.dx))
		yi := int(roundNearest((yNew + h.maxX) / h.dx))

		// a one dimensional voting space based on tree tables
		// xi * planes * direction + yi * direction + loop
		index := (xi * h.numX * h.numB) + (yi * h.numB) + j

		if index < len(h.votingSpace) {
			if add {
				h.votingSpace[index]++
			} else {
				h.votingSpace[index]--
			}
		}
	}
}

func roundNearest(n float64) float64 {
	if n > 0.0 {
		return math.Floor(n + 0.5)
	}
	return math.Ceil(n - 0.5)
}
