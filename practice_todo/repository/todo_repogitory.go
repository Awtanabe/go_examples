package repository

import (
	"fmt"
	"practice_todo/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITodoRepository interface {
	GetAllTodos(todos *[]models.Todo, userId uint) error
	GetTodoById(todo *models.Todo, userId uint, todoId uint) error
	CreateTodo(todo *models.Todo) error
	UpdateTodo(todo *models.Todo, userId uint, todoId uint) error
	DeleteTodo(userId uint, todoId uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository{
	return &todoRepository{db}
}

func (tr *todoRepository) GetAllTodos(todos *[]models.Todo, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(todos).Error; err != nil{
		return err
	}
	return nil
}

func (tr *todoRepository) GetTodoById(todo *models.Todo, userId uint, todoId uint) error{
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(todo, todoId).Error; err != nil{
		return err
	}
	return nil
}

func (tr *todoRepository) CreateTodo(todo *models.Todo) error {
	if err := tr.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) UpdateTodo(todo * models.Todo, userId uint, todoId uint) error {
	result := tr.db.Model(todo).Clauses(clause.Returning{}).Where("id=? AND user_id=?",todoId, userId).Update("title", todo.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *todoRepository) DeleteTodo(userId uint, todoId uint) error {
	result := tr.db.Where("id=? AND user_id=?", todoId, userId).Delete(&models.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
