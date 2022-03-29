import (
	"fmt"
	"log"
	"time"
	"strconv"
	"os"
	"path/filepath"
)

func main(){
	var (
		path = flag.String("path", "/mqtthost/commands.config", "open a file with the commands to the app")
	)

	flag.Parse()

	extInput := filepath.Ext(*path)

	if (extInput == '.config'){
		file, _ := os.Open(*path)	
	}else{
		fmt.print("Erro ao abrir arquivo")
	}
}