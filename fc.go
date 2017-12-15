package gnn

import (
	"math"

	"github.com/jacsmith21/gnn/mat"
	"github.com/jacsmith21/gnn/vec"
)

// FC FC
type FC struct {
	In      int
	Out     int
	Weights vec.Vector
	Biases  vec.Vector
}

// NewFC NewFC
func NewFC(in, out int) *FC {
	fc := FC{
		In:      in,
		Out:     out,
		Weights: vec.Make(in * out),
		Biases:  vec.Make(out),
	}

	vec.Rand(fc.Weights, vec.Gaussian)
	fc.Weights.Scale(1 / math.Sqrt(float64(in)))

	return &fc
}

// Forward Forward
func (f FC) Forward(a mat.Matrix) {
	for i := 0; i < f.Out; i++ {

	}
}

// Backward Backward
func (f FC) Backward() {

}