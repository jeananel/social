package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jeananel/social.git/middleware"
	"github.com/jeananel/social.git/routers"
	"github.com/rs/cors"
)

//Managements for set server and route listener ports
func Managements() {
	router := mux.NewRouter()

	//Routes main login and register users
	router.HandleFunc("/Register", middleware.CheckConnectionToDatabase(routers.Register)).Methods("POST")
	router.HandleFunc("/Login", middleware.CheckConnectionToDatabase(routers.Login)).Methods("POST")

	//Profile user
	router.HandleFunc("/ViewProfile", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/UpdateProfile", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.UpdateProfile))).Methods("PUT")

	//Management Tweets
	router.HandleFunc("/SaveTweet", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/GetTweets", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/RemoveTweet", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.RemoveTweet))).Methods("DELETE")
	router.HandleFunc("/GetTweetsFollowes", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.GetTweetsFollowes))).Methods("GET")

	//Management files in server routes
	router.HandleFunc("/UploadAvatar", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/UploadBanner", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/GetFileAvatar", middleware.CheckConnectionToDatabase(routers.GetFileAvatar)).Methods("GET")
	router.HandleFunc("/GetFileBanner", middleware.CheckConnectionToDatabase(routers.GetFileBanner)).Methods("GET")

	//Follow and unfollow routes
	router.HandleFunc("/FollowUser", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.FollowUser))).Methods("POST")
	router.HandleFunc("/UnfollowUser", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.UnfollowUser))).Methods("DELETE")
	router.HandleFunc("/CheckFollowing", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.CheckFollowing))).Methods("GET")

	//Get users following and followers
	router.HandleFunc("/GetUsers", middleware.CheckConnectionToDatabase(middleware.ValidationToken(routers.GetUsers))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	//Access to any users to API
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
