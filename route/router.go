package route

import (
  "app/controller/v1"
  "github.com/labstack/echo"
)

func Init() *echo.Echo {
	
  e := echo.New()

  api := e.Group("/api/v1")
  {
    api.GET("/samples", controller.GetSamples())
    api.GET("/heartBeat", controller.GetHeartBeat())
  }
  return e
}
