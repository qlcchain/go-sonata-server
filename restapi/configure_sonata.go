// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/ecdsa"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"os/user"
	"path"
	"path/filepath"
	osRuntime "runtime"

	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product_offering_qualification"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product_order"
	"github.com/qlcchain/go-sonata-server/restapi/operations/quote"

	"github.com/go-openapi/swag"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"github.com/qlcchain/go-sonata-server/auth"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/operations"
	"github.com/qlcchain/go-sonata-server/restapi/operations/cancel_product_order"
	"github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address"
	"github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address_validation"
	"github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"
	"github.com/qlcchain/go-sonata-server/restapi/operations/notification"
)

var cfg = &ServerCfg{}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

func logDir() string {
	home := homeDir()
	logDir := "Gsonata"

	if home != "" {
		if osRuntime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "Application Support", logDir)
		} else if osRuntime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", logDir)
		} else {
			return filepath.Join(home, logDir)
		}
	}
	return ""
}

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		return err
	}
	return nil
}

func init() {
	dir := logDir()
	if dir == "" {
		dir = "/tmp"
	}
	_ = createDirIfNotExist(dir)
	fn := path.Join(dir, "sonata.log")

	lw, _ := rotatelogs.New(
		fn+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fn),
	)

	lh := lfshook.NewHook(
		lw,
		&logrus.JSONFormatter{},
	)
	logrus.AddHook(lh)

	logrus.SetLevel(logrus.ErrorLevel)
}

//go:generate swagger generate server --target ../../go-sonata-server --name Sonata --spec ../spec/all.yaml --principal models.Principal

type ServerCfg struct {
	Key        string            `json:"key" short:"K" long:"key" description:"private key(elliptic P521) to sign JWT token"`
	Verbose    bool              `json:"verbose" short:"V" long:"verbose" description:"Show verbose debug information"`
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	PublicKey  *ecdsa.PublicKey  `json:"-"`
}

