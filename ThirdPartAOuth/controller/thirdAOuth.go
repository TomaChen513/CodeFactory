package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"thirdPartLogin/lib"

	"github.com/gin-gonic/gin"
)

// Login page
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// detail see github docs:https://docs.github.com/en/developers/apps/building-oauth-apps/authorizing-oauth-apps
func HandlerGithubLogin(c *gin.Context) {
	// GET https://github.com/login/oauth/authorize

	conf := lib.LoadServerConfig()
	// state string:An unguessable random string.
	// It is used to protect against cross-site request forgery attacks.
	state := "xxxxxxx"

	url := "https://github.com/login/oauth/authorize?client_id=" + conf.AppId + "&redirect_uri=" + conf.RedirectURI + "&state=" + state

	c.Redirect(http.StatusMovedPermanently, url)
}

// get access_token
func GetGithubToken(c *gin.Context) {
	// POST https://github.com/login/oauth/access_token
	conf := lib.LoadServerConfig()
	code := c.Query("code")

	loginUrl := "https://github.com/login/oauth/access_token?client_id=" + conf.AppId + "&client_secret=" + conf.AppKey + "&redirect_uri=" + conf.RedirectURI + "&code=" + code

	response, err := http.PostForm(loginUrl, url.Values{
		"client_id":     {conf.AppId},
		"client_secret": {conf.AppKey},
		"redirect_uri":  {conf.RedirectURI},
		"code":          {code},
	})

	if err != nil {
		fmt.Println("post error!", err.Error())
		return
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	body := string(bs)
	resultMap := lib.ConvertToMap(body)

	accessToken := resultMap["access_token"]

	GetGithubUserMessage(accessToken, c)
}

// get user data
func GetGithubUserMessage(access_token string, c *gin.Context) {
	// 	Authorization: Bearer OAUTH-TOKEN
	// GET https://api.github.com/user

	client := &http.Client{}
	reqest, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		panic(err)
	}

	reqest.Header.Add("Authorization", "token "+access_token)
	resp, _ := client.Do(reqest)

	if err != nil {
		fmt.Println("GetMessage Err", err.Error())
		return
	}

	defer resp.Body.Close()


	// use for your website
	message, _ := lib.ParseResponse(resp)

	c.JSON(http.StatusOK,message)
}
