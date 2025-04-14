package variable

import "fmt"

type Base struct {
	Num int
}

func (b Base) Describe() string {
	return fmt.Sprintf("Base with num=%v", b.Num)
}

type Container struct {
	Base
	Str string
}
