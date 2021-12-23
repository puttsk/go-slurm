package cli_test

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/puttsk/go-slurm/cli"
	"github.com/puttsk/go-slurm/cli/mocks"
)

func TestListAssoc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sacctmgrOutput := sacctmgrAssocOutput

	m := mocks.NewMockSacctMgrCLIHander(ctrl)
	m.EXPECT().ListAssoc().Return(sacctmgrOutput, nil)

	cli.SetSacctmgrHander(m)

	a, err := cli.ListAssoc()
	if err != nil {
		t.Error(err)
	}

	assocCount := len(strings.Split(strings.TrimSpace(sacctmgrOutput), "\n")) - 1

	if len(a) != assocCount {
		t.Errorf("Invalid number of accounts. Expect: %d, Actual: %d", assocCount, len(a))
	}

	if a[3].User != "jsuwanim" {
		t.Errorf("Invalid assoc name parsing. Expect %s, Actual %s", "jsuwanim", a[3].User)
	}

	if a[10].Acct != "pre0001" {
		t.Errorf("Invalid assoc account parsing. Expect %s, Actual %s", "pre0001", a[10].Acct)
	}
}

const sacctmgrAssocOutput = `ID|User|Account|Cluster|Def QOS|Share|GrpTRESMins|GrpTRESRunMins|GrpTRES|GrpJobs|GrpJobsAccrue|GrpSubmit|GrpWall|MaxTRES|MaxTRESMins|MaxTRESPerNode|MaxJobs|MaxJobsAccrue|MaxSubmit|MaxWall|QOS|QOS_RAW|Par ID|Par Name|Partition|Priority|LFT|RGT
1||root|tara||1|||||||||||||||normal|1|0||||1|2002
2|root|root|tara||1|||||||||||||||normal|1|1||||2000|2001
306||intern|tara|intern|1|||||||||||||||intern|99|1|root|||1198|1209
308|jsuwanim|intern|tara|intern|1|||||||||||||||intern|99|306||||1205|1206
840|nnampan|intern|tara|intern|1|||||||||||||||intern|99|306||||1199|1200
307|pteerata|intern|tara|intern|1|||||||||||||||intern|99|306||||1207|1208
309|vjarerat|intern|tara|intern|1|||||||||||||||intern|99|306||||1203|1204
812|wtongpra|intern|tara|intern|1|||||||||||||||intern|99|306||||1201|1202
901||jcsse2021|tara|jcsse2021|1|||||||||||||||jcsse2021|274|1|root|||144|225
231||pre0001|tara|pre0001|1|||||||||||||||pre0001|86|1|root|||1418|1429
232|ichatnun|pre0001|tara|pre0001|1|||||||||||||||pre0001|86|231||||1427|1428
235|kpimchar|pre0001|tara|pre0001|1|||||||||||||||pre0001|86|231||||1421|1422
234|mkunaset|pre0001|tara|pre0001|1|||||||||||||||pre0001|86|231||||1423|1424
236|mliangru|pre0001|tara|pre0001|1|||||||||||||||pre0001|86|231||||1419|1420
233|steerapi|pre0001|tara|pre0001|1|||||||||||||||pre0001|86|231||||1425|1426
289||pre0002|tara|pre0002|1|||||||||||||||pre0002|94|1|root|||1286|1293
302|kthangth|pre0002|tara|pre0002|1|||||||||||||||pre0002|94|289||||1289|1290
290|sboonkla|pre0002|tara|pre0002|1|||||||||||||||pre0002|94|289||||1291|1292
364|vchunwij|pre0002|tara|pre0002|1|||||||||||||||pre0002|94|289||||1287|1288
510||pre0003|tara|pre0003|1|||||||||||||||pre0003|141|1|root|||900|909
511|psakdhna|pre0003|tara|pre0003|1|||||||||||||||pre0003|141|510||||907|908
832|pyoovidh|pre0003|tara|pre0003|1|||||||||||||||pre0003|141|510||||901|902
512|rausavar|pre0003|tara|pre0003|1|||||||||||||||pre0003|141|510||||905|906
513|ssrisawa|pre0003|tara|pre0003|1|||||||||||||||pre0003|141|510||||903|904
763||pre0004|tara|pre0004|1|||||||||||||||pre0004|253|1|root|||402|407
779|njatusri|pre0004|tara|pre0004|1|||||||||||||||pre0004|253|763||||403|404
764|pchoopan|pre0004|tara|pre0004|1|||||||||||||||pre0004|253|763||||405|406
773||pre0005|tara|pre0005|1|||||||||||||||pre0005|256|1|root|||378|385
813|apiyatum|pre0005|tara|pre0005|1|||||||||||||||pre0005|256|773||||379|380
774|nwongwae|pre0005|tara|pre0005|1|||||||||||||||pre0005|256|773||||383|384
775|smarukat|pre0005|tara|pre0005|1|||||||||||||||pre0005|256|773||||381|382
784||pre0006|tara|pre0006|1|||||||||||||||pre0006|259|1|root|||360|365
787|srujikie|pre0006|tara|pre0006|1|||||||||||||||pre0006|259|784||||363|364
788|ssooksat|pre0006|tara|pre0006|1|||||||||||||||pre0006|259|784||||361|362
816||pre0007|tara|pre0007|1|||||||||||||||pre0007|266|1|root|||300|309
987|ssitdhip|pre0007|tara|pre0007|1|||||||||||||||pre0007|266|816||||301|302
833|tlodkaew|pre0007|tara|pre0007|1|||||||||||||||pre0007|266|816||||305|306
817|wchonnap|pre0007|tara|pre0007|1|||||||||||||||pre0007|266|816||||307|308
968|wnarongw|pre0007|tara|pre0007|1|||||||||||||||pre0007|266|816||||303|304
472||pre0008|tara|pre0008|1|||||||||||||||pre0008|131|1|root|||962|967
995|vjarerat|pre0008|tara|pre0008|1|||||||||||||||pre0008|131|472||||963|964
251||pre5001|tara|pre5001|1|||||||||||||||pre5001|89|1|root|||1344|1371
265|bnutho|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1361|1362
255|crungnim|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1367|1368
1028|kkanjama|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1345|1346
950|ksanacha|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1353|1354
951|nkongtaw|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1351|1352
952|npattara|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1349|1350
268|nsuvanna|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1355|1356
267|pmahalap|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1357|1358
953|rmendez|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1347|1348
264|shannong|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1363|1364
252|trungrot|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1369|1370
266|tsomboon|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1359|1360
256|vjarerat|pre5001|tara|pre5001|1|||||||||||||||pre5001|89|251||||1365|1366
269||pre5002|tara|pre5002|1|||||||||||||||pre5002|92|1|root|||1304|1327
288|ahuang|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1305|1306
275|awangwiw|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1315|1316
273|ebatty|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1319|1320
274|isensorn|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1317|1318
285|kjoonlas|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1309|1310
276|kkumporn|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1313|1314
272|nkotanan|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1321|1322
270|tchookaj|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1325|1326
282|tchookaj1|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1311|1312
271|tkochaka|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1323|1324
286|wmanasat|pre5002|tara|pre5002|1|||||||||||||||pre5002|92|269||||1307|1308
299||pre5003|tara|pre5003|1|||||||||||||||pre5003|96|1|root|||1260|1267
446|njatusr|pre5003|tara|pre5003|1|||||||||||||||pre5003|96|299||||1261|1262
440|njatusri|pre5003|tara|pre5003|1|||||||||||||||pre5003|96|299||||1263|1264
301|pchoopan|pre5003|tara|pre5003|1|||||||||||||||pre5003|96|299||||1265|1266
300||pre5004|tara|pre5004|1|||||||||||||||pre5004|97|1|root|||1226|1259
339|achaiyar|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1233|1234
338|dwichada|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1235|1236
335|ieksinch|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1241|1242
342|krojvibo|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1227|1228
341|ksripani|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1229|1230
336|ltongsin|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1239|1240
311|ncheewar|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1255|1256
329|nnupairo|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1251|1252
332|nsethasa|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1247|1248
333|rlaohasu|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1245|1246
337|spiwluan|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1237|1238
312|spromon|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1253|1254
334|ssomsaku|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1243|1244
310|tachalak|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1257|1258
340|tchalida|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1231|1232
331|vmuangsi|pre5004|tara|pre5004|1|||||||||||||||pre5004|97|300||||1249|1250
377||pre5005|tara|pre5005|1|||||||||||||||pre5005|107|1|root|||1096|1127
383|amahawan|pre5005|tara|pre5005|1|||||||||||||||pre5005|107|377||||1115|1116
`
