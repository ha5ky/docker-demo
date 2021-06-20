/**
 * @Author: hu5ky
 * @Description:
 * @File:  memory
 * @Version: 0.0.1
 * @Date: 2020/12/7 12:22 上午
**/

package subsystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct {
}

func (m *MemorySubSystem) Name() string {
	return "memory"
}

func (m *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(m.Name(), cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(
				path.Join(subsysCgroupPath, LimitInBytesFile),
				[]byte(res.MemoryLimit),
				0644);
				err != nil {
				return fmt.Errorf("set cgroup memory limit fail %v", err)
			}
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %v error %v",cgroupPath,err)
	}
}

func (m *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(m.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(
			path.Join(subsysCgroupPath, Tasks),
			[]byte(strconv.Itoa(pid)),
			0644);
			err != nil {
			return fmt.Errorf("apply process to cgroup fail %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %v error %v",cgroupPath,err)
	}
}

func (m *MemorySubSystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(m.Name(), cgroupPath, true); err == nil {
		return os.Remove(path.Join(subsysCgroupPath))
	} else {
		return err
	}
}
