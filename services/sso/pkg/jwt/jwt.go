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

func (j *JwtService) GenerateAccessToken(user_id string, roles []string) (string, error) {
	return j.jwtGenerator.GenerateAccessToken(user_id, roles)
}

func (j *JwtService) GenerateRefreshToken() (string, error) {
	return j.jwtGenerator.GenerateRefreshToken()
}

func (j *JwtService) ValidateToken(token string) (map[string]interface{}, error) {
	return j.jwtValidator.ValidateToken(token)
}
