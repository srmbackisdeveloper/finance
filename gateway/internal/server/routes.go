package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

)

func (s *Server) RegisterRoutes() http.Handler {
    routes := gin.Default()

    // add cors validations 
    routes.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           6 * time.Hour,
    }))

    public := routes.Group("")
    {
        public.POST("/register", s.registerHandler)
        public.POST("/register/verify", s.verifyHandler)
        public.POST("/login", s.loginHandler)
        public.POST("/refresh", s.refreshTokenHandler)
    }

    private := routes.Group("")
    private.Use(s.authMiddleware)
    {
        private.GET("/user", s.getUserHandler)
        private.PATCH("/user", s.updateUserHandler)
        private.DELETE("/user", s.deleteUserHandler)
    }

    return routes
}