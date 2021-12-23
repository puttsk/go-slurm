package cli_test

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/puttsk/go-slurm"
	"github.com/puttsk/go-slurm/cli"
	"github.com/puttsk/go-slurm/cli/mocks"
)

func TestListUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sacctmgrOutput := sacctmgrUserOutput

	m := mocks.NewMockSacctMgrCLIHander(ctrl)
	m.EXPECT().ListUser().Return(sacctmgrOutput, nil)

	cli.SetSacctmgrHander(m)

	u, err := cli.ListUser()
	if err != nil {
		t.Error(err)
	}

	userCount := len(strings.Split(strings.TrimSpace(sacctmgrOutput), "\n")) - 1

	if len(u) != userCount {
		t.Errorf("Invalid number of users. Expect: %d, Actual: %d", userCount, len(u))
	}

	if u[3].AdminLevel != slurm.SlurmDBAdminNone {
		t.Errorf("Invalid user admin level parsing. Expect %d, Actual %d", slurm.SlurmDBAdminNone, u[3].AdminLevel)
	}

	if u[4].AdminLevel != slurm.SlurmDBAdminOperator {
		t.Errorf("Invalid user admin level parsing. Expect %d, Actual %d", slurm.SlurmDBAdminOperator, u[4].AdminLevel)
	}

	if u[9].AdminLevel != slurm.SlurmDBAdminSuperUser {
		t.Errorf("Invalid user admin level parsing. Expect %d, Actual %d", slurm.SlurmDBAdminSuperUser, u[9].AdminLevel)
	}

	if u[11].DefaultAcct != "thaisc" {
		t.Errorf("Invalid user default account parsing. Expect %s, Actual %s", "thaisc", u[11].DefaultAcct)
	}

	if u[11].Name != "rtuchind" {
		t.Errorf("Invalid user default account parsing. Expect %s, Actual %s", "rtuchind", u[11].Name)
	}

}

const sacctmgrUserOutput = `Admin|Coord Accounts|Def Acct|Def WCKey|User
None||pre5006||aapai
None||trial0003||aaroonsr
None||proj5013||aasanakh
None||superai022||aauppaka
Operator||thaisc||abusaran
None||pre5004||achaiyar
None||superai061||adaowrae
None||proj5015||rmendez
None||trial0008||rmethach
Administrator||root||root
None||proj0108||rsalaeh
None||thaisc||rtuchind
None||proj0113||rwasitth
None||intern||wtongpra
None||thaisc||wtubtimt
Operator||thaisc||wudomsir
None||proj5014||yraksri
None||proj0133||ythu
None||proj5006||ywongcha
`
