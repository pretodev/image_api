package server

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/replit/database-go"
	"net/http"
	"strconv"
	"strings"
)

type Controllers struct{}

// UpsertUserPhoto
// @Summary 	 Salva foto de usuário
// @Description  Salva foto de usuário utilizando o id
// @Tags 		 users
// @Accept 		 mpfd
// @Produce 	 json
// @Param 		 id   	path     	int  true  "ID do usuário"
// @Param 		 file 	formData 	file true  "Foto do usuário"
// @Success 	 200 	{object} 	models.UpsertUserPhotoResonse
// @Failure 	 400 	{object} 	models.ErrorResponse
// @Failure      500  	{object}  	models.ErrorResponse
// @Router /users/{id}/photo [patch]
func (c *Controllers) UpsertUserPhoto(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid file",
		})
		return
	}

	fileNameParts := strings.Split(file.Filename, ".")
	fileExtension := fileNameParts[len(fileNameParts)-1]
	if fileExtension != "jpg" && fileExtension != "jpeg" && fileExtension != "png" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Only jpg, jpeg and png files are allowed",
		})
		return
	}
	fileName := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d.%s", id, file.Filename)))
	err = ctx.SaveUploadedFile(file, "./static/"+fileName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error saving photo file",
		})
		return
	}

	err = database.Set(fmt.Sprintf("user_%d_photo", id), fileName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error saving photo",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"photo_url": fmt.Sprintf("%s/static/%s", ctx.Request.Host, fileName),
		"message":   fmt.Sprintf("File %s uploaded successfully", file.Filename),
	})
}

// GetUserPhoto
// @Summary 	 Recupera foto de usuário
// @Description  Recupera foto de usuário utilizando o id
// @Tags 		 users
// @Produce 	 json
// @Param 		 id   	path     	int  true  "ID do usuário"
// @Success 	 200 	{object} 	models.UpsertUserPhotoResonse
// @Failure 	 400 	{object} 	models.ErrorResponse
// @Failure 	 404 	{object} 	models.ErrorResponse
// @Failure      500  	{object}  	models.ErrorResponse
// @Router /users/{id}/photo [get]
func (c *Controllers) GetUserPhoto(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
		return
	}
	photo, err := database.Get(fmt.Sprintf("user_%d_photo", id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User photo not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"photo_url": fmt.Sprintf("%s/static/%s", ctx.Request.Host, photo),
	})
}
