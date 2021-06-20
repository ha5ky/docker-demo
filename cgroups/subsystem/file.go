/**
 * @Author: hu5ky
 * @Description:
 * @File:  file
 * @Version: 0.0.1
 * @Date: 2020/12/7 1:03 上午
**/

package subsystem

const (
	CpuShare = "cpu.shares"
	CpuSet = "cpuset.cpus" // centos7 要先设置这个参数，不然会报write /sys/fs/cgroup/cpuset/dockerDome/tasks: no space left on device
	MountInfo = "/proc/self/mountinfo"
	LimitInBytesFile = "memory.limit_in_bytes"
	Tasks = "tasks"
)
