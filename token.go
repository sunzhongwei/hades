/*
实现 token 的生成和解析相关的接口。

token 可以用于下载文件的临时授权，或者其他需要临时授权的场景。

可以设置 token 的过期时间，确保安全性。

Example usage:

package main

import (
	"fmt"
	"time"

	// replace with your actual module import path
	"your/module/path/backend/utils"
)

func main() {
	// Set a stable secret at startup (optional if using TOKEN_SECRET env var)
	utils.SetSecret([]byte("super-secret-key"))

	// Generate a token valid for 15 minutes for resource "file:123"
	tok, err := utils.GenerateToken("file:123", 15*time.Minute)
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", tok)

	// Validate token (checks signature and expiration)
	payload, err := utils.ValidateToken(tok)
	if err != nil {
		fmt.Println("invalid token:", err)
		return
	}
	fmt.Printf("valid token for resource: %s, expires at %d\n", payload.Resource, payload.Exp)

	// ParseToken can be used when you only want to verify signature (no expiry check)
	// p, err := utils.ParseToken(tok)
	// ...
}
*/

package hades

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

var (
	// ErrInvalidToken token format invalid
	ErrInvalidToken = errors.New("invalid token format")
	// ErrSignatureMismatch token signature verification failed
	ErrSignatureMismatch = errors.New("token signature mismatch")
	// ErrExpired token expired
	ErrExpired = errors.New("token expired")
)

var defaultSecret []byte

func init() {
	if s := os.Getenv("TOKEN_SECRET"); s != "" {
		defaultSecret = []byte(s)
		return
	}
	// fallback: generate a random secret (note: will change across restarts)
	secret := make([]byte, 32)
	_, _ = rand.Read(secret)
	defaultSecret = secret
}

// SetSecret sets the secret used to sign/verify tokens.
// Call this during application startup if you want a stable secret.
func SetSecret(secret []byte) {
	if len(secret) == 0 {
		return
	}
	defaultSecret = make([]byte, len(secret))
	copy(defaultSecret, secret)
}

// TokenPayload is the JSON payload embedded in the token.
type TokenPayload struct {
	Resource string `json:"res"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Nonce    string `json:"nonce"`
}

// GenerateToken creates a signed token for the given resource that expires after ttl.
func GenerateToken(resource string, ttl time.Duration) (string, error) {
	now := time.Now()
	payload := TokenPayload{
		Resource: resource,
		Iat:      now.Unix(),
		Exp:      now.Add(ttl).Unix(),
		Nonce:    randomNonce(12),
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	sig := hmacSHA256(defaultSecret, b)
	encPayload := base64.RawURLEncoding.EncodeToString(b)
	encSig := base64.RawURLEncoding.EncodeToString(sig)
	return encPayload + "." + encSig, nil
}

// ParseToken verifies the token signature and returns the decoded payload.
// It does NOT check expiration; use ValidateToken to also enforce expiry.
func ParseToken(token string) (TokenPayload, error) {
	var payload TokenPayload
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return payload, ErrInvalidToken
	}
	encPayload, encSig := parts[0], parts[1]
	b, err := base64.RawURLEncoding.DecodeString(encPayload)
	if err != nil {
		return payload, ErrInvalidToken
	}
	sig, err := base64.RawURLEncoding.DecodeString(encSig)
	if err != nil {
		return payload, ErrInvalidToken
	}
	expected := hmacSHA256(defaultSecret, b)
	if !hmac.Equal(sig, expected) {
		return payload, ErrSignatureMismatch
	}
	if err := json.Unmarshal(b, &payload); err != nil {
		return payload, ErrInvalidToken
	}
	return payload, nil
}

// ValidateToken parses the token, verifies signature and expiration.
// Returns the payload if valid.
func ValidateToken(token string) (TokenPayload, error) {
	payload, err := ParseToken(token)
	if err != nil {
		return payload, err
	}
	if time.Now().Unix() > payload.Exp {
		return payload, ErrExpired
	}
	return payload, nil
}

// helpers

func hmacSHA256(key, msg []byte) []byte {
	m := hmac.New(sha256.New, key)
	_, _ = m.Write(msg)
	return m.Sum(nil)
}

func randomNonce(n int) string {
	if n <= 0 {
		n = 12
	}
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		// fallback to timestamp-derived nonce when crypto/rand fails
		return hex.EncodeToString([]byte(time.Now().String()))
	}
	return hex.EncodeToString(b)
}
