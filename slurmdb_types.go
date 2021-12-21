package slurm

// slurmdb_admin_level_t
// slurm/slurmdb.h:47
const (
	SlurmDBAdminNotset    uint16 = 0
	SlurmDBAdminNone      uint16 = 1
	SlurmDBAdminOperator  uint16 = 2
	SlurmDBAdminSuperUser uint16 = 3
)

// SLURMDB_USER_FLAG_
// slurm/slurmdb.h: 1182
const (
	SlurmDBUserFlagNone    uint32 = 0
	SlurmDBUserFlagDeleted uint32 = 1
)

// slurmdb_user_rec
// slurm/slurmdb.h:879
type UserRecord struct {
	AdminLevel   uint16 // uint16_t admin_level;
	DefaultAcct  string // char *default_acct;
	DefaultWCKey string // char *default_wckey;
	Flags        uint32 // uint32_t flags;		/* SLURMDB_USER_FLAG_* */
	Name         string // char *name;
	OldName      string // char *old_name;
	UID          uint32 // uint32_t uid;
	//AssocList  	   // List assoc_list; /* list of slurmdb_assoc_rec_t *'s */
	//BFUsage 		   // slurmdb_bf_usage_t *bf_usage;
	//CoordAccts	   // List coord_accts; /* list of slurmdb_coord_rec_t *'s */
	//WCKeyList  //List wckey_list; /* list of slurmdb_wckey_rec_t *'s */
}

// SLURMDB_ACCT_FLAG_
// slurm/slurmdb.h:364
const (
	SlurmDBAcctFlagNone    uint32 = 0
	SlurmDBAcctFlagDeleted uint32 = uint32(1) << 0
)

// slurmdb_account_rec_t
// slurm/slurmdb.h:369
type AccountRecord struct {
	AssocList    []string //List assoc_list; /* list of slurmdb_assoc_rec_t *'s */
	Coordinators []string //List coordinators; /* list of slurmdb_coord_rec_t *'s */
	Description  string   //char *description;
	Flags        uint32   //uint32_t flags; /* SLURMDB_ACCT_FLAG_* */
	Name         string   //char *name;
	Organization string   //char *organization;
}

