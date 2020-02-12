package main

import (
	"github.com/shikanon/IoTOrbHub/config"
	"github.com/shikanon/IoTOrbHub/pkg/apiserver"
	"github.com/shikanon/IoTOrbHub/pkg/influxdb"
	mqttBus "github.com/shikanon/IoTOrbHub/pkg/mqttbus"
)

func init() {
	config.ConfigInit()

	hub := &mqttBus.Client{
		MQTTUrl: mqttBus.MQTTUrl,
	}
	mqttBus.MQTTHub = hub
	hub.InitSubClient()
	hub.InitPubClient()
	c := influxdb.NewInfluxdbClient()
	influxdb.InfluxClient = c
}

func main() {
	apiserver.ApiRegister()
	defer influxdb.InfluxClient.Close()
}
