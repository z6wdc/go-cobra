package bot

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Echo() {
	reader := bufio.NewReader(os.Stdin)
	testCount, _ := reader.ReadString('\n')
	fmt.Println(strings.TrimSpace(testCount))
}
