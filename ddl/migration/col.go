package migration

import "strings"

type Col struct {
	query strings.Builder
}
func (c *Col) BigSerial() *Col {
	c.query.WriteString(" BIGSERIAL")
	return c
}
func (c *Col) Primarykey() *Col {
	c.query.WriteString(" PRIMARY KEY")
	return c
}
func (c *Col) Name(name string) *Col {
	c.query.WriteString(" ")
	c.query.WriteString(name)
	return c
}
func (c *Col) End() *Col {
	c.query.WriteString(" );")
	return c
}
