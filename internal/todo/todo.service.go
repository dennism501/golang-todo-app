package todo

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

type TodoService interface {
	GetTodo(ID uint) (Todo, error)
	GetAllTodos() ([]Todo, error)
	CreateTodo(todo Todo) (Todo, error)
	UpdateTodo(ID uint, newTodo Todo) (Todo, error)
	DeleteTodo(ID uint) error
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

//Method returns a single Todo by its ID
func (s *Service) GetTodo(ID uint) (Todo, error) {
	var todo Todo

	//add Postgres call here
	return todo, nil
}

func (s *Service) CreateTodo(todo Todo) (Todo, error) {

	return todo, nil
}

func (s *Service) UpdateTodo(ID uint, newTodo Todo) (Todo, error) {
	// add err for error handling when adding DB call
	todo, _ := s.GetTodo(ID)

	return todo, nil

}

func (s *Service) GetAllTodos() ([]Todo, error) {
	var todos []Todo

	return todos, nil
}
