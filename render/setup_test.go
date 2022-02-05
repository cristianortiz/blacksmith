package render

import (
	"os"
	"testing"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./testdata/views"),
	jet.InDevelopmentMode(),
)

//type to mock a Renderer  for testing
var testRenderer = Render{
	Renderer: "",
	RootPath: "",
	JetViews: views,
}

//TestMain is the entry point function to run tests
func TestMain(m *testing.M) {
	//run tests before exit...
	os.Exit(m.Run())
}
