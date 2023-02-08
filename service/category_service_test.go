package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"unittest/entity"
	"unittest/repository"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, _ := categoryService.Get("1")

	assert.Nil(t, category)
}

func TestCategoryExist(t *testing.T) {
	category := entity.Category{Id: "2", Name: "Makanan"}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")

	assert.NotNil(t, result)

	assert.Nil(t, err)
}
