package middlewareAuth

import "errors"

var (
	ErrorAuthHeaderEmpty   = errors.New("Invalid Access Token")
	ErrorAuthNotHaveBearer = errors.New("authorization header doesn't have bearer format")
	ErrorAuthNotHaveToken  = errors.New("authorization header doesn't have access token value")
	ErrorUserTokenEmpty    = errors.New("failed getting user info from nil user")
	ErrorUserFromContext   = errors.New("failed getting user info from context")
)
