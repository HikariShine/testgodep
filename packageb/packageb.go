package packagea

import (
	"fmt"
	"github.com/golang/glog"
)

func main() {
	fmt.Println("Hello, World!")
}

func Test() {
	glog.Info("Hello, World!")
}
