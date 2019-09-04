package webform

//go:generate go get github.com/kevinburke/go-bindata
//go:generate go-bindata -pkg=webform webform-sample.css webform-template.html webform.js

// WebFormScript --
func WebFormScript() []byte {
	name := "webform.js"
	return MustAsset(name)
}

// WebFormSampleCSS --
func WebFormSampleCSS() []byte {
	name := "webform-sample.css"
	return MustAsset(name)
}

// WebFormTemplate --
func WebFormTemplate() []byte {
	name := "webform-template.html"
	return MustAsset(name)
}
