package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRender_Page(t *testing.T) {
	r, err := http.NewRequest("GET", "some-url", nil)
	if err != nil {
		t.Error(err)
	}
	//mock a ResponseWritter
	w := httptest.NewRecorder()

	//populate fields of Renderer type
	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"
	//Test rendering a go template
	err = testRenderer.Page(w, r, "hometest", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}
	//Test rendering a Jet file
	testRenderer.Renderer = "jet"
	err = testRenderer.Page(w, r, "homejet", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)

	}
	//Test non existing go template file
	testRenderer.Renderer = "go"
	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("error rendering non-existent go template", err)

	}
	//Test non existing jet file
	testRenderer.Renderer = "jet"
	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("error rendering not existent jet template ", err)

	}

	//Test non configured rendering engine
	testRenderer.Renderer = ""
	err = testRenderer.Page(w, r, "hometest", nil, nil)
	if err == nil {
		t.Error("No error returned while renderer is not configured", err)

	}
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
