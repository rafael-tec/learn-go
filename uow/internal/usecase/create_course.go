package usecase

import (
	"context"
	"uow/internal/entity"
	"uow/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type CreateCourseUseCase struct {
	CourseRepository   repository.CourseRepositoryInterface
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewCreateCourseUseCase(
	courseRepository repository.CourseRepositoryInterface,
	categoryRepository repository.CategoryRepositoryInterface,
) *CreateCourseUseCase {
	return &CreateCourseUseCase{
		CourseRepository:   courseRepository,
		CategoryRepository: categoryRepository,
	}
}

func (uc *CreateCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := uc.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	err = uc.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
