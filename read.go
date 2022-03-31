package main

import (
	"fmt"
	"io/ioutil"
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
		content, _ := ioutil.ReadFile(*path)
		fmt.Print(string(content))	
	}else{
		fmt.Print("Erro ao abrir arquivo")
	}
}