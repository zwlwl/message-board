package dao

import (
	"message-board/model"
)

func UpdatePassword(username,newPassword string) error  {
	_, err := dB.Exec("UPDATE user SET password = ? WHERE username = ?", newPassword,username)
	return err
}

func SelectUserByUsername(username string) (model.User,error)  {
	user := model.User{}

	row := dB.QueryRow("SELECT id,password FROM user WHERE username = ? ",username)
	if row.Err() != nil {
		return user,row.Err()
	}

	err := row.Scan(&user.Id,&user.Password)
	if err != nil {
		return user,err
	}
	return user,nil
}

func InsertUser(user model.User) error  {
	_, err := dB.Exec("INSERT INTO user(username, password)"+"values(?,?);",user.Username,user.Password)
	return err
}