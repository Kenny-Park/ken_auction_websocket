package connectors

type ConnectUserIdManager struct {
	id int64
}

func (c *ConnectUserIdManager) New() int64 {
	c.id += int64(1)
	return c.id
}
