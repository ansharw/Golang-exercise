package services

import (
	"MyGram/helpers"
	"MyGram/model"
	"MyGram/repository"
	"context"
	"log"

	"gorm.io/gorm"
)

type commentService struct {
	db          *gorm.DB
	repoComment repository.CommentRepository
}

func NewCommentService(db *gorm.DB, repoComment repository.CommentRepository) *commentService {
	return &commentService{
		db:          db,
		repoComment: repoComment,
	}
}

func (service *commentService) FindAllByPhotoId(ctx context.Context, photoID uint) ([]model.Comments, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if photos, err := service.repoComment.FindAllByPhotoId(ctx, tx, photoID); err != nil {
		log.Println("Data not found")
		return photos, err
	} else {
		return photos, nil
	}
}

func (service *commentService) FindByPhotoId(ctx context.Context, photoID uint, id uint) (model.Comments, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if photo, err := service.repoComment.FindByPhotoId(ctx, tx, photoID, id); err != nil {
		log.Println("Data not found")
		return photo, err
	} else {
		return photo, nil
	}
}

func (service *commentService) Create(ctx context.Context, req model.RequestComments, userID, photoID uint) (model.Comments, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if photo, err := service.repoComment.Create(ctx, tx, req, userID, photoID); err != nil {
		log.Println("Failed to create photo")
		return photo, err
	} else {
		return photo, nil
	}
}

func (service *commentService) Update(ctx context.Context, req model.RequestComments, id, userID, photoID uint) (model.Comments, error) {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if photo, err := service.repoComment.Update(ctx, tx, req, id, userID, photoID); err != nil {
		log.Println("Failed to update photo")
		return photo, err
	} else {
		return photo, nil
	}
}

func (service *commentService) Delete(ctx context.Context, id, userID, photoID uint) error {
	tx := service.db.Begin()
	defer helpers.CommitOrRollback(tx)

	if err := service.repoComment.Delete(ctx, tx, id, userID, photoID); err != nil {
		log.Println("Failed to delete photo")
		return err
	} else {
		return nil
	}
}
