package route

import (
	"P1/handler"
	"P1/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	api := r.Group("/books")
	{
		api.POST("", server.CreateBook)
		api.GET("", server.ListBook)
		api.GET(":id", server.GetBook)
		api.PUT(":id", server.UpdateBook)
		api.DELETE(":id", server.DeleteBook)
	}
}
