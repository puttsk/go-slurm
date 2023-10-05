// Package slurm provides an interface for interacting with Slurm
package slurm

import "time"

// A JobInformation contains information of job in partitions, based on squeue command
type JobInformation struct {
	Account                string    `json:"account"`            // Account associated with the job
	AccrueTime             time.Time `json:"accrue_time"`        // Time job is eligible for running
	AdminComment           string    `json:"admin_comment"`      // Administrator's arbitrary comment
	AllocNode              string    `json:"alloc_node"`         // Local node making resource alloccation
	AllocSID               string    `json:"alloc_sid"`          // Local SID making resource alloc
	ArrayJobID             uint32    `json:"array_job_id"`       // Job ID of a job array or 0 if N/A
	ArrayTaskID            uint32    `json:"array_task_id"`      // Task ID of a job array
	AssociationID          uint32    `json:"assoc_id"`           // Association id for job
	BatchFlag              bool      `json:"batch_flag"`         // true if batch: queued job with script
	BatchHost              string    `json:"batch_host"`         // Name of host running batch script
	BoardsPerNode          uint16    `json:"boards_per_node"`    // Boards per node required by job
	BurstBuffer            string    `json:"burst_buffer"`       // Burst buffer specifications
	BurstBufferState       string    `json:"burst_buffer_state"` // Burst buffer state info
	Cluster                string    `json:"cluster"`            // Name of cluster that the job is on
	ClusterFeatures        string    `json:"cluster_features"`   // Comma separated list of required cluster features
	Command                string    `json:"command"`            // Command to be executed, built from submittedjob's argv
	Comment                string    `json:"comment"`            // Arbitrary comment
	Container              string    `json:"container"`          // OCI Container bundle path
	ContainerID            string    `json:"container_id"`       // OCI Container ID
	Contiguous             bool      `json:"contiguous"`         // true if job requires contiguous nodes
	CoresPerSocket         uint16    `json:"cores_per_socket"`   // Cores per socket required by job
	CoreSpec               uint16    `json:"core_spec"`          // Specialized core count
	CPUFrequency           uint32    `json:"cpu_freq"`           // Frequency of the allocated CPUs
	CPUsPerTask            uint16    `json:"cpus_per_task"`      // Number of processors required for each task
	CPUPerTRES             string    `json:"cpus_per_tres"`      // Semicolon delimited list of TRES=# values
	Deadline               time.Time `json:"deadline"`           // Deadline
	DelayBoot              uint32    `json:"delay_boot"`         // Delay boot for desired node state
	Dependency             string    `json:"dependency"`         // Job dependency
	DerivedEC              uint32    `json:"derived_ec"`         // Highest exit code of all job steps
	EligibleTime           time.Time `json:"eligible_time"`      // Time job is eligible for running
	EndTime                time.Time `json:"end_time"`           // Time of termination, actual or expected
	ExitCode               uint32    `json:"exit_code"`          // Exit code for job (status from wait call)
	Features               string    `json:"features"`           // Comma separated list of required features
	GroupID                uint32    `json:"group_id"`           // Group ID of the job
	GroupName              string    `json:"group_name"`         // Group Name of the job
	HeterogeneousJobID     uint32    `json:"het_job_id"`         // Job ID of the heterogeneous job leader
	HeterogeneousJobIDSet  string    `json:"het_job_id_set"`     // Expression identifying all components job IDs within a heterogeneous job
	HeterogeneousJobOffset uint32    `json:"het_job_offset"`     // Zero origin offset within a collection of heterogeneous job components.
	JobID                  uint32    `json:"job_id"`             // Job ID
	JobState               uint32    `json:"job_state"`          // State of the job, see enum job_states
	LastSchedEval          time.Time `json:"last_sched_eval"`    // Last time job was evaluated for scheduling
	Licenses               string    `json:"licenses"`           // Licenses required by the job
	MaxCPUs                uint32    `json:"max_cpus"`           // Maximum number of cpus usable by job
	MaxNodes               uint32    `json:"max_nodes"`          // Maximum number of nodes usable by job
	MinMemory              uint64    `json:"min_memory"`         // Minimum size of memory (in MB) requested by the job
	MinCPUs                uint32    `json:"min_cpus"`           // Minimum number of CPUs (processors) per node requested by the job.
	MinTmpDisk             uint32    `json:"min_tmp_disk"`       // Minimum size of temporary disk space (in MB) requested by the job.
	MinTime                uint32    `json:"min_time"`           // Minimum run time in minutes or INFINITE
	MCSLabel               string    `json:"mcs_label"`          // MCS label if mcs plugin in use
	MemPerTRES             string    `json:"mem_per_tres"`       // semicolon delimited list of TRES=# values
	Name                   string    `json:"name"`               // Name of the job
	Network                string    `json:"network"`            // Network specification
	Nodes                  string    `json:"nodes"`              // List of nodes allocated to job
	Nice                   uint32    `json:"nice"`               // Nice value (adjustment to a job's scheduling priority)
	TasksPerCore           uint16    `json:"ntasks_per_core"`    // Number of tasks to invoke on each core
	TasksPerTRES           uint16    `json:"ntasks_per_tres"`    // Number of tasks that can access each gpu
	TasksPerNode           uint16    `json:"ntasks_per_node"`    // Number of tasks to invoke on each node
	TasksPerSocket         uint16    `json:"ntasks_per_socket"`  // Number of tasks to invoke on each socket
	TasksPerBoard          uint16    `json:"ntasks_per_board"`   // Number of tasks to invoke on each board
	NumCPUs                uint32    `json:"num_cpus"`           // Minimum number of cpus required by job
	NumNodes               uint32    `json:"num_nodes"`          // Minimum number of nodes required by job
	NumTasks               uint32    `json:"num_tasks"`          // Requested task count
	Partition              string    `json:"partition"`          // Name of assigned partition
	Prefer                 string    `json:"prefer"`             // Comma separated list of soft features
	PreemptTime            time.Time `json:"preempt_time"`       // The preempt time for the job.
	Priority               uint64    `json:"priority"`           // Priority of the job
	Profile                uint32    `json:"profile"`            // Profile of the job
	QoS                    string    `json:"qos"`                // Quality of Service
	Reboot                 uint8     `json:"reboot"`             // Node reboot requested before start
	ReqNodes               string    `json:"req_nodes"`          // Comma separated list of required nodes
	ReqSwitch              uint32    `json:"req_switch"`         // Minimum number of switches
	Requeue                uint16    `json:"requeue"`            // Enable or disable job requeue option
	ResizeTime             time.Time `json:"resize_time"`        // Time of latest size change
	RestartCount           uint16    `json:"restart_cnt"`        // Count of job restarts
	Reservation            string    `json:"resv_name"`          // Reservation name
	SchedNodes             string    `json:"sched_nodes"`        // List of nodes scheduled to be used for job
	SocketsPerBoards       uint16    `json:"sockets_per_board"`  // Sockets per board required by job
	SocketsPerNodes        uint16    `json:"sockets_per_node"`   // Sockets per node required by job
	StartTime              time.Time `json:"start_time"`         // Time execution begins, actual or expected
	StateDescription       string    `json:"state_desc"`         // Optional details for state_reason
	StateReason            uint32    `json:"state_reason"`       // Reason job still pending or failed, see slurm.h:enum job_state_reason
	StdErr                 string    `json:"std_err"`            // Pathname of job's stderr file
	StdIn                  string    `json:"std_in"`             // Pathname of job's stdin file
	StdOut                 string    `json:"std_out"`            // Pathname of job's stdout file
	SubmitTime             time.Time `json:"submit_time"`        // Time of job submission
	SubpendTime            time.Time `json:"suspend_time"`       // Time job last suspended or resumed
	SystemComment          string    `json:"system_comment"`     // slurmctld's arbitrary comment
	TimeLimit              uint32    `json:"time_limit"`         // Maximum run time in minutes or INFINITE
	ThreadsPerCore         uint16    `json:"threads_per_core"`   // Threads per core required by job
	TRESBind               string    `json:"tres_bind"`          // Task to TRES binding directives
	TRESFrequency          string    `json:"tres_freq"`          // TRES frequency directives
	TRESPerJob             string    `json:"tres_per_job"`       // Semicolon delimited list of TRES=# values
	TRESPerNode            string    `json:"tres_per_node"`      // Semicolon delimited list of TRES=# values
	TRESPerSocket          string    `json:"tres_per_socket"`    // Semicolon delimited list of TRES=# values
	TRESPerTask            string    `json:"tres_per_task"`      // Semicolon delimited list of TRES=# values
	TRESRequest            string    `json:"tres_req_str"`       // TRES requested in the job
	TRESAlloc              string    `json:"tres_alloc_str"`     // TRES used in the job
	UserID                 uint32    `json:"user_id"`            // User the job runs as
	Username               string    `json:"user_name"`          // User name for a job or job step.
	Wait4Switch            uint32    `json:"wait4switch"`        // Maximum time to wait for minimum switches
	WCKey                  string    `json:"wckey"`              // WCKey for job
	WorkDIR                string    `json:"work_dir"`           // Pathname of working directory
}
