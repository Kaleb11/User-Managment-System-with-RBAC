package handler

import (
	//"Auth/internal/database/gorm"
	//"Auth/internal/user"

	gorm "Auth/internal/database/gorm/user"
	user "Auth/internal/user/model"
	"Auth/internal/user/pagination"
	service "Auth/internal/user/service"
	tok "Auth/internal/user/service/token"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func AddUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body := req.Body
		usr := gorm.NewUser()
		err := json.NewDecoder(body).Decode(usr)
		fmt.Println(err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(err)
		defer body.Close()
		_, err = gorm.Createuser(usr)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func GetUsers(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	pg := pagination.NewPagination(req)
	gorm.Getusers(req.Context(), &pg)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Data{
		Data:    pg,
		Message: "success",
		Status:  200,
	})

}
func View(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	http.Handle("/*", http.FileServer(http.Dir("C:\\Users\\Tilefamily\\Documents\\AuthHexa\\cmd\\rest\\router\\views")))

}

func GetUsersByID(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	userID := ps.ByName("id")
	usrID, err := strconv.Atoi(userID)
	usr, err := gorm.Getusersbyid(usrID)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(err)
	err = json.NewEncoder(w).Encode(usr)

	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func GetUsersByusername(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	username := ps.ByName("username")

	//usrID, err := strconv.Atoi(userID)
	usr, err := gorm.Getusersbyname(username)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(err)
	err = json.NewEncoder(w).Encode(usr)

	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func UpdateUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		var usr user.User
		err := json.NewDecoder(req.Body).Decode(&usr)
		defer req.Body.Close()
		fmt.Print(err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = gorm.Updateuser(usr)
		fmt.Print(err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Print(err)
		w.WriteHeader(http.StatusOK)
	})
}
func DeleteUsers(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	Idd := ps.ByName("id")
	//var usr gorm.User
	usrID, err := strconv.Atoi(Idd)

	err = gorm.DeleteUser(usrID)
	fmt.Print(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print(err)
	w.WriteHeader(http.StatusOK)
}
func Signup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	//user := gorm.User
	body := req.Body
	var usr user.User
	//ur, _ := gorm.CreateToken(gorm.User{})
	//var usr gorm.User
	//u := gorm.NewUser()

	//json.NewEncoder(body).Encode(&ur)

	//var usr gorm.User
	//u := gorm.NewUser()
	err := json.NewDecoder(body).Decode(&usr)
	userr, err := service.Signup(usr)
	//user, err := user.Signup(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := tok.CreateToken(&userr)
	err = json.NewEncoder(w).Encode(token)
	fmt.Println(token)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func Signin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	//user := gorm.User
	body := req.Body
	var usr user.User
	//user := &usr
	//ur, _ := gorm.CreateToken(gorm.User{})
	//var usr gorm.User
	//u := gorm.NewUser()

	//json.NewEncoder(body).Encode(&ur)

	//var usr gorm.User
	//u := gorm.NewUser()
	err := json.NewDecoder(body).Decode(&usr)

	token, err := service.Signin(&usr)
	if token == nil {
		http.Error(w, "Incorrect password or username", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(token)
	//rr, err := gorm.TokenDetails(&usrr)
	//token, err := gorm.CreateToken(user)
	//err = json.NewEncoder(w).Encode(token)
	//fmt.Println(token)
	//fmt.Println(err)

	w.WriteHeader(http.StatusOK)
}
func Photoupload(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseMultipartForm(40 << 20)
	fmt.Sprintf("%v", req.Form)
	//n := gorm.NewUser()
	Id := ps.ByName("id")
	//var usr gorm.User
	usrID, err := strconv.Atoi(Id)
	//var usr user.User
	var photo user.User

	profileImageFile, handler, err := req.FormFile("photo")
	if err != nil {
		return
	}
	profilephotoPath, err := uploadFile(profileImageFile, handler)
	if err != nil {
		return
	}
	fmt.Println(profilephotoPath)
	photo.Photo = profilephotoPath

	// defer body.Close()
	err = gorm.Photoupload(photo.Photo, usrID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Hey I am usr photo", photo.Photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

const (
	imageDir = "Images"
)

func uploadFile(file multipart.File, handler *multipart.FileHeader) (string, error) {
	fmt.Printf("Uploading File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile(imageDir, fmt.Sprintf("profile-*%s", handler.Filename))
	uploadPath := tempFile.Name()
	fmt.Println(uploadPath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	tempFile.Write(fileBytes)
	return uploadPath, err
}
