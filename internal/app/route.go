package app

import (
	"pro-link-api/docs"
	mdw "pro-link-api/internal/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *ServerHttp) Rounte() {

	r := s.server
	s.swagger(r)
	s.middleware(r)

	v1 := r.Group("/api/v1")
	eg := v1.Group("/example")
	eg.GET("/helloworld", s.adapter.Helloworld)

}

func (s *ServerHttp) swagger(route *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *ServerHttp) middleware(route *gin.Engine) {
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-type"},
		AllowCredentials: true,
	}))

	route.Use(mdw.DBTransactionMdw(s.database.GetDB()))
}
