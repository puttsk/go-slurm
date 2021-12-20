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

	cli.SetSacctmgrHander(m)

	u, err := cli.ListUser()
	if err != nil {
		t.Error(err)
	}

	userCount := len(strings.Split(strings.TrimSpace(sacctmgrUserOutput), "\n")) - 1

	if len(u) != userCount {
		t.Errorf("Invalid number of users. Expect: %d, Actual: %d", userCount, len(u))
	}
}

func TestListQOS(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSacctMgrCLIHander(ctrl)
	m.EXPECT().ListQOS().Return(sacctmgrQOSOutput, nil)

	cli.SetSacctmgrHander(m)

	q, err := cli.ListQOS()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", q)

	qosCount := len(strings.Split(strings.TrimSpace(sacctmgrQOSOutput), "\n")) - 1

	if len(q) != qosCount {
		t.Errorf("Invalid number of QoS. Expect: %d, Actual: %d", qosCount, len(q))
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

const sacctmgrQOSOutput = `ID|Name|Descr|GraceTime|GrpTRESMins|GrpTRESRunMins|GrpTRES|GrpJobs|GrpJobsAccrue|GrpSubmit|GrpWall|MaxTRESMins|MaxTRESPA|MaxTRES|MaxTRESPerNode|MaxTRESPU|MaxJobsPA|MaxJobsAccruePA|MaxJobsPU|MaxJobsAccruePU|MaxSubmitPA|MaxSubmitPU|MaxWall|MinPrioThres|MinTRES|Preempt|PreemptExemptTime|PreemptMode|Priority|UsageFactor|Flags
1|normal|Normal QOS default|00:00:00||||||||||||||||||||||||cluster|0|1.000000|
7|proj0104|backen|00:00:00|billing=22222||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
8|thaisc|thaisc|00:00:00|billing=100001002|||||||||||||||||||||||cluster|0|1.000000|NoDecay
10|proj0102|parai|00:00:00|billing=71111111||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
12|proj0103|tial|00:00:00|billing=3500000||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
13|proj0106|nocat|00:00:00|billing=244494443||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
14|proj0105|hccat|00:00:00|billing=66666666||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
15|proj0108|hdo|00:00:00|billing=88888888||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
16|proj0107|mofcat|00:00:00|billing=122222221||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
51|proj0112|egatnp|00:00:00|billing=31915009||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
52|proj0111|horn|00:00:00|billing=22222222||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
53|proj0114|sgnome|00:00:00|billing=22222222||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
54|proj0110|boemd|00:00:00|billing=88888888||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
55|proj0115|objdet|00:00:00|billing=22222222||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
56|proj0113|nbt|00:00:00|billing=1108050099||||||||cpu=1500||||100|10||||||||||cluster|0|1.000000|NoDecay
57|proj0116|aiuav|00:00:00|billing=23333333||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
58|proj0117|vistts|00:00:00|billing=144444443||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
59|tutorial|tutorial|00:00:00|billing=613388||||||||cpu=1000|||||||||0||||||cluster|0|1.000000|NoDecay
60|proj0118|ptassy|00:00:00|billing=15565855||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
61|proj0120|lstdto|00:00:00|billing=11111111||||||||cpu=1000||||100|||||||||||cluster|0|1.000000|NoDecay
62|proj0119|astif|00:00:00|billing=6666666||||||||cpu=1000||||100|||||0||||||cluster|0|1.000000|NoDecay
63|proj0121|lstdic|00:00:00|billing=22222222||||||||cpu=1000||||100|||||||||||cluster|0|1.000000|NoDecay
65|proj0122|master|00:00:00|billing=11111111||||||||cpu=1000||||100|||||||||||cluster|0|1.000000|NoDecay
66|proj0123|snp|00:00:00|billing=22222221||||||||cpu=1000||||100|||||0||||||cluster|0|1.000000|NoDecay
67|proj0124|anno|00:00:00|billing=11111110||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
68|proj0125|muomic|00:00:00|billing=4444444||||||||cpu=1000||||100|||||0||||||cluster|0|1.000000|NoDecay
71|proj0126|fishp|00:00:00|billing=4444444||||||||cpu=1000||||100|||||||||||cluster|0|1.000000|NoDecay
72|proj0127|mopcat|00:00:00|billing=77777777||||||||cpu=1000||||100|||||||||||cluster|0|1.000000|NoDecay
77|proj0128|ceramd|00:00:00|billing=2222222||||||||cpu=1000||||100|||||0||||||cluster|0|1.000000|NoDecay
78|proj0129|ftco|00:00:00|billing=166666666||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
80|devel|devel|00:00:00|||||||||cpu=2000|||||||||||||||cluster|0|1.000000|
82|proj0130|isadb|00:00:00|billing=666666||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
83|proj0131|busdlt|00:00:00|billing=133309221||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
84|proj0132|exasc|00:00:00|billing=33333333||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
85|proj0133|nlpmtt|00:00:00|billing=2222222||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
86|pre0001|cvdimg|00:00:00|billing=5000000||||||||cpu=1000||||100|10||||0||||||cluster|0|1.000000|NoDecay
87|proj5001|oniomx|00:00:00|billing=18196261||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
88|proj5002|scisut|00:00:00|billing=156607110||||||||cpu=1000||||100|10||||||||||cluster|0|1.000000|NoDecay
`
