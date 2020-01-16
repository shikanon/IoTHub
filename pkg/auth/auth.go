package auth

import "github.com/shikanon/IoTOrbHub/pkg/hub"

type DeviceAuthen interface {
	CheckDevice(hub.Device) bool
}
