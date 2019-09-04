package standalone

//go:generate go get github.com/kevinburke/go-bindata
//go:generate go-bindata -pkg=standalone favicon.png grpc-logo.svg index-template.html jquery-3.3.1.min.js jquery-ui-1.12.1.min.css jquery-ui-1.12.1.min.js

// GetResources is an exported accessor for resource file contents.
func GetResources() map[string]func() []byte {
	m := map[string]func() []byte{}
	names := AssetNames()
	for _, name := range names {
		k := "/" + name
		v := MustAsset(name)
		m[k] = func() []byte {
			return v
		}
	}
	return m
}

// GetIndexTemplate is a dedicated exported accessor for index-template.html
func GetIndexTemplate() []byte {
	return MustAsset("index-template.html")
}
