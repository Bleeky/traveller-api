package userservice

import (
	"github.com/emicklei/go-restful"
)

type User struct {
	ID, Name, Token string
}

// type UserRessource struct {
// 	Name     string `json: "name"`
// 	Password string `json: "password`
// }

func New() *restful.WebService {
	service := new(restful.WebService)
	service.Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	service.Route(service.GET("/{user-id}").To(FindUser))
	service.Route(service.GET("").To(GetUsers))
	return service
}
func GetUsers(request *restful.Request, response *restful.Response) {

}
func FindUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	// here you would fetch user from some persistence system
	usr := &User{ID: id, Name: "John Doe"}
	response.WriteEntity(usr)
}
