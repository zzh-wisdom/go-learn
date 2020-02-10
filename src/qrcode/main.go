package main
import(
"github.com/skip2/go-qrcode"
"image/color"
"log"
)
func main() {
	qr,err:=qrcode.New("https://docs.qq.com/doc/DWlVXeHBTdFZNYWdS",qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.White
		qr.ForegroundColor = color.Black
		qr.WriteFile(256,"./golang_qrcode.png")
	}
	//qrcode.WriteFile("http://c.biancheng.net/",qrcode.Medium,256,"./golang_qrcode.png")
}