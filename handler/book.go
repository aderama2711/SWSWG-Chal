package handler

import (
	"P1/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	var newBook model.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	res, err := h.app.CreateBook(newBook)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) ListBook(c *gin.Context) {

	res, err := h.app.ListBook()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) GetBook(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	res, err := h.app.GetBook(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var newBook model.Book

	newBook.ID = id

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	res, err := h.app.UpdateBook(newBook)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	temp_id := c.Param("id")

	id, err := strconv.Atoi(temp_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id param must be integer",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var newBook model.Book
	newBook.ID = id

	err = h.app.DeleteBook(newBook)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
				"status":  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book deleted successfully"),
	})
}
