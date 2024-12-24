package server

import (
	"math/rand"
	"net/http"

	_ "northal.com/docs"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

func InitDocs(router *gin.Engine) {

	router.GET("/docs", func(c *gin.Context) {

		defaultToken := "your-default-token-here"

		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			Theme:   scalar.ThemeId([]string{string(scalar.ThemeBluePlanet), string(scalar.ThemeDeepSpace)}[rand.Intn(2)]), // 随机选择主题
			CustomOptions: scalar.CustomOptions{
				PageTitle: "North GO API DOCS",
			},
			MetaData:       "{{ meta_token }}",
			Authentication: defaultToken,
			DarkMode:       true,
		})

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, htmlContent)

	})
}
