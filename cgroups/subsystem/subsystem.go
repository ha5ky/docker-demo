/**
 * @Author: hu5ky
 * @Description:
 * @File:  subsystem
 * @Version: 0.0.1
 * @Date: 2020/12/7 12:14 上午
**/

package subsystem

type ResourceConfig struct {
	MemoryLimit string
	CPUShare    string
	CPUSet      string
}

type Subsystem interface {
	Name() string
	Set(path string, res *ResourceConfig) error
	Apply(path string, pid int) error
	Remove(path string) error
}

var (
	SubsystemIns = []Subsystem{
		&CpuSetSubSystem{},
		&MemorySubSystem{},
		&CpuShareSubSystem{},
	}
)
