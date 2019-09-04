package standalone

// all binaries of files under /resources are combined into a single file 'bindata.go'.

//go:generate ./go-bindata-all.sh

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
