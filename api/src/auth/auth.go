package auth

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/Hinkku-Company/aphrodite_monolite/config"
	"github.com/Hinkku-Company/aphrodite_monolite/logger"
	"github.com/Hinkku-Company/aphrodite_monolite/src/shared/models/tables"
	"github.com/Hinkku-Company/aphrodite_monolite/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	config config.Config
	log    *slog.Logger
}

func NewAuth(config config.Config) *Auth {
	decodeCert(&config)
	return &Auth{
		config: config,
		log:    logger.Log(),
	}
}

func decodeCert(config *config.Config) {
	cert, err := utils.DecodeB64(config.AccessTokenPrivateKey)
	if err != nil {
		panic(err)
	}
	config.AccessTokenPrivateKey = cert
	cert, err = utils.DecodeB64(config.RefreshTokenPrivateKey)
	if err != nil {
		panic(err)
	}
	config.RefreshTokenPrivateKey = cert
	cert, err = utils.DecodeB64(config.AccessTokenPublicKey)
	if err != nil {
		panic(err)
	}
	config.AccessTokenPublicKey = cert
	cert, err = utils.DecodeB64(config.RefreshTokenPublicKey)
	if err != nil {
		panic(err)
	}
	config.RefreshTokenPublicKey = cert
}

func (a *Auth) CreateToken(user tables.User) (*tables.Access, error) {
	token, err := a.generateToken(user)
	if err != nil {
		return nil, err
	}
	rToken, err := a.generateRefreshToken(user)
	if err != nil {
		return nil, err
	}
	return &tables.Access{
		Token:        token,
		TokenRefresh: rToken,
	}, nil
}

func (a *Auth) ValidToken(token string) (tables.User, error) {
	resp := tables.User{}
	key, err := jwt.ParseECPublicKeyFromPEM([]byte(a.config.AccessTokenPublicKey))
	if err != nil {
		a.log.Error("parse private key", "error", err)
		return resp, err
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			a.log.Error("parse error algorithm", "error", err)
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		a.log.Error("validate error", "error", err)
		return resp, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		a.log.Error("invalid token")
		return resp, fmt.Errorf("validate: invalid token")
	}
	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		a.log.Error("expired token")
		return resp, fmt.Errorf("validate: expired token")
	}

	resp.ID, err = uuid.Parse(claims["sub"].(string))
	if err != nil {
		a.log.Error("uuid parse error", "error", err)
		return resp, err
	}
	// resp.TypeUser.Name = claims["role"].(string)

	return resp, nil
}

func (a *Auth) ValidRefreshToken(token string) (tables.User, error) {
	resp := tables.User{}
	key, err := jwt.ParseECPublicKeyFromPEM([]byte(a.config.RefreshTokenPublicKey))
	if err != nil {
		a.log.Error("parse private key", "error", err)
		return resp, err
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			a.log.Error("parse error algorithm", "error", err)
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		a.log.Error("validate error", "error", err)
		return resp, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		a.log.Error("invalid token")
		return resp, fmt.Errorf("validate: invalid token")
	}
	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		a.log.Error("expired token")
		return resp, fmt.Errorf("validate: expired token")
	}

	resp.ID, err = uuid.Parse(claims["sub"].(string))
	if err != nil {
		a.log.Error("uuid parse error", "error", err)
		return resp, err
	}
	// resp.TypeUser.Name = claims["role"].(string)

	return resp, nil
}

func (a *Auth) CreatePasswordHash(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		a.log.Error("generate password hash", "error", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func (a *Auth) ValidPasswordHash(password, hash string) (bool, error) {
	passwordBytes := []byte(password)
	hashedPassword := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	if err != nil {
		a.log.Error("password hash compare error", "error", err)
		return false, nil
	}
	return true, err
}

func (a *Auth) generateToken(user tables.User) (string, error) {
	now := time.Now().UTC()
	exp, err := strconv.Atoi(a.config.AccessTokenExpiredMin)
	if err != nil {
		a.log.Error("parse expired", err)
		return "", err
	}
	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(a.config.AccessTokenPrivateKey))
	if err != nil {
		a.log.Error("parse private key", err)
		return "", err
	}

	atClaims := make(jwt.MapClaims)
	atClaims["sub"] = user.ID
	roles := []string{}
	for _, role := range user.AccessRols {
		roles = append(roles, role.AccessRol.Name)
	}
	atClaims["role"] = strings.Join(roles, ",")
	atClaims["exp"] = now.Add(time.Duration(exp) * time.Minute).Unix()
	atClaims["iat"] = now.Unix()
	atClaims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, atClaims).SignedString(key)
	if err != nil {
		a.log.Error("generate token", err)
		return "", err
	}

	return token, nil
}

func (a *Auth) generateRefreshToken(user tables.User) (string, error) {
	now := time.Now().UTC()
	exp, err := strconv.Atoi(a.config.RefreshTokenExpiredMin)
	if err != nil {
		a.log.Error("parse expired", err)
		return "", err
	}
	key, err := jwt.ParseECPrivateKeyFromPEM([]byte(a.config.RefreshTokenPrivateKey))
	if err != nil {
		a.log.Error("parse private key", err)
		return "", err
	}

	atClaims := make(jwt.MapClaims)
	atClaims["sub"] = user.ID
	atClaims["role"] = user.Name
	atClaims["exp"] = now.Add(time.Duration(exp) * time.Minute).Unix()
	atClaims["iat"] = now.Unix()
	atClaims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, atClaims).SignedString(key)
	if err != nil {
		a.log.Error("generate token", err)
		return "", err
	}

	return token, nil
}
