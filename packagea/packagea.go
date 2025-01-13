package packagea

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello, World!")
}

func Test() {
	logrus.Errorf("configs.GetInstance().WorkspacePos.Strategies.FindStrategy err: %s", "agfdfs")
}
