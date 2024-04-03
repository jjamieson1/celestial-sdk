package utilities

import (
	"strings"

	"github.com/revel/revel"
)

func HandleValidationError(errors []*revel.ValidationError, statusCode int, c *revel.Controller) revel.Result {
	c.Response.Status = statusCode
	return c.RenderJSON(errors)
}

func HandleBadRequestError(key, message string, statusCode int, c *revel.Controller) revel.Result {
	c.Validation.Error("error").Key(key).Message(message)
	c.Response.Status = statusCode
	return c.RenderJSON(c.Validation.Errors)
}

func HandleInternalError(key, message string, statusCode int, c *revel.Controller) revel.Result {
	c.Validation.Error("error").Key(key).Message(message)
	c.Response.Status = statusCode
	return c.RenderJSON(c.Validation.Errors)
}

func HandleAuthorizationError(key, message string, statusCode int, c *revel.Controller) revel.Result {
	c.Validation.Error("error").Key(key).Message(message)
	c.Response.Status = statusCode
	return c.RenderJSON(c.Validation.Errors)
}

func HandleHeaderTenantId(c *revel.Controller) (string, []*revel.ValidationError) {
	tenantId := c.Request.Header.Get("tenantId")
	c.Validation.Required(tenantId).Key("tenantId").Message("tenantId in header is required")
	c.Validation.Length(tenantId, 36).Key("tenantId").Message("tenantId wrong length expected 36 chars")
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		return "", c.Validation.Errors
	}
	return tenantId, nil
}

func HandelHeaderJWT(c *revel.Controller) (jwt string) {
	authorization := c.Request.Header.Get("Authorization")
	authFragments := strings.Split(authorization, " ")
	if len(authFragments) == 0 {
		c.Validation.Error("error").Key("header").Message("missing authorization in request")
	}

	if authFragments[0] == "Bearer" {
		if len(authFragments) == 2 {
			jwt = authFragments[1]
		} else {
			c.Validation.Error("error").Key("header").Message("Bearer found but no token")
		}
	} else {
		c.Validation.Error("error").Key("header").Message("String does not start with Bearer")
	}
	return jwt
}

func HandelJWTCookie(c *revel.Controller) (jwt string) {
	jwtCookie, err := c.Request.Cookie("jwt")

	if err != nil {
		c.Validation.Error("error").Key("cookie").Message("missing jwt cookie")
		return ""
	}

	return jwtCookie.GetValue()
}
