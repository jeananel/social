package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/jeananel/social.git/bd"
	"github.com/jeananel/social.git/models"
)

var absolutePathAvatar string = "upload/avatars/"
var absolutePathBanner string = "upload/banners/"

//Abstraer codigo repetido para cargar archivos en Utils

/*GetFileAvatar send  Avatar  HTTP */
func GetFileAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	perfil, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open(absolutePathAvatar + perfil.Avatar)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copy file", http.StatusBadRequest)
	}
}

/*GetFileBanner send  Avatar  HTTP */
func GetFileBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	perfil, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open(absolutePathBanner + perfil.Avatar)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copy file", http.StatusBadRequest)
	}
}

/*UploadAvatar upload file avatar to server */
func UploadAvatar(write http.ResponseWriter, request *http.Request) {

	file, handler, err := request.FormFile("Avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileUpload string = absolutePathAvatar + IDUser + "." + extension

	f, err := os.OpenFile(fileUpload, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(write, "Upload file error ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(write, "Copy file error ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = bd.UpdateUser(user, IDUser) //Update avatar of user
	if err != nil || status == false {
		http.Error(write, "An error ocurred while save avatar in database! "+err.Error(), http.StatusBadRequest)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
}

/*UploadBanner upload file avatar to server */
func UploadBanner(write http.ResponseWriter, request *http.Request) {

	file, handler, err := request.FormFile("Banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileUpload string = absolutePathBanner + IDUser + "." + extension

	f, err := os.OpenFile(fileUpload, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(write, "Upload file error ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(write, "Copy file error ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = bd.UpdateUser(user, IDUser) //Update banner of user
	if err != nil || status == false {
		http.Error(write, "An error ocurred while save banner in database! "+err.Error(), http.StatusBadRequest)
		return
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
}
