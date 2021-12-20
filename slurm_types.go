package slurm

// INFINITE is used to identify unlimited configurations,
// eg. the maximum count of nodes any job may use in some partition
//slurm.h:130
const (
	Infinite8  uint8  = 0xff               //#define	INFINITE8  (0xff)
	Infinite16 uint16 = 0xffff             //#define	INFINITE16 (0xffff)
	Infinite   uint32 = 0xffffffff         //#define	INFINITE   (0xffffffff)
	Infinite64 uint64 = 0xffffffffffffffff //#define	INFINITE64 (0xffffffffffffffff)

	NoVal8       uint8  = 0xfe               //#define NO_VAL8    (0xfe)
	NoVal16      uint16 = 0xfffe             //#define NO_VAL16   (0xfffe)
	NoVal        uint32 = 0xfffffffe         //#define NO_VAL     (0xfffffffe)
	NoVal64      uint64 = 0xfffffffffffffffe //#define NO_VAL64   (0xfffffffffffffffe)
	NoConsumeVal uint64 = 0xfffffffffffffffd //#define NO_CONSUME_VAL64 (0xfffffffffffffffd)
)

//slurm.h:2699
const (
	PreemptModeOff     uint16 = 0x0000 /* disable job preemption */
	PreemptModeSuspend uint16 = 0x0001 /* suspend jobs to preempt */
	PreemptModeRequeue uint16 = 0x0002 /* requeue or kill jobs to preempt */

	PreemptModeCancel  uint16 = 0x0008 /* always cancel the job */
	PreemptModeCondOff uint16 = 0x0010 /* represents PREEMPT_MODE_OFF in list*/
	PreemptModeGang    uint16 = 0x8000 /* enable gang scheduling */
)

//Define QOS flags
//slurm.h: 123
const (
	QOSFlagBase   uint64 = 0x0fffffff //#define	QOS_FLAG_BASE                0x0fffffff
	QOSFlagNotset uint64 = 0x10000000 //#define	QOS_FLAG_NOTSET              0x10000000
	QOSFlagAdd    uint64 = 0x20000000 //#define	QOS_FLAG_ADD                 0x20000000
	QOSFlagRemove uint64 = 0x40000000 //#define	QOS_FLAG_REMOVE              0x40000000

	QOSFlagPartMinNode      uint64 = uint64(1) << 0 //#define	QOS_FLAG_PART_MIN_NODE       SLURM_BIT(0)
	QOSFlagPartMaxNode      uint64 = uint64(1) << 1 // #define	QOS_FLAG_PART_MAX_NODE       SLURM_BIT(1)
	QOSFlagPartTimeLimit    uint64 = uint64(1) << 2 //#define	QOS_FLAG_PART_TIME_LIMIT     SLURM_BIT(2)
	QOSFlagEnforceUsageTRES uint64 = uint64(1) << 3 //#define	QOS_FLAG_ENFORCE_USAGE_THRES SLURM_BIT(3)
	QOSFlagNoReserve        uint64 = uint64(1) << 4 // #define	QOS_FLAG_NO_RESERVE          SLURM_BIT(4)
	QOSFlagReqResv          uint64 = uint64(1) << 5 //#define	QOS_FLAG_REQ_RESV            SLURM_BIT(5)
	QOSFlagDenyLimit        uint64 = uint64(1) << 6 //#define	QOS_FLAG_DENY_LIMIT          SLURM_BIT(6)
	QOSFlagOverPartQOS      uint64 = uint64(1) << 7 // #define	QOS_FLAG_OVER_PART_QOS       SLURM_BIT(7)
	QOSFlagNoDecay          uint64 = uint64(1) << 8 // #define	QOS_FLAG_NO_DECAY            SLURM_BIT(8)
	QOSFlagUsageFactorSafe  uint64 = uint64(1) << 9 // #define	QOS_FLAG_USAGE_FACTOR_SAFE   SLURM_BIT(9)
)

//#define SLURM_BIT(offset) ((uint64_t)1 << offset)
//slurm.h: 260
func SlurmBit(offset int) uint64 {
	var bit uint64 = 1
	return (bit << uint64(offset))
}
