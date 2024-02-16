package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tadi-chat-service/configs"

	"github.com/gin-gonic/gin"
)

func TokenMiddleware(c *gin.Context) {

    token := c.GetHeader("Authorization")

	 req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth-check", configs.Envs.UserServiceBaseURI), nil)
	 if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	 }

	 req.Header.Set("Authorization", token)
	 req.Header.Set("x-api-key", configs.Envs.UserServiceApiKey)

	 client := http.Client{}
	 resp, err := client.Do(req)
	 if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"status": "failure", "message": "Unable Error in Validating User"})
		return
	 }

	 defer resp.Body.Close()

	 if resp.StatusCode != http.StatusOK {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"status": "failure", "message": "Token or API Key Invalid: User Unauthorized"})
		return
	 }

	 body, err := io.ReadAll(resp.Body)
	 if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	 }

	 var data struct {
		Data interface{}		`json:"data"`
		Status string	`json:"status"`
		StatusCode int	`json:"statusCode"`
	 }

	 if err := json.Unmarshal(body, &data); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	 }

	 c.Set("User", data.Data)
	//  c.Set("User", resp.Body)
    c.Next()
}
