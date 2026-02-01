package governance

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAdminAction_IsEffective(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name   string
		action AdminAction
		now    time.Time
		want   bool
	}{
		{
			name: "currently effective",
			action: AdminAction{
				EffectiveAt: now.Add(-1 * time.Hour),
			},
			now:  now,
			want: true,
		},
		{
			name: "not yet effective",
			action: AdminAction{
				EffectiveAt: now.Add(1 * time.Hour),
			},
			now:  now,
			want: false,
		},
		{
			name: "expired",
			action: AdminAction{
				EffectiveAt: now.Add(-2 * time.Hour),
				ExpiresAt:   ptr(now.Add(-1 * time.Hour)),
			},
			now:  now,
			want: false,
		},
		{
			name: "not yet expired",
			action: AdminAction{
				EffectiveAt: now.Add(-1 * time.Hour),
				ExpiresAt:   ptr(now.Add(1 * time.Hour)),
			},
			now:  now,
			want: true,
		},
		{
			name: "invalidated",
			action: AdminAction{
				EffectiveAt:   now.Add(-1 * time.Hour),
				InvalidatedAt: ptr(now.Add(-30 * time.Minute)),
			},
			now:  now,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.action.IsEffective(tt.now))
		})
	}
}

func ptr(t time.Time) *time.Time {
	return &t
}
