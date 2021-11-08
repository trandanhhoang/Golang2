package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func popLine(f *os.File) ([]byte, error) {
	//Stat take fileInfo, that have Size() to take length in bytes for regular files
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	// len 0, capacity fi.Size()
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))
	// SeekStart   = 0 // seek relative to the origin of the file
	// SeekCurrent = 1 // seek relative to the current offset
	// SeekEnd     = 2 // seek relative to the end
	// offset -1 mean go back 1, 0 mean stay here, 1 go forward 1
	_, err = f.Seek(0, io.SeekStart) // curPos is the head of file
	if err != nil {
		return nil, err
	}
	// Copy all from f to buf
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}

	// ReadBytes(delim) stop when meet delim '\n', and all byte readed is got rid of
	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	// curPos, err = f.Seek(0, io.SeekCurrent)
	// fmt.Print("curPos", curPos), curPos will in end of file.

	// because offset in the end of file after ReadBytes()
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	nw, err := io.Copy(f, buf)
	if err != nil {
		return nil, err
	}

	err = f.Truncate(nw)
	if err != nil {
		return nil, err
	}
	err = f.Sync()
	if err != nil {
		return nil, err
	}

	return line, nil
}

func main() {
	fname := `popline.txt`
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()
	line, err := popLine(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("pop:", string(line))
}
