package cmdmanager

import "fmt"
import "os/exec"
import "strings"

var(
	searchCmd = "docker search --filter is-official=true --limit 1 --format '{{.Name}}'"
	listCmd = "docker ps --format '{{.Image}}/{{.ID}}'"
)

func Search(str string){
	searchCmd := searchCmd+ " "+str
	cmd := createCommand(searchCmd)
	out, err := cmd.Output()
	if err != nil {
		// TODO: handle error more gracefully
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func List() {
	cmd := createCommand(listCmd)
	out, err := cmd.Output()
	containers := strings.Fields(string(out))
	if err != nil {
		// TODO: handle error more gracefully
		fmt.Println(err)
	}
	fmt.Println(containers)
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

func createCommand(str string) *exec.Cmd{
	cmd := strings.Fields(str)
	if len(cmd) < 2{
		fmt.Println("Invalid command")
	}
	return exec.Command(cmd[0], cmd[1:]...)
}