package main

import (
	"net/http"
	"strconv"
	"time"

	"example.com/captable"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCertificates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, captable.GetCertificates())
}

func GetCertificateByID(c *gin.Context) {
	uuid, _ := uuid.Parse(c.Param("id"))
	c.IndentedJSON(http.StatusOK, captable.GetCertificateByID(uuid))
}

type NewCertInput struct {
	Owner     string `json:"owner"`
	Quantity  string `json:"quantity"`
	IssueDate string `json:"issueDate"`
}

func AddCertificate(c *gin.Context) {
	var newCertificateData NewCertInput
	err := c.BindJSON(&newCertificateData)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	issueDate, _ := time.Parse(time.RFC3339, newCertificateData.IssueDate)
	quantity, _ := strconv.Atoi(newCertificateData.Quantity)

	newCertificate := captable.NewCertificate{
		Owner:     newCertificateData.Owner,
		Quantity:  quantity,
		IssueDate: issueDate,
	}
	c.IndentedJSON(http.StatusOK, captable.AddCertificate(newCertificate))

}

func GetCapTable(c *gin.Context) {
	asOf := c.Query("as_of")
	parsedTime := time.Now()
	if asOf != "" {
		parsedTime, _ = time.Parse(time.RFC3339, asOf)
	}
	c.IndentedJSON(http.StatusOK, captable.GetCapTable(parsedTime))
}

func main() {
	router := gin.Default()
	router.GET("/certificates", GetCertificates)
	router.POST("/certificates", AddCertificate)
	router.GET("/certificates/:id", GetCertificateByID)
	router.GET("/captable", GetCapTable)

	router.Run("localhost:8080")
}
