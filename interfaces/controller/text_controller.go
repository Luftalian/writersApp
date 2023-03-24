package controller

import (
	"log"
	"net/http"

	"github.com/Luftalian/writers_app/domain"
	"github.com/Luftalian/writers_app/interfaces/database"
	"github.com/Luftalian/writers_app/usecase"
	"github.com/google/uuid"
)

type TextController struct {
	Interactor usecase.TextInteractor
}

func NewTextController(handler database.DbHandler) *TextController {
	return &TextController{
		Interactor: usecase.TextInteractor{
			TextRepository: &database.TextRepository{
				DbHandler: handler,
			},
		},
	}
}

func (controller *TextController) Create(c Context) {
	t := domain.Text{}
	if err := c.Bind(&t); err != nil {
		log.Printf("Error in TextController.Create: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// if t.TextID != uuid.Nil || t.Title == "" || t.Content == "" || t.UserID == uuid.Nil || t.UserName == "" || !t.CreatedAt.IsZero() || !t.ChangedAt.IsZero() || t.GoodCount != 0 || t.BadCount != 0 {
	// 	log.Printf("Error in TextController.Create: %v", t)
	// 	c.JSON(http.StatusBadRequest, "Error in TextID, Title, Content, UserID, UserName, CreatedAt, UpdatedAt, GoodCount, BadCount")
	// 	return
	// }
	t, err := controller.Interactor.Add(t)
	if err != nil {
		log.Printf("Error in TextController.Create: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("TextController.Create: %v", t)
	c.JSON(http.StatusCreated, t)
}

func (controller *TextController) Index(c Context) {
	texts, err := controller.Interactor.Texts()
	if err != nil {
		log.Printf("Error in TextController.Index: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("TextController.Index: %v", texts)
	c.JSON(http.StatusOK, texts)
}

func (controller *TextController) Show(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error in TextController.Show: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	text, err := controller.Interactor.TextByID(id)
	if err != nil {
		log.Printf("Error in TextController.Show: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("TextController.Show: %v", text)
	c.JSON(http.StatusOK, text)
}

func (controller *TextController) Change(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error in TextController.Change: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	text, err := controller.Interactor.TextByID(id)
	if err != nil {
		log.Printf("Error in TextController.Change: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if err := c.Bind(&text); err != nil {
		log.Printf("Error in TextController.Change: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// if text.TextID == uuid.Nil || text.Title == "" || text.Content == "" || text.UserID == uuid.Nil || text.UserName == "" || text.CreatedAt.IsZero() {
	// 	log.Printf("Error in TextController.Change: %v", text)
	// 	c.JSON(http.StatusBadRequest, "Error in TextID, Title, Content, UserID, UserName, CreatedAt, UpdatedAt, GoodCount, BadCount")
	// 	return
	// }
	err = controller.Interactor.TextChange(text)
	if err != nil {
		log.Printf("Error in TextController.Change: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("TextController.Change: %v", text)
	c.JSON(http.StatusOK, text)
}

func (controller *TextController) Delete(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error in TextController.Delete: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if id == uuid.Nil {
		log.Printf("Error in TextController.Delete: %v", id)
		c.JSON(http.StatusBadRequest, "Error in TextID")
		return
	}
	err = controller.Interactor.TextDelete(id)
	if err != nil {
		log.Printf("Error in TextController.Delete: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("TextController.Delete: %v", id)
	c.JSON(http.StatusOK, nil)
}
