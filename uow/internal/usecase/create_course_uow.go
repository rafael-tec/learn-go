package usecase

import (
	"context"
	"uow/internal/entity"
	"uow/internal/repository"
	"uow/pkg/uow"
)

type CreateCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewCreateCourseUseCaseUow(
	uow uow.UowInterface,
) *CreateCourseUseCaseUow {
	return &CreateCourseUseCaseUow{
		Uow: uow,
	}
}

func (uc *CreateCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return uc.Uow.Do(ctx, func(uow *uow.UOW) error {
		category := entity.Category{
			Name: input.CategoryName,
		}

		repoCategory := uc.GetCategoryRepository(ctx)
		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		repoCourse := uc.GetCourseCategoryRepository(ctx)
		err = repoCourse.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (uc *CreateCourseUseCaseUow) GetCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	rep, err := uc.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}

	return rep.(repository.CategoryRepositoryInterface)
}

func (uc *CreateCourseUseCaseUow) GetCourseCategoryRepository(ctx context.Context) repository.CourseRepositoryInterface {
	rep, err := uc.Uow.GetRepository(ctx, "CourseCategoryRepository")
	if err != nil {
		panic(err)
	}

	return rep.(repository.CourseRepositoryInterface)
}
