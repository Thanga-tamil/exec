package handler

import (
	"errors"
	"net/http"
	"strconv"

	cal "github.com/Thanga-tamil/lib"
	"github.com/gin-gonic/gin"

	"github.com/Thanga-tamil/exec/internal/response"
	logger "github.com/Thanga-tamil/exec/internal/utils"
)

func Calculate(ginCtx *gin.Context){
	
	valType, err := validAndGetParsedVal(ginCtx, &ginCtx.Params, "type")
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}

	a, err := validAndGetParsedVal(ginCtx, &ginCtx.Params, "a")
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}

	b, err := validAndGetParsedVal(ginCtx, &ginCtx.Params, "b")
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}

	valA, err := strconv.Atoi(a)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}
	
	valB, err := strconv.Atoi(b)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return 
	}

	logger.Info("Invoking lib Do function for: ", valType)
	result := cal.Do(valA, valB, valType)

	ginCtx.JSON(http.StatusOK, response.Success("Calculation Completed", 200, result))
}

func validAndGetParsedVal(ginCtx *gin.Context, params *gin.Params, key string) (string, error) {

	val := ginCtx.Query(key)

	if val == ""  {
		logger.Error("Input '" + key + "' param must not be null or empty")
		return "", errors.New("Input '" + key + "' must not be null or empty")
	}

	return val, nil
}

