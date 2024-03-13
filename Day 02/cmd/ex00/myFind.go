package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	dirFlag     = flag.Bool("d", false, "Print directories")
	fileFlag    = flag.Bool("f", false, "Print files")
	symlinkFlag = flag.Bool("sl", false, "Print symbolic links")
	extFlag     = flag.String("ext", "", "Print files with this extension")
)

func parseFlags() (string, error) {
	flag.Parse()

	if !*fileFlag && *extFlag != "" {
		return "", fmt.Errorf("to use -ext you need -f")
	}

	if len(flag.Args()) < 1 {
		return "", fmt.Errorf("enter a directory")
	}

	if !*dirFlag && !*fileFlag && !*symlinkFlag {
		*dirFlag = true
		*fileFlag = true
		*symlinkFlag = true
	}

	return flag.Args()[0], nil
}

func main() {
	root, err := parseFlags()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	err = filepath.Walk(root, visit)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func visit(path string, info fs.FileInfo, err error) error {
	if err != nil {
		if os.IsPermission(err) {
			return filepath.SkipDir
		}

		return err
	}

	if !info.IsDir() && !info.Mode().IsRegular() && !isSymlink(info.Mode()) {
		return nil
	}

	if *dirFlag && info.IsDir() {
		fmt.Println(path)
	}

	if *fileFlag && info.Mode().IsRegular() {
		if *extFlag != "" && !strings.HasSuffix(info.Name(), "."+*extFlag) {
			return nil
		}

		fmt.Println(path)
	}

	if *symlinkFlag && isSymlink(info.Mode()) {
		link, err := filepath.EvalSymlinks(path)
		if err != nil {
			fmt.Printf("%s -> [broken]\n", path)
		} else {
			fmt.Printf("%s -> %s\n", path, link)
		}
	}

	return nil
}

func isSymlink(mode os.FileMode) bool {
	return mode&os.ModeSymlink != 0
}
