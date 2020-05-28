package main

import (
	"github.com/wcharczuk/go-chart"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {

	n := 12
	w := 1100
	N := 64

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var y = make([]float64, N, N)
	var x = make([]float64, N, N)

	var Freal = make([]float64, N, N)
	var Fimag = make([]float64, N, N)
	var F = make([]float64, N, N)

	var sum float64

	for i := 0; i < N; i++ {

		var ytemp float64
		for j := 0; j < n; j++ {
			ytemp += r.Float64() *
				(math.Sin(float64((w/n)*(j+1)*i) + r.Float64()))
		}
		sum += ytemp

		y[i] = ytemp
		x[i] = float64(i)

	}

	graph := chart.Chart{
		Width:  N * 10,
		Height: 500,
		XAxis: chart.XAxis{},
		YAxis: chart.YAxis{},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}

	os.Chdir("/home/tinguen/Desktop/embedded2.1-2.2-master/embedded2.1")
	f, _ := os.Create("graph2_1_1.png")
	graph.Render(chart.PNG, f)
	f.Close()

	for p := 0; p < N; p++ {
		for i := 0; i < N-1; i++ {
			Freal[p] += y[i] * math.Cos(math.Pi*2*float64(p*i)/float64(N))
			Fimag[p] += y[i] * math.Sin(math.Pi*2*float64(p*i)/float64(N))
		}
		F[p] = math.Sqrt(Freal[p]*Freal[p] + Fimag[p]*Fimag[p])
	}

	graph2 := chart.Chart{
		Width:  N * 10,
		Height: 500,
		XAxis: chart.XAxis{},
		YAxis: chart.YAxis{},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: F,
			},
		},
	}

	f, _ = os.Create("graph2_1_2.png")
	graph2.Render(chart.PNG, f)
	f.Close()
}
