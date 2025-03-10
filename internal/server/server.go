package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"github.com/sirupsen/logrus"
	"github.com/wickedv43/yd-diploma/internal/config"
	"github.com/wickedv43/yd-diploma/internal/logger"
	"github.com/wickedv43/yd-diploma/internal/storage"
)

type Server struct {
	echo    *echo.Echo
	cfg     *config.Config
	storage storage.DataKeeper
	logger  *logrus.Entry
}

func NewServer(i do.Injector) (*Server, error) {
	s, err := do.InvokeStruct[Server](i)
	if err != nil {
		return nil, errors.Wrap(err, "invoke struct error")
	}

	//init
	s.echo = echo.New()
	s.cfg = do.MustInvoke[*config.Config](i)
	s.logger = do.MustInvoke[*logger.Logger](i).WithField("component", "server")

	s.storage = do.MustInvoke[*storage.PostgresStorage](i)

	//middleware
	s.echo.Use(middleware.Recover(), middleware.Gzip(), s.logHandler, s.CORSMiddleware)

	//free routes
	s.echo.POST(`/api/user/register`, s.onRegUser)
	s.echo.POST(`/api/user/login`, s.onLogin)

	//authorized users
	user := s.echo.Group(``, s.authMiddleware)
	user.POST(`/api/user/orders`, s.onPostOrders)
	user.GET(`/api/user/orders`, s.onGetOrders)
	user.GET(`/api/user/balance`, s.onGetUserBalance)
	user.POST(`/api/user/balance/withdraw`, s.onProcessPayment)
	user.GET(`/api/user/withdrawals`, s.GetUserBills)

	return s, nil
}

func (s *Server) Start() {
	s.logger.Info("server started...")
	err := s.echo.Start(s.cfg.Server.RunAddress)
	if err != nil {
		s.logger.Fatal(errors.Wrap(err, "start server"))
	}
}
