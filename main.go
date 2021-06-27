package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ijunyu/gee/cache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *cache.Group {
	return cache.NewGroup("scores", 2<<10, cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}

			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, group *cache.Group) {
	peers := cache.NewHTTPPool(addr)
	peers.Set(addrs...)
	group.RegisterPeers(peers)
	log.Println("cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, group *cache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := group.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		}))

	log.Println("frontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func main() {
	var port = 8001
	var api bool
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	group := createGroup()
	if api {
		go startAPIServer(apiAddr, group)
	}
	startCacheServer(addrMap[port], addrs, group)
}

// func main() {
// 	r := engine.Default()

// 	r.GET("/", func(c *engine.Context) {
// 		c.String(http.StatusOK, "welcome")
// 	})

// 	v1 := r.Group("/v1")
// 	{
// 		v1.GET("/hello", func(c *engine.Context) {
// 			c.String(http.StatusOK, "hello, you're at %s\n", c.Path)
// 		})
// 	}

// 	v2 := r.Group("/v2")
// 	{
// 		v2.GET("/hello/:name", func(c *engine.Context) {
// 			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
// 		})

// 		v2.GET("/assets/*filepath", func(c *engine.Context) {
// 			c.JSON(http.StatusOK, engine.H{"filepath": c.Param("filepath")})
// 		})
// 	}

// 	r.GET("/panic", func(c *engine.Context) {
// 		s := []string{"test"}
// 		c.String(http.StatusOK, s[100])
// 	})

// 	r.Run(":8080")
// }
