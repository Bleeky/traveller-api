package placeservice

import (
	"traveller-api/utils"

	"github.com/emicklei/go-restful"
)

type PlaceResource struct {
	Id     int    `db:"id" json:"id"`
	Coords string `db:"coords" json:"coords"`
}

func New() *restful.WebService {
	service := new(restful.WebService)
	service.Path("/places").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	service.Route(service.GET("").To(GetPlaces))
	service.Route(service.GET("/{place-id}").To(FindPlace))
	service.Route(service.POST("").To(CreatePlace))
	service.Route(service.PUT("/{place-id}").To(UpdatePlace))
	service.Route(service.DELETE("/{place-id}").To(DeletePlace))
	return service
}
func GetPlaces(request *restful.Request, response *restful.Response) {
	session := utils.NewSession()
	rows, err := session.DB.Query("SELECT id, coords FROM place")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var results []PlaceResource
	var place PlaceResource
	for rows.Next() {
		err := rows.Scan(&place.Id, &place.Coords)
		if err != nil {
			panic(err)
		}
		results = append(results, place)
	}
	session.Close()
	response.WriteAsJson(results)
}
func UpdatePlace(request *restful.Request, response *restful.Response) {

}
func DeletePlace(request *restful.Request, response *restful.Response) {

}
func FindPlace(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("travel-id")
	session := utils.NewSession()
	rows, err := session.DB.Query("SELECT id, coords FROM travel WHERE id = " + id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var place PlaceResource
	for rows.Next() {
		err := rows.Scan(&place.Id, &place.Coords)
		if err != nil {
			panic(err)
		}
	}
	session.Close()
	response.WriteAsJson(place)
}
func CreatePlace(request *restful.Request, response *restful.Response) {
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
