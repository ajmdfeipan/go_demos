package main

import ("flag"
	"fmt"
)

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	fmt.Println(*cfg)
	fmt.Println(*version)
	fmt.Println(*check)
//	flag.Parse()


}
