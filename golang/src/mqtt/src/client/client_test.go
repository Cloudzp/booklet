package client_test

import (
	"mqtt/client"
	"testing"
	proto "github.com/huin/mqtt"
)

func TestNewClient(t *testing.T) {
	tqs := proto.TopicQos{
		Topic: "cloudzp-test",
		Qos:proto.QosAtMostOnce,
	}
	c, err := client.NewClient("localhost:8888",tqs)
	if err != nil {
		t.Error(err)
		return
	}
	c.Run()

	c.Stop()
}
