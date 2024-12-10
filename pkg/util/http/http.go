package httputil

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	dbutil "github.com/Storm-cpu/core/pkg/util/db"

	"github.com/imdatngo/gowhere"
	"github.com/labstack/echo/v4"
)

func ReqID(c echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return id, nil
}

// ListRequest holds data of listing request from react-admin
// Note: To add these parameters to swagger:operation, check the file /internal/util/swagger
// swagger:ignore
type ListRequest struct {
	// Number of records per page
	// default: 25
	Limit int `json:"l,omitempty" query:"l" validate:"max=300"`
	// Current page number
	// default: 1
	Page int `json:"p,omitempty" query:"p"`
	// Field name for sorting
	// default:
	Sort string `json:"s,omitempty" query:"s"`
	// Sort direction, must be one of ASC, DESC
	// default:
	Order string `json:"o,omitempty" query:"o"`
	// JSON string of filter. E.g: {"field_name":"value"}
	// default:
	Filter string `json:"f,omitempty" query:"f"`
}

// ReqListQuery parses url query string for listing request
func ReqListQuery(c echo.Context) (*dbutil.ListQueryCondition, error) {
	isValidParams := regexp.MustCompile(`^[a-zA-Z0-9._"]*$`).MatchString

	lr := &ListRequest{}
	if err := c.Bind(lr); err != nil {
		return nil, err
	}

	lq := &dbutil.ListQueryCondition{
		Page:    lr.Page,
		PerPage: lr.Limit,
		Filter:  gowhere.WithConfig(gowhere.Config{Strict: true}),
	}

	if lr.Filter != "" {
		var filter interface{}
		err := json.Unmarshal([]byte(lr.Filter), &filter)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err)
		}

		if err := lq.Filter.Where(filter).Build().Error; err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}

	if lr.Sort != "" {
		if !isValidParams(lr.Sort) || len(lr.Sort) > 50 {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid params for sort")
		}
		sortField := lr.Sort
		sortOrder := "ASC" // default
		if lr.Order != "" && strings.ToLower(lr.Order) == "desc" {
			sortOrder = "DESC"
		}
		lq.Sort = []string{sortField + " " + sortOrder}
	}

	return lq, nil
}
