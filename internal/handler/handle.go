package handler

import (
	"errors"
	"fmt"
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
	
	valType, err := validAndGetParsedVal(c, &c.Params, "type")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}

	a, err := validAndGetParsedVal(c, &c.Params, "a")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}

	b, err := validAndGetParsedVal(c, &c.Params, "b")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}


	valA, err := strconv.Atoi(a)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}
	
	valB, err := strconv.Atoi(b)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return 
	}

	cal.Do(valA, valB, valType)
}

func validAndGetParsedVal(c *gin.Context, params *gin.Params, key string) (string, error) {

	fmt.Println("key: ", key)
	val := c.Query(key)

	if val == ""  {
		logger.Error("type input param must not be null or empty")
		return "", errors.New("Input '" + key + "' must not be null or empty")
	}

	return val, nil
}

func Sub(c *gin.Context){
	log.Println("calling lib Do function for subtraction")
	// cal.Do()
}
