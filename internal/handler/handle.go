package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	cal "github.com/Thanga-tamil/lib"
	"github.com/gin-gonic/gin"

	"github.com/Thanga-tamil/exec/internal/response"
	logger "github.com/Thanga-tamil/exec/internal/utils"
)

func Calculate(ginCtx *gin.Context){
	
	valType, err := validAndGetParsedVal(ginCtx, "type"); errs(ginCtx, err)

	a, err := validAndGetParsedVal(ginCtx, "a"); errs(ginCtx, err)
	b, err := validAndGetParsedVal(ginCtx, "b"); errs(ginCtx, err)

	valA, err := strconv.Atoi(a); errs(ginCtx, err)
	valB, err := strconv.Atoi(b); errs(ginCtx, err)

	time.Sleep(time.Second*30)
	logger.Info("Invoking lib Do function for: ", valType)
	result := cal.Do(valA, valB, valType)

	ginCtx.JSON(http.StatusOK, response.Success("Calculation Completed", 200, result))
}

func errs(ginCtx *gin.Context, err error){
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Error(err.Error(), 400)); return
	}
}

func validAndGetParsedVal(ginCtx *gin.Context, key string) (string, error) {

	val := ginCtx.Query(key)

	if val == ""  {
		logger.Error("Input '" + key + "' param must not be null or empty")
		return "", errors.New("Input '" + key + "' must not be null or empty")
	}

	return val, nil
}

