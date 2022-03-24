import (
	"fmt"
	"log"
	"time"
	"strconv"
	"os"
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
		benckmark = flag.Int("benckmark", 1, "Selected benckmark software:\n 1 - Soft 1\n 2 - Soft 2")
	)

	flag.Parse()

	commands := []string {
		*broker,
		*topic,
		string(*fin),
		string(*fout),
		*username,
		*password,
		string(*pubqos),
		string(*subqos),
		string(*size),
		string(*count),
		string(*clients),
		string(*keepalive),
		*format,
		string(*quiet),
		string(*frequency),
		string(*benckmark)
	}

    opts := mqtt.NewClientOptions().AddBroker(*broker).SetAutoReconnect(true).SetKeepAlive(ka).SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
    }).SetConnectionLostHandler(func(client mqtt.Client, reason error) {
        log.Printf("SUBSCRIBER Orquestrador lost connection to the broker: %v. Will reconnect...\n", reason.Error())
    })

    client := mqtt.NewClient(opts)

    if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("SUBSCRIBER Orquestrador had error connecting to the broker: %v\n", token.Error())
        os.Exit(3)
	}

	token := client.Subscribe("status_clients", byte(1), func(Client client, Message m) {
		// [0:8] command, [8:16] result 
		message := string(m.Payload())
		commad_channel := message[0:8]
		resul_channel := message[8:16]

		var command_message := strings.Join(commands," ")

		// enviar comando
		client.Publish(commad_channel, byte(1), true, command_message)
		// se inscrever no canal que o cliente enviou
		client.Subscribe(resul_channel, byte(1), func(Client client, Message m){
			result := string(m.Payload())
			// adicionar em uma estrutura de resultados
		})
	}) // ver se eh isso mesmo
}