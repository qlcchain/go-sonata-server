// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/ecdsa"
	"crypto/tls"
	"net/http"
	"path"

	"github.com/dgrijalva/jwt-go"

	"github.com/qlcchain/go-sonata-server/config"
	"github.com/qlcchain/go-sonata-server/util"

	"github.com/go-openapi/swag"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/restapi/handler/mock"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"github.com/qlcchain/go-sonata-server/auth"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/operations"
)

var cfg = &ServerCfg{}

func init() {
	dir := config.LogDir()
	fn := path.Join(dir, "sonata.log")

	lw, _ := rotatelogs.New(
		fn+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fn),
	)

	lh := lfshook.NewHook(
		lw,
		&log.JSONFormatter{},
	)
	log.AddHook(lh)
	//log.SetFormatter(&log.TextFormatter{
	//	FullTimestamp: true,
	//	DisableColors: true,
	//})
	log.SetReportCaller(true)
	log.SetLevel(log.ErrorLevel)
}

//go:generate swagger generate server --target ../../go-sonata-server --name Sonata --spec ../spec/all.yaml --principal models.Principal

type ServerCfg struct {
	Key        string            `json:"key" short:"K" long:"key" description:"private key(elliptic P521) to sign JWT token"`
	Verbose    bool              `json:"verbose" short:"V" long:"verbose" description:"Show verbose debug information"`
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	PublicKey  *ecdsa.PublicKey  `json:"-"`
}

func (s *ServerCfg) String() string {
	return util.ToString(s)
}

func configureFlags(api *operations.SonataAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "jwt",
			LongDescription:  "jwt options",
			Options:          cfg,
		},
	}
}

func configureAPI(api *operations.SonataAPI) http.Handler {
	if cfg.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	if cfg.Key != "" {
		var err error
		if cfg.PrivateKey, err = auth.FromBase64(cfg.Key); err != nil {
			log.Fatal(err)
		}
	}

	if cfg.PrivateKey == nil {
		log.Debug("use default key...")
		cfg.PrivateKey = auth.Decode(auth.DefaultKey)
		cfg.PublicKey = &cfg.PrivateKey.PublicKey
	}

	// configure the api here
	api.ServeError = errors.ServeError
	api.Logger = log.Printf
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.BearerAuth = func(token string, scopes []string) (*models.Principal, error) {
		// TODO: verify scopes???
		if claims, err := auth.ParseAndCheckToken(token, cfg.PublicKey); err == nil {
			return &models.Principal{
				Code:   0,
				Reason: "",
				Roles:  claims.Roles,
			}, nil
		} else {
			switch vErr, ok := err.(*jwt.ValidationError); ok {
			case vErr.Errors&jwt.ValidationErrorMalformed != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorUnverifiable != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorSignatureInvalid != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorClaimsInvalid != 0:
				return &models.Principal{
					Code: 41, Reason: "Invalid credentials",
				}, nil
			case vErr.Errors&jwt.ValidationErrorExpired != 0:
				return &models.Principal{
					Code: 42, Reason: "Expired credentials",
				}, nil
			default:
				return &models.Principal{
					Code: 400, Reason: err.Error(),
				}, nil
			}
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	mock.Bind(api)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {
		if mock.Store != nil {
			if err := mock.Store.Close(); err != nil {
				log.Error(err)
			}
		}
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	writer := &logger{}
	return alice.New(handlers.CORS(handlers.AllowedOrigins([]string{"*"}), handlers.AllowCredentials()),
		func(handler http.Handler) http.Handler {
			return handlers.LoggingHandler(writer, handler)
		},
		handlers.ProxyHeaders,
		handlers.CompressHandler,
		handlers.RecoveryHandler(
			handlers.RecoveryLogger(writer),
			handlers.PrintRecoveryStack(true),
		)).Then(handler)
}

type logger struct {
}

func (l *logger) Write(p []byte) (n int, err error) {
	log.Debugln(string(p))
	return len(p), nil
}

func (l *logger) Println(args ...interface{}) {
	log.Error(args)
}
