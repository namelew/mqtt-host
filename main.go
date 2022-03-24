import (
	"fmt"
	"log"
	"time"
	"strconv"
)

import (
	"github.com/GaryBoone/GoStats/stats"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)


func main(){
	var (
		broker    = flag.String("broker", "tcp://localhost:1883", "MQTT broker endpoint as scheme://host:port")
		topic     = flag.String("topic", "/test", "MQTT topic for outgoing messages")
		fin       = flag.Bool("fin", false, "Allows fan-in tests")
		fout      = flag.Bool("fout", false, "Allows fan-out tests")
		username  = flag.String("username", "", "MQTT username (empty if auth disabled)")
		password  = flag.String("password", "", "MQTT password (empty if auth disabled)")
		pubqos    = flag.Int("pubqos", 1, "QoS for published messages")
		subqos    = flag.Int("subqos", 1, "QoS for subscribed messages")
		size      = flag.Int("size", 100, "Size of the messages payload (bytes)")
		count     = flag.Int("count", 100, "Number of messages to send per pubclient")
		clients   = flag.Int("clients", 10, "Number of clients pair to start")
		keepalive = flag.Int("keepalive", 60, "Keep alive period in seconds")
		format    = flag.String("format", "text", "Output format: text|json")
		quiet     = flag.Bool("quiet", false, "Suppress logs while running")
		frequency = flag.Int("frequency", 1, "Number of publications per second")
	)

	flag.Parse()
    ka, _ := time.ParseDuration(strconv.Itoa(*keepalive) + "s")

    opts := mqtt.NewClientOptions().AddBroker(*broker).SetAutoReconnect(true).SetKeepAlive(ka).SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
        // ler mensagem e achar canal de envio de comandos e canal de recebimento de resultados
    }).SetConnectionLostHandler(func(client mqtt.Client, reason error) {
        log.Printf("SUBSCRIBER Orquestrador lost connection to the broker: %v. Will reconnect...\n", reason.Error())
    })

    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("SUBSCRIBER Orquestrador had error connecting to the broker: %v\n", token.Error())
        // exit
	}

	if token := client.Subscribe("status_clients", byte(1), nil); token.Wait() && token.Error() != nil {
		log.Printf("SUBSCRIBER Orquestrador had error subscribe with topic: %v\n", token.Error())
		// exit
	}

    // enviar commandos aos cliente encontrados
    // esperar respostas e calcular resultados

}