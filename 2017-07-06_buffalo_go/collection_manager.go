package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//
// InitJob
//

type InitJob struct {
	DatabaseName   string
	CollectionName string
	index          []mgo.Index
}

//
// Executer
//

type Executer interface {
	Execute(fn ExecuterFunc) error
}

type ExecuterFunc func(collection *mgo.Collection) error

//
// CollectionManager
//

type CollectionManager interface {
	Start() error
	Stop() error
	GetExecuter(databaseName, collectionName string) (Executer, error)
}

type MongoCollectionManager struct {
	mongohost string
	timeout   int
	initJobs  []InitJob
}

func NewCollectionManager(mongohost string, timeout int, initJobs []InitJob) CollectionManager {
	return &MongoCollectionManager{
		mongohost: mongohost,
		timeout:   timeout,
		intiJobs:  initJobs,
	}
}

func (self *MongoCollectionManager) Start() error {
	return nil
	// execute init jobs
	// start db
}

func (self *MongoCollectionManager) Stop() error {
	return nil
}

func (self *MongoCollectionManager) GetExecuter(databaseName, collectionName string) (Executer, error) {
	return nil, nil
}

//
// Example Operators/Layers
//

func Abc() error {
	var err error
	//
	// config
	//

	mongohost := "localhost:9090"
	timeout := 60

	//
	// collection manager
	//

	jobs := []InitJob{
		{
			DatabaseName:   "demo-service",
			CollectionName: "stammdaten",
			index: []mgo.Index{
				{
					Name:       "sgkIndex",
					Key:        []string{"strukturgruppenkennung"},
					Unique:     true,  // Doppelte Einträge sind nicht erlaubt
					DropDups:   false, // doppelte Einträge führen zum Fehler
					Background: true,  // Index wird im Hintergrund gebaut, andere Sessions sind nicht betroffen
					Sparse:     true,  // nur Dokumente mit den Index-Feldern werden beachtet (und beim Sort zurückgeliefert)
				},
			},
		},
	}

	collectionManager := NewCollectionManager(mongohost, timeout, jobs)
	err = collectionManager.Start()
	if err != nil {
		return err
	}

	//
	// stamdaten aggregat
	//

	databaseName := "demo-service"
	collectionName := "stammdaten"
	stammdatenExecuter, err := collectionManager.GetExecuter(databaseName, collectionName)
	if err != nil {
		return err
	}

	stammdatenReplacer := NewKeyValueReplacerMongo(stammdatenExecuter) // multiple keys
	stammdatenUpdater := NewKeyValueUpdaterMongo(stammdatenExecuter)   // unique keys
	stammdatenAggregat := NewStammdatenAggregat(stammdatenExecuter, stammdatenReplacer, stammdatenUpdater)
	stammdatenAggregat = stammdatenAggregat

	//
	// Execute
	//

	artikelSlice := []Artikel{}

	// closure
	findeArtikel := func(collection *mgo.Collection) error {
		return collection.Find(bson.M{"uuid": "my_uuid"}).All(&artikelSlice)
	}

	// execute
	err = stammdatenExecuter.Execute(findeArtikel)
	if err != nil {
		return err
	}

	// use
	for _, a := range artikelSlice {
		fmt.Printf("Artikel Name: %s", a.Name)
	}

	return nil
}

//
// Beispiel: StammdatenAggregat
//

type StammdatenAggregat struct {
	StammdatenExecuter Executer
	StammdatenReplacer KeyValueReplacer
	StammdatenUpdater  KeyValueUpdater
}

func NewStammdatenAggregat(executer Executer, replacer KeyValueReplacer, updater KeyValueUpdater) *StammdatenAggregat {
	return nil
}

func (self *StammdatenAggregat) Get(uuid string) (Stammdatum, error) {
	selector := Stammdatum{UUID: uuid}
	stammdatum := Stammdatum{}

	err := self.StammdatenReplacer.Get(selector, &stammdatum)
	if err != nil {
		return Stammdatum{}, err
	}

	return stammdatum, nil
}

//
// Stammdatum
//

type Stammdatum struct {
	UUID string
}
