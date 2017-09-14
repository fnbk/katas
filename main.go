package main

func main() {
	productProvider := ProductProvider{
		Products: []Product{
			{
				ID:   "1",
				Name: "Name1",
			},
			{
				ID:   "2",
				Name: "Name2",
			},
			{
				ID:   "3",
				Name: "Name3",
			},
		},
	}
	structureProvider := StructureProvider{
		Structures: []Structure{
			{
				ID:   "123",
				Name: "Name123",
			},
		},
	}
	httpPortal := HTTPPortal{productProvider, structureProvider}
	app := App{httpPortal}
	app.Run()
}
