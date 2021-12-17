package slurm

// slurmdb_admin_level_t
// slurm/slurmdb.h: 47
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
// slurm/slurmdb.h: 1187
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
