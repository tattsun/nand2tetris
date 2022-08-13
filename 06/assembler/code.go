package main

type Code struct{}

func NewCode() *Code {
	return &Code{}
}

func (c *Code) Dest(mnemonic string) string {
	panic("not implemented")
}

func (c *Code) Comp(mnemonic string) string {
	panic("not implemented")
}

func (c *Code) Jump(mnemonic string) string {
	panic("not implemented")
}
