package main

import (
	_ "encoding/json"
	"github.com/BurntSushi/toml"
	// "github.com/amulya.bl/demoCRUDApp/config"
	"github.com/amulya.bl/demoCRUDApp/handlers"
	"github.com/amulya.bl/demoCRUDApp/repositories"
	"github.com/amulya.bl/demoCRUDApp/routes"
	"github.com/amulya.bl/demoCRUDApp/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	_ "log"
	"net/http"
	_ "strconv"
)

type Config struct {
	Database struct {
		User     string `toml:"User"`
		Password string `toml:"Password"`
		Name     string `toml:"Name"`
		Host     string `toml:"Host"`
		Port     string `toml:"Port"`
	} `toml:"Database"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(filename, &config)
	return config, err
}
func main() {
	r := mux.NewRouter()
	config, err := LoadConfig("/Users/amulya.bl/demoCRUDApp/config/config.toml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dsn := config.Database.User +
		":" + config.Database.Password +
		"@tcp(" + config.Database.Host +
		":" + config.Database.Port +
		")/" + config.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepository := repositories.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	routes.RegisterRoutes(r, userHandler)
	http.ListenAndServe(":8090", r)
}
