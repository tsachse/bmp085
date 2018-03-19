package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
	"github.com/kidoman/embd/sensor/bmp085"
)

var m map[string]int

var (
	bmp085Temperature = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bmp085_temperature",
			Help: "BMP085 Temperature",
		})

	bmp085Pressure = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "bmp085_pressure",
			Help: "BMP085 Pressure",
		})
)

func init() {
	prometheus.MustRegister(bmp085Temperature)
	prometheus.MustRegister(bmp085Pressure)
}

func main() {
	bus := embd.NewI2CBus(1)
	baro := bmp085.New(bus)

	go func() {
		for {

			temp, err := baro.Temperature()
			if err != nil {
				temp = 0.0
			}
			press, err := baro.Pressure()
			if err != nil {
				press = 0
			}
			bmp085Temperature.Set(temp)
			bmp085Pressure.Set(float64(press))

			time.Sleep(550 * time.Millisecond)
		}
	}()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))

}
