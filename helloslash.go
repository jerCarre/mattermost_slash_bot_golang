package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type botResponse struct {
	Response_type string `json:"response_type" xml:"response_type"`
	Text          string `json:"text" xml:"text"`
}

type slashRequest struct {
	Channel_id   string `form:"channel_id"`
	Channel_name string `form:"channel_name"`
	Command      string `form:"command"`
	Response_url string `form:"response_url"`
	Team_domain  string `form:"team_domain"`
	Team_id      string `form:"team_id"`
	Text         string `form:"text"`
	Token        string `form:"token"`
	User_id      string `form:"user_id"`
	User_name    string `form:"user_name"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/", func(c echo.Context) error {
		br := new(botResponse)
		br.Response_type = "in_channel"

		slash := new(slashRequest)
		if err := c.Bind(slash); err != nil {
			br.Text = "***Bad Request***"
		} else {
			br.Text = `---
#### Your Request

| Request field                    | Value |
|:---------------------------------|-------|
| Channel_id | ` + slash.Channel_id + ` |
| Channel_name | ` + slash.Channel_name + ` |
| Command | ` + slash.Command + ` |
| Response_url | ` + slash.Response_url + ` |
| Team_domain | ` + slash.Team_domain + ` |
| Team_id | ` + slash.Team_id + ` |
| Text | ` + slash.Text + ` |
| Token | ` + slash.Token + ` |
| User_id | ` + slash.User_id + ` |
| User name | ` + slash.User_name + ` |
---`
		}
		return c.JSON(http.StatusOK, br)
	})
	e.Logger.Fatal(e.Start(addr))
}
