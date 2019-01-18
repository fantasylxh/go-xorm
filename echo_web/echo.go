package echo_web

import (
	"fmt"
	"gameCityService/controllers"
	"gameCityService/global"
	"gameCityService/routers"
	"html/template"
	"io"
	"strings"

	"github.com/json-iterator/go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func EchoInit() {
	e := echo.New()
	e.Static("/__ADMIN_JS__", "staticres/admin/js")
	e.Static("/__ADMIN_IMG__", "staticres/admin/img")
	e.Static("/__ADMIN_CSS__", "staticres/admin/css")
	e.Static("/__LIBS__", "staticres/libs")
	e.Static("/fonts", "staticres/admin/fonts")
	//renderer:= &TemplateRenderer{ templates: template.Must(template.ParseGlob("templates/*/*.html")), }
	//e.Renderer = renderer
	domain := global.Cfg.Section("domain")
	origins := make([]string, len(domain.Keys()))
	for i := 0; i < len(domain.Keys()); i++ {
		origins[i] = domain.Keys()[i].Value()
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{

		AllowCredentials : true,
		AllowOrigins: origins,
		//AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Use(ServerHeader)
	/*
		e.Use(middleware.KeyAuth(func(s string, context echo.Context) (bool, error) {
			return true, nil
		}))*/

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controllers.PublicsLogin)
	e.GET("/checkCode", controllers.PublicCheckCode)
	e.POST("/runLogin", controllers.PublicsRunLogin)
	e.POST("/runLogout",controllers.PublicsRunLogout)

	admin := e.Group("/admin", middleware.KeyAuth(func(s string, context echo.Context) (bool, error) {
		userInfo, err := global.RedisConn.Get("Auth_" + s).Result()
		fmt.Println(userInfo)
		if err != nil {
			return false, nil
		}
		if len(userInfo) > 0 {
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			decoder := json.NewDecoder(strings.NewReader(userInfo))
			params := make(map[string]interface{})
			err := decoder.Decode(&params)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%+v\n", params)
			}
			return true, nil
		}
		//验证账户

		return false, nil
	}))
	//后台路由注册
	routers.AdminRegisterRouter(admin)
	//Api路由注册
	routers.ApiRegisterRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
