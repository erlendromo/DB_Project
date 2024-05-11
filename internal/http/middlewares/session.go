package middlewares

import (
	"DB_Project/internal/http/dependencies"
	"DB_Project/internal/utils"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"time"
)

// GenerateRandomID generates a random ID of the specified length
func GenerateRandomID(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

// Used to validate sessions retrieved by cookies
var sessions map[string]SessionData

func init() {
	sessions = make(map[string]SessionData)
}

// SessionData Session data
//
//	@title			SessionData
//	@summary		Used to store session data
//	@description	Used to store session data
type SessionData struct {
	ID        int
	Username  string
	LoginTime time.Time
	Admin     bool
}

// SetSession Set a session
//
//	@title			SetSession
//	@summary		Set a session
//	@description	Set a session with a username
func SetSession(w http.ResponseWriter, username string) error {
	sessionID, err := GenerateRandomID(32)
	if err != nil {
		return err
	}

	cd := dependencies.Dependencies.CustomerAddressDeps.PSQLCustomer

	c, err := cd.GetCustomerByUsername(context.Background(), username)
	if err != nil {
		return err
	}

	var a bool
	if c.Role == 1 {
		a = true
	}

	expiration := time.Now().Add(12 * time.Hour)

	cookie := &http.Cookie{
		Name:    "session",
		Value:   sessionID,
		Expires: expiration,
	}

	http.SetCookie(w, cookie)

	sessions[sessionID] = SessionData{
		ID:        c.ID,
		Username:  c.Username,
		LoginTime: time.Now(),
		Admin:     a,
	}

	return nil
}

// ClearSession Clear a session
//
//	@title			ClearSession
//	@summary		Clear a session
//	@description	Clear a session if it exists
func ClearSession(w http.ResponseWriter, r *http.Request) (int, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return http.StatusUnauthorized, utils.NewUnauthorizedError(err)
	}

	delete(sessions, cookie.Value)

	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
	return http.StatusOK, nil
}

// GetUserFromSession Get a user from a session
//
//	@title			GetUserFromSession
//	@summary		Get a user from a session
//	@description	Get a user from a session if it exists
func GetUserFromSession(r *http.Request) (SessionData, int, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return SessionData{}, http.StatusUnauthorized, errors.New("session not found")
	}

	sessionData, exists := sessions[cookie.Value]
	if !exists {
		return SessionData{}, http.StatusUnauthorized, errors.New("invalid session ID")
	}

	return sessionData, http.StatusOK, nil
}

// Explicitly define the type of the key to avoid collisions
type SessionKey string

const sessionKey SessionKey = "sessionData"

// SessionMiddleware Middleware to handle sessions
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err == nil {
			if sessionData, exists := sessions[cookie.Value]; exists {
				ctx := context.WithValue(r.Context(), sessionKey, sessionData)
				r = r.WithContext(ctx)
			}
		}

		next.ServeHTTP(w, r)
	})
}