func (s *ServerCfg) String() string {
	if data, err := json.Marshal(s); err == nil {
		return string(data)
	}
	return ""
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
	if cfg.Key == "" {
		cfg.Key = auth.DefaultKey
	}
	if cfg.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
	cfg.PrivateKey = auth.Decode(cfg.Key)
	cfg.PublicKey = &cfg.PrivateKey.PublicKey

	// configure the api here
	api.ServeError = errors.ServeError
	api.Logger = logrus.Printf
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.BearerAuth = func(token string, scopes []string) (*models.Principal, error) {
		// TODO: verify scopes???
		if claims, err := auth.ParseAndCheckToken(token, cfg.PublicKey); err != nil {
			return nil, err
		} else {
			return &models.Principal{
				Roles: claims.Roles,
			}, nil
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.CancelProductOrderCancelProductOrderCreateHandler = cancel_product_order.CancelProductOrderCreateHandlerFunc(func(params cancel_product_order.CancelProductOrderCreateParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderCreate has not yet been implemented")
	})
	api.CancelProductOrderCancelProductOrderFindHandler = cancel_product_order.CancelProductOrderFindHandlerFunc(func(params cancel_product_order.CancelProductOrderFindParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderFind has not yet been implemented")
	})
	api.CancelProductOrderCancelProductOrderGetHandler = cancel_product_order.CancelProductOrderGetHandlerFunc(func(params cancel_product_order.CancelProductOrderGetParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderGet has not yet been implemented")
	})
	api.GeographicAddressGeographicAddressGetHandler = geographic_address.GeographicAddressGetHandlerFunc(func(params geographic_address.GeographicAddressGetParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation geographic_address.GeographicAddressGet has not yet been implemented")
	})

	api.GeographicAddressValidationGeographicAddressValidationCreateHandler = geographic_address_validation.GeographicAddressValidationCreateHandlerFunc(func(params geographic_address_validation.GeographicAddressValidationCreateParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation geographic_address_validation.GeographicAddressValidationCreate has not yet been implemented")
	})
	api.GeographicSiteGeographicSiteFindHandler = geographic_site.GeographicSiteFindHandlerFunc(func(params geographic_site.GeographicSiteFindParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation geographic_site.GeographicSiteFind has not yet been implemented")
	})
	api.GeographicSiteGeographicSiteGetHandler = geographic_site.GeographicSiteGetHandlerFunc(func(params geographic_site.GeographicSiteGetParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation geographic_site.GeographicSiteGet has not yet been implemented")
	})
	api.NotificationNotificationProductOrderAttributeValueChangeNotificationHandler = notification.NotificationProductOrderAttributeValueChangeNotificationHandlerFunc(func(params notification.NotificationProductOrderAttributeValueChangeNotificationParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation notification.NotificationProductOrderAttributeValueChangeNotification has not yet been implemented")
	})
	api.NotificationNotificationProductOrderCreationNotificationHandler = notification.NotificationProductOrderCreationNotificationHandlerFunc(func(params notification.NotificationProductOrderCreationNotificationParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation notification.NotificationProductOrderCreationNotification has not yet been implemented")
	})
	api.NotificationNotificationProductOrderInformationRequiredNotificationHandler = notification.NotificationProductOrderInformationRequiredNotificationHandlerFunc(func(params notification.NotificationProductOrderInformationRequiredNotificationParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation notification.NotificationProductOrderInformationRequiredNotification has not yet been implemented")
	})
	api.NotificationNotificationProductOrderStateChangeNotificationHandler = notification.NotificationProductOrderStateChangeNotificationHandlerFunc(func(params notification.NotificationProductOrderStateChangeNotificationParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation notification.NotificationProductOrderStateChangeNotification has not yet been implemented")
	})
	api.NotificationNotificationQuoteAttributeValueChangeNotificationHandler = notification.NotificationQuoteAttributeValueChangeNotificationHandlerFunc(func(params notification.NotificationQuoteAttributeValueChangeNotificationParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation notification.NotificationQuoteAttributeValueChangeNotification has not yet been implemented")
	})
	api.NotificationNotificationQuoteCreationNotificationHandler = notification.NotificationQuoteCreationNotificationHandlerFunc(func(params notification.NotificationQuoteCreationNotificationParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation notification.NotificationQuoteCreationNotification has not yet been implemented")
	})
	api.NotificationNotificationQuoteInformationRequiredNotificationHandler = notification.NotificationQuoteInformationRequiredNotificationHandlerFunc(func(params notification.NotificationQuoteInformationRequiredNotificationParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation notification.NotificationQuoteInformationRequiredNotification has not yet been implemented")
	})
	api.NotificationNotificationQuoteStateChangeNotificationHandler = notification.NotificationQuoteStateChangeNotificationHandlerFunc(func(params notification.NotificationQuoteStateChangeNotificationParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation notification.NotificationQuoteStateChangeNotification has not yet been implemented")
	})
	api.ProductProductFindHandler = product.ProductFindHandlerFunc(func(params product.ProductFindParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation product.ProductFind has not yet been implemented")
	})
	api.ProductProductGetHandler = product.ProductGetHandlerFunc(func(params product.ProductGetParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation product.ProductGet has not yet been implemented")
	})
	api.ProductOfferingQualificationProductOfferingQualificationCreateHandler = product_offering_qualification.ProductOfferingQualificationCreateHandlerFunc(func(params product_offering_qualification.ProductOfferingQualificationCreateParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation product_offering_qualification.ProductOfferingQualificationCreate has not yet been implemented")
	})
	api.ProductOfferingQualificationProductOfferingQualificationFindHandler = product_offering_qualification.ProductOfferingQualificationFindHandlerFunc(func(params product_offering_qualification.ProductOfferingQualificationFindParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation product_offering_qualification.ProductOfferingQualificationFind has not yet been implemented")
	})
	api.ProductOfferingQualificationProductOfferingQualificationGetHandler = product_offering_qualification.ProductOfferingQualificationGetHandlerFunc(func(params product_offering_qualification.ProductOfferingQualificationGetParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation product_offering_qualification.ProductOfferingQualificationGet has not yet been implemented")
	})
	api.HubProductOfferingQualificationManagementHubDeleteHandler = hub.ProductOfferingQualificationManagementHubDeleteHandlerFunc(func(params hub.ProductOfferingQualificationManagementHubDeleteParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubDelete has not yet been implemented")
	})
	api.HubProductOfferingQualificationManagementHubGetHandler = hub.ProductOfferingQualificationManagementHubGetHandlerFunc(func(params hub.ProductOfferingQualificationManagementHubGetParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubGet has not yet been implemented")
	})
	api.HubProductOfferingQualificationManagementHubPostHandler = hub.ProductOfferingQualificationManagementHubPostHandlerFunc(func(params hub.ProductOfferingQualificationManagementHubPostParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubPost has not yet been implemented")
	})

	api.ProductOrderProductOrderCreateHandler = product_order.ProductOrderCreateHandlerFunc(func(params product_order.ProductOrderCreateParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation product_order.ProductOrderCreate has not yet been implemented")
	})

	api.ProductOrderProductOrderFindHandler = product_order.ProductOrderFindHandlerFunc(func(params product_order.ProductOrderFindParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation product_order.ProductOrderFind has not yet been implemented")
	})

	api.ProductOrderProductOrderGetHandler = product_order.ProductOrderGetHandlerFunc(func(params product_order.ProductOrderGetParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation product_order.ProductOrderGet has not yet been implemented")
	})

	api.HubProductOrderManagementHubCreateHandler = hub.ProductOrderManagementHubCreateHandlerFunc(func(params hub.ProductOrderManagementHubCreateParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.ProductOrderManagementHubCreate has not yet been implemented")
	})

	api.HubProductOrderManagementHubDeleteHandler = hub.ProductOrderManagementHubDeleteHandlerFunc(func(params hub.ProductOrderManagementHubDeleteParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.ProductOrderManagementHubDelete has not yet been implemented")
	})

	api.HubProductOrderManagementHubFindHandler = hub.ProductOrderManagementHubFindHandlerFunc(func(params hub.ProductOrderManagementHubFindParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.ProductOrderManagementHubFind has not yet been implemented")
	})

	api.QuoteQuoteCreateHandler = quote.QuoteCreateHandlerFunc(func(params quote.QuoteCreateParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation quote.QuoteCreate has not yet been implemented")
	})

	api.QuoteQuoteFindHandler = quote.QuoteFindHandlerFunc(func(params quote.QuoteFindParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation quote.QuoteFind has not yet been implemented")
	})

	api.QuoteQuoteGetHandler = quote.QuoteGetHandlerFunc(func(params quote.QuoteGetParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation quote.QuoteGet has not yet been implemented")
	})

	api.HubQuoteManagementHubCreateHandler = hub.QuoteManagementHubCreateHandlerFunc(func(params hub.QuoteManagementHubCreateParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.QuoteManagementHubCreate has not yet been implemented")
	})

	api.HubQuoteManagementHubDeleteHandler = hub.QuoteManagementHubDeleteHandlerFunc(func(params hub.QuoteManagementHubDeleteParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.QuoteManagementHubDelete has not yet been implemented")
	})

	api.HubQuoteManagementHubFindHandler = hub.QuoteManagementHubFindHandlerFunc(func(params hub.QuoteManagementHubFindParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation hub.QuoteManagementHubFind has not yet been implemented")
	})

	api.QuoteQuoteRequestStateChangeHandler = quote.QuoteRequestStateChangeHandlerFunc(func(params quote.QuoteRequestStateChangeParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation quote.QuoteRequestStateChange has not yet been implemented")
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

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
	logrus.Debugln(string(p))
	return len(p), nil
}

func (l *logger) Println(args ...interface{}) {
	logrus.Error(args)
}
