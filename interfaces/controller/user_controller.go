package controller

import (
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
	"github.com/Luftalian/writers_app/interfaces/database"
	"github.com/Luftalian/writers_app/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(handler database.DbHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				DbHandler: handler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	log.Println("Create user")
	u := domain.User{}
	if err := c.Bind(&u); err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if u.UserID != uuid.Nil || !u.CreatedAt.IsZero() {
		log.Printf("UserID and CreatedAt must be empty")
		c.JSON(http.StatusBadRequest, "UserID and CreatedAt must be empty")
		return
	}
	user, err := controller.Interactor.Add(u)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("User created: %v", user)
	c.JSON(http.StatusCreated, user)
}

func (controller *UserController) Index(c Context) {
	log.Println("Index user by log")
	users, err := controller.Interactor.Users()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Users: %v", users)
	c.JSON(http.StatusOK, users)
}

func (controller *UserController) Show(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error getting user: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := controller.Interactor.UserByID(id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("User: %v", user)
	c.JSON(http.StatusOK, user)
}
