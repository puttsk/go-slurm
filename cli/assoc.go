package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/puttsk/go-slurm"
)

func parseAssocList(s string) ([]slurm.AssocRecord, error) {
	// Assume sacctmgr comes with headers
	// Parse sacctmgr show user output and return a list of UserRecords

	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid input")
	}

	var alist []slurm.AssocRecord

	// Sacctmgr format and output field mapping
	// format=ID,User,Account,Cluster,DefaultQOS,Fairshare,GrpTRESMins,
	//        ID|User|Account|Cluster|Def QOS   |Share    |GrpTRESMins|
	// format=GrpTRESRunMins,GrpTRES,GrpJobs,GrpJobsAccrue,GrpSubmitJobs,
	//        GrpTRESRunMins|GrpTRES|GrpJobs|GrpJobsAccrue|GrpSubmit    |
	// format=GrpWall,MaxTRESPerJob,MaxTRESMinsPerJob,MaxTRESPerNode,
	//        GrpWall|MaxTRES      |MaxTRESMins      |MaxTRESPerNode|
	// format=MaxJobs,MaxJobsAccrue,MaxSubmitJobs,MaxWallDurationPerJob,
	//        MaxJobs|MaxJobsAccrue|MaxSubmit    |MaxWall              |
	// format=Qos,QosRaw ,ParentID,ParentName,Partition,Priority,LFT,RGT
	//        QOS|QOS_RAW|Par ID  |Par Name  |Partition|Priority|LFT|RGT
	headers := strings.Split(lines[0], SlurmDelimiter)
	for _, l := range lines[1:] {
		var a slurm.AssocRecord
		for i, v := range strings.Split(l, SlurmDelimiter) {
			switch headers[i] {
			case "ID":
				o, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return nil, err
				}
				a.ID = uint32(o)
			case "User":
				a.User = v
			case "Account":
				a.Acct = v
			case "Cluster":
				a.Cluster = v
			case "Def QOS":
				// TODO
			case "Share":
				o, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return nil, err
				}
				a.SharesRaw = uint32(o)
			case "GrpTRESMins":
				a.GrpTRESMins = v
			case "GrpTRES":
				a.GrpTRES = v
			case "GrpJobs":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.GrpJobs = val
			case "GrpJobsAccrue":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.GrpJobsAccrue = val
			case "GrpSubmit":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.GrpSubmitJobs = val
			case "GrpWall":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.GrpWall = val
			case "MaxTRES":
				a.MaxTRESPJ = v
			case "MaxTRESMins":
				a.MaxTRESMinsPJ = v
			case "MaxTRESPerNode":
				a.MaxTRESPN = v
			case "MaxJobs":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.MaxJobs = val
			case "MaxJobsAccrue":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.MaxJobsAccrue = val
			case "MaxSubmit":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.MaxSubmitJobs = val
			case "MaxWall":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.MaxWallPJ = val
			case "QOS":
				q := strings.TrimSpace(v)
				a.QOSList = strings.Split(q, ",")
			case "QOS_RAW":
				o, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return nil, err
				}
				a.DefQOSID = uint32(o)
			case "Par ID":
				o, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return nil, err
				}
				a.ParentID = uint32(o)
			case "Par Name":
				a.ParentAcct = v
			case "Partition":
				a.Partition = v
			case "Priority":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				a.Priority = val
			case "LFT":
				o, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return nil, err
				}
				a.Lft = uint32(o)
			case "RGT":
				o, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return nil, err
				}
				a.Rgt = uint32(o)
			}
		}
		alist = append(alist, a)
	}

	return alist, nil
}

func ListAssoc() ([]slurm.AssocRecord, error) {
	output, err := sacctmgrHander.ListAssoc()
	if err != nil {
		return nil, err
	}

	return parseAssocList(output)
}
