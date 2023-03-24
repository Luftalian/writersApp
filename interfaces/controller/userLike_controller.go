package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
	"github.com/Luftalian/writers_app/interfaces/database"
	"github.com/Luftalian/writers_app/usecase"
)

type UserLikeController struct {
	Interactor usecase.UserLikeInteractor
}

func NewUserLikeController(handler database.DbHandler) *UserLikeController {
	return &UserLikeController{
		Interactor: usecase.UserLikeInteractor{
			UserLikeRepository: &database.UserLikeRepository{
				DbHandler: handler,
			},
		},
	}
}

func (controller *UserLikeController) Create(c Context) {
	ul := domain.UserLike{}
	if err := c.Bind(&ul); err != nil {
		log.Printf("Error while binding user like: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ul, err := controller.Interactor.Add(ul)
	if err != nil {
		log.Printf("Error while creating user like: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Created user like: %v", ul)
	c.JSON(http.StatusCreated, ul)
}

func (controller *UserLikeController) Index(c Context) {
	userLikes, err := controller.Interactor.UserLikes()
	if err != nil {
		log.Printf("Error while getting user likes: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Got user likes: %v", userLikes)
	c.JSON(http.StatusOK, userLikes)
}

func (controller *UserLikeController) Check(c Context) {
	ul := domain.UserLike{}
	if err := c.Bind(&ul); err != nil {
		log.Printf("Error while binding user like: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	checkStruct := domain.Check{}
	check, err := controller.Interactor.CheckLike(ul)
	checkStruct.Check = check
	if err == sql.ErrNoRows {
		log.Printf("UserLike already exists")
		c.JSON(http.StatusOK, checkStruct)
		return
	}
    if err != nil {
        log.Printf("Error while getting user likes: %v", err)
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    log.Printf("Got user likes: %v", checkStruct)
    c.JSON(http.StatusOK, checkStruct)
}

func (controller *UserLikeController) ShowByUserID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error while parsing user id: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userLike, err := controller.Interactor.UserLikeByUserID(id)
	if err != nil {
		log.Printf("Error while getting user like: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Got user like: %v", userLike)
	c.JSON(http.StatusOK, userLike)
}

func (controller *UserLikeController) ShowByTextID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error while parsing user id: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userLike, err := controller.Interactor.UserLikeByTextID(id)
	if err != nil {
		log.Printf("Error while getting user like: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Got user like: %v", userLike)
	c.JSON(http.StatusOK, userLike)
}

func (controller *UserLikeController) Change(c Context) {
	ul := domain.UserLike{}
	if err := c.Bind(&ul); err != nil {
		log.Printf("Error while binding user like: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := controller.Interactor.Change(ul)
	if err != nil {
		log.Printf("Error while changing user like: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Changed user like: %v", ul)
	c.JSON(http.StatusOK, nil)
}

func (controller *UserLikeController) Delete(c Context) {
	ul := domain.UserLike{}
	if err := c.Bind(&ul); err != nil {
		log.Printf("Error while binding user like: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := controller.Interactor.Delete(ul)
	if err != nil {
		log.Printf("Error while deleting user like: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Deleted user like: %v", ul)
	c.JSON(http.StatusOK, nil)
}

func (controller *UserLikeController) DeleteByUserID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error while parsing user id: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = controller.Interactor.DeleteByUserID(id)
	if err != nil {
		log.Printf("Error while deleting user like: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Deleted user like: %v", id)
	c.JSON(http.StatusOK, nil)
}

func (controller *UserLikeController) DeleteByTextID(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Printf("Error while parsing user id: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = controller.Interactor.DeleteByTextID(id)
	if err != nil {
		log.Printf("Error while deleting user like: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	log.Printf("Deleted user like: %v", id)
	c.JSON(http.StatusOK, nil)
}
