package api

import (
	_ "github.com/mokh1rbek/Book_CRUD/api/docs"
	"github.com/mokh1rbek/Book_CRUD/api/handler"
	"github.com/mokh1rbek/Book_CRUD/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(storage)

	r.POST("/book", handlerV1.CreateBook)
	r.GET("/book/:id", handlerV1.GetBookById)
	r.GET("/book", handlerV1.GetBookList)
	r.PUT("/book/:id", handlerV1.UpdateBook)
	r.DELETE("/book/:id", handlerV1.DeleteBook)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
