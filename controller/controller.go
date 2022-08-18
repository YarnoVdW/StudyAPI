package controller

import (
	"net/http"

	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/config"
	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/model"
	jwtapple2 "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// checks if jsonobject with user values is inside the request, then checks
// if the user already exists and sends a response with a message
func RegisterEndPoint(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userCheck model.User
	config.GetDb().First(&userCheck, "username = ?", user.Username)

	if userCheck.ID > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	}

	config.GetDb().Save(&user)

	c.JSON(http.StatusCreated, gin.H{"message: ": "User created successfully!"})
}

// creates study item associated with a User
func CreateStudyItem(c *gin.Context) {
	claims := jwtapple2.ExtractClaims(c)

	var user model.User
	config.GetDb().Where("id = ?", claims[config.IdentityKey]).First(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var studyItem model.StudyItem
	if err := c.ShouldBindJSON(&studyItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studyItem.ID = user.ID
	config.GetDb().Save(&studyItem)
	c.JSON(http.StatusCreated, gin.H{"message": "Study Item created succesfully", "task": studyItem})

}

func FetchAllStudyitems(c *gin.Context) {
	claims := jwtapple2.ExtractClaims(c)

	var user model.User
	config.GetDb().Where("id = ?", claims[config.IdentityKey]).First(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var studyItems []model.StudyItem
	config.GetDb().Where("user_id = ?", user.ID).Order("created_at desc").Find(&studyItems)

	if len(studyItems) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No studyItems found!", "data": studyItems})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": studyItems})
}

func FetchSingleTask(c *gin.Context) {
	studyItemID := c.Param("id")

	if len(studyItemID) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var studyItem model.StudyItem
	config.GetDb().First(&studyItem, studyItemID)

	if studyItem.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No studyItem found!"})
		return
	}

	c.JSON(http.StatusOK, studyItem)
}
