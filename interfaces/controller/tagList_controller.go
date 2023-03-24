package controller

import (
	"log"
	"net/http"

	"github.com/Luftalian/writers_app/domain"
	"github.com/Luftalian/writers_app/interfaces/database"
	"github.com/Luftalian/writers_app/usecase"
	"github.com/google/uuid"
)

type TagListController struct {
	Interactor usecase.TagListInteractor
}

func NewTagListController(handler database.DbHandler) *TagListController {
	return &TagListController{
		Interactor: usecase.TagListInteractor{
			TagListRepository: &database.TagListRepository{
				DbHandler: handler,
			},
		},
	}
}

func (controller *TagListController) Create(c Context) {
	t := domain.TagList{}
	if err := c.Bind(&t); err != nil {
		log.Printf("Error creating tag list: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t, err := controller.Interactor.Add(t)
	if err != nil {
		log.Printf("Error creating tag list: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Created tag list: %v", t)
	c.JSON(http.StatusCreated, t)
}

func (controller *TagListController) Index(c Context) {
	tags, err := controller.Interactor.TagLists()
	if err != nil {
		log.Printf("Error getting tag list: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Tag list: %v", tags)
	c.JSON(http.StatusOK, tags)
}

func (controller *TagListController) ShowByTextID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error parsing tag list id: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	tags, err := controller.Interactor.TagListByTextID(id)
	if err != nil {
		log.Printf("Error getting tag list: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Tag list: %v", tags)
	c.JSON(http.StatusOK, tags)
}

func (controller *TagListController) ShowByTagID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
    if err!= nil {
        log.Printf("Error parsing tag list id: %v", err)
        c.JSON(http.StatusBadRequest, err)
        return
    }
    tags, err := controller.Interactor.TagListByTagID(id)
    if err!= nil {
        log.Printf("Error getting tag list: %v", err)
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    log.Printf("Tag list: %v", tags)
    c.JSON(http.StatusOK, tags)
}

func (controller *TagListController) ShowByName(c Context) {
	name := c.Param("tagName")
    texts, err := controller.Interactor.TagListByName(name)
    if err!= nil {
        log.Printf("Error getting tag list: %v", err)
        c.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Printf("Text list: %v", texts)
    c.JSON(http.StatusOK, texts)
}

func (controller *TagListController) Change(c Context) {
	t := domain.TagList{}
	if err := c.Bind(&t); err != nil {
		log.Printf("Error changing tag list: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t, err := controller.Interactor.Update(t)
	if err != nil {
		log.Printf("Error changing tag list: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Updated tag list: %v", t)
	c.JSON(http.StatusOK, t)
}

func (controller *TagListController) Delete(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error parsing tag list id: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := controller.Interactor.Delete(id); err != nil {
		log.Printf("Error deleting tag list: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Deleted tag list: %v", id)
	c.JSON(http.StatusOK, nil)
}

