package main

import (
	"fmt"
	"os"
	"path/filepath"
	"flag"
)

func main(){
	var (
		path = flag.String("path", "commands.config", "open a file with the commands to the app")
	)

	flag.Parse()

	extInput := filepath.Ext(*path)

	if (extInput == ".config"){
		file, _ := os.Open(*path)
		fmt.Print(file)
		file.Close()	
	}else{
		fmt.Print("Erro ao abrir arquivo")
	}
}