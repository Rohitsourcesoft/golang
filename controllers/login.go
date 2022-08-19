package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
)
func LoginUser(c *gin.Context) {
  c.JSON(200, gin.H{
      "message": "hello world",
  })
}
