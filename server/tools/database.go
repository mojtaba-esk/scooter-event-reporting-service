package tools

import (
	"log"
	"net/http"
	"scootin/global"
	"strconv"
)

/*------------------------------*/

type LimitOffset struct {
	Limit  int
	Offset int
	Page   int
}

func GetLimitOffset(req *http.Request) LimitOffset {
	qryParams := req.URL.Query()

	page := 1
	if _, ok := qryParams["page"]; ok {

		var err error
		page, err = strconv.Atoi(qryParams["page"][0])
		if err != nil {
			log.Printf("Error in page number: %v", err)
			page = 1
		}
		if page <= 0 {
			page = 1
		}
	}

	limit := global.RowsPerPage
	offset := (page - 1) * limit

	return LimitOffset{limit, offset, page}
}

/*------------------------------*/
