package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/xiaosongluo/dashboard/src/app/controllers"
	"fmt"
	"time"
)

// Load returns the routes and middleware
func Load() http.Handler {
	return middleware(routes())
}

// LoadHTTPS returns the HTTP routes and middleware
func LoadHTTPS() http.Handler {
	return middleware(routes())
}

// LoadHTTP returns the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	return middleware(routes())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

// Optional method to make it easy to redirect from HTTP to HTTPS
func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}


// *****************************************************************************
// Routes
// *****************************************************************************
func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/doc", controllers.DocumentHandller).Methods("GET")
	router.HandleFunc("/{dashid}.json", controllers.GetDashboardJsonHandler).Methods("GET")
	router.HandleFunc("/{dashid}", controllers.GetDashboardHandller).Methods("GET")
	router.HandleFunc("/{dashid}", controllers.DeleteDashboardHandller).Methods("DELETE")
	router.HandleFunc("/{dashid}/{metricid}", controllers.PutMetricHandller).Methods("PUT")
	router.HandleFunc("/{dashid}/{metricid}", controllers.DeleteMetricHandller).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./statics/")))

	return router
}

func middleware(h http.Handler) http.Handler {

	// Log every request
	h=handler(h)

	return h
}

func handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}