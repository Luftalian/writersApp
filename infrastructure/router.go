package infrastructure

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Luftalian/writers_app/interfaces/controller"
)

var Router *echo.Echo

func init() {
	dbHandler := NewDbHandler()
	store, err := mysqlstore.NewMySQLStoreFromConnection(dbHandler.Conn.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"https://writersapp.trap.games"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
        AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
    }))
	
	userController := controller.NewUserController(dbHandler)

	// e := e.Group("/api")
	e.POST("/users", func(c echo.Context) error {
		userController.Create(c)
		return nil
	})
	e.GET("/users", func(c echo.Context) error {
		userController.Index(c)
		return nil
	})
	e.GET("/users/:id", func(c echo.Context) error {
		userController.Show(c)
		return nil
	})

	textController := controller.NewTextController(dbHandler)
	e.POST("/texts", func(c echo.Context) error {
		textController.Create(c)
		return nil
	})
	e.GET("/texts", func(c echo.Context) error {
		textController.Index(c)
		return nil
	})
	e.GET("/texts/:id", func(c echo.Context) error {
		textController.Show(c)
		return nil
	})
	e.PUT("/texts/:id", func(c echo.Context) error {
		textController.Change(c)
		return nil
	})
	e.DELETE("/texts/:id", func(c echo.Context) error {
		textController.Delete(c)
		return nil
	})

	tagController := controller.NewTagController(dbHandler)
	e.POST("/tags", func(c echo.Context) error {
		tagController.Create(c)
		return nil
	})
	e.GET("/tags", func(c echo.Context) error {
		tagController.Index(c)
		return nil
	})
	e.GET("/tags/:id", func(c echo.Context) error {
		tagController.Show(c)
		return nil
	})

	tagListController := controller.NewTagListController(dbHandler)
	e.POST("/tag", func(c echo.Context) error {
		tagListController.Create(c)
		return nil
	})
	e.GET("/tag", func(c echo.Context) error {
		tagListController.Index(c)
		return nil
	})
	e.GET("/tag/text/:id", func(c echo.Context) error {
		tagListController.ShowByTextID(c)
		return nil
	})
	e.GET("/tag/tag/:id", func(c echo.Context) error {
        tagListController.ShowByTagID(c)
        return nil
    })
	e.GET("/tag/name/:tagName", func(c echo.Context) error {
		tagListController.ShowByName(c)
        return nil
    })
	e.PUT("/tag/:id", func(c echo.Context) error {
		tagListController.Change(c)
		return nil
	})
	e.DELETE("/tag/:id", func(c echo.Context) error {
		tagListController.Delete(c)
		return nil
	})

	userLikeController := controller.NewUserLikeController(dbHandler)
	e.POST("/like", func(c echo.Context) error {
		userLikeController.Create(c)
		return nil
	})
	e.GET("/like", func(c echo.Context) error {
		userLikeController.Index(c)
		return nil
	})
	e.POST("/like/check", func(c echo.Context) error {
		userLikeController.Check(c)
        return nil
    })
	e.GET("/like/user/:id", func(c echo.Context) error {
		userLikeController.ShowByUserID(c)
		return nil
	})
	e.GET("/like/text/:id", func(c echo.Context) error {
		userLikeController.ShowByTextID(c)
		return nil
	})
	e.PUT("/like/:id", func(c echo.Context) error {
		userLikeController.Change(c)
		return nil
	})
	e.DELETE("/like/user/:id", func(c echo.Context) error {
		userLikeController.DeleteByUserID(c)
		return nil
	})
	e.DELETE("/like/text/:id", func(c echo.Context) error {
		userLikeController.DeleteByTextID(c)
		return nil
	})
	e.POST("/like/delete", func(c echo.Context) error {
		userLikeController.Delete(c)
		return nil
	})

	userCreateController := controller.NewUserCreateController(dbHandler)
	e.POST("/create", func(c echo.Context) error {
		userCreateController.Create(c)
		return nil
	})
	e.GET("/create", func(c echo.Context) error {
		userCreateController.Index(c)
		return nil
	})
	e.GET("/create/user/:id", func(c echo.Context) error {
		userCreateController.ShowByUserID(c)
		return nil
	})
	e.GET("/create/text/:id", func(c echo.Context) error {
		userCreateController.ShowByTextID(c)
		return nil
	})
	e.PUT("/create/:id", func(c echo.Context) error {
		userCreateController.Change(c)
		return nil
	})
	e.DELETE("/create/user/:id", func(c echo.Context) error {
		userCreateController.DeleteByUserID(c)
		return nil
	})
	e.DELETE("/create/text/:id", func(c echo.Context) error {
		userCreateController.DeleteByTextID(c)
		return nil
	})
	e.POST("/create/delete", func(c echo.Context) error {
		userCreateController.Delete(c)
		return nil
	})

	Router = e
}
