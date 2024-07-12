package flags

import (
	"flag"
	"fmt"
	"os"
)

// Creating Required Flags

func Main() {
	var input = flag.String("i", "", "path to input file (required)")
	var compressed = flag.Bool("c", false, "compress output")
	var copies = flag.Int("x", 1, "how many copies")
	var output = flag.String("o", "./path.json", "output path")

	flag.Parse()

	if *input == "" {
		fmt.Println("The input flag is required!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println(*compressed)
	fmt.Println(*copies)
	fmt.Println(*output)
}
