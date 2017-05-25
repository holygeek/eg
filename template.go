package P

import (
	"context"
	"lib"
)

func before() {
	lib.Lib()
}

func after() {
	lib.LibContext(context.TODO())
}
