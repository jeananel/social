package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jeananel/social.git/bd"
	"github.com/jeananel/social.git/models"
)

/*ResponseConsultUserFollower get status between relation two users  */
type ResponseConsultUserFollower struct {
	Status bool `json:"status"`
}

/*FollowUser follow user, relation between users */
func FollowUser(write http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(write, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	var object models.UsersFollowers
	object.UserID = IDUser
	object.FollowerID = ID

	status, err := bd.InsertUsersFollowers(object)
	if err != nil {
		http.Error(write, "An error ocurred while tried insert users follower  "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(write, "Dont insert user follower "+err.Error(), http.StatusBadRequest)
		return
	}
	write.WriteHeader(http.StatusCreated)
}

/*UnfollowUser  unfollow user, relation between users */
func UnfollowUser(write http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	var t models.UsersFollowers
	t.UserID = IDUser
	t.FollowerID = ID

	status, err := bd.DeleteUsersFollowers(t)
	if err != nil {
		http.Error(write, "An error ocurred while tried remove users follower "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(write, "Dont remove user follower "+err.Error(), http.StatusBadRequest)
		return
	}
	write.WriteHeader(http.StatusCreated)
}

/*CheckFollowing verify relation of following between users */
func CheckFollowing(write http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")

	var object models.UsersFollowers
	object.UserID = IDUser
	object.FollowerID = ID

	var response ResponseConsultUserFollower

	status, err := bd.CheckFollowing(object)
	if err != nil || status == false {
		log.Fatal(err.Error())
		response.Status = false
	} else {
		response.Status = true
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(response)
}

/*GetUsers list users relations followers, following */
func GetUsers(write http.ResponseWriter, request *http.Request) {

	typeUser := request.URL.Query().Get("type")
	page := request.URL.Query().Get("page")
	search := request.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(write, "Page value should be greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.GetUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(write, "Error GetUsers", http.StatusBadRequest)
		return
	}
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(result)
}
