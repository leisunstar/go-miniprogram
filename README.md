# go-miniprogram

##用法

	miniProgram := NewMiniProgramImpl().AddAppId("xx").AddAppSecret("xx")
	//1 获取session
	session, err := miniProgram.GetSession(code)
	if err != nil{
		return 
	}
	//2 根据session 解密
	user := &User{}
	err = miniProgram.Decode(encryptedData, iv, session,user) 
	if err != nil{
		return 
	}
	// 获取小程序二维码
	err := miniProgram.GetWXacodeunLimitToFile("scene", "pages/index", 400, false, "./qr.png")
	if err != nil {
		t.Fatalf("err %v", err)
	}
	f, err := os.Create("qr1.png")
	if err != nil {
		t.Fatalf("err %v", err)
	}
	err = miniProgram.GetWXacodeunLimitWriter("scene", "pages/index", 400, false, f)
	if err != nil {
		t.Fatalf("err %v", err)
	}