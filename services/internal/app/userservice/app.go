package userservice

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/starter-microservices/service/internal/app/userservice/module/auth"
	"github.com/tanveerprottoy/starter-microservices/service/internal/app/userservice/module/user"
	"github.com/tanveerprottoy/starter-microservices/service/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-microservices/service/internal/pkg/middleware"
	"github.com/tanveerprottoy/starter-microservices/service/internal/pkg/router"
	"github.com/tanveerprottoy/starter-microservices/service/pkg/data/sql/sqlxpkg"
	"github.com/tanveerprottoy/starter-microservices/service/pkg/validatorpkg"
	// "go.uber.org/zap"
)

// App struct
type App struct {
	DBClient *sqlxpkg.Client
	router           *router.Router
	Middlewares      []any
	AuthModule       *auth.Module
	UserModule       *user.Module
	Validate         *validator.Validate
}

func NewApp() *App {
	a := new(App)
	// run db script
	// will run in a goroutine, channel is be used
	// to get the output, waitgroup is used to make
	// sure it completes execution
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)
	go a.runScript(ch, &wg)
	wg.Wait()
	o := <-ch
	fmt.Println(o)
	a.initComponents()
	return a
}

func (a *App) runScript(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// cmd, err := exec.Command("/bin/sh", "../../scripts/init_db.sh").Output()
	// cmd, err := exec.Command("chmod +x ./scripts/init_db.sh").Output()
	cmd, err := exec.Command("./scripts/init_db.sql").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	ch <- output
}

func (a *App) initDB() {
	a.DBClient = sqlxpkg.GetInstance()
}

func (a *App) initMiddlewares() {
	authMiddleWare := middleware.NewAuthMiddleware(a.AuthModule.Service)
	a.Middlewares = append(a.Middlewares, authMiddleWare)
}

func (a *App) initModules() {
	a.UserModule = user.NewModule(a.DBClient.DB, a.Validate)
	a.AuthModule = auth.NewModule(a.UserModule.Service)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.AuthMiddleware)
	router.RegisterUserRoutes(a.router, constant.V1, a.UserModule, m)
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
