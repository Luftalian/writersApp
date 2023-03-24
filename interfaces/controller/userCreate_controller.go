package controller

import (
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
	"github.com/Luftalian/writers_app/interfaces/database"
	"github.com/Luftalian/writers_app/usecase"
)

type UserCreateController struct {
	Interactor usecase.UserCreateInteractor
}

func NewUserCreateController(handler database.DbHandler) *UserCreateController {
	return &UserCreateController{
		Interactor: usecase.UserCreateInteractor{
			UserCreateRepository: &database.UserCreateRepository{
				DbHandler: handler,
			},
		},
	}
}

func (controller *UserCreateController) Create(c Context) {
	ul := domain.UserCreate{}
	if err := c.Bind(&ul); err != nil {
		log.Printf("Error creating userCreate: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ul, err := controller.Interactor.Add(ul)
	if err != nil {
		log.Printf("Error creating userCreate: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Created userCreate: %v", ul)
	c.JSON(http.StatusCreated, ul)
}

func (controller *UserCreateController) Index(c Context) {
	userCreates, err := controller.Interactor.UserCreates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, userCreates)
}

func (controller *UserCreateController) ShowByUserID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userCreate, err := controller.Interactor.UserCreateByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, userCreate)
}

func (controller *UserCreateController) ShowByTextID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userCreate, err := controller.Interactor.UserCreateByTextID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, userCreate)
}

func (controller *UserCreateController) Change(c Context) {
	ul := domain.UserCreate{}
	if err := c.Bind(&ul); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := controller.Interactor.Change(ul)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (controller *UserCreateController) Delete(c Context) {
	ul := domain.UserCreate{}
	if err := c.Bind(&ul); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := controller.Interactor.Delete(ul)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (controller *UserCreateController) DeleteByUserID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = controller.Interactor.DeleteByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (controller *UserCreateController) DeleteByTextID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = controller.Interactor.DeleteByTextID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
