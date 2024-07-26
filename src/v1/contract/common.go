package contract

import (
	"context"
	//"fmt"
	"net/http"
	"strconv"

	//"strings"

	"github.com/go-chi/chi/v5"
	//"github.com/go-sql-driver/mysql"
)

type GetListParam struct {
	Page    int    `json:"page" db:"page"`
	Limit   int    `json:"limit" db:"limit"`
	Offset  int    `json:"offset" db:"offset"`
	Keyword string `json:"keyword" db:"keyword"`
	Sort    string `json:"sort"`
}

// ValidateQuery return common converted parameter from query parameter for get list data
// common query parameter is keyword, page, limit, and offset
// page is number page where the data is now, keyword is for search data by string keyword,
// limit is limit data loaded per page, offset is number data skiped when loaded data
// data page and limit from query parameter is always number in string
// its need to converted to int, it will return error if page and limit is not a number
func ValidateAndBuildGetListRequest(r *http.Request) (getListParam *GetListParam, err error) {
	// default value for page and limit
	page, limit := 1, 10

	// get data from query parameter
	queryParams := r.URL.Query()
	limitQuery := queryParams.Get("limit")
	pageQuery := queryParams.Get("page")
	keyword := queryParams.Get("keyword")
	sort := queryParams.Get("sort")

	// query param validation
	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			return
		}
	}

	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			return
		}
	}

	// offset for OFFSET in get list query
	offset := (page - 1) * limit
	getListParam = &GetListParam{
		Page:    page,
		Limit:   limit,
		Offset:  offset,
		Keyword: keyword,
		Sort:    sort,
	}

	return
}

func ValidateIDParamRequest(ctx context.Context) (id int, err error) {
	idParam := chi.URLParamFromCtx(ctx, "id")
	id, err = strconv.Atoi(idParam)
	return
}

// func AddPercentToStrings(value string) (arrsql interface{}) {
// 	strings := strings.Split(value, ",")

// 	var result []string
// 	for _, str := range strings {
// 		result = append(result, fmt.Sprintf("%%%s%%", str))
// 	}

// 	arrsql = pg.Array(result)
// 	return
// }
