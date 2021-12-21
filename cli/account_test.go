package cli_test

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/puttsk/go-slurm/cli"
	"github.com/puttsk/go-slurm/cli/mocks"
)

func TestListAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sacctmgrOutput := sacctmgrAccountOutput

	m := mocks.NewMockSacctMgrCLIHander(ctrl)
	m.EXPECT().ListAccount().Return(sacctmgrOutput, nil)

	cli.SetSacctmgrHander(m)

	a, err := cli.ListAccount()
	if err != nil {
		t.Error(err)
	}

	accountCount := len(strings.Split(strings.TrimSpace(sacctmgrOutput), "\n")) - 1

	if len(a) != accountCount {
		t.Errorf("Invalid number of accounts. Expect: %d, Actual: %d", accountCount, len(a))
	}

	if a[4].Name != "pre0003" {
		t.Errorf("Invalid account name parsing. Expect %s, Actual %s", "pre0003", a[4].Name)
	}

	if a[10].Description != "ncovcu" {
		t.Errorf("Invalid account description parsing. Expect %s, Actual %s", "ncovcu", a[10].Description)
	}

	if a[8].Organization != "pre0007" {
		t.Errorf("Invalid account organization parsing. Expect %s, Actual %s", "pre0007", a[10].Organization)
	}
}

const sacctmgrAccountOutput = `Account|Descr|Org|Coord Accounts
intern|intern|intern|
jcsse2021|jcsse2021|jcsse2021|
pre0001|cvdimg|pre0001|
pre0002|ffasr|pre0002|
pre0003|graph|pre0003|
pre0004|tarcfd|pre0004|
pre0005|medseg|pre0005|
pre0006|autest|pre0006|
pre0007|ufprsr|pre0007|
pre0008|emul|pre0008|
pre5001|ncovcu|pre5001|
pre5002|covnet|pre5002|
pre5003|aiir|pre5003|
pre5004|ctm|pre5004|
pre5005|hpcai|pre5005|
pre5006|afs|pre5006|
pre5007|furero|pre5007|
pre5008|ngscov|pre5008|
pre5009|bvoc|pre5009|
proj0102|parai|proj0102|
proj0103|tial|proj0103|
proj0104|backen|proj0104|
proj0105|hccat|proj0105|
proj0106|nocat|proj0106|
proj0107|mofcat|proj0107|
proj0108|hdo|proj0108|
proj0110|boemd|proj0110|
proj0111|horn|proj0111|
proj0112|egatnp|proj0112|
proj0113|nbt|proj0113|
proj0114|sgnome|proj0114|
proj0115|objdet|proj0115|
proj0116|aiuav|proj0116|
proj0117|vistts|proj0117|
proj0118|ptassy|proj0118|
proj0119|astif|proj0119|
proj0120|lstdto|proj0120|
proj0121|lstdic|proj0121|
proj0122|master|proj0122|
proj0123|snp|proj0123|
proj0124|anno|proj0124|
proj0125|muomic|proj0125|
proj0126|fishp|proj0126|
proj0127|mopcat|proj0127|
proj0128|ceramd|proj0128|
proj0129|ftco|proj0129|
proj0130|isadb|proj0130|
proj0131|busdlt|proj0131|
proj0132|exasc|proj0132|
proj0133|nlpmtt|proj0133|
proj0134|infra|proj0134|
proj0135|vmgni|proj0135|
proj0136|laccat|proj0136|
proj0137|gstest|proj0137|
proj0138|smim|proj0138|
proj0139|mnps|proj0139|
proj0140|pmucrr|proj0140|
proj0141|ffasr|proj0141|
proj0142|abdul|proj0142|
proj0143|qcdcs|proj0143|
proj0144|boropn|proj0144|
proj0145|carban|proj0145|
proj0146|gemang|proj0146|
`
