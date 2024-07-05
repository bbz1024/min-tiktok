package uid

import (
	"context"
	"github.com/google/uuid"
)

func GenUid(ctx context.Context, userID int) string {
	return uuid.NewString()
}
