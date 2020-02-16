package main

import (
	"fmt"
	"time"

	"github.com/solsson/go-conbee/sensors"

	"github.com/prometheus/client_golang/prometheus"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type deconzCollector struct {
	temperatureMetric *prometheus.Desc
	humidityMetric *prometheus.Desc
	pressureMetric *prometheus.Desc
}

var (
	ss *sensors.Sensors = nil
)

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newDeconzCollector(sss *sensors.Sensors) *deconzCollector {
	ss = sss
	sensors, err := ss.GetAllSensors()
	if err != nil {
		fmt.Println("sensors.GetAllSensors() ERROR: ", err)
	} else {
		fmt.Println()
		fmt.Println("Sensors")
		fmt.Println("------")
		for _, l := range sensors {
			//fmt.Printf("Sensor:\n%s\n", l.StringWithIndentation("  "))
	
			if l.Type == "ZHATemperature" {
				fmt.Printf("%s Temperature %d '%s' %d\n", l.State.LastUpdated, l.ID, l.Name, l.State.Temperature)
			}
			if l.Type == "ZHAHumidity" {
				fmt.Printf("%s Humidity %d '%s' %d\n", l.State.LastUpdated, l.ID, l.Name, l.State.Humidity)
			}
			if l.Type == "ZHAPressure" {
				fmt.Printf("%s Pressure %d '%s' %d\n", l.State.LastUpdated, l.ID, l.Name, l.State.Pressure)
			}
		}
	}

	return &deconzCollector{
		temperatureMetric: prometheus.NewDesc(
			"climate_temperature",
			"Temperature C",
			[]string{"name"}, nil,
		),
		humidityMetric: prometheus.NewDesc(
			"climate_humidity",
			"Humidity %",
			[]string{"name"},
			nil,
		),
		pressureMetric: prometheus.NewDesc(
			"climate_pressure",
			"Pressure",
			[]string{"name"},
			nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *deconzCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.temperatureMetric
	ch <- collector.humidityMetric
	ch <- collector.pressureMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *deconzCollector) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()

	sensors, err := ss.GetAllSensors()
	if err != nil {
		fmt.Println("sensors.GetAllSensors() ERROR: ", err)
		return
	}

	oldest := "2999-01-01T00:00:00"

	for _, l := range sensors {
		//fmt.Printf("Sensor:\n%s\n", l.StringWithIndentation("  "))
		if l.Type == "ZHATemperature" {
			ch <- prometheus.MustNewConstMetric(
				collector.temperatureMetric,
				prometheus.GaugeValue,
				float64(l.State.Temperature) / 100,
				l.Name)
			if oldest > l.State.LastUpdated {
				oldest = l.State.LastUpdated
			}
		}
		if l.Type == "ZHAHumidity" {
			ch <- prometheus.MustNewConstMetric(
				collector.humidityMetric,
				prometheus.GaugeValue,
				float64(l.State.Humidity) / 100,
				l.Name)
			if oldest > l.State.LastUpdated {
				oldest = l.State.LastUpdated
			}
		}
		if l.Type == "ZHAPressure" {
			ch <- prometheus.MustNewConstMetric(
				collector.pressureMetric,
				prometheus.GaugeValue,
				float64(l.State.Pressure),
				l.Name)
			if oldest > l.State.LastUpdated {
				oldest = l.State.LastUpdated
			}
		}
	}

	fmt.Printf("%v Metrics collected, oldest is from: %sZ\n", start.Format("2006-01-02T15:04:05.000Z"), oldest)
}
