package main

import (
	"fmt"
	"os"

	"github.com/copperium/fractals/internal/api"
)

func main() {
	err := api.Setup(":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
	}
}
