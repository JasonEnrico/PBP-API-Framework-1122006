package controllers

import (
	m "echo/models"
	"net/http"
	"strconv"
)

func GetAllUsers() (m.GlobalResponse, error) {
	var user m.User
	var users []m.User
	var response m.GlobalResponse

	data := Connect()
	defer data.Close()

	query := "SELECT * FROM users"

	rows, err := data.Query(query)

	if err != nil {
		return response, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender); err != nil {
			return response, err
		}

		users = append(users, user)
	}

	response.Status = http.StatusAccepted
	response.Message = "Success!"
	response.Data = users

	return response, err
}

func GetUserByID(id string) (m.GlobalResponse, error) {
	var user m.User
	var response m.GlobalResponse

	data := Connect()
	defer data.Close()

	query := "SELECT * FROM users WHERE id=" + id

	rows, err := data.Query(query)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender); err != nil {
			return response, err
		}
	}

	if user.ID == "0" {
		return response, err
	}

	response.Status = http.StatusAccepted
	response.Message = "Success!"
	response.Data = user

	return response, err
}

func InsertUser(name string, age string, gender string) (m.GlobalResponse, error) {
	var user m.User
	var response m.GlobalResponse

	data := Connect()
	defer data.Close()

	_, err := data.Exec("INSERT INTO users VALUES (null, ?, ?, ?)", name, age, gender)
	if err == nil {
		response.Status = http.StatusAccepted
		response.Message = "Data has been successfully inserted!"
		user.Name = name
		user.Age, _ = strconv.Atoi(age)
		user.Gender = gender
		response.Data = user
	} else {
		println(err.Error())
		return response, err
	}

	return response, err
}

func UpdateUser(id string, name string, age string, gender string) (m.GlobalResponse, error) {
	var user m.User
	var response m.GlobalResponse

	data := Connect()
	defer data.Close()

	query := "UPDATE users SET name = ?, age = ?, gender = ? WHERE id = ?"

	rows, err := data.Prepare(query)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	_, err = rows.Exec(name, age, gender, id)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	response.Status = http.StatusOK
	response.Message = "User updated successfully"
	user.ID = id
	user.Name = name
	user.Age, _ = strconv.Atoi(age)
	user.Gender = gender
	response.Data = user

	return response, err
}

func DeleteUser(id string) (m.GlobalResponse, error) {
	var user m.User
	var response m.GlobalResponse

	data := Connect()
	defer data.Close()

	_, err := data.Exec("DELETE FROM users WHERE id=?", id)
	if err == nil {
		response.Status = http.StatusAccepted
		response.Message = "User deleted successfully"
		user.ID = id
		response.Data = user
	} else {
		println(err.Error())
		return response, err
	}

	return response, nil
}
