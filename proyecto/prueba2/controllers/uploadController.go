package controllers

import (
	"bufio"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"proyecto/prueba2/model"
	"proyecto/prueba2/shared"
)

func UploadController(c *gin.Context)  {
	response := shared.Context{}
	file,_:=c.FormFile("file")
	f,_ := file.Open()

	read:= csv.NewReader(bufio.NewReader(f))
	linen := 0

	var upload []model.Upload
	for {
		line, err :=read.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if len(line) > 0 && linen > 0  {
			data := model.Upload{}
			data.Organization = line[0]
			data.User.Name = line[1]
			if data.User.Name == line[1] {
				data.User.Roles = append(data.User.Roles, line[2])
			}
			upload = append(upload, data)
		}
		linen ++
	}
	c.JSON(200, response.ResponseData(shared.StatusSuccess,"",upload))

}