package cli

import (
	"bytes"
	"log"
	"os/exec"
)

const SacctMgr string = "sacctmgr"

type SacctMgrCLIHander interface {
	ListUser() (string, error)
	ListQOS() (string, error)
	ListAccount() (string, error)
	ListAssoc() (string, error)
}

type SacctMgrCLI struct {
}

func (s SacctMgrCLI) executeCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	log.Println(cmd)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func (s SacctMgrCLI) ListUser() (string, error) {
	return s.executeCommand(SacctMgr,
		"-noheader",
		"-P",
		"show user format=adminlevel,coordinators,defaultaccount,defaultwckey,user")

}

func (s SacctMgrCLI) ListQOS() (string, error) {
	return s.executeCommand(SacctMgr,
		"-noheader",
		"-P",
		"show qos format=ID,Name,Description,GraceTime,GrpTRESMins,GrpTRESRunMins,GrpTRES,GrpJobs,GrpJobsAccrue,GrpSubmitJobs,GrpWall,MaxTRESMinsPerJob,MaxTRESPerAccount,MaxTRESPerJob,MaxTRESPerNode,MaxTRESPerUser,MaxJobsPerAccount,MaxJobsAccruePerAccount,MaxJobsPerUser,MaxJobsAccruePerUser,MaxSubmitJobsPerAccount,MaxSubmitJobsPerUser,MaxWallDurationPerJob,MinPrioThreshold,MinTRESPerJob,Preempt,PreemptExemptTime,PreemptMode,Priority,UsageFactor,Flags")
}

func (s SacctMgrCLI) ListAccount() (string, error) {
	return s.executeCommand(SacctMgr,
		"-noheader",
		"-P",
		"show account format=Account,Description,Organization,Coordinators")
}

func (s SacctMgrCLI) ListAssoc() (string, error) {
	return s.executeCommand(SacctMgr,
		"-noheader",
		"-P",
		"show assoc format=ID,User,Account,Cluster,DefaultQOS,Fairshare,GrpTRESMins,GrpTRESRunMins,GrpTRES,GrpJobs,GrpJobsAccrue,GrpSubmitJobs,GrpWall,MaxTRESPerJob,MaxTRESMinsPerJob,MaxTRESPerNode,MaxJobs,MaxJobsAccrue,MaxSubmitJobs,MaxWallDurationPerJob,Qos,QosRaw,ParentID,ParentName,Partition,Priority,LFT,RGT")
}
