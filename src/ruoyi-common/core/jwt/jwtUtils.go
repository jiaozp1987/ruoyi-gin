package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/golang-jwt/jwt/v4"
	iniUtils "ruoyi-gin/src/ruoyi-common/core/ini"
	redisUtils "ruoyi-gin/src/ruoyi-common/core/redis"
	"ruoyi-gin/src/ruoyi-domain/domainFactory"
	"ruoyi-gin/src/ruoyi-domain/loginBody"
	lu "ruoyi-gin/src/ruoyi-domain/loginUser"
	"time"
)

// GenToken 生成JWT
func GenToken(body *loginBody.LoginBody) (string, error) {

	// jwt token 的过期时间
	TokenExpireDuration := time.Duration(iniUtils.JWT_DURATION) * time.Minute
	sysUser := domainFactory.NewSysUserWithId(body.UserName)
	permissions := make([]string, 1)
	for _, v := range sysUser.Role {
		permissions = slice.Union(v.Permissions, permissions)
	}
	loginUser := domainFactory.NewLoginUser(sysUser.UserId, sysUser.DeptId, sysUser, permissions)
	// 创建一个我们自己的声明
	loginUser.User = *sysUser
	loginUser.ExpireTime = int(time.Now().Add(TokenExpireDuration).Unix())
	loginUser.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		Issuer:    iniUtils.JWT_ISSUER, // 签发人
	}
	loginUser.Token, _ = random.UUIdV4()

	loginUser.Token = strutil.ReplaceWithMap(loginUser.Token, map[string]string{"-": ""})
	// 使用指定的签名方法创建签名对象
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, loginUser)
	token, err := jwtToken.SignedString([]byte(iniUtils.JWT_KEY))
	if err != nil {
		fmt.Println(err.Error())
	}
	loginUserJson, _ := json.Marshal(loginUser)
	//loginUserStr, _ := convertor.ToJson(loginUser)
	err = redisUtils.Set(token, string(loginUserJson), iniUtils.REDIS_EXPIRE)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*lu.LoginUser, error) {
	// 解析token
	sysuser := lu.LoginUser{}
	loginUserJson, err := redisUtils.Get(tokenString)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("invalid token")
	}
	err = json.Unmarshal([]byte(loginUserJson), &sysuser)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &lu.LoginUser{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(iniUtils.JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*lu.LoginUser); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
