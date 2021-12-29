package slurm

import (
	"log"
)

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

	initialized bool //Check if QOSRecord is initalized
}

//extern void slurmdb_init_qos_rec(slurmdb_qos_rec_t *qos, bool free_it, uint32_t init_val)
//src/common/slurmdb_defs.c
func (q *QOSRecord) Init() {
	q.initialized = true
	q.Flags = QOSFlagNotset
	q.GraceTime = NoVal
	q.PreemptMode = NoVal16
	q.PreemptExemptTime = NoVal
	q.Priority = NoVal

	q.GrpJobs = NoVal
	q.GrpJobsAccrue = NoVal
	q.GrpSubmitJobs = NoVal
	q.GrpWall = NoVal

	q.MaxJobsPA = NoVal
	q.MaxJobsPU = NoVal
	q.MaxJobsAccruePA = NoVal
	q.MaxJobsAccruePU = NoVal
	q.MinPrioThresh = NoVal
	q.MaxSubmitJobsPA = NoVal
	q.MaxSubmitJobsPU = NoVal
	q.MaxWallPJ = NoVal

	q.UsageFactor = float64(NoVal)
	//qos->usage_thres = (double)init_val;
	q.LimitFactor = float64(NoVal)
}

func (q *QOSRecord) SetFlagOp(op string) {
	if !q.initialized {
		log.Panicf("QOSRecord is used without initialization")
	}
	switch op {
	case "+":
		q.Flags = q.Flags & ^QOSFlagRemove
		q.Flags = q.Flags | QOSFlagAdd
	case "-":
		q.Flags = q.Flags & ^QOSFlagAdd
		q.Flags = q.Flags | QOSFlagRemove
	case "":
		q.Flags = q.Flags & ^QOSFlagRemove
		q.Flags = q.Flags & ^QOSFlagAdd
	default:
		log.Panicf("Unknown operation: %s", op)
	}
}

func (q QOSRecord) GetFlagOp() string {
	if !q.initialized {
		log.Panicf("QOSRecord is used without initialization")
	}
	if (q.Flags & QOSFlagAdd) > 0 {
		return "+"
	}
	if (q.Flags & QOSFlagRemove) > 0 {
		return "-"
	}
	return ""
}

func (q QOSRecord) GetFlag(flag string) bool {
	if !q.initialized {
		log.Panicf("QOSRecord is used without initialization")
	}

	switch flag {
	case "DenyOnLimit":
		return (q.Flags & QOSFlagDenyLimit) > 0
	case "EnforceUsageThreshold":
		return (q.Flags & QOSFlagEnforceUsageTRES) > 0
	case "NoDecay":
		return (q.Flags & QOSFlagNoDecay) > 0
	case "NoReserve":
		return (q.Flags & QOSFlagNoReserve) > 0
	case "PartitionMaxNodes":
		return (q.Flags & QOSFlagPartMaxNode) > 0
	case "PartitionMinNodes":
		return (q.Flags & QOSFlagPartMinNode) > 0
	case "OverPartQOS":
		return (q.Flags & QOSFlagOverPartQOS) > 0
	case "PartitionTimeLimit":
		return (q.Flags & QOSFlagPartTimeLimit) > 0
	case "RequiresReservation":
		return (q.Flags & QOSFlagReqResv) > 0
	case "UsageFactorSafe":
		return (q.Flags & QOSFlagUsageFactorSafe) > 0
	}
	return false
}

// #define SLURMDB_FS_USE_PARENT 0x7FFFFFFF
// slurm/slurmdb.h:219
const SlurmDBFSUseParent uint32 = 0x7FFFFFFF

