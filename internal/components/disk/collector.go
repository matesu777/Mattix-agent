package disk

import (
	"golang.org/x/sys/unix"
)

func (d *Disk) Collect() error {
	var stat unix.Statfs_t

	err := unix.Statfs("/", &stat)
	if err != nil {
		return err
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)
	used := total - free

	d.Total = total
	d.Free = free
	d.Used = used

	return nil
}
