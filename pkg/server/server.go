package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bloomingbug/depublic/configs"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/middlewares"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/bloomingbug/depublic/pkg/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
	cfg *configs.Config
}

func NewServer(cfg *configs.Config, binder *binder.Binder, publicRoutes, privateRoutes []*route.Route) *Server {
	e := echo.New()
	e.HideBanner = true
	e.Binder = binder

	e.Use(
		middleware.Logger(),
		//middleware.Recover(),
		middleware.CORS(),
	)

	authMiddleware := middlewares.NewMiddleware(cfg.JWT.SecretKey)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "API Depublic", nil))
	})

	endpoint := e.Group("/api")

	if len(publicRoutes) > 0 {
		for _, r := range publicRoutes {
			endpoint.Add(r.Method, r.Path, r.Handler)
		}
	}

	if len(privateRoutes) > 0 {
		for _, r := range privateRoutes {
			endpoint.Add(r.Method, r.Path, r.Handler, authMiddleware.For(r.Roles...))
		}
	}

	return &Server{e, cfg}
}

func (s *Server) Run() {
	splash()
	runServer(s, s.cfg.Port)
	gracefulShutdown(s)
}

func runServer(s *Server, port string) {
	go func() {
		err := s.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func gracefulShutdown(s *Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	go func() {
		if err := s.Shutdown(ctx); err != nil {
			s.Logger.Fatal("Server Shutdown: ", err)
		}
	}()
}

func splash() {
	colorReset := "\033[0m"

	splashText := `      %##%#             #%##%       %%%%%%  *#%%%#.       #%%%%%#%* .%%%%%%%%%%%%%%%%%%%%# *%%%%%. 
     .%%%%%#%,         ,%%%%%%%      %%%%%%  *%%%%%.    #%%%%%%#   .%%%%%%%%%%%%%%%%%%#%%   *%%%%%. 
     %%%%#%%%##       #%%%%%%%%%     %%%%%%  *%%%%%. ##%%%%%(              (%%%%%           *%%%%%. 
    ##%%%(*%#%%#    .%%%%%.%%%%##    %%%%%%  *%%%%#%%%##%.                 (%%%%%           *%%%%%. 
   /#%%#%   ##%%%( #%%#%%   %%%%%,   %%%%%%  *%%%##*%%%%%%%,               (%%%%%           *%%%%%. 
  .%#%%#     ##%%%%%#%#(    *%%%%%   %%%%%%  *%%%%%.  (%%%%%%%*            (%%%%%           *%%%%%. 
  %%%%%(      *%#%%%%%.      #%%%%%  %%%%%%  *%%%%%.     (#%%%#%%(         (%%%%%           *%%%#%. 
 (#####         /%%#*         #####/ ######  *#####.        (#######(      (#####           *#####. `
	fmt.Println(colorReset, splashText)
}
