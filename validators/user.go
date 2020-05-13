package validators

import (
	"errors"
	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"go-rest-api/models"
	"io/ioutil"
	"regexp"
)

func CreateOrUpdateUser(isCreate bool) func(*gin.Context) {

	return func(ctx *gin.Context) {
		var user models.User

		body, _ := ioutil.ReadAll(ctx.Request.Body)

		var stringValidator = jio.String()

		if isCreate {
			stringValidator = stringValidator.Required()
		} else {
			stringValidator = stringValidator.Optional()
		}

		var keys = map[string]jio.Schema{}

		keys["firstName"] = jio.String().Optional().Min(2).Max(32).Trim().Lowercase()
		keys["lastName"] = jio.String().Optional().Min(2).Max(32).Trim().Lowercase()
		keys["username"] = func() jio.Schema {
			var schema = jio.String()
			isRequired(schema, isCreate)
			return schema.Min(2).Max(32).Check(func(s string) error {
				// Not Allowed Characters
				if regexp.MustCompile("^[a-z0-9_]{2,32}$").Match([]byte(s)) {
					return nil
				}
				return errors.New("can only contains 'a-z0-9_-'")
			})
		}()
		keys["email"] = func() jio.Schema {
			var schema = jio.String()
			isRequired(schema, isCreate)
			return schema.Check(func(s string) error {
				// Valid Email
				if regexp.MustCompile("[^@]+@[^.]+\\..+").Match([]byte(s)) {
					return nil
				}
				return errors.New("must be a valid email address")
			})
		}()
		keys["password"] = func() jio.Schema {
			var schema = jio.String()
			isRequired(schema, isCreate)
			return schema.Min(6).Max(32)
		}()

		_, err := jio.ValidateJSON(&body, jio.Object().Keys(keys))

		_ = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(body, &user)

		setValidationErrors(ctx, err)

		if isCreate {
			ctx.Set("create-user", user)
		} else {
			ctx.Set("update-user", user)
		}

		ctx.Next()
	}

}
