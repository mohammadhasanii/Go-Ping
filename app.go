package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	var ip string
	var count int

	fmt.Print("Enter the IP address :  ")
	fmt.Scanln(&ip)

	fmt.Print("Enter the number of pings : ")
	fmt.Scanln(&count)

	//This command is only for Windows and on Linux this command is different
	cmd := exec.Command("cmd", "/C", "ping", "-n", strconv.Itoa(count), ip)
	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	defer output.Close()

	if err := cmd.Start(); err != nil {
		fmt.Println("Error executing the command ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error executing the command ", err)
		os.Exit(1)
	}
}
