package alpaca

import (
	"fmt"
	"os"
)

func HandleError(e error) {
	if e != nil {
		fmt.Println("Encountered an error while running!")
		fmt.Println()
		fmt.Println(e)
		fmt.Println("\nIf you are unable to solve it, please report at")
		fmt.Println("https://github.com/pksunkara/alpaca/issues")
		os.Exit(1)
	}
}
