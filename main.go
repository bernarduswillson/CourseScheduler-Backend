package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"singleservice/controllers"
	"singleservice/initializers"
	"singleservice/migrate"
	// "fmt"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	// run the migration
	migrate.Migrate()

	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins: []string{"http://localhost:3000", "https://ohl-fe.vercel.app", "http://localhost:5173", "https://monolith-labpro.up.railway.app/", "http://127.0.0.1:8000/"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Accept", "Accept-Encoding", "Content-Length", "X-CSRF-Token", "Authorization"},
	}))
	r.POST("/add", controllers.AddMatkul)
	r.GET("/get", controllers.GetMatkul)
	r.DELETE("/delete/:id", controllers.DeleteMatkul)
	r.POST("/schedule", controllers.ScheduleCourses)

	r.Run()
}
