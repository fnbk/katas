package main

import "fmt"

type Rider struct {
	UUID   string `bson:"uuid"`
	Bottel string `bson:"bottel"`
	Age    int    `bson:"age"`
	// Names []string `bson:"names"`
	Addresses []Address `bson:"addresses"`
}

type Address struct {
	Street string `bson:"street"`
}

func main() {
	//
	// setup
	//

	executer := CollectionExecuter{
		Mongohost:      "127.0.0.1:27017",
		DatabaseName:   "buffalo",
		CollectionName: "rider",
	}
	err := executer.Start()
	if err != nil {
		panic(err)
	}

	storer := MongoStorer{executer}

	//
	//
	//

	// r1 := Rider{"abc", 30, []string{"a", "b"}}
	r1 := Rider{"abc", "water", 1, []Address{{"A"}, {"B"}}}
	selector := Rider{UUID: r1.UUID}

	//
	//
	//

	err = storer.Put(selector, &r1)
	if err != nil {
		panic(err)
	}

	//
	//
	//

	// selector = Rider{Age: 0}
	returnValue := []Rider{}
	err = storer.Get(nil, &returnValue)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GET:%+v\n", returnValue)
}
