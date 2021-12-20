package cli

import (
	"bytes"
	"fmt"
	"os/exec"
)

const SacctMgr string = "sacctmgr"

type SacctMgrCLIHander interface {
	ListUser() (string, error)
	ListQOS() (string, error)
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

func (s SacctMgrCLI) ListQOS() (string, error) {
	cmd := exec.Command(SacctMgr, "-noheader", "-P", "show qos format=sacctmgr -P show qos format=ID,Name,Description,GraceTime,GrpTRESMins,GrpTRESRunMins,GrpTRES,GrpJobs,GrpJobsAccrue,GrpSubmitJobs,GrpWall,MaxTRESMinsPerJob,MaxTRESPerAccount,MaxTRESPerJob,MaxTRESPerNode,MaxTRESPerUser,MaxJobsPerAccount,MaxJobsAccruePerAccount,MaxJobsPerUser,MaxJobsAccruePerUser,MaxSubmitJobsPerAccount,MaxSubmitJobsPerUser,MaxWallDurationPerJob,MinPrioThreshold,MinTRESPerJob,Preempt,PreemptExemptTime,PreemptMode,Priority,UsageFactor,Flags")
	fmt.Println(cmd)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
