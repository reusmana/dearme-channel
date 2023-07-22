package authRepository

import (
	"fmt"

	"github.com/rals/dearme-channel/config"
	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/models"
)

func CheckLogin(username string, password string) (models.User, []models.Modulars, error) {
	// var objAuth models.Auth
	var objUser models.User
	var objModulars models.Modulars
	var arrObjModulars []models.Modulars
	// var pwd string

	db := config.ConnectionPg()
	//set schema postgresql
	db.Exec(`set search_path='authorization'`)

	sqlAuth := `SELECT id_master_users, id_master_roles, username,
						fullname, email, email_verified_at, password,
						url_photo, login_count, remember_token, sequence,
						is_active, created_by, created_date
					 FROM "authorization".master_users
					 WHERE username = $1`
	err := db.Raw(sqlAuth, username).Scan(&objUser).Error

	if err != nil {
		fmt.Println("Email not found")
		return objUser, arrObjModulars, err
	}

	match, errPass := helpers.VerifyPassword(password, objUser.Password)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return objUser, arrObjModulars, errPass
	}

	sqlModulars := `SELECT id_master_modulars, id_master_applications,
						modular_name, modular_icon, sequence, is_active,
						created_by, created_date, updated_by, updated_date
					FROM master_modulars
					WHERE 1=1`
	rows, errMod := db.Raw(sqlModulars).Rows()
	defer rows.Close()
	if errMod != nil {
		fmt.Println("The Data Modulars is Not Found.")
		return objUser, arrObjModulars, errMod
	}

	for rows.Next() {
		rows.Scan(&objModulars.IdMasterModulars, &objModulars.IdMasterApplications,
			&objModulars.ModularName, &objModulars.ModularIcon, &objModulars.Sequence,
			&objModulars.IsActive, &objModulars.CreatedBy, &objModulars.CreatedDate,
			&objModulars.UpdatedBy, &objModulars.UpdatedDate)

		// arrObjModulars = append(arrObjModulars, objModulars)
		arrObjModulars = append(arrObjModulars, objModulars)
	}

	return objUser, arrObjModulars, nil
}
