package user_usecase

import (
	"context"
)

// Delete xóa user
func (u *usecase) Delete(ctx context.Context, id uint) error {
	return u.userRepo.Delete(ctx, id)
}
