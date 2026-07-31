package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	cal "github.com/Thanga-tamil/lib"
	"github.com/gin-gonic/gin"

	"github.com/Thanga-tamil/exec/internal/response"
	logger "github.com/Thanga-tamil/exec/internal/utils"
)

func Add(c *gin.Context){
	log.Println("calling lib Do function for addition")
	
	Type, err := validAndGetParsedVal(c, &c.Params, "type")

	a, err := validAndGetParsedVal(c, &c.Params, "a")
	b, err := validAndGetParsedVal(c, &c.Params, "b")

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}

	// func Do (x, y int, arg string) 
	valA, err := strconv.Atoi(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error(), 400))
		return
	}
	
	valB, err := strconv.Atoi(b)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error(), 400))
		return 
	}

	cal.Do(valA, valB, Type)
}

func Sub(c *gin.Context){
	log.Println("calling lib Do function for subtraction")
	// cal.Do()
}

func validAndGetParsedVal(c *gin.Context, params *gin.Params, key string) (string, error) {

	val, exist := c.Params.Get(key)

	if !exist {
		logger.Error("type input param must not be null or empty")
		c.JSON(http.StatusBadRequest, response.Error("Input 'type' must not be null or empty", 400))
		return "", errors.New("null/empty")
	}

	return val, nil
}
