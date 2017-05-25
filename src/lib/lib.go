package lib

import (
	"context"
	"fmt"
)

func Lib() {
	fmt.Println("Lib")
}

func LibContext(ctx context.Context) {
	fmt.Println("LibContext")
}
