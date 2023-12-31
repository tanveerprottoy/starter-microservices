package gateway

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/auth"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/app/gateway/module/user"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/pkg/global"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/pkg/middleware"
	"github.com/tanveerprottoy/starter-microservices/gateway/internal/pkg/router"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/config"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/cryptopkg"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/data/nosql/mongodb"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/data/sql/postgres"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/file"
	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/validatorpkg"

	"github.com/go-playground/validator/v10"
	// "go.uber.org/zap"
)

// App struct
type App struct {
	MongoDBClient    *mongodb.Client
	PostgresDBClient *postgres.Client
	router           *router.Router
	Middlewares      []any
	Validate         *validator.Validate
	AuthModule       *auth.Module
	UserModule       *user.Module
	ContentModule    *content.Module
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) getConfigValues() {
	global.UserServiceBaseUrl = config.GetEnvValue("USER_SERVICE_BASE_URL")
}

func (a *App) initDB() {
	a.MongoDBClient = mongodb.GetInstance()
	a.PostgresDBClient = postgres.GetInstance()
}

func (a *App) initMiddlewares() {
	authMiddleWare := middleware.NewAuthMiddleware(a.AuthModule.Service)
	a.Middlewares = append(a.Middlewares, authMiddleWare)
}

func (a *App) initModules() {
	a.UserModule = user.NewModule(a.MongoDBClient.DB, a.PostgresDBClient.DB, a.Validate)
	a.AuthModule = auth.NewModule(a.UserModule.Service)
	a.ContentModule = content.NewModule(a.PostgresDBClient.DB)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.AuthMiddleware)
	router.RegisterUserRoutes(a.router, constant.V1, a.UserModule, m)
	router.RegisterContentRoutes(a.router, constant.V1, a.ContentModule)
}

func (a *App) initValidators() {
	a.Validate = validator.New()
	_ = a.Validate.RegisterValidation("notempty", validatorpkg.NotEmpty)
}

/* func (a *App) initLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"proxy.log",
	}
	cfg.Build()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	task := "taskName"
	logger.Info("failed to do task",
		// Structured context as strongly typed Field values.
		zap.String("url", task),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
} */

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.router = router.NewRouter()
	a.initModules()
	a.initMiddlewares()
	a.initModuleRouters()
	a.initValidators()
	// a.initLogger()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(
		":8080",
		a.router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}

// Run app
func (a *App) RunTLSSimpleConfig() {
	err := http.ListenAndServeTLS(":443", "cert.crt", "key.key", a.router.Mux)
	if err != nil {
		log.Fatal(err)
	}
}

// use mutual TLS and not just one-way TLS,
// we must instruct it to require client authentication to ensure clients present a certificate from our CA when they connect
func (a *App) RunTLSMutual() {
	caCert, _ := file.ReadFile("ca.crt")
	cp, _ := cryptopkg.AppendCertsFromPEM(caCert)
	tlsConf := &tls.Config{
		ClientCAs:  cp,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConf.BuildNameToCertificate()
	srv := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConf,
		Handler:   a.router.Mux,
	}
	srv.ListenAndServeTLS("cert.crt", "key.key")
}

func (a *App) RunDisableHTTP2() {
	srv := &http.Server{
		Handler:      a.router.Mux,
		Addr:         "127.0.0.1:8080",
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	log.Fatal(srv.ListenAndServe())
}
