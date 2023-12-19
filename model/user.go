package model

func QueryUserPwd(uid string) string {
	var user1 User
	DB.Where("email = ?",uid).First(&user1)
	return user1.Password
}
func UpdateCode(em string,cod string)  {
	var user1 UserTmp
	if resu:= DB.Where("email = ?",em).First(&user1);resu.Error != nil{
		user1.Email=em;
		user1.Code=cod;
		DB.Create(&user1)
	}else{//可查，说明注册过了
		DB.Model(&user1).Where("email=?",em).Update("code",cod)
	}
}
func UpdateUser(em string,newpwd string)  {
	var user1 User
	if resu:= DB.Where("email = ?",em).First(&user1);resu.Error != nil{
		user1.Email=em;
		user1.Password=newpwd;
		DB.Create(&user1)
	}else{//可查，说明注册过了
		DB.Model(&user1).Where("email=?",em).Update("password",newpwd)
	}
}
