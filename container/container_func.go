/**
 * @Author: hu5ky
 * @Description:
 * @File:  container_func
 * @Version: 0.0.1
 * @Date: 2020/12/6 9:01 下午
**/

package container

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := ioutil.ReadAll(pipe)
	if err != nil {
		log.Errorf("read pipe error %v", err)
		return nil
	}
	return strings.Split(string(msg), " ")
}

func RunContainerInitProcess() error {
	comArray := readUserCommand()
	if comArray == nil || len(comArray) == 0 {
		return fmt.Errorf("Run container get user command error, cmdArray is nil")
	}
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	path, err := exec.LookPath(comArray[0])
	if err != nil {
		log.Errorf("Exec path %v error %v", comArray, err)
		return err
	}
	log.Infof("Find path %v", path)
	if err := syscall.Exec(path, comArray[0:], os.Environ()); err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func NewPipe() (read *os.File, write *os.File, err error) {
	if read, write, err = os.Pipe(); err != nil {
		return nil, nil, err
	} else {
		return
	}
}

func NewParentProcess(tty bool) (cmd *exec.Cmd, writePipe *os.File) {
	readPipe, writePipe, err := NewPipe()
	if err != nil {
		log.Errorf("create pipe error %v", err)
		return nil, nil
	}
	cmd = exec.Command("/proc/self/exe", "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.ExtraFiles = []*os.File{readPipe}
	return
}

func Run(tty bool, comArray []string /*, res *subsystem.ResourceConfig*/) {
	parentCmd, writePipe := NewParentProcess(tty)
	if parentCmd == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parentCmd.Start(); err != nil {
		log.Error(err)
	}
	//cgroupM := cgroups.NewCgroupManager("dockerDome")
	//defer cgroupM.Destroy()
	//err := cgroupM.Set(res)
	//if err != nil {
	//	log.Errorf("cgroup set error %v",err)
	//	return
	//}
	//err = cgroupM.Apply(parentCmd.Process.Pid)
	//if err != nil {
	//	log.Errorf("cgroup apply error %v",err)
	//	return
	//}

	sendInitCommand(comArray, writePipe)
	_ = parentCmd.Wait()
	os.Exit(-1)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	cmd := strings.Join(comArray, " ")
	log.Infof("command all is %v", cmd)
	_, err := writePipe.WriteString(cmd)
	if err != nil {
		log.Error(err)
	}
	err = writePipe.Close()
	if err != nil {
		log.Error(err)
	}
}

func pivotRoot(root string) error {
	_ = syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("mount rootfs to itself error %v", err)
	}
	pivotDir := path.Join(root, ".pivot_root")
	if err := os.Mkdir(pivotDir, 0777); err != nil {
		return fmt.Errorf("make pivot dir error %v", err)
	}
	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return fmt.Errorf("pivot_root error %v", err)
	}

	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("change dir error %v", err)
	}

	pivotDir = path.Join("/", ".pivot_root")
	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("umount error %v", err)
	}

	return os.Remove(pivotDir)
}

func setUpMount() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Errorf("get local work dir error %v", err)
		return
	}
	log.Info("current location dir is %v", pwd)
	if err := pivotRoot(pwd); err != nil {
		log.Errorf("pivotRoot error %v", err)
	}
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	_ = syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	_ = syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	_ = syscall.Mount("tmpfs", "/dev", "tmpfs",
		uintptr(syscall.MS_NOSUID|syscall.MS_STRICTATIME), "mode=755")

}
