package middlewareAuth

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"belajar/package/config"

	jwtAuth "belajar/package/auth/jwt"

	jsonHelper "belajar/helper/json"

	"github.com/rs/zerolog/log"

	conversionHelper "belajar/helper/conversion"
)

type Middleware interface {
	InitLog(next http.Handler) http.Handler
	CheckToken(next http.Handler) http.Handler
	GetUserInfoFromContext(ctx context.Context) (*UserData, error)
}

type Options struct {
	jwt jwtAuth.Jwt
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(Options)
	opt.jwt = jwtAuth.NewJwt(cfg)
	return opt
}

func (o *Options) getTokenInHeader(r *http.Request) (string, error) {
	authzHeader := r.Header.Get("Authorization")
	if authzHeader == "" {
		return "", ErrorAuthHeaderEmpty
	}

	accessToken := strings.Split(authzHeader, " ")
	if accessToken[0] != "Bearer" {
		err := ErrorAuthNotHaveBearer
		return "", err
	}

	if len(accessToken) == 1 {
		err := ErrorAuthNotHaveToken
		return "", err
	}

	return accessToken[1], nil
}

func (o *Options) InitLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set body
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body.Close() //  must close
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// set start time
		tnow := time.Now()

		ctx := context.WithValue(r.Context(), config.ContextKey("body"), bodyBytes)
		ctx = context.WithValue(ctx, config.ContextKey("startTime"), tnow)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (o *Options) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := o.getTokenInHeader(r)
		if err != nil {
			code := "[Middleware] Middleware-1"
			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponse(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		jwtData, err := o.jwt.VerifyAccessToken(accessToken)
		if err != nil {
			code := "[Middleware] Middleware-2"
			if err.Error() == "Token is expired" {
				err = errors.New("token is expired")
				log.Error().Err(err).Msg(code)
				jsonHelper.ErrorResponse(w, r, http.StatusUnauthorized, err.Error())
				return
			}

			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponse(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		if jwtData == nil {
			code := "[Middleware] Middleware-3"
			log.Error().Err(err).Msg(code)
			jsonHelper.ErrorResponse(w, r, http.StatusUnauthorized, ErrorUserTokenEmpty.Error())
			return
		}

		ctx := context.WithValue(r.Context(), config.ContextKey("userInfo"), jwtData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (o *Options) GetUserInfoFromContext(ctx context.Context) (*UserData, error) {
	userInfo, ok := ctx.Value(config.ContextKey("userInfo")).(*jwtAuth.JwtData)
	if !ok || userInfo == nil {
		code := "[Middleware] GetUserInfoFromContext-1"
		log.Error().Err(ErrorUserFromContext).Msg(code)
		return nil, ErrorUserFromContext
	}

	newUserID, _ := conversionHelper.StrToInt64(userInfo.UserID)

	userData := &UserData{
		UserID:   newUserID,
		UserName: userInfo.UserName,
		Role:     userInfo.Role,
	}

	return userData, nil
}
