package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"flag"
	"strings"
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
		fmt.Println(commands)
		fmt.Println(oq_commands)

	}else{
		fmt.Print("Erro ao abrir arquivo")
	}
}