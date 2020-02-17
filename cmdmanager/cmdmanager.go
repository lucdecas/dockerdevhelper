package cmdmanager

import (
	"fmt"
	"os/exec"
)

var(
	d = "docker" 
	s = "search"
	r = "run"
	filter = "--filter is-official=true --limit 1"
)

func Search(str string){
	cmd := exec.Command("docker", "ps")
	fmt.Printf("Running command and waiting for it to finish...", cmd)
	err := cmd.Run()
	fmt.Printf("Command finished with error: %v", err)
}

func Start(str string){
	fmt.Println(str)
	fmt.Println("Start Command Issued")
}

func Stop(str string){
	fmt.Println(str)
	fmt.Println("Stop Command Issued")
}

func Reset(str string){
	fmt.Println(str)
	fmt.Println("Reset command issued")
}