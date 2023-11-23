package jiraticket

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	name := "test"
	fmt.Println(name)
	r.Run()
}
