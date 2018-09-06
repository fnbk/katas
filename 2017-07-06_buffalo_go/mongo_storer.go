package main

//
// CollectionExecuter
//

// type CollectionExecuterFunc func(collection *mgo.Collection) error
//
// type CollectionExecuter interface {
// 	Execute(*mgo.Collection) error
// }

// type CollectionExecuter struct {
// 	session        *mgo.Session
// 	Mongohost      string
// 	DatabaseName   string
// 	CollectionName string
// }
//
// func NewCollectionExecuter(mongoHost, databaseName, collectionName string) *CollectionExecuter {
// 	return &CollectionExecuter{
// 		Mongohost:      mongoHost,
// 		DatabaseName:   databaseName,
// 		CollectionName: collectionName,
// 	}
// }
//
// func (self *CollectionExecuter) Start() error {
// 	var err error
//
// 	// MongoDB Verbindung aufbauen
// 	addrs := strings.Split(self.Mongohost, ",")
// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:   addrs,
// 		Timeout: 2 * time.Second,
// 	}
// 	self.session, err = mgo.DialWithInfo(mongoDBDialInfo)
// 	if err != nil {
// 		return err
// 	}
//
// 	// Strong muss gesetzt werden
// 	// ansonsten kann es vorkommen, dass Read Zugriffe alte Daten liefern, weil die neuen Daten
// 	// noch nicht geschrieben wurden
// 	self.session.SetMode(mgo.Strong, true)
//
// 	return nil
// }
//
// func (self *CollectionExecuter) Execute(fn CollectionExecuterFunc) error {
// 	sessionCopy := self.session.Copy()
// 	defer sessionCopy.Close()
// 	collection := sessionCopy.DB(self.DatabaseName).C(self.CollectionName)
//
// 	err := fn(collection)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

//
// IndexJob
//

// type IndexJob struct {
// 	DatabaseName   string
// 	CollectionName string
// 	index          []mgo.Index
// }
//
// //
// // CollectionManager
// //
//
// type CollectionManager struct {
// }
//
// func NewCollectionManager(mongohost string, timeout int, indexJobs []IndexJob) *CollectionManager {
// }
//
// func (self *CollectionManager) Start() error {
// 	return nil
// }
//
// func (self *CollectionManager) Stop() error {
// 	return nil
// }
//
// func (self *CollectionManager) GetExecuter(databaseName, collectionName string) CollectionExecuter {
// 	return nil
// }
//
// //
// // CollectionExecuter
// //
//
// type CollectionExecuterFunc func(collection *mgo.Collection) error
//
// type CollectionExecuter interface {
// 	Execute(*mgo.Collection) error
// }
//
// //
// // Operators/Layers
// //
//
// func Abc() {
// 	cm := CollectionManager{}
// 	executer := cm.GetExecuter()
//
// 	KeyValueReplacerPut(executer)
// }
//
// //
// // KeyValueReplacer
// //
//
// type KeyValueReplaceOperator struct {
// 	executer CollectionExecuter
// }
//
// func KeyValueReplacerPut(executer Executer) {
//
// }
//
// //
// // MongoStorer
// //
//
// type MongoStorer struct {
// 	executer CollectionExecuter
// }
//
// func NewMongoStorer(executer CollectionExecuter) *MongoStorer {
// 	return &MongoStorer{
// 		executer: executer,
// 	}
// }
//
// // selector==nil => get all
// // selector==bson.M{} => fehler (selector leer), um unklare abfragen, wie Blue{UUID:""} abzufangen (leerer selector, wahrscheinlich falsch; um alle abzufragen nil nutzen)
// // returnValue: slice of values
// // check: returnValue muss pointer slice sein
// // Limitationen: Zero-Werte können nicht abgefragt werden
// func (self *MongoStorer) Get(selector, returnValue interface{}) error {
// 	var err error
//
// 	//
// 	// selector
// 	//
//
// 	var findSelector bson.M
// 	if selector != nil {
// 		findSelector, err = buildUpsertSelector(selector)
// 		if err != nil {
// 			return err
// 		}
// 		if reflect.DeepEqual(findSelector, bson.M{}) {
// 			return fmt.Errorf("selector ist leeres Struct")
// 		}
// 	} else {
// 		findSelector = bson.M{}
// 	}
//
// 	//
// 	// query
// 	//
//
// 	fn := func(collection *mgo.Collection) error {
// 		err := collection.Find(findSelector).All(returnValue)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
//
// 	//
// 	// execute
// 	//
//
// 	err = self.executer.Execute(fn)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// // (selector==nil => duplicate upsert TODO: eigener storer, duplicate upsert storer)
// // value: value (SINGLE)
// // value: slice of values (MULTIPLE) => pairs???
// // check: value muss pointer sein
//
// //
// // FEATURES
// //
//
// //
// // BsonSelectorBuilder()
// //
// //
//
// // consumebBson annotated struct, produce Bson.M{}
//
// //
// // replace-key-value-store/executer/operator
// //
//
// // GET
// //  single selector, single value, found, error
// //  Beispiel: found, err = Get(selector, &value)
// //  guards/expections:
// //   selector==nil, selector==bson.M{}
// //   value==Gpd{}, value==nil (need: pointer of Gpd)
// //
// // PUT
// //  single selector, single value, found, error
// //  Beispiel: found, err = Put(selector, &value)
// //
//
// //
// // partial-update-key-value-store/executer/operator
// //
//
// // GET
// //  multiple selector, multiple value, count, error
// //  Beispiel: count, err = Get(selectors, &values)
// //  guards/expections:
// //   selector==nil, selector==bson.M{}
// //   value==Gpd{}, value==nil (need: pointer of slice of Gpd)
// //   SelectAll() not possible
// //
// // PUT
// //  single selector, single value, count, error
// //  Beispiel: count, err = Put(selector, &value)
//
// // exceptions:
// // selector == nil
// // selector == bson.M{} // empty
// func (self *MongoStorer) Put(selector, value interface{}) error {
// 	//
// 	// guard
// 	//
//
// 	if reflect.DeepEqual(selector, nil) {
// 		return fmt.Errorf("kein gueltiger selector: nil")
// 	}
//
// 	//
// 	// selector
// 	//
//
// 	upsertSelector, err := buildUpsertSelector(selector)
// 	if err != nil {
// 		return err
// 	}
// 	if reflect.DeepEqual(upsertSelector, bson.M{}) {
// 		return fmt.Errorf("kein selector angegeben")
// 	}
//
// 	updateValue := bson.M{"$set": value}
//
// 	//
// 	// query
// 	//
//
// 	fn := func(collection *mgo.Collection) error {
// 		changeInfo, err := collection.Upsert(upsertSelector, updateValue)
// 		if err != nil {
// 			return err
// 		}
// 		// bei einem Put sollte nur Dokument betroffen sein
// 		if changeInfo.Matched > 1 {
// 			errMsg := "Anzahl der gefunden Dokumente is größer als 1. Daten zum entsprechenden Key sind mehrfach vorhanden."
// 			return fmt.Errorf(errMsg)
// 		}
// 		return nil
// 	}
//
// 	//
// 	// execute
// 	//
//
// 	err = self.executer.Execute(fn)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
