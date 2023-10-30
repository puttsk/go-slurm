package cli

import "github.com/puttsk/go-slurm"

type SlurmCLIBackend struct{}

func (s *SlurmCLIBackend) ListCurrentJobs(flags slurm.ListJobFlags) ([]slurm.Job, error) {
	return nil, nil
}
