package main

import (
	"log"
	"net/http"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

var Schema = `
	schema {
		query: Query
	}
	type Person{
		id: ID!
		firstName: String!
		lastName: String
	}
	type Query{
		person(id: ID!): Person
	}
`

type person struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}

var people = []*person{
	{
		ID:        "1000",
		FirstName: "Pedro",
		LastName:  "Marquez",
	},

	{
		ID:        "1001",
		FirstName: "John",
		LastName:  "Doe",
	},
}

type personResolver struct {
	p *person
}

func (r *personResolver) ID() graphql.ID {
	return r.p.ID
}

func (r *personResolver) FirstName() string {
	return r.p.FirstName
}

func (r *personResolver) LastName() *string {
	return &r.p.LastName
}

type Resolver struct{}

func (r *Resolver) Person(args struct{ ID graphql.ID }) *personResolver {
	if p := peopleData[args.ID]; p != nil {
		log.Print("Found in resolver!/n")
		return &personResolver{p}
	}
	return nil
}

var peopleData = make(map[graphql.ID]*person)

var mainSchema *graphql.Schema

func init() {
	for _, p := range people {
		peopleData[p.ID] = p
	}

	mainSchema = graphql.MustParseSchema(Schema, &Resolver{})
}

func main() {
	http.Handle("/query", &relay.Handler{Schema: mainSchema})
	log.Print("Starting to listen 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}