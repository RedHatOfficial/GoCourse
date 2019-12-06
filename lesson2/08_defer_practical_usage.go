package main

import (
	"fmt"
	"io"
	"os"
)

func closeFile(file *os.File) {
	fmt.Printf("Closing file '%s'\n", file.Name())
	file.Close()
}

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("Cannot open file '%s' for reading\n", srcName)
		return
	}
	fmt.Printf("File '%s' opened for reading\n", srcName)
	defer closeFile(src)

	dst, err := os.Create(dstName)
	if err != nil {
		fmt.Printf("Cannot create destination file '%s'\n", dstName)
		return
	}
	fmt.Printf("File '%s' opened for writing\n", dstName)
	defer closeFile(dst)

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
