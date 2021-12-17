package cli_test

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/puttsk/go-slurm/cli"
	"github.com/puttsk/go-slurm/cli/mocks"
)

func TestListUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSacctMgrCLIHander(ctrl)
	m.EXPECT().ListUser().Return(sacctmgrUserOutput, nil)

	u, err := cli.ListUser(m)
	if err != nil {
		t.Error(err)
	}

	userCount := len(strings.Split(strings.TrimSpace(sacctmgrUserOutput), "\n")) - 1

	if len(u) != userCount {
		t.Errorf("Invalid number of users. Expect: %d, Actual: %d", userCount, len(u))
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
