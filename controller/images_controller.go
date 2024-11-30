package controller

import (
	"fmt"
	"myapp/model"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UploadImage handles POST requests to upload a new image.
func UploadImage(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	file, err := c.FormFile("image") // make sure that the form field is named 'image'

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the file to the server's local storage
	filePath := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save the image metadata to the database
	image := model.Image{
		Title:       title,
		Description: description,
		Path:        filePath,
	}
	if err := model.CreateImage(&image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully!", "image": image})
}

// GetImage handles GET requests to retrieve an image's metadata.
func GetImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	image, err := model.GetImageByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image": image})
}

// UpdateImage handles POST requests to update an existing image's metadata.
func UpdateImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")

	image, err := model.GetImageByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	image.Title = title
	image.Description = description
	if err := model.UpdateImage(image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image updated successfully!", "image": image})
}

// DeleteImage handles DELETE requests to delete an image.
func DeleteImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	if err := model.DeleteImage(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully!"})
}

// RegisterImageRoutes sets up the routing for image operations.
func RegisterImageRoutes(router *gin.Engine) {
	router.POST("/images", UploadImage)
	router.GET("/images/:id", GetImage)
	router.POST("/images/:id", UpdateImage)
	router.DELETE("/images/:id", DeleteImage)
}

