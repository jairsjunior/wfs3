package handlers

import (
	"net/http"

	"github.com/boundlessgeo/wt/model"
	"github.com/boundlessgeo/wt/ogc"
	"github.com/gin-gonic/gin"
)

func (h *HTTPServer) makeCollectionHandlers(d *model.DB) {
	h.router.GET("/collections", getCollectionsInfo(d))
	h.router.GET("/collections/:collid/schema", getCollectionInfo(d))
	// h.router.PUT("/collections/:collid/schema", updateCollectionInfo())
	// h.router.POST("/collections", createCollectionInfo(d))
	// h.router.DELETE("/collections/:collid", deleteCollection())
}

func getCollectionsInfo(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		cidbs := db.AllCollectionInfos()
		cis := make([]*ogc.CollectionInfo, 0)
		for _, v := range cidbs {
			cis = append(cis, v.CollectionInfo)
		}
		c.JSON(http.StatusOK, gin.H{"collections": cis})
	}
}

func getCollectionInfo(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collInfo := db.FindCollection(c.Param(":collid"))
		c.JSON(http.StatusOK, gin.H{"collection": collInfo})
	}
}
