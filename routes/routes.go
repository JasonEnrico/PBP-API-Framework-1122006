package routes

import (
	"echo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	res, err := controllers.GetAllUsers()
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"message": err.Error()}, "	")
	}

	return c.JSONPretty(http.StatusAccepted, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, " ")
}

func GetUserByID(c echo.Context) error {
	res, err := controllers.GetUserByID(c.Param("id"))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"message": err.Error()}, "	")
	}

	return c.JSONPretty(http.StatusAccepted, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, " ")
}

func InsertUser(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	gender := c.QueryParam("gender")
	if name == "" || age == "" || age == "0" || gender == "" {
		return c.JSONPretty(http.StatusBadRequest, map[string]string{"message": "Invalid Variables"}, "	")
	}

	res, _ := controllers.InsertUser(name, age, gender)

	return c.JSONPretty(http.StatusAccepted, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, " ")
}

func UpdateUser(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	gender := c.QueryParam("gender")

	res, err := controllers.UpdateUser(c.Param("id"), name, age, gender)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"message": err.Error()}, "	")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, " ")
}

func DeleteUser(c echo.Context) error {
	res, err := controllers.DeleteUser(c.Param("id"))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"message": err.Error()}, "	")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"data": res.Data, "message": res.Message, "status": res.Status}, " ")
}
