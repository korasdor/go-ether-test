package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenBinding struct {
	IPAddr    string `json:"ip_addr"`
	UserAgent string `json:"user_agent"`
}

func (t *TokenBinding) Parse(req *http.Request) (*TokenBinding, error) {
	// Get client IP address from the request
	ipAddr := req.Header.Get("X-Forwarded-For")
	if ipAddr == "" {
		ipAddr = req.RemoteAddr
	}
	ipAddr = strings.Split(ipAddr, ",")[0]
	if ipAddr == "" {
		return nil, errors.New("empty ip address")
	}

	// Get user agent from the request
	userAgent := req.UserAgent()
	if userAgent == "" {
		return nil, errors.New("empty user-agent")
	}

	tokenBinding := &TokenBinding{
		IPAddr:    ipAddr,
		UserAgent: userAgent,
	}

	return tokenBinding, nil
}

type TokenManager interface {
	NewJWT(userId string, tokenBinding *TokenBinding, ttl time.Duration) (string, error)
	ParseJWT(token string) (*TokenData, error)
}

type TokenData struct {
	UserId       string
	TokenBinding *TokenBinding
}

type Manager struct {
	signingKey []byte
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{
		signingKey: []byte(signingKey),
	}, nil
}

func (m *Manager) NewJWT(userId string, tokenBinding *TokenBinding, ttl time.Duration) (string, error) {
	tbJson, err := json.Marshal(tokenBinding)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
		"tb":  string(tbJson),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(m.signingKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (m *Manager) ParseJWT(tokenString string) (*TokenData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return m.signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error get user claims from token")
	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("error while geting userId from jwt claims")
	}

	tbJson, ok := claims["tb"].(string)
	if !ok {
		return nil, errors.New("error while geting token binding from jwt claims")
	}

	var tb TokenBinding
	err = json.Unmarshal([]byte(tbJson), &tb)
	if err != nil {
		return nil, errors.New("error while unmarshal json string to token binding struct")
	}

	tokenData := &TokenData{
		UserId:       userId,
		TokenBinding: &tb,
	}

	return tokenData, nil
}
