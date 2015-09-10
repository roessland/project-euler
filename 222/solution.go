package main

import "os"
import "fmt"
import "math"
import "text/template"

const templ = `\documentclass{article}
\usepackage{tikz}
\usepackage{tikzscale}
\usepackage{fullpage}

\begin{document}
\begin{tikzpicture}[scale=0.10, x=1mm, y=1mm]
    \draw (0,0) -- (0, {{.L}});
    \draw (0,0) -- (100, 0);
    \draw (100,0) -- (100, {{.L}});
	{{range .Circles}}\draw ({{.Y}}, {{.X}}) circle [radius={{.R}}] node { {{.R}} };
{{end}}
	\draw (0, {{.L}}) -- (100, {{.L}});
	\node at (50, 1800) {L = {{.L}} };
\end{tikzpicture}
\end{document}
`

type Data struct {
	Circles []Circle
	L       float64
}

type Circle struct {
	R, X, Y float64
}

func CreateFigure(filename string, R float64, xx, yy, rr []float64, L float64) {
	circles := []Circle{}
	for i, _ := range rr {
		circles = append(circles, Circle{rr[i], xx[i], yy[i]})
	}

	t := template.New("Balls template")
	t, err := t.Parse(templ)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
	err = t.Execute(os.Stdout, Data{circles, L})
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}

}

func AddBall(R float64, xx, yy, rr []float64, Qr float64) ([]float64, []float64, []float64) {
	if len(xx) == 0 {
		xx = append(xx, Qr)
		yy = append(yy, Qr)
		rr = append(rr, Qr)
		return xx, yy, rr
	}
	Px := xx[len(xx)-1]
	Py := yy[len(yy)-1]
	Pr := rr[len(rr)-1]
	var flipped = false
	if Py > R {
		Py = 2*R - Py
		flipped = true
	}

	Qy := 2*R - Qr
	dx := math.Sqrt((Pr+Qr)*(Pr+Qr) - (Qy-Py)*(Qy-Py))
	Qx := Px + dx

	if flipped {
		Qy = 2*R - Qy
	}

	xx = append(xx, Qx)
	yy = append(yy, Qy)
	rr = append(rr, Qr)

	return xx, yy, rr
}

func main() {
	R := float64(50)
	var xx, yy, rr []float64

	// Observe that 50 and 49 should be in oppsite ends, to save as much as
	// space as possible. If we continue the same pattern with the remaining
	// balls 30..48, we end up with the even balls on the bottom and odd balls
	// on the top. This way the small balls end up in the middle, which saves a
	// lot of space.
	rs := []float64{50, 48, 46, 44, 42, 40, 38, 36, 34, 32, 30, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49}

	for _, r := range rs {
		xx, yy, rr = AddBall(R, xx, yy, rr, float64(r))
	}

	L := xx[len(xx)-1] + rr[len(rr)-1]

	CreateFigure("fig.tex", R, xx, yy, rr, L)
}
