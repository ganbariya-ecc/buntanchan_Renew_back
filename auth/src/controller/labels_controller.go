package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Labels struct {
	Labels   []string    `json:"labels"`
}

func UpdateLabels(ctx echo.Context) error {
	// adminid 取得
	adminid := ctx.Get("adminid").(string)

	// userid 取得
	userid := ctx.Request().Header.Get("userid")

	// ユーザーID がない場合
	if userid == "" {
		// ユーザーIDがない場合
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "User ID not specified",
		})
	}

	// ラベルを取得
	var label_data Labels
	if err := ctx.Bind(&label_data); err != nil {
		log.Println("Label acquisition error : " + err.Error())
		return ctx.JSON(http.StatusBadRequest,echo.Map{
			"result" : "Label acquisition error",
		})
	}

	log.Println(adminid)
	log.Println("labels : " + strings.Join(label_data.Labels, ","))

	return ctx.JSON(http.StatusOK,echo.Map{
		"result" : "success",
	})
}