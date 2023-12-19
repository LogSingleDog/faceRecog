package service
import(
	"time"
	"faceRecog/model"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)
// GenToken 生成JWT
func GenerateToken(UserId string) (string, error) {
	// 创建一个我们自己的声明
	expireTime := time.Now().Add(30 * 24 * time.Hour)
	claims := model.Claims{
		ID:UserId, // 自定义字段
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "43.138.61.49", // 签名颁发者
			Subject:   "人脸识别",         //签名主题人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte("Hello"))
}

// 解析token
func GetToken(c *gin.Context) string {
	tokenString := c.GetHeader("Authorization")
	//vcalidate token formate
	if tokenString == "" {
		return "0"
	}
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return "0"
	}
	return claims.ID
}
func ParseToken(tokenString string) (*jwt.Token, *model.Claims, error) {
	Claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("Hello"), nil
	})
	return token, Claims, err
}