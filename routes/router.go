package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"leaguesimulator/league"

	"leaguesimulator/prediction"
)

var manager league.LeagueManager

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/init-league", func(c *gin.Context) {
		manager.InitLeague()
		c.JSON(http.StatusOK, gin.H{"message": "League initialized"})
	})

	router.POST("/next-week", func(c *gin.Context) {
		matches := manager.PlayNextWeek()
		if matches == nil {
			c.JSON(http.StatusOK, gin.H{"message": "League finished"})
			return
		}
		c.JSON(http.StatusOK, matches)
	})

	router.GET("/standings", func(c *gin.Context) {
		c.JSON(http.StatusOK, manager.GetStandings())
	})

	router.GET("/matches", func(c *gin.Context) {
		c.JSON(http.StatusOK, manager.GetMatches())
	})

	router.GET("/predict", func(c *gin.Context) {
		predictions, err := prediction.RunPrediction()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, predictions)
	})

	return router
}
