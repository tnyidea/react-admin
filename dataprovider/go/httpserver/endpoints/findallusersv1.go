package endpoints

import (
	"github.com/tnyidea/go-sample-userdata/models"
	"github.com/tnyidea/react-admin/dataprovider/go/httpserver/response"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func FindAllUsersV1(r *http.Request, db models.DB) response.DataProviderResponseData {
	log.Println("=== Executing FindAllUsersV1 ===")
	defer log.Println("=== FindAllUsersV1 Execution Complete ===")

	// https://jsonplaceholder.typicode.com/users?_end=10&_order=ASC&_sort=id&_start=0
	query := r.URL.Query()
	listStart, err := strconv.Atoi(query.Get("_start"))
	if err != nil {
		// 400 error
	}
	listEnd, err := strconv.Atoi(query.Get("_end"))
	if err != nil {
		// 400 error
	}
	listSortColumn := query.Get("_sort")
	if listSortColumn == "" {
		// 400 error
	}
	listSortOrder := query.Get("_order")
	if listSortOrder == "" {
		// 400 error
	}

	users, err := db.FindAllUsers()
	if err != nil {
		// 500 error
	}
	totalCount, err := db.Count()
	if err != nil {
		// 500 error
	}

	// TODO determine best way to sort on other struct fields
	if listSortOrder == "ASC" {
		sort.Slice(users, func(i, j int) bool {
			return users[i].Id < users[j].Id
		})
	}
	if listSortOrder == "DESC" {
		sort.Slice(users, func(i, j int) bool {
			return users[i].Id > users[j].Id
		})
	}

	return response.DataProviderResponseData{
		Status:     http.StatusOK,
		Items:      users[listStart:listEnd],
		TotalCount: totalCount,
	}
}
