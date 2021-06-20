/**
 * @Author: hu5ky
 * @Description:
 * @File:  utils
 * @Version: 0.0.1
 * @Date: 2020/12/10 2:29 上午
**/

package subsystem

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

func GetCgroupPath(subsystemType, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountPoint(subsystemType)
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if errMkdir := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); errMkdir != nil {
				return "", fmt.Errorf("error create cgroup %v", errMkdir)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}

func FindCgroupMountPoint(subsystemName string) string {
	f, err := os.Open(MountInfo)
	if err != nil {
		logrus.Errorf("read mountinfo fail %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txtRow := scanner.Text()
		info := strings.Split(txtRow, " ")
		for _, opt := range strings.Split(info[len(info)-1], ",") {
			if opt == subsystemName {
				return info[4]
			}
		}

	}
	if scanner.Err() != nil {
		logrus.Errorf("read mountinfo fail %v", err)
	}
	return ""
}
