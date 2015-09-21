package rabbitmq

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"
	. "bitbucket.org/makeusmobile/makeus-golang-framework/src/utility"

	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	Conn    *amqp.Connection // Dial
	Channel *amqp.Channel    // Channel
}

var Instance = &Rabbitmq{Conn: nil, Channel: nil}

func PublisherInit() *amqp.Channel {

	conn, err := amqp.Dial(config.RABBITMQ_ADDR)
	ErrCheck(err)

	ch, err := conn.Channel()
	ErrCheck(err)

	_, err = ch.QueueDeclare(
		config.RABBITMQ_QUEUENAME, // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	ErrCheck(err)

	Instance = &Rabbitmq{Conn: conn, Channel: ch}

	return ch
}

func SubscriberInit() <-chan amqp.Delivery {

	conn, err := amqp.Dial(config.RABBITMQ_ADDR)
	ErrCheck(err)

	ch, err := conn.Channel()
	ErrCheck(err)

	msgs, err := ch.Consume(
		config.RABBITMQ_QUEUENAME, // queue name
		"",    // consumer   	consumer에 대한 식별자를 지정합니다. consumer tag는 로컬에 channel이므로, 두 클라이언트는 동일한 consumer tag를 사용할 수있다.
		false, // autoAck    	false는 명시적 Ack를 해줘야 메시지가 삭제되고 true는 메시지를 빼면 바로 삭제
		false, // exclusive	현재 connection에만 액세스 할 수 있으며, 연결이 종료 할 때 Queue가 삭제됩니다.
		false, // noLocal    	필드가 설정되는 경우 서버는이를 published 연결로 메시지를 전송하지 않을 것입니다.
		false, // noWait		설정하면, 서버는 Method에 응답하지 않습니다. 클라이언트는 응답 Method를 기다릴 것이다. 서버가 Method를 완료 할 수 없을 경우는 채널 또는 연결 예외를 발생시킬 것입니다.
		nil,   // arguments	일부 브로커를 사용하여 메시지의 TTL과 같은 추가 기능을 구현하기 위해 사용된다.
	)

	ErrCheck(err)

	Instance = &Rabbitmq{Conn: conn, Channel: ch}

	return msgs
}

func Close() {
	Instance.Channel.Close()
	Instance.Conn.Close()
}
