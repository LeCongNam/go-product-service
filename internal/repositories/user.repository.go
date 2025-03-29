package repositories

import (
	"product_srv/internal/database"
	"product_srv/internal/domain/users/models"

	"gorm.io/gorm"
)

// UserRepository kế thừa BaseRepository
type UserRepository struct {
	DB *gorm.DB
}

type GetList struct {
	Limit   int
	Offset  int
	Filters map[string][]string
}

// NewUserRepository - Inject DB vào BaseRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{DB: database.DB}
}

// GetListUser - Lấy danh sách user từ DB
func (ur *UserRepository) GetListUser(filter GetList) ([]models.User, int64, error) {
	var users []models.User
	var count int64
	query := ur.DB

	for key, values := range filter.Filters {
		for _, value := range values {
			query = query.Where(key+" = ?", value)
		}
	}

	query.Model(&users).Count(&count)      // Đếm số lượng trước
	query.Limit(10).Offset(0).Find(&users) // Lấy danh sách user

	return users, count, nil
}

func (ur *UserRepository) CreateUser(user models.User) (models.User, error) {
	ur.DB.Create(&user)
	return user, nil
}
