package entities

type ConnectEntityId struct {
	id int64
}

func (c *ConnectEntityId) New() int64 {
	c.id += int64(1)
	return c.id
}
