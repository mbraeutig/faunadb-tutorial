package main

import (
	"fmt"
	"log"

	f "github.com/fauna/faunadb-go/faunadb"
)

func main() {

	// Instantiate the client
	client := f.NewFaunaClient("YOUR_FAUNADB_SECRET")

	// Create a database
	result, err := client.Query(f.CreateDatabase(f.Obj{"name": "my_app"}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Accessing the database
	result, err = client.Query(
		f.CreateKey(
			f.Obj{"database": f.Database("my_app"), "role": "server"},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Set up a collection
	result, err = client.Query(f.CreateCollection(f.Obj{"name": "posts"}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Create an index
	result, err = client.Query(
		f.CreateIndex(
			f.Obj{
				"name":   "posts_by_title",
				"source": f.Collection("posts"),
				"terms":  f.Arr{f.Obj{"field": f.Arr{"data", "title"}}},
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	result, err = client.Query(
		f.CreateIndex(
			f.Obj{
				"name":   "posts_by_tags_with_title",
				"source": f.Collection("posts"),
				"terms":  f.Arr{f.Obj{"field": f.Arr{"data", "tags"}}},
				"values": f.Arr{f.Obj{"field": f.Arr{"data", "title"}}},
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Create a post
	result, err = client.Query(
		f.Create(
			f.Collection("posts"),
			f.Obj{"data": f.Obj{"title": "What I had for breakfast .."}},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Create several posts
	result, err = client.Query(
		f.Map(
			f.Arr{
				"My cat and other marvels",
				"Pondering during a commute",
				"Deep meanings in a latte",
			},
			f.Lambda(
				"post_title",
				f.Create(
					f.Collection("posts"),
					f.Obj{"data": f.Obj{"title": f.Var("post_title")}},
				),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Retrieve posts
	result, err = client.Query(f.Get(f.RefCollection(f.Collection("posts"), "192903209792046592")))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	result, err = client.Query(
		f.Get(
			f.MatchTerm(
				f.Index("posts_by_title"),
				"My cat and other marvels",
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Update posts
	result, err = client.Query(
		f.Update(
			f.RefCollection(f.Collection("posts"), "192903209792046592"),
			f.Obj{"data": f.Obj{"tags": f.Arr{"pet", "cute"}}},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Replace posts
	result, err = client.Query(
		f.Replace(
			f.RefCollection(f.Collection("posts"), "192903209792046592"),
			f.Obj{"data": f.Obj{"title": "My dog and other marvels"}},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Delete a post
	result, err = client.Query(
		f.Delete(f.RefCollection(f.Collection("posts"), "192903209792045568")),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	_, err = client.Query(f.Get(f.RefCollection(f.Collection("posts"), "192903209792045568")))
	if err != nil {
		log.Fatal(err)
	}
}
