package pr

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/pr/estimate", estimateHandler)
}

func estimateHandler(c *gin.Context) {
    var req EstimateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, ErrResp{Error: err.Error()})
        return
    }

    oneRM, err := Brzycki1RM(req.Weight, req.Reps)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, ErrResp{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, EstimateResponse{
        Input:   req,
        OneRM:   round2(oneRM),
        Table:   RepTable(oneRM, 15),
        Formula: "brzycki",
    })
}
