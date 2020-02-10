package main
import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)
func main() {
	f, err := os.Open("output.tar")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer f.Close()
	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		f, err := os.Create("123" + fileinfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}