// Package slurm provides an interface for interacting with Slurm
package slurm

// SQueueJobState defines slurm job state
type JobState string

const JobStateRunning JobState = "RUNNING" // Job is running

// A JobResourceAllocatedCores defines CPU cores allocated by a job
type JobResourceAllocatedCores struct {
	Cores map[string]string `json:"cores"`
}

// A JobResourceAllocatedNode defines node resources allocated by a job
type JobResourceAllocatedNode struct {
	NodeName        string                                 `json:"nodename"`
	CPUsUsed        uint32                                 `json:"cpus_used"`
	MemoryUsed      uint64                                 `json:"memory_used"`
	MemoryAllocated uint64                                 `json:"memory_allocated"`
	Sockets         []map[string]JobResourceAllocatedCores `json:"sockets"`
}

// A JobResources defines allocated resources of a job
type JobResources struct {
	Nodes          string                     `json:"nodes"`
	AllocatedCores uint32                     `json:"allocated_cores"`
	AllocatedHost  uint32                     `json:"allocated_hosts"`
	AllocatedNodes []JobResourceAllocatedNode `json:"allocated_nodes"`
}

// A Job contains job information. The designed is based on JSON structure returned by squeue command
type Job struct {
	Account                  string        `json:"account"`                    // Slurm account running this job
	AccrueTime               int64         `json:"accrue_time"`                // Unix timestamp of job accrue time
	AdminComment             string        `json:"admin_comment"`              // Comment from administator
	ArrayJobID               uint32        `json:"array_job_id"`               // Job ID of a job array or 0 if N/A
	ArrayTaskID              *uint32       `json:"array_task_id"`              // Task ID of a job array
	ArrayMaxTasks            uint32        `json:"array_max_tasks"`            //
	ArrayTaskString          string        `json:"array_task_string"`          //
	AssociationID            uint32        `json:"association_id"`             // Association id for job
	BatchFeatures            string        `json:"batch_features"`             // Batch features
	BatchFlag                bool          `json:"batch_flag"`                 // true if batch: queued job with script
	BatchHost                string        `json:"batch_host"`                 // Executing (batch) host.
	Flags                    []string      `json:"flags"`                      // Possible values: JOB_ACCRUE_OVER, JOB_WAS_RUNNING
	BurstBuffer              string        `json:"burst_buffer"`               //
	BurstBufferState         string        `json:"burst_buffer_state"`         //
	Cluster                  string        `json:"cluster"`                    // Cluster name
	ClusterFeatures          []string      `json:"cluster_features"`           // Cluster features
	Command                  string        `json:"command"`                    // Job submission command
	Comment                  string        `json:"comment"`                    // Job comment
	Container                string        `json:"container"`                  // OCI container bundle path
	Contiguous               bool          `json:"contiguous"`                 // true if contiguous nodes are requested by the job.
	CoreSpec                 *uint16       `json:"core_spec"`                  // Cores per socket required by job. nil if not set
	ThreadSpec               interface{}   `json:"thread_spec"`                // TODO: fix type
	CoresPerSocket           *uint16       `json:"cores_per_socket"`           // Cores per socket required by job. nil if not set
	BillableTRES             float64       `json:"billable_tres"`              // TRES billing value
	CPUsPerTask              *uint16       `json:"cpus_per_task"`              // CPUs per task required by job. nil if not set
	CPUFrequencyMinimum      *uint32       `json:"cpu_frequency_minimum"`      // Minimum CPU frequency required by job. nil if not set
	CPUFrequencyMaximum      *uint32       `json:"cpu_frequency_maximum"`      // Maximum CPU frequency required by job. nil if not set
	CPUFrequencyGovernor     interface{}   `json:"cpu_frequency_governor"`     // TODO: fix type
	CPUsPerTRES              string        `json:"cpus_per_tres"`              // CPUs per TRES required by job.
	Deadline                 int64         `json:"deadline"`                   // Job deadline in Unix timestamp
	DelayBoot                int64         `json:"delay_boot"`                 // Delay boot time in Unix timestamp
	Dependency               string        `json:"dependency"`                 // Job dependency
	DerivedExitCode          uint32        `json:"derived_exit_code"`          // Job highest exit code
	EligibleTime             int64         `json:"eligible_time"`              // Time the job is eligible for running in Unix timestamp
	EndTime                  int64         `json:"end_time"`                   // Job ending time (actaul or expected) in Unix timestamp
	ExcludedNodes            string        `json:"excluded_nodes"`             // List of excluded nodes for running job
	ExitCode                 string        `json:"exit_code"`                  // Job exit code
	Features                 string        `json:"features"`                   // Features required by the job
	FederationOrigin         string        `json:"federation_origin"`          //
	FederationSiblingsActive string        `json:"federation_siblings_active"` //
	FederationSiblingsViable string        `json:"federation_siblings_viable"` //
	GRESDetail               []interface{} `json:"gres_detail"`                // TODO: fix type
	GroupID                  string        `json:"group_id"`                   // Group ID of the job
	GroupName                string        `json:"group_name"`                 // Group name of the job
	JobID                    string        `json:"job_id"`                     // Job ID
	JobResources             JobResources  `json:"job_resources"`              // Job resources
	JobState                 JobState      `json:"job_state"`                  // Job state
	LastScheduleEvaluation   int64         `json:"last_sched_evaluation"`      //
	Licenses                 string        `json:"licenses"`                   // Licenses required by job
	MaxCPUs                  uint32        `json:"max_cpus"`                   // Maximum number of CPUs
	MaxNodes                 uint32        `json:"max_nodes"`                  // Maximum number of nodes
	MCSlabel                 string        `json:"mcs_label"`                  //
	MemoryPerTRES            string        `json:"memory_per_tres"`            // Memory allocated per TRES
	Name                     string        `json:"name"`                       // Job name
	Nodes                    string        `json:"nodes"`                      // Nodes allocated for the job
	Nice                     *uint32       `json:"nice"`                       // Nice value (adjustment to a job's scheduling priority).
	TasksPerCore             *uint16       `json:"tasks_per_core"`             // Tasks per core, nil if not set
	TasksPerNode             *uint16       `json:"tasks_per_node"`             // Tasks per node, nil if not set
	TasksPerSocket           *uint16       `json:"tasks_per_socket"`           // Tasks per socket, nil if not set
	TasksPerBoard            *uint16       `json:"tasks_per_board"`            // Tasks per board, nil if not set
	CPUs                     uint32        `json:"cpus"`                       // CPUs count
	NodeCount                uint32        `json:"node_count"`                 // Node count
	Tasks                    uint32        `json:"tasks"`                      // Tasks count
	HeterogeneousJobID       uint32        `json:"het_job_id"`                 // Job ID of the heterogeneous job leader
	HeterogeneousJobIDSet    string        `json:"het_job_id_set"`             // Expression identifying all components job IDs within a heterogeneous job
	HeterogeneousJobOffset   uint32        `json:"het_job_offset"`             // Zero origin offset within a collection of heterogeneous job components.
	Partition                string        `json:"partition"`                  // Partition running job
	Prefer                   string        `json:"prefer"`                     //
	MemoryPerNode            *uint64       `json:"memory_per_node"`            // Memory requested per node. nil if not set
	MemoryPerCPU             *uint64       `json:"memory_per_cpu"`             // Memory requested per CPU. nil if not set
	MinimumCPUsPerNode       uint32        `json:"minimum_cpus_per_node"`      // Minimum CPUs per node
	MinimumTMPDiskPerNode    uint32        `json:"minimum_tmp_disk_per_node"`  // Minimum TMP disk per node
	PreemptTime              int64         `json:"preempt_time"`               // The preempt time for the job
	PreSusTime               int64         `json:"pre_sus_time"`               //
	Priority                 int64         `json:"priority"`                   // Job priority
	Profile                  *uint32       `json:"profile"`                    // Profile of the job
	QOS                      string        `json:"qos"`                        // Job QoS
	Reboot                   bool          `json:"reboot"`                     // true if the allocated nodes should be rebooted before starting the job.
	RequiredNodes            string        `json:"required_nodes"`             // List of required nodes
	Requeue                  bool          `json:"requeue"`                    // true if job will requeue on failure
	ResizeTime               int64         `json:"resize_time"`                // The amount of time changed for the job to run.
	RestartCount             uint32        `json:"restart_cnt"`                // The number of restarts for the job
	ReservationName          string        `json:"resv_name"`                  // Reservation used by the job
	Shared                   interface{}   `json:"shared"`                     // TODO: fix type
	ShowFlags                []string      `json:"show_flags"`                 //
	SocketsPerBoard          uint16        `json:"sockets_per_board"`          // Number of socket per board
	SocketsPerNode           *uint16       `json:"sockets_per_node"`           //
	StartTime                int64         `json:"start_time"`                 // Job start time in Unix timestamp
	StateDescription         string        `json:"state_description"`          //
	StateReason              string        `json:"state_reason"`               //
	StandardError            string        `json:"standard_error"`             // STDERR for the job
	StandardInput            string        `json:"standard_input"`             // STDIN for the job
	StandardOutput           string        `json:"standard_output"`            // STDOUT for the job
	SubmitTime               int64         `json:"submit_time"`                // Job submission time in Unit timestamp
	SuspendTime              int64         `json:"suspend_time"`               // Job suspend time in Unit timestamp
	SystemComment            string        `json:"system_comment"`             // System comment
	TimeLimit                int64         `json:"time_limit"`                 // Job time limit in seconds
	TimeMinimum              int64         `json:"time_minimum"`               //
	ThreadsPerCore           *uint16       `json:"threads_per_core"`           // Threads per core requested by the job. nil if not set
	TRESBind                 string        `json:"tres_bind"`                  // Task to TRES binding directives
	TRESFrequency            string        `json:"tres_freq"`                  // TRES frequency directives
	TRESPerJob               string        `json:"tres_per_job"`               // Semicolon delimited list of TRES=# values
	TRESPerNode              string        `json:"tres_per_node"`              // Semicolon delimited list of TRES=# values
	TRESPerSocket            string        `json:"tres_per_socket"`            // Semicolon delimited list of TRES=# values
	TRESPerTask              string        `json:"tres_per_task"`              // Semicolon delimited list of TRES=# values
	TRESRequest              string        `json:"tres_req_str"`               // TRES requested in the job
	TRESAlloc                string        `json:"tres_alloc_str"`             // TRES used in the job
	UserID                   uint32        `json:"user_id"`                    // User the job runs as
	Username                 string        `json:"user_name"`                  // User name for a job or job step.
	WCKey                    string        `json:"wckey"`                      // WCKey for job
	WorkDIR                  string        `json:"current_working_directory"`  // Pathname of working directory
}
