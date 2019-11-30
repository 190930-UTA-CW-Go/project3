package downloadfile

import (
	"fmt"
	"os/exec"
	"strings"
)

//DownloadAWS download a file form AWS
func DownloadAWS(name string) {

	folder := strings.Title(strings.ToLower(name))
	file := strings.ToLower(name)

	cmd := exec.Command("scp", "-i", "~/go/src/github.com/190930-UTA-CW-Go/project3/grader/downloadfile/rego.pem", "ec2-user@ec2-18-188-174-65.us-east-2.compute.amazonaws.com:/home/ec2-user/Portfolios/"+folder+"/"+file+".json", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + "; " + string(output))
		return
	}
	fmt.Println(fmt.Sprint(err) + ": " + string(output))

	fmt.Println("Downloading file....")

}
