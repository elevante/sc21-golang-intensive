package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Find() {

	flagF := flag.Bool("f", false, "")
	flagSl := flag.Bool("sl", false, "")
	flagD := flag.Bool("d", false, "")
	flagExt := flag.String("ext", "", "")

	flag.Parse()

	err := filepath.Walk(flag.Args()[0], func(path string, info os.FileInfo, err error) error {

		if os.IsPermission(err) {
			return filepath.SkipDir
		} else if err != nil {
			fmt.Println(err)
			return nil
		}

		if !*flagF && !*flagSl && !*flagD {
			if info.Mode()&os.ModeSymlink != 0 {
				linkPath, err := filepath.EvalSymlinks(path)
				if os.IsNotExist(err) {
					fmt.Printf("/%s -> broken\n", path)
				}
				if err == nil {
					fmt.Printf("/%s -> /%s\n", path, linkPath)
				}
			} else {
				fmt.Printf("/%s\n", path)
			}
		}
		if *flagD {
			if info.IsDir() {
				fmt.Printf("/%s\n", path)
			}
		}

		if *flagSl {
			if info.Mode()&os.ModeSymlink != 0 {
				linkPath, err := filepath.EvalSymlinks(path)
				if os.IsNotExist(err) {
					fmt.Printf("/%s -> broken\n", path)
				}
				if err == nil {
					fmt.Printf("/%s -> /%s\n", path, linkPath)
				}
			}
		}

		if *flagF && *flagExt != "" {
			if info.Mode().IsRegular() && strings.HasSuffix(info.Name(), *flagExt) {
				fmt.Printf("%s\n", path)
			}
		} else if *flagF && *flagExt == "" {
			if info.Mode().IsRegular() {
				fmt.Printf("%s\n", path)
			}
		}
		return err
	})
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	Find()
}
