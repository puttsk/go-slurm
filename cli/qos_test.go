package cli_test

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/puttsk/go-slurm"
	"github.com/puttsk/go-slurm/cli"
	"github.com/puttsk/go-slurm/cli/mocks"
)

func TestGenerateQOSParams(t *testing.T) {
	type QOSTest struct {
		q      slurm.QOSRecord
		expect string
	}

	var q1 slurm.QOSRecord
	q1.Init()
	q1.Description = "test001"
	q1.Flags = slurm.QOSFlagNoDecay

	var q2 slurm.QOSRecord
	q2.Init()
	q2.Description = "test002"
	q2.Flags = slurm.QOSFlagNoReserve | slurm.QOSFlagPartMaxNode | slurm.QOSFlagNoDecay
	q2.GraceTime = slurm.Infinite

	var q3 slurm.QOSRecord
	q3.Init()
	q3.Description = "test003"
	q3.Flags = slurm.QOSFlagNoDecay
	q3.MaxJobsAccruePA = 100
	q3.MaxSubmitJobsPA = 1000
	q3.GrpTRESMins = "billing=1000000"

	testcases := []QOSTest{
		{
			q:      q1,
			expect: "Description=test001 Flags=NoDecay",
		},
		{
			q:      q2,
			expect: "Description=test002 Flags=NoDecay,NoReserve,PartitionMaxNodes GraceTime=-1",
		},
		{
			q:      q3,
			expect: "Description=test003 Flags=NoDecay GrpTRESMins=billing=1000000 MaxJobsAccruePerAccount=1000 MaxSubmitJobsPerAccount=1000",
		},
	}

	for _, c := range testcases {
		actual, err := cli.GenerateQOSCmdParams(c.q)
		if err != nil {
			t.Error(err)
		}
		if actual != c.expect {
			t.Errorf("Invalid QOS params: expected: \"%s\", actual: \"%s\"\nQOSRecord:%+v\n", c.expect, actual, c.q)
		}
	}
}

func TestListQOS(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sacctmgrOutput := sacctmgrQOSOutput

	m := mocks.NewMockSacctMgrCLIHander(ctrl)
	m.EXPECT().ListQOS().Return(sacctmgrOutput, nil)

	cli.SetSacctmgrHander(m)

	q, err := cli.ListQOS()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v\n", q)

	qosCount := len(strings.Split(strings.TrimSpace(sacctmgrOutput), "\n")) - 1

	if len(q) != qosCount {
		t.Errorf("Invalid number of QoS. Expect: %d, Actual: %d", qosCount, len(q))
	}

	if q[3].Name != "proj0102" {
		t.Errorf("Invalid qos name parsing. Expect %s, Actual %s", "proj0102", q[3].Name)
	}

	if q[5].GraceTime != 0 {
		t.Errorf("Invalid qos grace time parsing. Expect %d, Actual %d", 0, q[5].GraceTime)
	}

	if q[10].GrpTRESMins != "billing=22222222" {
		t.Errorf("Invalid qos GrpTRESMins parsing. Expect %s, Actual %s", "billing=22222222", q[10].GrpTRESMins)
	}

	if q[11].GrpJobs != slurm.Infinite {
		t.Errorf("Invalid qos GrpJobs parsing. Expect 0x%d, Actual 0x%d", slurm.Infinite, q[11].GrpJobs)
	}

	if (q[15].Flags & slurm.QOSFlagNoDecay) == 0 {
		t.Errorf("Invalid qos flags parsing. Expect %X, Actual %X", slurm.QOSFlagNoDecay, q[15].Flags)
	}
}

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