// slurmdb_qos_rec_t
// slurm/slurmdb.h: 1187
type QOSRecord struct {
	Description        string //char *description;
	ID                 uint32 //uint32_t id;
	Flags              uint32 //uint32_t flags; /* flags for various things to enforce or override other limits */
	GraceTime          uint32 //uint32_t grace_time; /* preemption grace time in seconds*/
	GrpJobsAccrue      uint32 //uint32_t    grp_jobs_accrue /* max number of jobs this qos can have accruing priority time */
	GrpJobs            uint32 //uint32_t grp_jobs /* max number of jobs this qos can run at one time */
	GrpSubmitJobs      uint32 //uint32_t grp_submit_jobs /* max number of jobs this qos can submit at one time */
	GrpTRES            string //char          *grp_tres /* max number of tres this qos can allocate at one time */
	GrpTRESCtld        uint64 //uint64_t      *grp_tres_ctld /* grp_tres broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	GrpTRESMins        string //char *grp_tres_mins /* max number of tres minutes this qos can run for */
	GrpTRESMinsCtld    int64  //uint64_t      *grp_tres_mins_ctld /* grp_tres_mins broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	GrpTRESRunMins     string //char            *grp_tres_run_mins /* max number of tres minutes this qos can have running at one time */
	GrpTRESRunMinsCtld uint64 //uint64_t *grp_tres_run_mins_ctld /* grp_tres_run_mins broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	GrpWall            uint32 //uint32_t           grp_wall /* total time in hours this qos can run for */

	LimitFactor          float64 //double limit_factor /* factor to apply to tres_cnt for associations using this qos */
	MaxJobsPA            uint32  //uint32_t max_jobs_pa /* max number of jobs an account can run with this qos at one time */
	MaxJobsPU            uint32  //uint32_t max_jobs_pu /* max number of jobs a user can run with this qos at one time */
	MaxJobsAccruePA      uint32  //uint32_t max_jobs_accrue_pa /* max number of jobs an account can have accruing priority time */
	MaxJobsAccruePU      uint32  //uint32_t max_jobs_accrue_pu /* max number of jobs a user can have accruing priority time */
	MaxSubmitJobsPA      uint32  //uint32_t max_submit_jobs_pa /* max number of jobs an account can submit with this qos at once */
	MaxSubmitJobsPU      uint32  //uint32_t max_submit_jobs_pu /* max number of jobs a user can submit with this qos at once */
	MaxTRESMinsPJ        string  //char *max_tres_mins_pj /* max number of tres minutes this qos can have per job */
	MaxTRESMinsPJCtld    uint64  //uint64_t *max_tres_mins_pj_ctld /* max_tres_mins broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	MaxTRESPA            string  //char     *max_tres_pa /* max number of tres this QOS can allocate per account */
	MaxTRESPACtld        uint64  //uint64_t *max_tres_pa_ctld /* max_tres_pa broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	MaxTRESPJ            string  //char     *max_tres_pj /* max number of tres this qos can allocate per job */
	MaxTRESPJCtld        uint64  //uint64_t *max_tres_pj_ctld /* max_tres_pj broken out in an array based off the ordering of the total number of TRES in the system(DON'T PACK) */
	MaxTRESPN            string  //char     *max_tres_pn /* max number of tres this qos can allocate per job */
	MaxTRESPNCtld        uint64  //uint64_t *max_tres_pn_ctld /* max_tres_pj broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	MaxTRESPU            string  //char     *max_tres_pu /* max number of tres this QOS can allocate per user */
	MaxTRESPUCtld        uint64  //uint64_t *max_tres_pu_ctld /* max_tres broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	MaxTRESRunMinsPA     string  //char     *max_tres_run_mins_pa /* max number of tres minutes this qos can having running at one time per account, currently this doesn't do anything.*/
	MaxTRESRunMinsPACtld uint64  //uint64_t *max_tres_run_mins_pa_ctld /* max_tres_run_mins_pa broken out in an array based off the ordering of the total number of TRES in the system, currently this doesn't do anything. (DON'T PACK) */
	MaxTRESRunMinsPU     string  //char     *max_tres_run_mins_pu /* max number of tres minutes this qos can having running at one time, currently this doesn't do anything.*/
	MaxTRESRunMinsPUCtld uint64  //uint64_t *max_tres_run_mins_pu_ctld /* max_tres_run_mins_pu broken out in an array based off the ordering of the total number of TRES in the system, currently this doesn't do anything. (DON'T PACK) */
	MaxWallPJ            uint32  //uint32_t max_wall_pj /* longest time this qos can run a job */
	MinPrioThresh        uint32  //uint32_t min_prio_thresh /* Don't reserve resources for pending jobs unless they have a priority equal to or higher than this. */
	MinTRESPJ            string  //char     *min_tres_pj /* min number of tres a job can allocate with this qos */
	MinTRESPJCtld        uint64  //uint64_t *min_tres_pj_ctld /* min_tres_pj broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */

	Name              string   //char     *name
	PreemptBitstr     []uint32 //bitstr_t *preempt_bitstr /* other qos' this qos can preempt */
	PreemptList       []string //List     preempt_list    /* list of char *'s only used to add or change the other qos' this can preempt, when doing a get use the preempt_bitstr */
	PreemptMode       uint16   //uint16_t      preempt_mode        /* See PREEMPT_MODE_* in slurm/slurm.h */
	PreemptExemptTime uint32   //uint32_t      preempt_exempt_time /* Job run time before becoming eligible for preemption */
	Priority          uint32   //uint32_t priority /* ranged int needs to be a unint for heterogeneous systems */
	// Usage //slurmdb_qos_usage_t *usage       /* For internal use only, DON'T PACK */
	UsageFactor float64 //double  usage_factor /* factor to apply to usage in this qos */
	// UsageThres  float64 //double  usage_thres  /* percent of effective usage of an association when breached will deny pending and new jobs */
	// BlockUntil uint64 time_t blocked_until /* internal use only, DON'T PACK  */
}
