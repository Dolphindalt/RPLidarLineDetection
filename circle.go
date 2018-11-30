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
	// holy fucking fuck
	c.triangles = append(c.triangles, 0)
	c.triangles = append(c.triangles, 1)
	c.triangles = append(c.triangles, 2) // 1
	c.triangles = append(c.triangles, 0)
	c.triangles = append(c.triangles, 1)
	c.triangles = append(c.triangles, 3) // 2
	c.triangles = append(c.triangles, 0)
	c.triangles = append(c.triangles, 2)
	c.triangles = append(c.triangles, 4) // 3
	c.triangles = append(c.triangles, 0)
	c.triangles = append(c.triangles, 4)
	c.triangles = append(c.triangles, 6) // 4
	c.triangles = append(c.triangles, 0)
	c.triangles = append(c.triangles, 3)
	c.triangles = append(c.triangles, 6) // 5
	c.triangles = append(c.triangles, 1)
	c.triangles = append(c.triangles, 2)
	c.triangles = append(c.triangles, 5) // 6
	c.triangles = append(c.triangles, 1)
	c.triangles = append(c.triangles, 3)
	c.triangles = append(c.triangles, 7) // 7
	c.triangles = append(c.triangles, 1)
	c.triangles = append(c.triangles, 5)
	c.triangles = append(c.triangles, 7) // 8
	c.triangles = append(c.triangles, 2)
	c.triangles = append(c.triangles, 4)
	c.triangles = append(c.triangles, 8) // 9
	c.triangles = append(c.triangles, 2)
	c.triangles = append(c.triangles, 5)
	c.triangles = append(c.triangles, 8) // 10
	c.triangles = append(c.triangles, 3)
	c.triangles = append(c.triangles, 6)
	c.triangles = append(c.triangles, 9) // 11
	c.triangles = append(c.triangles, 3)
	c.triangles = append(c.triangles, 7)
	c.triangles = append(c.triangles, 9) // 12
	c.triangles = append(c.triangles, 4)
	c.triangles = append(c.triangles, 8)
	c.triangles = append(c.triangles, 10) // 13
	c.triangles = append(c.triangles, 8)
	c.triangles = append(c.triangles, 10)
	c.triangles = append(c.triangles, 11) // 14
	c.triangles = append(c.triangles, 5)
	c.triangles = append(c.triangles, 8)
	c.triangles = append(c.triangles, 11) // 15
	c.triangles = append(c.triangles, 5)
	c.triangles = append(c.triangles, 7)
	c.triangles = append(c.triangles, 11) // 16
	c.triangles = append(c.triangles, 7)
	c.triangles = append(c.triangles, 9)
	c.triangles = append(c.triangles, 11) // 17
	c.triangles = append(c.triangles, 9)
	c.triangles = append(c.triangles, 10)
	c.triangles = append(c.triangles, 11) // 18
	c.triangles = append(c.triangles, 6)
	c.triangles = append(c.triangles, 9)
	c.triangles = append(c.triangles, 10) // 19
	c.triangles = append(c.triangles, 4)
	c.triangles = append(c.triangles, 6)
	c.triangles = append(c.triangles, 10) // 20
}

// To the compiler: no no no somebody else do it I will not
func (cir *Circle) subDivide() {
	verticesLength := len(cir.vertices)
	var norm float64
	num := len(cir.triangles) / 3
	// subdividing those triangles
	for i := 0; i < num; i++ {
		var a, b, c, d, e, f *Vector2f
		var ai, bi, ci, di, ei, fi int
		ai = cir.triangles[0]
		bi = cir.triangles[1]
		ci = cir.triangles[2]
		cir.triangles = cir.triangles[3:] // deque first 3
		a = cir.vertices[ai]
		b = cir.vertices[bi]
		c = cir.vertices[ci]
		// d = a + b
		d = VectorAdd(a, b)
		norm = d.Magnitude()
		d = ScalarQuotient(d, norm)
		// e = b + c
		e = VectorAdd(b, c)
		norm = e.Magnitude()
		e = ScalarQuotient(e, norm)
		// f = c + a
		f = VectorAdd(c, a)
		norm = f.Magnitude()
		f = ScalarQuotient(f, norm)
		// add new stuff to triangles
		foundD := false
		foundE := false
		foundF := false
		for j := verticesLength; j < len(cir.vertices); j++ {
			if VectorEquals(cir.vertices[j], d) {
				foundD = true
				di = j
				continue
			}
			if VectorEquals(cir.vertices[j], e) {
				foundE = true
				continue
			}
			if VectorEquals(cir.vertices[j], f) {
				foundF = true
				fi = j
				continue
			}
		}

		if !foundD {
			di = len(cir.vertices)
			cir.vertices = append(cir.vertices, d)
		}
		if !foundE {
			ei = len(cir.vertices)
			cir.vertices = append(cir.vertices, e)
		}
		if !foundF {
			fi = len(cir.vertices)
			cir.vertices = append(cir.vertices, f)
		}

		cir.triangles = append(cir.triangles, ai)
		cir.triangles = append(cir.triangles, di)
		cir.triangles = append(cir.triangles, fi)

		cir.triangles = append(cir.triangles, di)
		cir.triangles = append(cir.triangles, bi)
		cir.triangles = append(cir.triangles, ei)

		cir.triangles = append(cir.triangles, fi)
		cir.triangles = append(cir.triangles, ei)
		cir.triangles = append(cir.triangles, ci)

		cir.triangles = append(cir.triangles, fi)
		cir.triangles = append(cir.triangles, di)
		cir.triangles = append(cir.triangles, ei)
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
