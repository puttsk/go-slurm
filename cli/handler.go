package cli

import (
	"bytes"
	"fmt"
	"os/exec"
)

const SacctMgr string = "sacctmgr"

type SacctMgrCLIHander interface {
	ListUser() (string, error)
}

type SacctMgrCLI struct {
}

func (s SacctMgrCLI) ListUser() (string, error) {
	cmd := exec.Command(SacctMgr, "-noheader", "-P", "show user format=adminlevel,coordinators,defaultaccount,defaultwckey,user")
	fmt.Println(cmd)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
