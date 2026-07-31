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

func Calculate(c *gin.Context){
	log.Println("Invoking lib Do function for addition")
	
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

	result := cal.Do(valA, valB, valType)

	if result < 0 {
		msg := "bad result / may functionality not implemented yet"
		c.JSON(http.StatusBadRequest, response.Error(msg, 400)); return 
	}
 
	c.JSON(http.StatusOK, response.Success("Calculation Completed", 200, result))
}

func validAndGetParsedVal(c *gin.Context, params *gin.Params, key string) (string, error) {

	val := c.Query(key)

	if val == ""  {
		logger.Error("type input param must not be null or empty")
		return "", errors.New("Input '" + key + "' must not be null or empty")
	}

	return val, nil
}

