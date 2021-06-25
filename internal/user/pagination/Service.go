package pagination

import (
	"fmt"
	"math"
	"net/http"

	"strconv"

	"gorm.io/gorm"
)

func NewPagination(req *http.Request) Pagination {
	var limit int = 2
	var page int = 1
	sort := "created_at asc"

	limit, _ = strconv.Atoi(req.URL.Query().Get("limit"))

	page, _ = strconv.Atoi(req.URL.Query().Get("page"))

	if sort = req.URL.Query().Get("sort"); sort == "" {
		sort = "created_at asc"
	}

	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func (pag *Pagination) BuildLink(totalPages int) {

	pag.FirstPage = fmt.Sprint(pag.Limit, "/", 0)
	pag.LastPage = fmt.Sprint(pag.Limit, "/", totalPages)

	if pag.Page > 0 {
		// set previous page pagination response
		pag.PreviousPage = fmt.Sprint(pag.Limit, "/", pag.Page-1)
	}

	if pag.Page < totalPages {
		// set next page pagination response
		pag.NextPage = fmt.Sprint(pag.Limit, "/", pag.Page+1)
	}

	if pag.Page > totalPages {
		// reset previous page
		pag.PreviousPage = ""
	}

}

func (pag Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if pag.Page == 0 {
			pag.Page = 1
		}

		switch {
		case pag.Limit > 100:
			pag.Limit = 100
		case pag.Limit <= 0:
			pag.Limit = 10
		}

		offset := (pag.Page - 1) * pag.Limit
		return db.Offset(offset).Limit(pag.Limit)
	}
}

func (pg *Pagination) Builder(totalRows int64) error {

	totalPages, toRow, fromRow := 0, 0, 0

	pg.TotalRows = totalRows
	totalPages = int(math.Ceil(float64(totalRows)/float64(pg.Limit))) - 1

	if pg.Page == 0 {

		fromRow = 1
		toRow = pg.Limit
	} else {
		if pg.Page <= totalPages {

			fromRow = pg.Page*pg.Limit + 1
			toRow = (pg.Page + 1) * pg.Limit
		}
	}

	if toRow > int(totalRows) {

		toRow = int(totalRows)
	}

	pg.FromRow = fromRow
	pg.ToRow = toRow

	pg.BuildLink(totalPages)
	return nil
}
