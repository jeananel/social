package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/jeananel/social.git/bd"
	"github.com/jeananel/social.git/models"
)

/*SaveTweet save tweet in database */
func SaveTweet(write http.ResponseWriter, request *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(request.Body).Decode(&message)

	object := models.Tweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	//Validations
	if len(object.Message) == 0 {
		http.Error(write, "Message is required.", 400)
		return
	}

	_, status, err := bd.InsertTweet(object)
	if err != nil {
		http.Error(write, "An error occurred in insert register. Please, try again."+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(write, "Error in save tweet", 400)
		return
	}

	write.WriteHeader(http.StatusCreated)

}

/*GetTweets Leo los tweets */
func GetTweets(write http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(write, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	if len(request.URL.Query().Get("page")) < 1 {
		http.Error(write, "Parameter page is required", http.StatusBadRequest)
		return
	}

	pagin, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		http.Error(write, "Page value should be greater than 0", http.StatusBadRequest)
		return
	}

	page := int64(pagin)
	respuesta, isOk := bd.GetTweets(ID, page)
	if isOk == false {
		http.Error(write, "An error ocurred while read tweets", http.StatusBadRequest)
		return
	}

	write.Header().Set("Content-type", "application/json")
	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(respuesta)
}

/*RemoveTweet remove tweet by id and userID */
func RemoveTweet(write http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(write, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(write, "An error ocurred while tried deleted tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	write.Header().Set("Content-type", "application/json")
	write.WriteHeader(http.StatusCreated)
}

/*GetTweetsFollowes get tweets of followers */
func GetTweetsFollowes(write http.ResponseWriter, response *http.Request) {

	if len(response.URL.Query().Get("page")) < 1 {
		http.Error(write, "Page parameter is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(response.URL.Query().Get("page"))

	if err != nil {
		http.Error(write, "Page value should be greater than 0", http.StatusBadRequest)
		return
	}

	result, isOk := bd.GetTweetsFollowes(IDUser, page)

	if isOk == false {
		http.Error(write, "Error read tweets of followers", http.StatusBadRequest)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(result)
}
