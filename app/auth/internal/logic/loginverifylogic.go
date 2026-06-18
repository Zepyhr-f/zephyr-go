package logic

import (
	"context"
	"time"

	"zephyr-go/app/auth/internal/svc"
	"zephyr-go/app/auth/pb/pb"
	"zephyr-go/app/identity/identityservice"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginVerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginVerifyLogic {
	return &LoginVerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 【账号密码登录】
func (l *LoginVerifyLogic) LoginVerify(in *pb.LoginVerifyReq) (*pb.LoginVerifyResp, error) {
	// 1. 调用 Identity 域获取身份验证信息
	authInfo, err := l.svcCtx.IdentityRpc.GetUserAuthInfo(l.ctx, &identityservice.GetUserAuthInfoReq{
		Username:   in.Username,
		TenantCode: in.TenantCode,
	})
	if err != nil {
		l.Logger.Errorf("Failed to get auth info for user %s: %v", in.Username, err)
		return nil, status.Error(codes.Unauthenticated, "用户名或密码错误")
	}

	if authInfo.Status == 0 {
		return nil, status.Error(codes.PermissionDenied, "账号已被停用")
	}

	// 2. 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(authInfo.PasswordHash), []byte(in.Password))
	if err != nil {
		l.Logger.Errorf("Password mismatch for user %s", in.Username)
		return nil, status.Error(codes.Unauthenticated, "用户名或密码错误")
	}

	// 3. 签发 Token
	// 从 Config 里读取 JWT Secret 和 Expire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	refreshExpire := int64(86400 * 7) // 7天

	now := time.Now().Unix()
	accessJti := uuid.New().String()
	
	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["jti"] = accessJti
	claims["user_code"] = authInfo.UserCode
	claims["tenant_code"] = authInfo.TenantCode

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	accessToken, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, status.Error(codes.Internal, "生成 Access Token 失败")
	}

	// 生成 RefreshToken (也可以用 JWT，或者就是一个普通的不透明 UUID 存在 Redis 里)
	refreshToken := uuid.New().String()
	
	// 4. 将 RefreshToken 存入 Redis (供之后刷新使用)
	err = l.svcCtx.BizRedis.Setex("refresh_token:"+refreshToken, authInfo.UserCode, int(refreshExpire))
	if err != nil {
		return nil, status.Error(codes.Internal, "保存 Refresh Token 失败")
	}

	return &pb.LoginVerifyResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expire:       now + accessExpire,
		UserCode:     authInfo.UserCode,
		TenantCode:   authInfo.TenantCode,
	}, nil
}
