# go-miniprogram

##用法

	miniProgram := NewMiniProgramImpl().AddAppId("xx").AddAppSecret("xx")
	args := &GetUserArgs{}
	user, err := miniProgram.GetUser(args)
	if err != nil{
		return 
	}
