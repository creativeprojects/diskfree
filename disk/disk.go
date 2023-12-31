//go:build !windows

package disk

import "syscall"

// Disk contains usage data and provides user-friendly access methods
type Disk struct {
	stat *syscall.Statfs_t
}

// New returns an object holding the disk usage of volumePath
func New(volumePath string) (*Disk, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(volumePath, &stat)
	if err != nil {
		return nil, err
	}
	return &Disk{&stat}, nil
}

func (c *Disk) Refresh() error {
	return syscall.Statfs("/", c.stat)
}

// Free returns total free bytes on file system
func (d *Disk) Free() uint64 {
	return d.stat.Bfree * uint64(d.stat.Bsize)
}

// Available return total available bytes on file system to an unprivileged user
func (d *Disk) Available() uint64 {
	return d.stat.Bavail * uint64(d.stat.Bsize)
}

// Size returns total size of the file system
func (d *Disk) Size() uint64 {
	return uint64(d.stat.Blocks) * uint64(d.stat.Bsize)
}

// Used returns total bytes used in file system
func (d *Disk) Used() uint64 {
	return d.Size() - d.Free()
}

// Usage returns percentage of use on the file system
func (d *Disk) Usage() float64 {
	return float64(d.Used()) / float64(d.Size())
}
