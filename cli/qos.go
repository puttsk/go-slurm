package cli

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/puttsk/go-slurm"
)

var qosFlags []string = []string{
	"DenyOnLimit",
	"EnforceUsageThreshold",
	"NoDecay",
	"NoReserve",
	"PartitionMaxNodes",
	"PartitionMinNodes",
	"OverPartQOS",
	"PartitionTimeLimit",
	"RequiresReservation",
	"UsageFactorSafe"}

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
		q.Init()
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

func GenerateQOSCmdParams(q slurm.QOSRecord) (string, error) {
	var cmdParams []string

	if q.Description != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("Description=%s", q.Description))
	}

	if q.Flags != 0 {
		var flags []string
		op := q.GetFlagOp()
		for _, f := range qosFlags {
			if q.GetFlag(f) {
				flags = append(flags, f)
			}
		}
		//TODO: Need a way to unset flags
		if len(flags) > 0 {
			cmdParams = append(cmdParams, fmt.Sprintf("Flags%s=%s", op, strings.Join(flags, ",")))
		}
	}

	if q.GraceTime != slurm.NoVal {
		pName := "GraceTime"
		if q.GraceTime == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.GraceTime))
		}
	}

	if q.GrpJobsAccrue != slurm.NoVal {
		pName := "GrpJobsAccrue"
		if q.GrpJobsAccrue == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.GrpJobsAccrue))
		}
	}

	if q.GrpJobs != slurm.NoVal {
		pName := "GrpJobs"
		if q.GrpJobs == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.GrpJobs))
		}
	}

	if q.GrpSubmitJobs != slurm.NoVal {
		pName := "GrpSubmitJobs"
		if q.GrpSubmitJobs == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.GrpSubmitJobs))
		}
	}

	if q.GrpTRES != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("GrpTRES=%s", q.GrpTRES))
	}

	if q.GrpTRESMins != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("GrpTRESMins=%s", q.GrpTRESMins))
	}

	if q.GrpTRESRunMins != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("GrpTRESRunMins=%s", q.GrpTRESRunMins))
	}

	if q.GrpWall != slurm.NoVal {
		pName := "GrpWall"
		if q.GrpSubmitJobs == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.GrpWall))
		}
	}

	if q.MaxJobsPA != slurm.NoVal {
		pName := "MaxJobsPerAccount"
		if q.MaxJobsPA == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.MaxJobsPA))
		}
	}

	if q.MaxJobsPU != slurm.NoVal {
		pName := "MaxJobsPerUser"
		if q.MaxJobsPU == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.MaxJobsPU))
		}
	}

	if q.MaxSubmitJobsPA != slurm.NoVal {
		pName := "MaxJobsAccruePerAccount"
		if q.MaxSubmitJobsPA == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.MaxSubmitJobsPA))
		}
	}

	if q.MaxJobsAccruePU != slurm.NoVal {
		pName := "MaxJobsAccruePerUser"
		if q.MaxJobsAccruePU == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.MaxJobsAccruePU))
		}
	}

	if q.MaxSubmitJobsPA != slurm.NoVal {
		pName := "MaxSubmitJobsPerAccount"
		if q.MaxSubmitJobsPA == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.MaxSubmitJobsPA))
		}
	}

	if q.MaxSubmitJobsPU != slurm.NoVal {
		pName := "MaxSubmitJobsPerUser"
		if q.MaxSubmitJobsPU == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.MaxSubmitJobsPU))
		}
	}

	if q.MaxTRESMinsPJ != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("MaxTRESMinsPerJob=%s", q.MaxTRESMinsPJ))
	}

	if q.MaxTRESPA != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("MaxTRESPerAccount=%s", q.MaxTRESPA))
	}

	if q.MaxTRESPJ != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("MaxTRESPerJob=%s", q.MaxTRESPJ))
	}

	if q.MaxTRESPN != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("MaxTRESPerNode=%s", q.MaxTRESPN))
	}

	if q.MaxTRESPU != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("MaxTRESPerUser=%s", q.MaxTRESPU))
	}

	if q.MaxTRESRunMinsPA != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("MaxTRESRunMinsPerAccount=%s", q.MaxTRESRunMinsPA))
	}

	if q.MaxTRESRunMinsPU != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("MaxTRESRunMinsPerUser=%s", q.MaxTRESRunMinsPU))
	}

	if q.MaxWallPJ != slurm.NoVal {
		pName := "MaxWallDurationPerJob"
		if q.MaxWallPJ == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.MaxWallPJ))
		}
	}

	if q.MinPrioThresh != slurm.NoVal {
		pName := "MinPrioThreshold"
		if q.MinPrioThresh == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.MinPrioThresh))
		}
	}

	if q.MinTRESPJ != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("MinTRESPerJob=%s", q.MinTRESPJ))
	}

	if q.Name != "" {
		cmdParams = append(cmdParams, fmt.Sprintf("Name=%s", q.Name))
	}

	if len(q.PreemptList) > 0 {
		cmdParams = append(cmdParams, fmt.Sprintf("Name=%s", strings.Join(q.PreemptList, ",")))
	}

	if q.PreemptMode != slurm.NoVal16 {
		log.Panicf("Setting 0reemptMode is not implemented")
	}

	if q.PreemptExemptTime != slurm.NoVal {
		pName := "PreemptExemptTime"
		if q.PreemptExemptTime == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.PreemptExemptTime))
		}
	}

	if q.Priority != slurm.NoVal {
		pName := "Priority"
		if q.Priority == slurm.Infinite {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, -1))
		} else {
			cmdParams = append(cmdParams, fmt.Sprintf("%s=%d", pName, q.Priority))
		}
	}

	return strings.Join(cmdParams, " "), nil
}

func ListQOS() ([]slurm.QOSRecord, error) {
	output, err := sacctmgrHander.ListQOS()
	if err != nil {
		return nil, err
	}

	return parseQOSList(output)
}
