package main


import (
	"github.com/wcharczuk/go-chart"
	"math"
	"math/rand"
	"os"
	"time"
)

var N = 64
var from = 64
var to = 3264
var step = 32

var time1 = make([]float64, (to - from) / step, (to - from) / step)
var time2 = make([]float64, (to - from) / step, (to - from) / step)
var ind = make([]float64, (to - from) / step)


func compare(N, index int) {
	var Fr1 = make([]float64, N, N)
	var Fr2 = make([]float64, N, N)
	var Fi1 = make([]float64, N, N)
	var Fi2 = make([]float64, N, N)
	var y = make([]float64, N, N)

	var Wreal = make([]float64, N, N)
	var Wimag = make([]float64, N, N)
	DIT := func (Freal1, Fimag1, Wreal, Wimag []float64, b int) {

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

	calcW := func (Wreal, Wimag []float64) {
		for i := 0; i < N; i++ {
			Wreal[i] = math.Cos(math.Pi * 2 * float64(i) / float64(N))
			Wimag[i] = math.Sin(math.Pi * 2 * float64(i) / float64(N))
		}
	}

	n := 12
	w := 1100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := make([]float64, N, N)
	start1 := time.Now()
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
	elapsed1 := time.Since(start1)
	calcW(Wreal, Wimag)
	start2 := time.Now()
	Fr1 = make([]float64, N, N)
	Fi1 = make([]float64, N, N)

	DIT(Fr1, Fi1, Wreal, Wimag, 0)
	DIT(Fr2, Fi2, Wreal, Wimag, 1)

	for p := 0; p < N; p++ {
		Fr2[p] += Fr1[p]
		Fi2[p] += Fi1[p]
		F[p] = math.Sqrt(math.Pow(Fr2[p], 2) + math.Pow(Fi2[p], 2))
	}
	elapsed2 := time.Since(start2)
	time1[index] = float64(elapsed1)
	time2[index] = float64(elapsed2)

}


func main() {
	index := 0
	for i := from; i < to; i = i + step {
		ind[index] = float64(index)
		compare(i, index)
		index = index + 1
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
				Name: "DFT",
				XValues: ind,
				YValues: time1,
			},
			chart.ContinuousSeries{
				Name: "FFT",
				XValues: ind,
				YValues: time2,
			},
			chart.AnnotationSeries{
				Annotations: []chart.Value2{
					{XValue: float64(32), YValue: time1[32], Label: "DFT"},
					{XValue: float64(64), YValue: time2[64], Label: "FFT"},
				},
			},
		},
	}
	os.Chdir("/home/tinguen/Desktop/embedded2/task")
	f, _ := os.Create("graph.png")
	graph2.Render(chart.PNG, f)
	f.Close()
}

