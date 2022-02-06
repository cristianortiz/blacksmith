package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string //test name
	renderer      string //render engine tested
	template      string //file template using in the test, can be a non-existing one
	errorExpected bool   //if an error must be returned as to PASS the test
	errorMessage  string //error msg if an error was expected to PASS the test
}{ //Test on rendering go templates
	{"go_page", "go", "hometest", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "error must be occurs when rendering a non-existent go template"},
	//Test on rendering jet templates
	{"jet_page", "jet", "homejet", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no-file", true, "error must be occurs when rendering a non-existent jet template"},
	//Test on trying a non valid render option (jet or go)
	{"invalid_render_engine", "foo", "hometest", true, "error must be occurs with non engine template configured"},
}

//TestRender_Page test render.Page() using to test tables
func TestRender_Page(t *testing.T) {
	//run all the tests with the parameteres defined in PageData[]
	for _, e := range pageData {
		r, err := http.NewRequest("GET", "some-url", nil)
		if err != nil {
			t.Error(err)
		}
		//mock a ResponseWritter
		w := httptest.NewRecorder()
		//populate fields of Renderer type
		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"
		//test execution
		err = testRenderer.Page(w, r, e.template, nil, nil)
		//if an error is expected as part of passing the test
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}

		} else {
			//if an error is not expected running the test,
			//but the error occurs and the test FAILS
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}

		}
	}
	//******* OLD: Testing first version ****************************
	// //Test rendering a go template
	// err = testRenderer.Page(w, r, "hometest", nil, nil)
	// if err != nil {
	// 	t.Error("error rendering page", err)
	// }
	// //Test rendering a Jet file
	// testRenderer.Renderer = "jet"
	// err = testRenderer.Page(w, r, "homejet", nil, nil)
	// if err != nil {
	// 	t.Error("error rendering page", err)

	// }
	// //Test non existing go template file
	// testRenderer.Renderer = "go"
	// err = testRenderer.Page(w, r, "no-file", nil, nil)
	// //the error must exists
	// if err == nil {
	// 	t.Error("error rendering non-existent go template", err)

	// }
	// //Test non existing jet file
	// testRenderer.Renderer = "jet"
	// err = testRenderer.Page(w, r, "no-file", nil, nil)
	// //the error must exists
	// if err == nil {
	// 	t.Error("error rendering not existent jet template ", err)

	// }

	// //Test non configured rendering engine
	// testRenderer.Renderer = ""
	// err = testRenderer.Page(w, r, "hometest", nil, nil)
	// //the error must exists
	// if err == nil {
	// 	t.Error("No error returned while renderer is not configured", err)

	// }
}

//test the specific function to render Go Templates
func TestRender_GoPage(t *testing.T) {
	r, err := http.NewRequest("GET", "some-url", nil)
	if err != nil {
		t.Error(err)
	}
	//mock a ResponseWritter to test handlers
	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"
	//Test rendering a go template, REC- inside Page() GoPage() is invoked
	err = testRenderer.Page(w, r, "hometest", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

}

//test the specific function to test jet templates
func TestRender_JetPage(t *testing.T) {
	r, err := http.NewRequest("GET", "some-url", nil)
	if err != nil {
		t.Error(err)
	}
	//mock a ResponseWritter to test handlers
	w := httptest.NewRecorder()

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata"
	//Test rendering a jest template, REC-inside Page JetPage is invoked
	err = testRenderer.Page(w, r, "homejet", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

}
