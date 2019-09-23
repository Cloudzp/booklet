package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

func main() {

	//mqtt.DEBUG = log.New(os-tool.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	//connect mqtt-server and set clientID
	opts := mqtt.NewClientOptions().AddBroker("tcp://testpark.hzgxtc.com:1883").SetClientID("M00009")

	//set userName
	opts.SetUsername("M00009")
	opts.SetPassword("M00009")
	opts.SetProtocolVersion(4)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(false)

	//create object
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var str = `{"cipher":"jG+AOdtzkAJJoReHRaVBTSedx36d05PqNFtmgq0sBOAaiVaA1Msgw2DSVwq0gvcYaG230EfMnvv7YsYx/iIo1KBDOusgl7cOV2OdB/GX6NsPf2WwKTo04rs/N2TUFLu8Ze/YTMk9T55JXhZHAlOVYx2TLbf8pmIjwbXO+POFJXpBN/R/sRBYbqD+T/EpNYpN2Q4VNtcRRvTmCN2HJEeBh99SotBzW+8bpjueNWu+gtDS7teYJCG8FnWRK08yNLBC6aFQuTr3FocRalNOsRiNyST8yshf1PDfXOi+Qcwe4aMHjGVw18MXpLsYioOnXUrhtvc2a/GeM5uUctA1Z4UE9FW+9QxkLryt5Sh9xm5VmIl5VCikm7QXO1s01amQy0MDrosSDuASxCzEBbt4IegN4XmNTvR4QkZjOjJpTeqks7bjWalquQlL8u/z6KanSFwEqQx7GdDmC3D1X8L8Mtjko0VommaOAClpa1CWkJAykN5QwferhuK7V1UOlFOfeESs","requestId":71,"sign":"Pf91KuHpDz4Y7E1i/CEdHAF8PlqSTIDHFfJ/V78A13HemwVrYp6vTESUqKPcPlUPm9TXnX8LXYuEgRAD9kZbZGWF4S0rsCgZda/JuKqwETXs2LDhr7cAlaajlV+AINzKlEkYSIokUfpYezC4IkPBUoWfaYyAgTTgcTjAmbXoPWo=","timestamp":"2019-06-05 16:26:33"}`

	token := client.Publish("/hzcity/noticepay/M00009/M00009s", 0, false, str)
	if token.Wait(); token.Error() != nil {
		panic("token error" + token.Error().Error())
	}
}
