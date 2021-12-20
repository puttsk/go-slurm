package slurm_test

import (
	"testing"

	"github.com/puttsk/go-slurm"
)

func TestSlurmBit(t *testing.T) {

	for i := 0; i < 64; i++ {
		flags := slurm.SlurmBit(i)
		expect := uint64(1) << (i)

		if flags != expect {
			t.Errorf("Invalid value at SlurmBit(%d). Expect: %x, Actual: %x", i, expect, flags)
		}
	}
}
