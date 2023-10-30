package slurm

// SlurmBackend provides an interface to extract information from Slurm
type SlurmBackend interface {
	ListCurrentJobs(flags *ListJobFlags) ([]Job, error)
}

type ListJobFlags struct {
}
