package server

import (
	"net/http"

	"github.com/TroJanBoi/temparary/cmd/api/docs"
	"github.com/TroJanBoi/temparary/internal/conf"
	"github.com/TroJanBoi/temparary/internal/services/controller"
	"github.com/TroJanBoi/temparary/internal/services/repository"
	"github.com/TroJanBoi/temparary/internal/services/usecases"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-User-ID")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) Router() (http.Handler, func()) {
	config := conf.NewConfig()
	r := gin.Default()

	r.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v2"

	paymentRepository := repository.NewPaymentRepository(s.db)
	paymentUseCase := usecases.NewPaymentUseCases(paymentRepository)
	paymentController := controller.NewPayment(paymentUseCase)

	api := r.Group("/api/v2")
	{
		paymentGroup := api.Group("/payment")
		{
			paymentController.PaymentRoutes(paymentGroup)
		}
	}
	if config.ENV == "dev" || config.ENV == "uat" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return r, func() {
		// sqlDB, err := s.db.DB()
		// if err != nil {
		// 	panic("Failed to get sql.DB from gorm.DB")
		// }
		// sqlDB.Close()
	}
}
