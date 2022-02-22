// main()

// receber paramentros via terminar(como no mqtt-lantency anterior)

// chamar funções do connection.go

// gerar tratar resultados do teste


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
   if *clients < 1 {
       log.Fatal("Invalid arguments")
   }
 
   if *fin && *fout {
       log.Fatal("Invalid arguments")
   }

   pub_clients := *clients
   sub_clients := *clients
   if *fin{
       sub_clients = 1
   } else if *fout{
       pub_clients = 1
   }

   // checar clientes que desejam fazer conexão

   // conectar com esses clientes

   // os dados das flags para esses clientes
}