package service
import(
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"log"
	"github.com/spf13/viper"
)
func SendEmail(to string,code string) error{
	email := gomail.NewMessage()
	from:=viper.GetString("email.from")
	email.SetHeader("From", from)
	email.SetHeader("To", to)
	email.SetHeader("Subject", "这是标题")
	//你先别急，这里还得改
	text:="您的验证码为"+code+",验证码10分钟内有效，请及时验证"
	email.SetBody("text/html", text)
	
    // 创建SMTP客户端，连接到远程的邮件服务器，需要指定服务器地址、端口号、用户名、
	//密码，如果端口号为465的话，自动开启SSL，这个时候需要指定TLSConfig
	pwd:=viper.GetString("email.pwd")
	d := gomail.NewDialer("smtp.qq.com", 465, from, pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 跳过ssl验证
	
	if err := d.DialAndSend(email); err != nil {
		log.Println(err)
		return err
	}
	log.Println("发送成功")
	return nil
}