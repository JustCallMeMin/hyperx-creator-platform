package sessions

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJWTManager_GenerateAndVerify(t *testing.T) {
	secret := "test-secret"
	mgr := NewJWTManager(secret, 1*time.Hour)
	userID := uuid.New()

	token, err := mgr.Generate(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := mgr.Verify(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
}

func TestJWTManager_ExpiredToken(t *testing.T) {
	secret := "test-secret"
	// 1ms duration for expiry test
	mgr := NewJWTManager(secret, 1*time.Millisecond)
	userID := uuid.New()

	token, err := mgr.Generate(userID)
	assert.NoError(t, err)

	// Wait for expiry
	time.Sleep(2 * time.Millisecond)

	_, err = mgr.Verify(token)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "token is expired")
}
