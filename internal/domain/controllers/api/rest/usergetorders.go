package rest

import "net/http"

func (h RESTControllersImpl) UserGetOrders(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	h.userService.UserGetOrders()

	res := `
	[
		  {
			  "number": "9278923470",
			  "status": "PROCESSED",
			  "accrual": 500,
			  "uploaded_at": "2020-12-10T15:15:45+03:00"
		  },
		  {
			  "number": "12345678903",
			  "status": "PROCESSING",
			  "uploaded_at": "2020-12-10T15:12:01+03:00"
		  },
		  {
			  "number": "346436439",
			  "status": "INVALID",
			  "uploaded_at": "2020-12-09T16:09:53+03:00"
		  }
	  ]
	`

	w.Write([]byte(res))
	w.WriteHeader(http.StatusOK)
}
