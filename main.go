package main

/* Small prototype of the cart project to demo a new design pattern
While DDD is ok, if u-services is the proposed architecture why not use KISS and Interface based design
If go has packages, they are designed for a  purpose. Imposing DDD style project structure is actually making it un-usable across projects
Author : niranjan.awati@ntucenterprise.sg
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.org/niranjanawati/cart-mydesign/cache"
	"bitbucket.org/niranjanawati/cart-mydesign/cart"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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

// seed_cache : for the prototype we are seeding the cart data on redis
// for seeding json data from file refer to ./seed/carts.json
func seed_cache(client *redis.Client) {
	// this is an indicator to load the test data
	cart := cart.ScanGo{}
	// getting data  from seed json
	// https://tutorialedge.net/golang/parsing-json-with-golang/
	/* ++++++++++++++++
	Reading seeding file
	++++++++++++++++*/
	jsonF, err := os.Open("./seed/carts.json")
	if err != nil {
		log.Errorf("Failed to open seed file, %s", err)
	}
	byt, err := ioutil.ReadAll(jsonF)
	if err != nil {
		log.Errorf("Failed read json seed file, %s", err)
	}
	if err := json.Unmarshal(byt, &cart); err != nil {
		log.Errorf("failed to unmarshal test json data, %s", err)
	}
	// Notice how we cache the entire cart as the value in redis
	// convenient since writing and reading th cache then is that much more faster
	jsonStr, err := json.Marshal(cart)
	if err != nil {
		log.Errorf("failed to get json string for the items, %s", err)
	}
	/* ++++++++++++++++
	Pushing to cache
	++++++++++++++++*/
	// https://tutorialedge.net/golang/go-redis-tutorial/
	key := fmt.Sprintf("cart-%s", cart.UserID)
	status := client.Exists(key)
	if count, _ := status.Result(); count == 0 {
		// seed the cache only if the cart seed does not exists
		result := client.Set(key, jsonStr, 0)
		if result.Err() != nil {
			log.Errorf("failed to set cart value cache, %s", result.Err())
		}
	} else {
		log.Info("Cache is already seeded .. skipping")
	}
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
		// If that cart is already pushed to cart no need to push again
		seed_cache(client)
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
	api.GET("/users/:userid/cart", CachedCart(&cache.RedisCache{Client: client}), HndlUsrCart)
	api.PATCH("/users/:userid/cart", HndlUsrCart)
	api.DELETE("/users/:userid/cart", HndlUsrCart)
	api.POST("/users/:userid/cart", HndlUsrCart)
	log.Fatal(r.Run(":8080"))
}
