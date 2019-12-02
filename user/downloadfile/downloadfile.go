package downloadfile

import (
	"fmt"
	"os/exec"
)

//DownloadAWS download a file form AWS
func DownloadAWS(name string) {

	cmd := exec.Command("scp", "-i", "~/go/src/github.com/190930-UTA-CW-Go/project3/rego.pem", "ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com:/home/ec2-user/Portfolios/"+name+"/"+name+".json", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + "; " + string(output))
		return
	}
	fmt.Println(fmt.Sprint(err) + ": " + string(output))

	fmt.Println("Downloading file....")

}
