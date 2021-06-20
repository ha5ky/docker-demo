/**
 * @Author: hu5ky
 * @Description:
 * @File:  cgroups
 * @Version: 0.0.1
 * @Date: 2020/12/3 5:09 下午
**/

package cgroups

import (
	"dockerDemo/cgroups/subsystem"
	"github.com/sirupsen/logrus"
)

type CgroupManager struct {
	Path string
	Res  *subsystem.ResourceConfig
}

func NewCgroupManager(name string) *CgroupManager {
	return &CgroupManager{
		Path: name,
	}
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range subsystem.SubsystemIns {
		if err := subSysIns.Apply(c.Path, pid); err != nil {
			return err
		}
	}
	return nil
}

func (c *CgroupManager) Set(res *subsystem.ResourceConfig) error {
	for _, subSysIns := range subsystem.SubsystemIns {
		if err := subSysIns.Set(c.Path, res); err != nil {
			return err
		}
	}
	return nil
}

func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range subsystem.SubsystemIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			logrus.Warnf("remove cgroup fail %v", err)
			return err
		}
	}
	return nil
}

/*
const cgroupMemoryHierarckyMount = "/sys/fs/cgroup/memory"

func Cgroups() {
	fmt.Println(os.Args)
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid is %d", syscall.Getpid())
		fmt.Println()
		cmd := exec.Command("sh", "-c", "stress --vm-bytes 200m --vm-keep -m 1")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	} else {
		// 得到folk出来进程映射在外部命名空间的pid
		fmt.Printf("%v", cmd.Process.Pid)

		testMemoryLimit := path.Join(cgroupMemoryHierarckyMount, "testMemoryLimit")
		// 在系统默认创建挂在的memory subsystem的hierarchy上创建cgroup
		err := os.Mkdir(testMemoryLimit, 0755)
		if err != nil {
			fmt.Println(err)
		}

		// sh -c "echo 1 > memory.oom_control" 需要关闭oom_killer
		OOMControl := path.Join(testMemoryLimit, "memory.oom_control")
		err = ioutil.WriteFile(OOMControl, []byte(strconv.Itoa(1)), 0644)
		if err != nil {
			fmt.Println(err)
		}

		// 将容器进程加入到这个cgroup中
		err = ioutil.WriteFile(path.Join(testMemoryLimit, "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		if err != nil {
			fmt.Println(err)
		}

		// 给进程添加内存限制
		err = ioutil.WriteFile(path.Join(testMemoryLimit, "memory.limit_in_bytes"), []byte("100m"), 0644)
		if err != nil {
			fmt.Println(err)
		}
		_, err = cmd.Process.Wait()
		if err != nil {
			fmt.Println(err)
		}
	}
}
*/
