package main
import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)
func main() {
	f, err := os.Create("./output.tar") //创建一个 tar 文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	tw := tar.NewWriter(f)
	defer tw.Close()
	fileName := []string{"./tar.txt", "./tar1.txt"}
	for _, file := range fileName {
		fileinfo, err := os.Stat(file) //获取文件相关信息
		if err != nil {
			fmt.Println(err)
		}
		hdr, err := tar.FileInfoHeader(fileinfo, "")
		if err != nil {
			fmt.Println(err)
		}
		err = tw.WriteHeader(hdr) //写入头文件信息
		if err != nil {
			fmt.Println(err)
		}
		f1, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		m, err := io.Copy(tw, f1) //将main.exe文件中的信息写入压缩包中
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(file, m)
	}
}