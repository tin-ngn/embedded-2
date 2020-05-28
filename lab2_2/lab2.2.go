package main

import (
	"github.com/wcharczuk/go-chart"
	"math"
	"math/rand"
	"os"
	"time"
)

var N = 64
var y = make([]float64, N, N)

var Fr1 = make([]float64, N, N)
var Fr2 = make([]float64, N, N)
var Fi1 = make([]float64, N, N)
var Fi2 = make([]float64, N, N)

var Wreal = make([]float64, N, N)
var Wimag = make([]float64, N, N)

func main() {

	n := 12
	w := 1100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := make([]float64, N, N)

	F := make([]float64, N, N)
	Freal := make([]float64, N, N)
	Fimag := make([]float64, N, N)


	for i := 0; i < N; i++ {

		var ytemp float64
		for j := 0; j < n; j++ {
			ytemp += r.Float64() *
				(math.Sin(float64((w/n)*(j+1)*i) + r.Float64()))
		}

		y[i] = ytemp
		x[i] = float64(i)

	}

	for p := 0; p < N; p++ {
		for i := 0; i < N-1; i++ {
			Freal[p] += y[i] * math.Cos(math.Pi*2*float64(p*i)/float64(N))
			Fimag[p] += y[i] * math.Sin(math.Pi*2*float64(p*i)/float64(N))
		}
		F[p] = math.Sqrt(Freal[p]*Freal[p] + Fimag[p]*Fimag[p])
	}

	calcW()

	graph := chart.Chart{
		Width:  N * 10,
		Height: 500,
		XAxis: chart.XAxis{
		},
		YAxis: chart.YAxis{
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: F,
			},
		},
	}

	os.Chdir("/home/tinguen/Desktop/embedded2.1-2.2-master/embedded2.2")
	f, _ := os.Create("graph2_2_1.png")
	graph.Render(chart.PNG, f)
	f.Close()

	Fr1 = make([]float64, N, N)
	Fi1 = make([]float64, N, N)

	DIT(Fr1, Fi1, 0)
	DIT(Fr2, Fi2, 1)

	for p := 0; p < N; p++ {
		Fr2[p] += Fr1[p]
		Fi2[p] += Fi1[p]
		F[p] = math.Sqrt(math.Pow(Fr2[p], 2) + math.Pow(Fi2[p], 2))
	}

	graph2 := chart.Chart{
		Width:  N * 10,
		Height: 500,
		XAxis: chart.XAxis{
		},
		YAxis: chart.YAxis{
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: F,
			},
		},
	}
	os.Chdir("/home/tinguen/Desktop/embedded2.1-2.2-master/embedded2.2")
	f, _ = os.Create("graph2_2_2.png")
	graph2.Render(chart.PNG, f)
	f.Close()
}


func DIT(Freal1, Fimag1 []float64, b int) {

	for p := 0; p < N/2-1; p++ {
		for i := 0; i < N/2-1; i++ {
			Freal1[p] += y[2*i+b] * Wreal[2*p*i%N]
			Fimag1[p] += y[2*i+b] * Wimag[2*p*i%N]
		}

		if b == 1 {
			Freal1[p] = Freal1[p]*Wreal[p] - Fimag1[p]*Wimag[p]
			Fimag1[p] = Freal1[p]*Wimag[p] + Fimag1[p]*Wreal[p]

			Freal1[p+N/2] = - Freal1[p]
			Fimag1[p+N/2] = - Fimag1[p]

		} else {
			Freal1[p+N/2] = Freal1[p]
			Fimag1[p+N/2] = Fimag1[p]
		}
	}
}

func calcW() {
	for i := 0; i < N; i++ {
		Wreal[i] = math.Cos(math.Pi * 2 * float64(i) / float64(N))
		Wimag[i] = math.Sin(math.Pi * 2 * float64(i) / float64(N))
	}
}
