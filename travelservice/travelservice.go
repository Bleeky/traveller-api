package travelservice

import (
	// "fmt"
	"github.com/emicklei/go-restful"
	// "net/http"
	"traveller-api/utils"
)

// TravelResource
type TravelResource struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	DateBegin   string `db:"date_begin" json:"date_begin"`
	DateEnd     string `db:"date_end" json:"date_end"`
}

//New WebService
func New() *restful.WebService {
	service := new(restful.WebService)
	service.Path("/travels").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	service.Route(service.GET("").To(GetTravels))
	service.Route(service.GET("/{travel-id}").To(FindTravel))
	service.Route(service.POST("").To(CreateTravel))
	service.Route(service.PUT("/{travel-id}").To(UpdateTravel))
	service.Route(service.DELETE("/{travel-id}").To(DeleteTravel))
	return service
}
func GetTravels(request *restful.Request, response *restful.Response) {
	session := utils.NewSession()
	rows, err := session.DB.Query("SELECT id, name, description, date_begin, date_end FROM travel")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var results []TravelResource
	var travel TravelResource
	for rows.Next() {
		err := rows.Scan(&travel.ID, &travel.Name, &travel.Description, &travel.DateBegin, &travel.DateEnd)
		if err != nil {
			panic(err)
		}
		results = append(results, travel)
	}
	session.Close()
	response.WriteAsJson(results)
}
func UpdateTravel(request *restful.Request, response *restful.Response) {

}
func DeleteTravel(request *restful.Request, response *restful.Response) {

}
func FindTravel(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("travel-id")
	session := utils.NewSession()
	rows, err := session.DB.Query("SELECT id, name, description, date_begin, date_end FROM travel WHERE id = " + id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var travel TravelResource
	for rows.Next() {
		err := rows.Scan(&travel.ID, &travel.Name, &travel.Description, &travel.DateBegin, &travel.DateEnd)
		if err != nil {
			panic(err)
		}
	}
	session.Close()
	response.WriteAsJson(travel)
}
func CreateTravel(request *restful.Request, response *restful.Response) {
	// travel := new(Travel)
	// err := request.ReadEntity(&travel)
	// session := utils.NewSession()
	// c := session.DB("traveller").C("travels")
	// c.Insert(&travel)
	// session.Close()
	// if err == nil {
	// 	response.WriteEntity(travel)
	// } else {
	// 	response.WriteError(http.StatusInternalServerError, err)
	// }
}
