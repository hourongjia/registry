package registry

// Registry 注册接口
type Registry interface {
	Register(path string, data string) (bool, error)
	RegisterWithDeadLine(path string, data string, deadlineTime float64) (bool, error)
	DelRegister(path string) (bool, error)
	GetRegisterData(path string) (string, error)
	UpdateDeadLine(path string, deadlineTime float64) (bool, error)
}

type Discover interface {
	Subscribe(key string, listener *Listener)
}
