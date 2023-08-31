package middle

//func CheckCookie(AuthSrv ports.AuthService, next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		cook, err := c.Cookie("Authorization")
//		if err != nil {
//			log.Println("Middleware COOKIE -> ", err)
//			return c.Render(200, "page_login", nil)
//		}
//		log.Println("MDL CheckCookie : value cookie -> ", AuthSrv.ValidateTokenCookie(cook.Value))
//		if !AuthSrv.ValidateTokenCookie(cook.Value) {
//			return c.Render(200, "page_login", nil)
//		}
//		return next(c)
//	}
//}
//
//func CheckHeaderToken(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		tokenString := c.Request().Header.Get("Authorization")
//		if tokenString == "" {
//			tokenString = c.QueryParam("token")
//		}
//
//		if tokenString == "" {
//			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
//		}
//		return next(c)
//	}
//}
