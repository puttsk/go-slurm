package cli

import (
	"strconv"

	"github.com/puttsk/go-slurm"
)

func parseUint16Field(v string) (uint16, error) {
	if v == "" {
		return slurm.Infinite16, nil
	} else {
		o, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return slurm.NoVal16, err
		}
		return uint16(o), nil
	}
}

func parseUint32Field(v string) (uint32, error) {
	if v == "" {
		return slurm.Infinite, nil
	} else {
		o, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return slurm.NoVal, err
		}
		return uint32(o), nil
	}
}
