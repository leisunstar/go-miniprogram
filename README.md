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