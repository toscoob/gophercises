package main

import (
	"flag"
	"fmt"
	"github.com/gophercises/urlshort"
	"io/ioutil"
	"net/http"
	"github.com/boltdb/bolt"
)

// go run main/main.go

func main() {
	yamlFilename := flag.String("y", "", "a yaml file in the format of path: 'path', url: 'url'")
	jsonFilename := flag.String("j", "", "a yaml file in the format of path: 'path', url: 'url'")
	useDB := flag.Bool("db", false, "use db")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	db, err := setupDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = addUrlsToDB(db)
	if err != nil {
		panic(err)
	}

	var handler http.HandlerFunc

	if *jsonFilename != "" {
		content, err := ioutil.ReadFile(*jsonFilename)
		if err != nil {
			panic(err)
		}
		jsn := string(content)
		handler, err = urlshort.JSONHandler([]byte(jsn), mapHandler)
		if err != nil {
			panic(err)
		}
	} else if *yamlFilename != "" {
		content, err := ioutil.ReadFile(*yamlFilename)
		if err != nil {
			panic(err)
		}
		yaml := string(content)
		handler, err = urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}
	} else if *useDB {
		handler, err = urlshort.DBHandler(db, mapHandler)
		if err != nil {
			panic(err)
		}
	} else {
		var err error
		yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
		//repeated code, but ok
		handler, err = urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}
	}
/*
	fmt.Println(yaml)
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
*/
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
	//http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("URLS"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

func addUrlsToDB(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URLS"))
		err := b.Put([]byte("/y"), []byte("https://www.youtube.com/watch?v=P02eA2o2RUI"))
		if err != nil {
			return err
		}
		err = b.Put([]byte("/z"), []byte("https://zupzup.org/boltdb-example/"))
		return err
	})

	return err
}
