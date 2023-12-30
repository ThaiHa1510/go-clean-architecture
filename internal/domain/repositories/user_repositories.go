package repositories

import (
	"fmt"
	"context"
)

type UserRepository interface{
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, userID string) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, userID string) error
}

func NewMysqlUserRepository(){
	return &MysqlUserRepository{
		
	}
}
type MysqlUserRepository struct{
	users []*domain.User
	
}

type InMemoryUserRepository struct {
	users []*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: []*domain.User{}}
}

func (repo *InMemoryUserRepository) Create(ctx context.Context, user *domain.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *InMemoryUserRepository) GetByID(ctx context.Context, userID string) (*domain.User, error) {
	for _, user := range repo.users {
			if user.ID == userID {
					return user, nil
			}
	}
	return nil, errors.New("user not found")
}

func (repo *InMemoryUserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	return repo.users, nil
}

func (repo *InMemoryUserRepository) Update(ctx context.Context, user *domain.User) error {
	for i, existingUser := range repo.users {
			if existingUser.ID == user.ID {
					repo.users[i] = user
					return nil
			}
	}
	return errors.New("user not found")
}

func (repo *InMemoryUserRepository) Delete(ctx context.Context, userID string) error {
	for i, user := range repo.users {
			if user.ID == userID {
					repo.users = append(repo.users[:i], repo.users[i+1:])
					return nil
			}
	}
	return errors.New("user not found")
}

type MysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) *MysqlUserRepository {
	return &MysqlUserRepository{db: db}
}

func (repo *MysqlUserRepository) Create(ctx context.Context, user *domain.User) error {
	if err := repo.db.Create(user).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (repo *MysqlUserRepository) GetByID(ctx context.Context, userID string) (*domain.User, error) {
	user := &domain.User{}
	if err := repo.db.Where("id = ?", userID).First(user).Error; err != nil {
			if errors.Is(err, gorm.ErrNotFound) {
					return nil, errors.New("user not found")
			}
			return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return user, nil
}

func (repo *MysqlUserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	if err := repo.db.Find(&users).Error; err != nil {
			return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return users, nil
}

func (repo *MysqlUserRepository) Update(ctx context.Context, user *domain.User) error {
	if err := repo.db.Save(user).Error; err != nil {
			return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (repo *MysqlUserRepository) Delete(ctx context.Context, userID string) error {
	if err := repo.db.Where("id = ?", userID).Delete(&domain.User{}).Error; err != nil{
		return err
	}
	return nil
}