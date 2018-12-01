package main

import (
	"math"
)

// Circle is
type Circle struct {
	vertices  []*Vector2f
	triangles []int
}

// FromIcosahedron creates the directions by subdivisions of icosahedron
func (c *Circle) FromIcosahedron(subdivisions int) {
	c.getIcosahedron()
	for i := 0; i < subdivisions; i++ {
		c.subDivide()
	}
	c.makeUnique()
}

// getIcosahedron creates the nodes and edges
func (c *Circle) getIcosahedron() {
	c.vertices = c.vertices[:0]
	c.triangles = c.triangles[:0]
	norm := math.Sqrt(1 + math.Phi*math.Phi)
	v := 1 / norm
	tau := math.Phi / norm

	var vec *Vector2f
	vec.x = -v
	vec.y = tau

	c.vertices = append(c.vertices, vec) // 1
	vec.x = v
	vec.y = tau
	c.vertices = append(c.vertices, vec) // 2
	vec.x = 0
	vec.y = v
	c.vertices = append(c.vertices, vec) // 3
	vec.x = 0
	vec.y = v
	c.vertices = append(c.vertices, vec) // 4
	vec.x = -tau
	vec.y = 0
	c.vertices = append(c.vertices, vec) // 5
	vec.x = tau
	vec.y = 0
	c.vertices = append(c.vertices, vec) // 6
	vec.x = -tau
	vec.y = 0
	c.vertices = append(c.vertices, vec) // 7
	vec.x = tau
	vec.y = 0
	c.vertices = append(c.vertices, vec) // 8
	vec.x = 0
	vec.y = -v
	c.vertices = append(c.vertices, vec) // 9
	vec.x = 0
	vec.y = -v
	c.vertices = append(c.vertices, vec) // 10
	vec.x = -v
	vec.y = -tau
	c.vertices = append(c.vertices, vec) // 11
	vec.x = v
	vec.y = -tau
	c.vertices = append(c.vertices, vec) // 12
	c.triangles = append(c.triangles,
		0, 1, 2,
		0, 1, 3,
		0, 2, 4,
		0, 4, 6,
		0, 3, 6,
		1, 2, 5,
		1, 2, 7,
		1, 5, 7,
		2, 4, 8,
		2, 5, 8,
		3, 6, 9,
		3, 4, 9,
		4, 8, 10,
		8, 10, 11,
		5, 8, 11,
		5, 7, 11,
		7, 9, 11,
		9, 10, 11,
		6, 9, 10,
		4, 6, 10,
	)
}

// To the compiler: no no no somebody else do it I will not
func (c *Circle) subDivide() {
	verticesLength := len(c.vertices)
	var norm float64
	num := len(c.triangles) / 3
	// subdividing those triangles
	for i := 0; i < num; i++ {
		var A, B, C, D, E, F *Vector2f
		var ai, bi, ci, di, ei, fi int
		ai = c.triangles[0]
		bi = c.triangles[1]
		ci = c.triangles[2]
		c.triangles = c.triangles[3:] // deque first 3
		A = c.vertices[ai]
		B = c.vertices[bi]
		C = c.vertices[ci]
		// d = a + b
		D = VectorAdd(A, B)
		norm = D.Magnitude()
		D = ScalarQuotient(D, norm)
		// e = b + c
		E = VectorAdd(B, C)
		norm = E.Magnitude()
		E = ScalarQuotient(E, norm)
		// f = c + a
		F = VectorAdd(C, A)
		norm = F.Magnitude()
		F = ScalarQuotient(F, norm)
		// add new stuff to triangles
		foundD := false
		foundE := false
		foundF := false
		for j := verticesLength; j < len(c.vertices); j++ {
			if VectorEquals(c.vertices[j], D) {
				foundD = true
				di = j
				continue
			}
			if VectorEquals(c.vertices[j], E) {
				foundE = true
				continue
			}
			if VectorEquals(c.vertices[j], F) {
				foundF = true
				fi = j
				continue
			}
		}

		if !foundD {
			di = len(c.vertices)
			c.vertices = append(c.vertices, D)
		}
		if !foundE {
			ei = len(c.vertices)
			c.vertices = append(c.vertices, E)
		}
		if !foundF {
			fi = len(c.vertices)
			c.vertices = append(c.vertices, F)
		}

		c.triangles = append(c.triangles, ai)
		c.triangles = append(c.triangles, di)
		c.triangles = append(c.triangles, fi)

		c.triangles = append(c.triangles, di)
		c.triangles = append(c.triangles, bi)
		c.triangles = append(c.triangles, ei)

		c.triangles = append(c.triangles, fi)
		c.triangles = append(c.triangles, ei)
		c.triangles = append(c.triangles, ci)

		c.triangles = append(c.triangles, fi)
		c.triangles = append(c.triangles, di)
		c.triangles = append(c.triangles, ei)
	} // AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHHHHHHHHHHHHH
}

func (c *Circle) makeUnique() {
	for i := 0; i < len(c.vertices); i++ {
		// z is always zero
		if c.vertices[i].x < 0 {
			copy(c.vertices[i:], c.vertices[i+1:]) // delete at position i
			c.vertices[len(c.vertices)-1] = nil
			c.vertices = c.vertices[:len(c.vertices)-1]
			t := 0
			for j := 0; j < len(c.triangles); {
				if c.triangles[t] == i || c.triangles[t+1] == i || c.triangles[t+2] == i {
					copy(c.vertices[j:], c.vertices[j+1:])
					c.vertices[len(c.vertices)-1] = nil
					c.vertices = c.vertices[:len(c.vertices)-1]
					j--
					copy(c.vertices[j:], c.vertices[j+1:])
					c.vertices[len(c.vertices)-1] = nil
					c.vertices = c.vertices[:len(c.vertices)-1]
					j--
					copy(c.vertices[j:], c.vertices[j+1:])
					c.vertices[len(c.vertices)-1] = nil
					c.vertices = c.vertices[:len(c.vertices)-1]
					j--
				} else {
					j += 3
					t += 3
				}
			}
			for k := 0; k < len(c.triangles); k++ {
				if c.triangles[k] > i {
					c.triangles[k]--
				}
			}
			i--
		} else if c.vertices[i].x == 0 && c.vertices[i].y == -1 {
			copy(c.vertices[i:], c.vertices[i+1:])
			c.vertices[len(c.vertices)-1] = nil
			c.vertices = c.vertices[:len(c.vertices)-1]
			t := 0
			for j := 0; j < len(c.triangles); {
				if c.triangles[t] == i || c.triangles[i+1] == i || c.triangles[i+2] == i {
					copy(c.vertices[j:], c.vertices[j+1:])
					c.vertices[len(c.vertices)-1] = nil
					c.vertices = c.vertices[:len(c.vertices)-1]
					j--
					copy(c.vertices[j:], c.vertices[j+1:])
					c.vertices[len(c.vertices)-1] = nil
					c.vertices = c.vertices[:len(c.vertices)-1]
					j--
					copy(c.vertices[j:], c.vertices[j+1:])
					c.vertices[len(c.vertices)-1] = nil
					c.vertices = c.vertices[:len(c.vertices)-1]
					j--
				} else {
					j += 3
					t += 3
				}
			}
			for k := 0; k < len(c.triangles); k++ {
				if c.triangles[k] > i {
					c.triangles[k]--
				}
			}
		}
	}
}
