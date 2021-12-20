package cli

import (
	"fmt"
	"strings"

	"github.com/puttsk/go-slurm"
)

const SlurmDelimiter string = "|"

// Set default sacctmgr hander to `SacctMgrCLI`
var sacctmgrHander SacctMgrCLIHander = new(SacctMgrCLI)

// SetSacctmgrHander: set default handler for sacctmgr function
func SetSacctmgrHander(s SacctMgrCLIHander) {
	sacctmgrHander = s
}

func parseUserList(s string) ([]slurm.UserRecord, error) {
	// Assume sacctmgr comes with headers
	// Parse sacctmgr show user output and return a list of UserRecords

	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid input")
	}

	var ulist []slurm.UserRecord

	headers := strings.Split(lines[0], SlurmDelimiter)
	for _, l := range lines[1:] {
		var u slurm.UserRecord
		for i, v := range strings.Split(l, SlurmDelimiter) {
			switch headers[i] {
			case "Admin":
				switch v {
				case "None":
					u.AdminLevel = slurm.SlurmDBAdminNone
				case "Operator":
					u.AdminLevel = slurm.SlurmDBAdminOperator
				case "Administrator":
					u.AdminLevel = slurm.SlurmDBAdminSuperUser
				}
			case "Coord Accounts":
			case "Def Acct":
				u.DefaultAcct = v
			case "Def WCKey":
				u.DefaultWCKey = v
			case "User":
				u.Name = v
			}
		}
		ulist = append(ulist, u)
	}

	return ulist, nil
}

func ListUser() ([]slurm.UserRecord, error) {
	output, err := sacctmgrHander.ListUser()
	if err != nil {
		return nil, err
	}

	return parseUserList(output)
}
