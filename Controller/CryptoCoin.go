package Controller

import (
	"bitcoin-checker/Models"
	"bitcoin-checker/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPrices(c *gin.Context) {

	date := c.Query("date")
	limit := Utils.StringToInt(c.Query("limit"))
	offset := Utils.StringToInt(c.Query("offset"))
	if bitcoin, err := Models.FetchBitcoinInfo(date, limit, offset); err != nil {
		// Abort api call with error response
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		bitcoin.Url = fmt.Sprintf("<https://api.coingecko.com/api/v3/simple/price?date=%s&offset=%d&limit=%d>", date, offset, limit)
		bitcoin.Next = fmt.Sprintf("<https://api.coingecko.com/api/v3/simple/price?date=%s&offset=%d&limit=%d>", date, (offset + 100), limit)
		c.JSON(http.StatusOK, bitcoin)
	}
}
