package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func main() {
	receive()
}

func receive() {
	conn, err := amqp.Dial("amqp://guest:guest@116.62.132.145:5672")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msgChan, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer msgChan.Close()

	// qos 会保证对列中UNACK的消息数量不会超过10个，
	/*if err := msgChan.Qos(10,0,false); err != nil {
		panic(err)
	}*/
	q, err := msgChan.QueueDeclare(
		"arm_msg_broadcast_test_faq",
		true,
		false,
		false,
		false,
		nil,
	)

	// 绑定队列到交换机 以便从指定交换机获取数据
	if err := msgChan.QueueBind(
		q.Name, "arm_msg_broadcast_test", "arm_msg_broadcast_test", false, nil,
	); err != nil {
		panic(err)
	}

	msg, err := msgChan.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	go func() {
		for d := range msg {
			// 这里两个参数，第一个表示是否ack，第二个表示时候重新将消息返回队列 ；
			//d.Nack(false, true)
			d.Ack(true)
			var b map[string]interface{}

			json.Unmarshal(d.Body, &b)

			msg := fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), b["vpl_number"], b["park_code"])
			fmt.Println(msg)
		}
	}()
	<-make(chan bool)
}
