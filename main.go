package main

import (
	"log"

	"gonum.org/v1/gonum/mat"
)

// Line it is a god damn line
type Line struct {
	p1 Vector2f
	p2 Vector2f
}

func main() {
	lineSegmentDetection()
}

// holy shit
func orthogonalLSQ(pc *PointCloud, a *Vector2f, b *Vector2f) (complex128, *PointCloud, *Vector2f, *Vector2f) {
	a = pc.MeanValue()

	points := mat.NewDense(len(pc.points), 3, nil)
	for i := 0; i < len(pc.points); i++ {
		points.Set(i, 0, pc.points[i].x)
		points.Set(i, 1, pc.points[i].y)
		points.Set(i, 2, 0.0)
	}

	centered := mat.NewDense(3, 3, nil)
	for i := 0; i < 3; i++ {
		col := centered.ColView(i)
		sum := 0.0
		for k := 0; k < col.Len(); k++ {
			sum += col.AtVec(k)
		}
		sum /= float64(col.Len()) // sum is the mean of the column
		for j := 0; j < 3; j++ {
			centered.Set(j, i, centered.At(j, i)-sum)
		}
	}

	centered.Mul(centered, centered.T())

	var eig mat.Eigen
	ok := eig.Factorize(centered, true, false)
	if !ok { // what the hell is an eigendecomposition
		log.Fatal("Eigendecomposition failed")
	}
	eigvecs := eig.LeftVectors()

	b.x = eigvecs.At(0, 2)
	b.y = eigvecs.At(1, 2)
	rc := eig.Values(nil)
	return rc[2], pc, a, b
}

func lineSegmentDetection() []*Line {
	var lines []*Line

	optDx := 0.0
	optLinesToDetect := 20
	optMinVotes := 10
	// number of icosahedron subdivisions for direction discretization
	granularity := 4

	var cloud *PointCloud
	cloud = &PointCloud{NewVector2f(0, 0), make([]*Vector2f, 0)}
	cloud.ReadFromFile("test.txt")
	var minP, maxP, shiftedMinP, shiftedMaxP *Vector2f
	minP = NewVector2f(0, 0)
	maxP = NewVector2f(0, 0)
	shiftedMinP = NewVector2f(0, 0)
	shiftedMaxP = NewVector2f(0, 0)
	minP, maxP = cloud.minMaxPoints(minP, maxP)
	d := (VectorSubtract(maxP, minP).Magnitude())
	if d == 0.0 {
		log.Fatalf("All points in the point cloud are identical\n")
	}
	cloud.toImageSpace()
	shiftedMinP, shiftedMaxP = cloud.minMaxPoints(shiftedMinP, shiftedMaxP)

	if optDx == 0.0 {
		optDx = d / 64.0
	} else if optDx >= d {
		log.Fatalf("dx is too large\n")
	}

	hough := NewHough(shiftedMinP, shiftedMaxP, optDx, granularity)
	hough.add(cloud)

	// iterative hough transform aka algorithm 1
	linesDetected := 0
	voteCount := 0
	var rc complex128
	rc = 0.0
	Y := &PointCloud{NewVector2f(0, 0), []*Vector2f{}}
	for {
		a := NewVector2f(0, 0) // point
		b := NewVector2f(0, 0) // direction
		hough.subtract(Y)
		voteCount, a, b = hough.getLine(a, b)

		a, b, Y = cloud.PointsCloseToLine(a, b, optDx, Y)
		rc, Y, a, b = orthogonalLSQ(Y, a, b)
		if rc == 0.0 {
			break
		}

		a, b, Y = cloud.PointsCloseToLine(a, b, optDx, Y)
		voteCount = len(Y.points)
		if voteCount < optMinVotes {
			break
		}

		rc, Y, a, b = orthogonalLSQ(Y, a, b)
		if rc == 0.0 {
			break
		}
		a = VectorAdd(a, cloud.shift)

		linesDetected++

		lines = append(lines, &Line{*a, *b})

		cloud.RemovePoints(Y)

		if !(len(cloud.points) > 1 && (optLinesToDetect == 0 || optLinesToDetect > linesDetected)) {
			break
		}
	} // WHHHAHHHAAAAAAAAAAAAAAAAA

	return lines
}