// slurmdb_assoc_rec
// slurm/slurmdb.h:447
type AssocRecord struct {
	//List accounting_list; /* list of slurmdb_accounting_rec_t *'s */
	Acct string //char *acct;		   /* account/project associated to * assoc */
	//struct slurmdb_assoc_rec *assoc_next; /* next assoc with same hash index based off the account/user DOESN'T GET PACKED */
	//struct slurmdb_assoc_rec *assoc_next_id; /* next assoc with same hash index DOESN'T GET PACKED */
	//slurmdb_bf_usage_t *bf_usage; /* data for backfill scheduler, (DON'T PACK) */
	Cluster            string   //char *cluster;		   /* cluster associated to association */
	DefQOSID           uint32   //uint32_t def_qos_id;       /* Which QOS id is this associations default */
	Flags              uint16   //uint16_t flags;            /* various flags see ASSOC_FLAG_* */
	GrpJobs            uint32   //uint32_t grp_jobs;	   /* max number of jobs the underlying group of associations can run at one time */
	GrpJobsAccrue      uint32   //uint32_t grp_jobs_accrue;  /* max number of jobs the underlying group of associations can have accruing priority at one time */
	GrpSubmitJobs      uint32   //uint32_t grp_submit_jobs;  /* max number of jobs the underlying group of associations can submit at one time */
	GrpTRES            string   //char *grp_tres;            /* max number of cpus the underlying group of associations can allocate at one time */
	GrpTRESCtld        uint64   //uint64_t *grp_tres_ctld;   /* grp_tres broken out in an array based off the ordering of the total  number of TRES in the system (DON'T PACK) */
	GrpTRESMins        string   //char *grp_tres_mins;       /* max number of cpu minutes the underlying group of associations can run for */
	GrpTRESMinsCtld    uint64   //uint64_t *grp_tres_mins_ctld; /* grp_tres_mins broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	GrpTRESRunMins     string   //char *grp_tres_run_mins;   /* max number of cpu minutes the underlying group of assoiciations can having running at one time */
	GrpTRESRunMinsCtld uint64   //uint64_t *grp_tres_run_mins_ctld; /* grp_tres_run_mins broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	GrpWall            uint32   //uint32_t grp_wall;         /* total time in hours the underlying group of associations can run for */
	ID                 uint32   //uint32_t id;		   /* id identifing a combination of user-account-cluster(-partition) */
	IsDef              uint16   //uint16_t is_def;           /* Is this the users default assoc/acct */
	Lft                uint32   //uint32_t lft;		   /* lft used for grouping sub associations and jobs as a left most container used with rgt */
	MaxJobs            uint32   //uint32_t max_jobs;	   /* max number of jobs this association can run at one time */
	MaxJobsAccrue      uint32   //uint32_t max_jobs_accrue;  /* max number of jobs this association can have accruing priority time.*/
	MaxSubmitJobs      uint32   //uint32_t max_submit_jobs;  /* max number of jobs that can be submitted by association */
	MaxTRESMinsPJ      string   //char *max_tres_mins_pj;    /* max number of cpu minutes this association can have per job */
	MaxTRESMinsPJCtld  uint64   //uint64_t *max_tres_mins_ctld; /* max_tres_mins broken out in an array based off the ordering of the  total number of TRES in the system (DON'T PACK) */
	MaxTRESRunMins     string   //char *max_tres_run_mins;   /* max number of cpu minutes this association can having running at one time */
	MaxTRESRunMinsCtld uint64   //uint64_t *max_tres_run_mins_ctld; /* max_tres_run_mins broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	MaxTRESPJ          string   //char *max_tres_pj;         /* max number of cpus this association can allocate per job */
	MaxTRESPJCtld      uint64   //uint64_t *max_tres_ctld;   /* max_tres broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	MaxTRESPN          string   //char *max_tres_pn;         /* max number of TRES this association can allocate per node */
	MaxTRESPNCtld      uint64   //uint64_t *max_tres_pn_ctld;   /* max_tres_pn broken out in an array based off the ordering of the total number of TRES in the system (DON'T PACK) */
	MaxWallPJ          uint32   //uint32_t max_wall_pj;      /* longest time this association can run a job */
	MinPrioThresh      uint32   //uint32_t min_prio_thresh;  /* Don't reserve resources for pending jobs unless they have a priority equal to or higher than this. */
	ParentAcct         string   //char *parent_acct;	   /* name of parent account */
	ParentID           uint32   //uint32_t parent_id;	   /* id of parent account */
	Partition          string   //char *partition;	   /* optional partition in a cluster associated to association */
	Priority           uint32   //uint32_t priority;	   /* association priority */
	QOSList            []string //List qos_list;             /* list of char * */

	Rgt uint32 //uint32_t rgt;		   /* rgt used for grouping sub associations and jobs as a right most container used with lft */

	SharesRaw uint32 //uint32_t shares_raw;	   /* number of shares allocated to association */

	UID uint32 //uint32_t uid;		   /* user ID */
	//slurmdb_assoc_usage_t *usage;
	User string //char *user;		   /* user associated to assoc */
	//slurmdb_user_rec_t *user_rec; /* Cache of user record soft ref - mem not managed here (DON'T PACK)*/
}
