package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func testCopyFile(srcName, dstName string) {
	copied, err := copyFile(srcName, dstName)
	if err != nil {
		fmt.Printf("copyFile('%s', '%s') failed!!!\n", srcName, dstName)
	} else {
		fmt.Printf("Copied %d bytes\n", copied)
	}
	fmt.Println()
}

func main() {
	testCopyFile("08_defer_practical_usage.go", "new.go")
	testCopyFile("not_exists", "new.go")
	testCopyFile("08_defer_practical_usage.go", "")
	testCopyFile("08_defer_practical_usage.go", "/dev/full")
	testCopyFile("/dev/null", "new2.go")
}
