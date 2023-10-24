package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get_NotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_Get_Found(t *testing.T) {
	category := entity.Category{
		Name: "EXAMPLE_CATEGORY",
		Id:   "1234",
	}

	categoryRepository.Mock.On("FindById", "1234").Return(category)

	res, err := categoryService.Get("1234")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, category.Id, res.Id)
	assert.Equal(t, category.Name, res.Name)
}
