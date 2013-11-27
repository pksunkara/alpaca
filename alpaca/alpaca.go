package alpaca

type Data struct {
	Pkg map[string]interface{}
	Api map[string]interface{}
	Doc map[string]interface{}
}

func ReadData(directory string) *Data {
	var pkg, api, doc map[string]interface{}

	ReadJSON(directory+"/pkg.json", &pkg)
	ReadJSON(directory+"/api.json", &api)
	ReadJSON(directory+"/doc.json", &doc)

	return &Data{pkg, api, doc}
}
