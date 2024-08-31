package auth

type AuthService struct {
	Store UserStorage
}

func NewService(s UserStorage) *AuthService {
	return &AuthService{Store: s}
}
