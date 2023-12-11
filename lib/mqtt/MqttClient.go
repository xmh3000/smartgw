package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type (
	Client interface {
		Connect()
		Disconnect()
		ReportTelemetry(payload []byte)
		ReceiveMessage(msg mqtt.Message)
	}
)
