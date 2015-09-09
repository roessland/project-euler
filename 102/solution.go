// Runs in 8 ms -- not really CPU bound
package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

type Triangle struct {
	X1, Y1, X2, Y2, X3, Y3 float64
}

// Calculates the cross product for the vector u1->u2 and v1->v2
func Cross(u1x, u1y, u2x, u2y, v1x, v1y, v2x, v2y float64) float64 {
	Ux := u2x - u1x
	Uy := u2y - u1y
	Vx := v2x - v1x
	Vy := v2y - v1y
	return Ux*Vy - Uy*Vx
}

func (t Triangle) ContainsOrigin() bool {
	X1, Y1, X2, Y2, X3, Y3 := t.X1, t.Y1, t.X2, t.Y2, t.X3, t.Y3

	if Cross(X1, Y1, X2, Y2, X1, Y1, 0, 0) > 0 &&
		Cross(X2, Y2, X3, Y3, X2, Y2, 0, 0) > 0 &&
		Cross(X3, Y3, X1, Y1, X3, Y3, 0, 0) > 0 {
		return true
	} else {
		return false
	}
}

func ReadTriangles(filename string) []Triangle {
	triangles := []Triangle{}

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Oops!")
	}
	s := strings.TrimSpace(string(buf))
	trianglesStr := strings.Split(s, "\n")
	for _, triangleStr := range trianglesStr {
		coordsStr := strings.Split(triangleStr, ",")
		X1, _ := strconv.ParseFloat(coordsStr[0], 64)
		Y1, _ := strconv.ParseFloat(coordsStr[1], 64)
		X2, _ := strconv.ParseFloat(coordsStr[2], 64)
		Y2, _ := strconv.ParseFloat(coordsStr[3], 64)
		X3, _ := strconv.ParseFloat(coordsStr[4], 64)
		Y3, _ := strconv.ParseFloat(coordsStr[5], 64)

		// Fix wrongly oriented triangles
		if Cross(X1, Y1, X2, Y2, X1, Y1, X3, Y3) < 0 {
			X2, Y2, X3, Y3 = X3, Y3, X2, Y2
		}
		triangles = append(triangles, Triangle{X1, Y1, X2, Y2, X3, Y3})
	}
	return triangles
}

func main() {
	triangles := ReadTriangles("p102_triangles.txt")

	count := 0
	for _, triangle := range triangles {
		if triangle.ContainsOrigin() {
			count++
		}
	}
	fmt.Printf("%v\n", count)
}
