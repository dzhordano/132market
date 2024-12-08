package jwt

type JwtService struct {
	jwtGenerator TokenGenerator
	jwtValidator TokenValidator
}

func NewJwtService(jwtGenerator TokenGenerator, jwtValidator TokenValidator) TokenManager {
	return &JwtService{
		jwtGenerator, jwtValidator,
	}
}

func (j *JwtService) GenerateToken(user_id string, roles []string) (string, error) {
	return j.jwtGenerator.GenerateToken(user_id, roles)
}

func (j *JwtService) ValidateToken(token string) (map[string]interface{}, error) {
	return j.jwtValidator.ValidateToken(token)
}
