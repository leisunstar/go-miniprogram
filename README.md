# go-miniprogram

eg:

	package main

    import (
    	"fmt"
    	"github.com/leisunstar/go-miniprogram"
    	"os"
    )

    func main() {
    	m := miniprogram.NewMiniProgramImpl().AddAppId("wx57211e78df7eabca").
    		AddAppSecret("f652af82546a92b5a8cef879af6b998c")
    	//获取session
    	session, err := m.GetSession("code")
    	if err != nil {
    		fmt.Printf("err %v \n", err)
    		return
    	}
    	//2 根据session 解密
    	user := &miniprogram.User{}
    	err = m.Decode("encryptedData", "iv", session, user)
    	if err != nil {
    		fmt.Printf("err %v \n", err)
    		return
    	}
    	// 获取小程序二维码 写入文件
    	err = m.GetWXacodeunLimitToFile("scene", "pages/index", 400, false, "./qr.png")
    	if err != nil {
    		fmt.Printf("err %v \n", err)
    		return
    	}
    	f, err := os.Create("qr1.png")
    	if err != nil {
    		fmt.Printf("err %v \n", err)
    		return
    	}
    	// 获取小程序二维码 写入io.writer
    	err = m.GetWXacodeunLimitWriter("scene", "pages/index", 400, false, f)
    	if err != nil {
    		fmt.Printf("err %v \n", err)
    		return
    	}
    }
