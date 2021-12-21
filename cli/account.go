package cli

import (
	"fmt"
	"strings"

	"github.com/puttsk/go-slurm"
)

func parseAccountList(s string) ([]slurm.AccountRecord, error) {
	// Assume sacctmgr comes with headers
	// Parse sacctmgr show account output and return a list of AccountRecords

	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid input")
	}

	var alist []slurm.AccountRecord

	headers := strings.Split(lines[0], SlurmDelimiter)
	for _, l := range lines[1:] {
		var a slurm.AccountRecord
		for i, v := range strings.Split(l, SlurmDelimiter) {
			switch headers[i] {
			case "Account":
				a.Name = v
			case "Descr":
				a.Description = v
			case "Org":
				a.Organization = v
			case "Coord Accounts":
				c := strings.TrimSpace(v)
				a.Coordinators = strings.Split(c, ",")
			}
		}
		alist = append(alist, a)
	}

	return alist, nil
}

func ListAccount() ([]slurm.AccountRecord, error) {
	output, err := sacctmgrHander.ListAccount()
	if err != nil {
		return nil, err
	}

	return parseAccountList(output)
}
