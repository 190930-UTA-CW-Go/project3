package downloadfile

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)


func DownloadAWS(name string){
	
	folder := strings.Title(strings.ToLower(name))
	file := strings.ToLower(name)
	
	cmd, err := exec.Command("scp", "-i", "rego.pem", "ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com:/home/ec2-user/Portfolios/" + folder + "/" + file + ".json", ".").Output()
	if err != nil{
		log.Fatalf("cmd.Run() failed with %s\n", err)
		}

	fmt.Println(cmd)
	fmt.Println("Downloading file....")

	}

	