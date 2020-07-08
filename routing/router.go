package routing

import (
    "log"
    "time"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "github.com/justinas/alice"
    "github.com/casbin/casbin"
    "github.com/alexkarpovich/mowy-api-go/middlewares"
    "github.com/alexkarpovich/mowy-api-go/routing/root"
    "github.com/alexkarpovich/mowy-api-go/routing/accounts"
)

func router() http.Handler {
    baseRouter := mux.NewRouter().StrictSlash(true)
    root.Router(baseRouter)
    accounts.Router(baseRouter)

    return baseRouter
}

func ListenAndServe(address string) error {
    authEnforcer, err := casbin.NewEnforcerSafe("./config/auth_model.conf", "./config/policy.csv")
    if err != nil {
        log.Fatal(err)
    }

    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

    corsHandler := handlers.CORS(headersOk, originsOk, methodsOk)
    authorizeHandler := middlewares.Authorizer(authEnforcer)
    chain := alice.New(
        corsHandler,
        middlewares.CurrentUser,
        authorizeHandler,
    ).Then(router())

    server := &http.Server{
    	ReadTimeout: 15 * time.Second,
    	WriteTimeout: 15 * time.Second,
    	IdleTimeout: 60 * time.Second,
    	Handler: chain,
    	Addr: address,
    }

	return server.ListenAndServe()
}
