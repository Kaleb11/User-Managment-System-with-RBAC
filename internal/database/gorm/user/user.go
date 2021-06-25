package user

import (
	"Auth/internal/database/gorm/config"
	user "Auth/internal/user/model"
	"Auth/internal/user/pagination"
	tok "Auth/internal/user/service/token"
	"context"

	//"Auth/internal/user/service"
	"log"
)

type Users user.User

func init() {

	// passhashadmin, _ := serv.HashPassword("adminpass")
	// passhashmember, _ := serv.HashPassword("memberpass")
	// Createuser(&user.User{Username: "Kaleb Tilahun", Password: passhashadmin, Role: "admin"})
	// Createuser(&user.User{Username: "Bekele", Password: passhashmember, Role: "member"})

}

// NewUser retrun a pointer to a User
func NewUser() *user.User {
	return new(user.User)
}
func Createuser(usrr *user.User) (user.User, error) {
	db, err := config.GetDBcon()
	if err != nil {
		return user.User{}, err
	}
	dbc, err := db.DB()
	if err != nil {
		return user.User{}, err
	}
	defer dbc.Close()
	password := usrr.Password
	hash, err := tok.HashPassword(password)
	// Hash password to database and gives token
	usrr.Password = hash
	usr := db.Create(&usrr)
	//user.Id = usr.
	return *usrr, usr.Error
}
func Createphoto(id int) (user.User, error) {
	user := user.User{}
	db, err := config.GetDBcon()
	if err != nil {
		return user, err
	}
	dbc, err := db.DB()
	if err != nil {
		return user, err
	}
	defer dbc.Close()
	//usr := db.Create(&usrr)
	db.Model(&user).Where("id = ?", id).Create(user.Photo)
	//user.Id = usr.
	return user, nil
}
func Updateuser(newusr user.User) error {
	db, err := config.GetDBcon()

	if err != nil {
		return err
	}
	dbc, err := db.DB()
	defer dbc.Close()
	if err != nil {
		return err
	}
	log.Println("updated user role ", newusr.Role)
	updatedusr := newusr
	db.First(&newusr)
	updatedusr.Id = newusr.Id
	db.Save(&updatedusr)

	return err
}
func DeleteUser(id int) error {
	db, err := config.GetDBcon()
	db.Delete(&user.User{}, id)
	//db.Delete(&user)
	// db.Where("id = ?", id).Delete(&User{})
	return err
}
func Photoupload(pp string, id int) error {
	user := &user.User{}
	db, err := config.GetDBcon()
	db.Model(&user).Where("id = ?", id).Update("Photo", pp)
	return err
}
func Getusers(ctx context.Context, pg *pagination.Pagination) error {
	db, err := config.GetDBcon()
	if err != nil {
		return nil
	}
	ur := []*user.User{}
	var totalRows int64 = 0

	db.Scopes(pg.Paginate()).Preload("Addresses").Model(user.User{}).Find(&ur)

	if errCount := db.Model(&user.User{}).Count(&totalRows).Error; errCount != nil {
		log.Println(errCount)
	}

	pg.Rows = ur
	return pg.Builder(totalRows)

	//return ur, err

}

func Getusersbyid(Id int) (user.ResponseUser, error) {
	db, err := config.GetDBcon()
	if err != nil {
		return user.ResponseUser{}, err
	}
	ur := user.ResponseUser{}
	users := db.Model(user.User{}).First(&ur, Id)

	// if users.Error != nil {
	// 	return ur, err
	// }
	// err1 := errors.New("No rows in this id")
	// //Us := NewUser().Id
	// if Id != ur.Id {
	// 	errors.Is(err1, ErrRecordNotFound)
	//     return ResponseUser{}, err1
	// }
	return ur, users.Error
}
func Getusersbyname(usr string) (user.User, error) {
	db, err := config.GetDBcon()
	if err != nil {
		return user.User{}, err
	}
	ur := user.User{}
	result := db.Where("username = ?", usr).First(&ur)

	// if users.Error != nil {
	// 	return ur, err
	// }
	// err1 := errors.New("No rows in this id")
	// //Us := NewUser().Id
	// if Id != ur.Id {
	// 	errors.Is(err1, ErrRecordNotFound)
	//     return ResponseUser{}, err1
	// }
	return ur, result.Error
}
