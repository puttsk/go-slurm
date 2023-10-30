package cli

import "github.com/puttsk/go-slurm"

// SQueueMetadata defines metadata returned from squeue command
type SQueueMetadata struct {
	Plugin struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"plugin"`

	Slurm struct {
		Version struct {
			Major uint16 `json:"major"`
			Micro uint16 `json:"micro"`
			Minor uint16 `json:"minor"`
		} `json:"version"`
		Release string `json:"release"`
	} `json:"slurm"`
}

// A SQueueResult contains information returned from squeue command
type SQueueResult struct {
	Metadata SQueueMetadata `json:"meta"`
	Error    []interface{}  `json:"errors"`
	Jobs     []slurm.Job    `json:"jobs"`
}
