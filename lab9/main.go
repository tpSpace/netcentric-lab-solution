package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jaswdr/faker"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Dob       string `json:"dob"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Street    string `json:"street"`
	Address   string `json:"address"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the User model to create the users table
	db.AutoMigrate(&User{})

	// Create a new Gin router
	router := gin.Default()

	// Define routes and their handlers
	// Base URL
	// localhost:8080/v1/
	{
		v1 := router.Group("/v1")
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API"})
		})
		v1.POST("/users", createUser)
		v1.GET("/users", getUsers)
		v1.GET("/users/:id", getUser)
		v1.PUT("/users/:id", updateUser)
		v1.DELETE("/users/:id", deleteUser)
		v1.GET("/users/username/:username", findUserByUsername)
		v1.GET("/users/firstname/:firstname", findUserByFirstname)
		v1.GET("/users/lastname/:lastname", findUserByLastname)
		// create fake data to the database
		v1.GET("/users/fake/:count", func(c *gin.Context) {
			count, _ := strconv.Atoi(c.Param("count"))
			insertData(db, count)
			c.JSON(http.StatusOK, gin.H{"message": "Fake data created"})
		})

	}

	// Run the server
	router.Run(":8080")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created", "user": user})
}

func getUsers(c *gin.Context) {
	var users []User
	if result := db.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if result := db.First(&user, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if result := db.Model(&User{}).Where("id = ?", id).Updates(user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	if result := db.Delete(&User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// create a function find user base on username in the database
func findUserByUsername(c *gin.Context) {
	username := c.Param("username")
	var user User
	if result := db.Where("username = ?", username).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// create a function to find user by firstname
func findUserByFirstname(c *gin.Context) {
	firstname := c.Param("firstname")
	var user User
	if result := db.Where("firstname = ?", firstname).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// create a function to find user by lastname
func findUserByLastname(c *gin.Context) {
	lastname := c.Param("lastname")
	var user User
	if result := db.Where("lastname = ?", lastname).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// INSERT INTO users (username, firstname, lastname, email, avatar, phone, dob, country, city, street, address) VALUES ('john_doe', 'John', 'Doe', '
func insertData(db *gorm.DB, count int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		fake := faker.New()
		rand.New(rand.NewSource(time.Now().UnixNano()))
		now := time.Now()
		user := User{
			Username:  fake.Person().FirstName() + strconv.Itoa(rand.Intn(1000)),
			Firstname: fake.Person().FirstName(),
			Lastname:  fake.Person().LastName(),
			Email:     fake.Internet().Email(),
			Avatar:    fake.Internet().URL(),
			Phone:     fake.Phone().Number(),
			Dob:       now.Format(time.RFC3339),
			Country:   fake.Address().Country(),
			City:      fake.Address().City(),
			Street:    fake.Address().StreetName(),
			Address:   fake.Address().Address(),
		}
		db.Create(&user)
	}
}
