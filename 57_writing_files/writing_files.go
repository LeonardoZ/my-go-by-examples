package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// write all at once
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat2", d1, 0644)
	check(err)

	// controlled write
	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	// bytes
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// strings
	n3, err := f.WriteString("write\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// flush
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
