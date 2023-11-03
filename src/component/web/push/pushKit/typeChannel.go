package pushKit

type Channel interface {
	// Close 后端主动关闭通道.
	/*
		@param reason 关闭的原因
	*/
	Close(reason string) error
	IsClosed() (rst bool)
	SetClosed() (flag bool)

	Push(data []byte) error

	GetId() string
	GetBsid() string
	SetBsid(string)
	ClearBsid()
	GetUser() string
	SetUser(string)
	ClearUser()
	GetGroup() string
	SetGroup(string)
	ClearGroup()
	GetData() interface{}
	SetData(interface{})
	ClearData()

	BindGroup(group string)
	BindUser(user string)
	BindBsid(bsid string)
}
