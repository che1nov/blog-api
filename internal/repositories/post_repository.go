package repositories

import (
	"errors"
	"fmt"

	"blog-api/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("post with ID %d not found", id)
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetAll(term string) ([]models.Post, error) {
	var posts []models.Post
	query := r.DB
	if term != "" {
		query = query.Where("title LIKE ? OR content LIKE ? OR category LIKE ?", "%"+term+"%", "%"+term+"%", "%"+term+"%")
	}
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(id uint, updatedPost *models.Post) error {
	var post models.Post
	if err := r.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("post with ID %d not found", id)
		}
		return err
	}
	return r.DB.Model(&post).Updates(updatedPost).Error
}

func (r *PostRepository) Delete(id uint) error {
	var post models.Post
	if err := r.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("post with ID %d not found", id)
		}
		return err
	}
	return r.DB.Delete(&post).Error
}
