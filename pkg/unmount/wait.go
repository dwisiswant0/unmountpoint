package unmount

import (
	"fmt"

	"github.com/moby/sys/mountinfo"
)

// Wait for path to unmount and send to channel
func Wait(c chan<- bool, path string) error {
	var i int

	p, e := mountinfo.GetMounts(mountinfo.ParentsFilter(path))
	if e != nil {
		return e
	}

	for d := range p {
		i = d
	}

	if i == 0 {
		return fmt.Errorf(errPathIsNotMountPoint, path)
	}

	for {
		m, e := mountinfo.Mounted(path)
		if e != nil {
			return e
		}

		if !m {
			c <- true
			return nil
		}
	}
}
