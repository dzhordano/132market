package jwt

type TokenGenerator interface {
	GenerateToken(user_id string, roles []string) (string, error)
}

type TokenValidator interface {
	ValidateToken(token string) (map[string]interface{}, error)
}

type TokenManager interface {
	TokenGenerator
	TokenValidator
}
