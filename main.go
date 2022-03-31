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

func RemoveIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}

func main(){
	var (
		path = flag.String("path", "commands.config", "open a file with the commands to the app")
	)

	flag.Parse()

	extInput := filepath.Ext(*path)

	if (extInput == ".config"){
		content, _ := ioutil.ReadFile(*path)
		data := strings.Split(string(content), "\n")
		oq_commands := ""
		
		for i:=0; i < len(data); i++{
			if strings.Contains(data[i], "//"){
				data = RemoveIndex(data, i)
			}
		}

		for i:=0; i < len(data); i++{
			if strings.Contains(data[i], "@"){
				oq_commands += data[i] + " "
				data = RemoveIndex(data, i)
				i -= 1 
			}
		}
		
		commands := strings.Join(data," ")
	}else{
		fmt.Print("Erro ao abrir arquivo")
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

		// enviar comando
		client.Publish(commad_channel, byte(1), true, commands)
		// se inscrever no canal que o cliente enviou
		client.Subscribe(resul_channel, byte(1), func(Client client, Message m){
			result := string(m.Payload())
			// adicionar em uma estrutura de resultados
		})
	}) // ver se eh isso mesmo
}