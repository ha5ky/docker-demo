/**
 * @Author: hu5ky
 * @Description:
 * @File:  cpu_share
 * @Version: 0.0.1
 * @Date: 2020/12/7 12:24 上午
**/

package subsystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuSetSubSystem struct {
}

func (c *CpuSetSubSystem) Name() string {
	return "cpu"
}

func (c *CpuSetSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(c.Name(), cgroupPath, true); err == nil {
		fmt.Println(c.Name(),subsysCgroupPath)
		if res.CPUShare != "" {
			if err := ioutil.WriteFile(
				path.Join(subsysCgroupPath, CpuShare),
				[]byte(res.CPUShare),
				0755);
				err != nil {
				return fmt.Errorf("set cgroup cpu share limit fail %v", err)
			}
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %v error %v",cgroupPath,err)
	}
}

func (c *CpuSetSubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(c.Name(), cgroupPath, false); err == nil {
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

func (c *CpuSetSubSystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(c.Name(), cgroupPath, false); err == nil {
		return os.Remove(path.Join(subsysCgroupPath))
	} else {
		return err
	}
}
