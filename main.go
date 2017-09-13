package main

func main() {
	productProvider := ProductProvider{IDs: []string{"1", "2", "3"}}
	app := App{productProvider}
	app.Run()
}
