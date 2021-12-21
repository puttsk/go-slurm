package cli

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/puttsk/go-slurm"
)

func parseQOSList(s string) ([]slurm.QOSRecord, error) {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid input")
	}

	var qlist []slurm.QOSRecord

	//Sacctmgr format and output field mapping
	//format=ID,Name,Description,GraceTime,GrpTRESMins,GrpTRESRunMins,GrpTRES,GrpJobs,GrpJobsAccrue,GrpSubmitJobs,
	//       ID|Name|Descr      |GraceTime|GrpTRESMins|GrpTRESRunMins|GrpTRES|GrpJobs|GrpJobsAccrue|GrpSubmit    |
	//format=GrpWall,MaxTRESMinsPerJob,MaxTRESPerAccount,MaxTRESPerJob,MaxTRESPerNode,MaxTRESPerUser,MaxJobsPerAccount,
	//       GrpWall|MaxTRESMins      |MaxTRESPA        |MaxTRES      |MaxTRESPerNode|MaxTRESPU     |MaxJobsPA        |
	//format=MaxJobsAccruePerAccount,MaxJobsPerUser,MaxJobsAccruePerUser,MaxSubmitJobsPerAccount,MaxSubmitJobsPerUser,
	//       MaxJobsAccruePA        |MaxJobsPU     |MaxJobsAccruePU     |MaxSubmitPA            |MaxSubmitPU         |
	//format=MaxWallDurationPerJob,MinPrioThreshold,MinTRESPerJob,Preempt,PreemptExemptTime,PreemptMode,Priority,UsageFactor,Flags
	//       MaxWall              |MinPrioThres    |MinTRES      |Preempt|PreemptExemptTime|PreemptMode|Priority|UsageFactor|Flags
	headers := strings.Split(lines[0], SlurmDelimiter)
	for _, l := range lines[1:] {
		var q slurm.QOSRecord
		for i, v := range strings.Split(l, SlurmDelimiter) {
			switch headers[i] {
			case "ID":
				o, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return nil, err
				}
				q.ID = uint32(o)
			case "Name":
				q.Name = v
			case "Descr":
				q.Description = v
			case "GraceTime":
				o, err := time.Parse("15:04:05", v)
				if err != nil {
					return nil, err
				}
				z, _ := time.Parse("15:04:05", "00:00:00")
				t := o.Sub(z)

				q.GraceTime = uint32(t.Seconds())
			case "GrpTRESMins":
				q.GrpTRESMins = v
			case "GrpTRESRunMins":
				q.GrpTRESRunMins = v
			case "GrpTRES":
				q.GrpTRES = v
			case "GrpJobs":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.GrpJobs = val
			case "GrpJobsAccrue":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.GrpJobsAccrue = val
			case "GrpSubmit":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.GrpSubmitJobs = val
			case "GrpWall":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.GrpWall = val
			case "MaxTRESMins":
				q.MaxTRESMinsPJ = v
			case "MaxTRESPA":
				q.MaxTRESPA = v
			case "MaxTRES":
				q.MaxTRESPJ = v
			case "MaxTRESPerNode":
				q.MaxTRESPN = v
			case "MaxTRESPU":
				q.MaxTRESPU = v
			case "MaxJobsPA":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.MaxJobsPA = val
			case "MaxJobsAccruePA":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.MaxJobsAccruePA = val
			case "MaxJobsPU":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.MaxJobsPU = val
			case "MaxJobsAccruePU":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.MaxJobsAccruePU = val
			case "MaxSubmitPA":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.MaxSubmitJobsPA = val
			case "MaxSubmitPU":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.MaxSubmitJobsPU = val
			case "MaxWall":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.MaxWallPJ = val
			case "MinPrioThres":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.MinPrioThresh = val
			case "MinTRES":
				q.MinTRESPJ = v
			case "Preempt":
				q.PreemptList = strings.Split(v, ",")
			case "PreemptExemptTime":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.PreemptExemptTime = val
			case "PreemptMode":
				switch v {
				case "cluster":
					q.PreemptMode = slurm.PreemptModeOff
				default:
					log.Panicf("Not implemented")
				}
			case "Priority":
				val, err := parseUint32Field(v)
				if err != nil {
					return nil, err
				}
				q.Priority = val
			case "UsageFactor":
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return nil, err
				}
				q.UsageFactor = val
			case "Flags":
				v = strings.TrimSpace(v)
				flagsStr := strings.Split(v, ",")
				flag := uint32(0)

				for _, f := range flagsStr {
					switch f {
					case "NoDecay":
						flag = flag | slurm.QOSFlagNoDecay
					case "":
					default:
						return nil, fmt.Errorf("Unknown QOS flag: %s", f)
					}
				}
				q.Flags = flag
			default:
				return nil, fmt.Errorf("Unknown QOS field: %s", headers[i])
			}
		}
		qlist = append(qlist, q)
	}

	return qlist, nil
}

func ListQOS() ([]slurm.QOSRecord, error) {
	output, err := sacctmgrHander.ListQOS()
	if err != nil {
		return nil, err
	}

	return parseQOSList(output)
}
