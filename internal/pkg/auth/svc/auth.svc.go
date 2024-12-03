package svc

import (
	"app/internal/core/cfg"
	"app/internal/pkg/user/ent"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	db gossiper.Database
}

func NewAuthService(db gossiper.Database) *AuthService {
	return &AuthService{db: db}
}

// Generate JWT Token
func (s *AuthService) generateJWT(user *ent.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.String(),
		"email":  user.Email,
		"exp":    time.Now().Add(168 * time.Hour).Unix(),
	})

	secret := cfg.Inst().SecretKey
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &tokenString, nil
}

func (s *AuthService) ValidateToken(token string) (bool, *ent.User, error) {
	secret := cfg.Inst().SecretKey

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userId, ok := claims["userId"].(string)
		if !ok {
			return false, nil, errors.New("invalid token payload: missing userId")
		}

		var user ent.User
		if err := s.db.GetDB().Where("id = ? AND deleted_at IS NULL", userId).First(&user).Error; err != nil {
			return false, nil, fmt.Errorf("user not found: %w", err)
		}

		return true, &user, nil
	}

	return false, nil, errors.New("invalid or expired token")
}

func (s *AuthService) Login(email, password string) (*string, *ent.User, error) {
	var user ent.User

	if err := s.db.GetDB().Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return nil, nil, errors.New("invalid email or password")
	}

	// Check user state todo: implement this logic later
	//if user.State != ent.Active {
	//	return nil, nil, errors.New("account is not active")
	//}

	// Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := s.generateJWT(&user)
	if err != nil {
		return nil, nil, err
	}

	return token, &user, nil
}

func (s *AuthService) Register(user *ent.User) (*string, *ent.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.State = ent.Suspended

	// Save user to DB
	if err := s.db.GetDB().Create(user).Error; err != nil {
		return nil, nil, err
	}

	// Generate JWT token
	token, err := s.generateJWT(user)
	if err != nil {
		return nil, nil, err
	}

	return token, user, nil
}
