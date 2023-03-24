package controller

import (
	"log"
	"net/http"

	"github.com/Luftalian/writers_app/domain"
	"github.com/Luftalian/writers_app/interfaces/database"
	"github.com/Luftalian/writers_app/usecase"
	"github.com/google/uuid"
)

type TagController struct {
	Interactor usecase.TagInteractor
}

func NewTagController(handler database.DbHandler) *TagController {
	return &TagController{
		Interactor: usecase.TagInteractor{
			TagRepository: &database.TagRepository{
				DbHandler: handler,
			},
		},
	}
}

func (controller *TagController) Create(c Context) {
	t := domain.Tag{}
	if err := c.Bind(&t); err != nil {
		log.Printf("Error while binding tag: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t, err := controller.Interactor.Add(t)
	if err != nil {
		log.Printf("Error while adding tag: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Created tag: %v", t)
	c.JSON(http.StatusCreated, t)
}

func (controller *TagController) Index(c Context) {
	tags, err := controller.Interactor.Tags()
	if err != nil {
		log.Printf("Error while getting tags: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Got tags: %v", tags)
	c.JSON(http.StatusOK, tags)
}

func (controller *TagController) Show(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error while parsing id: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	tag, err := controller.Interactor.TagByID(id)
	if err != nil {
		log.Printf("Error while getting tag: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Got tag: %v", tag)
	c.JSON(http.StatusOK, tag)
}
