package cli

const SlurmDelimiter string = "|"

// Set default sacctmgr hander to `SacctMgrCLI`
var sacctmgrHander SacctMgrCLIHander = new(SacctMgrCLI)

// SetSacctmgrHander: set default handler for sacctmgr function
func SetSacctmgrHander(s SacctMgrCLIHander) {
	sacctmgrHander = s
}
