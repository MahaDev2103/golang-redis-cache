package golang_redis_cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"golang-redis/DBUtils"
	"golang-redis/dao"
	"net/http"
	"time"
)

func main() {

	// Run a dev server to handle GET method
	router := mux.NewRouter()
	router.HandleFunc("getProducts/", func(resp http.ResponseWriter, request *http.Request) {
		// Make a db call to fetch all products
		db, err := DBUtils.GetDbHandle()
		if err != nil {
			fmt.Printf("Unable to get DB handle: %v.", err)
			//return nil
		}
		ProductMap := dao.GetProductsFromDB(db)
		fmt.Println(ProductMap)
		var cache = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
		var ctx = context.Background()
		for key, element := range ProductMap {
			fmt.Println("Key:", key, "=>", "Element:", element)
			cacheErr := cache.Set(ctx, string(key), element, 120*time.Second).Err()
			if cacheErr != nil {
				fmt.Println(cacheErr)
			}
		}

	})

	http.ListenAndServe("localhost:8009", router)

}
