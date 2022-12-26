package main

/* Small prototype of the cart project to demo a new design pattern
While DDD is ok, if u-services is the proposed architecture why not use KISS and Interface based design
If go has packages, they are designed for a  purpose. Imposing DDD style project structure is actually making it un-usable across projects
Author : niranjan.awati@ntucenterprise.sg
*/
import (
	"encoding/json"
	"os"

	"bitbucket.org/niranjanawati/cart-mydesign/cache"
	"bitbucket.org/niranjanawati/cart-mydesign/cart"
	"bitbucket.org/niranjanawati/cart-mydesign/catalogue"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var (
	TEST_DATA = true // this can be inturn loaded from container environment
)

func init() {

	// flag.BoolVar(&FVerbose, "verbose", false, "Level of logging messages are set here")
	// flag.BoolVar(&FLogF, "flog", false, "Direction in which the log should output")
	// Setting up log configuration for the api
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetReportCaller(false)
	// By default the log output is stdout and the level is info
	log.SetOutput(os.Stdout)     // FLogF will set it main, but dfault is stdout
	log.SetLevel(log.DebugLevel) // default level info debug but FVerbose will set it main
	// logFile = os.Getenv("LOGF")

	/* We can think of setting the values here for all the environment variables*/
}

/* A demo webapi to show how the cart can be designed differently
 */
func main() {
	log.Info("Starting up the cart service..")
	defer log.Warn("Now closing the cart service..")
	// before we go on to run the api server we test connect the redis cache connection
	/*++++++++++++++++
	Connecting redis client
	*****************/
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Panic("failed to ping redis server..")
	}
	log.Info("Connected to redis server")
	if TEST_DATA {
		// this is an indicator to load the test data
		id := uuid.New()
		data := cart.ScanGo{
			UserID: id.String(),
			Items: []catalogue.Product{
				&catalogue.Grocery{Title: "Tomatoes", Vendor: "Parson&Sons", Unit: catalogue.Piece, PerUnit: 0.5},
				&catalogue.Grocery{Title: "Ginger Garlic paste", Vendor: "Chings", Unit: catalogue.Bottle, PerUnit: 0.5},
			},
		}
		byt, _ := json.Marshal(data)
		/* For now we are just setting the cart identification as cart-userid
		We are aware ofcourse in the actual scene this may not be the case
		*/
		// if err := client.Set(fmt.Sprintf("cart-test", id.String()), byt, 0); err != nil {
		// 	log.WithFields(log.Fields{
		// 		"id": id.String(),
		// 	}).Error("failed to set cache test data: %s", err)
		// }
		if status := client.Set("cart-test", byt, 0); status.Err() != nil {
			log.WithFields(log.Fields{
				"id": id.String(),
			}).Error("failed to set cache test data: %s", err)
		}
	}
	/*++++++++++++++++
	Starting api server
	*****************/
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// Load html glob here
	// Do not set CORS before html glob, that can cause problems loading all html/ static content
	// r.Use(mddlwr.CORS)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app": "omni cart",
		})
	}) // this can help you check if the server is up and running
	api := r.Group("/api")
	// A user has a single cart that can be appended, cleared and posted on checkout
	api.GET("/users/:userid/cart", CachedCart(&cache.RedisCache{}), HndlUsrCart)
	api.PATCH("/users/:userid/cart", HndlUsrCart)
	api.DELETE("/users/:userid/cart", HndlUsrCart)
	api.POST("/users/:userid/cart", HndlUsrCart)
	log.Fatal(r.Run(":8080"))
}
