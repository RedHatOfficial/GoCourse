package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

type Points struct {
	X []float64 `json:"x"`
	Y []float64 `json:"y"`
}

const npoints = 40
const periods = 3

func makePoints(npoints uint) Points {
	var points Points
	points.X = make([]float64, npoints)
	points.Y = make([]float64, npoints)
	for i := uint(0); i < npoints; i++ {
		t := float64(i) * periods * 2.0 * math.Pi / float64(npoints)
		points.X[i] = t
		// limita
		if t == 0.0 {
			points.Y[i] = 1.0
		} else {
			points.Y[i] = math.Sin(t) / t
		}
	}
	return points
}

func valuesHandler(writer http.ResponseWriter, request *http.Request) {
	points := makePoints(npoints)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(points)
}

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir("chart-js/")))
	http.HandleFunc("/values", valuesHandler)
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
