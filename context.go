
package liter

type Context struct{
	c *Conn
}

func NewContext(conn *Conn)(*Context){
	return &Context{
		c: conn,
	}
}
