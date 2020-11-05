package utils

import (
        _"fmt"
	"os/exec"
)

func RunRESTCommand(restServer string, str string) ([]uint8, error) {
        cmd := "curl -s -XGET " +restServer +str +" -H \"accept:application/json\""
        out, err := exec.Command("/bin/bash", "-c", cmd).Output()

        return out, err
}

