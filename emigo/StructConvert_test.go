package emigo

import "testing"

func TestConvertViaJSON_MapsMatchingJsonTagsAcrossDifferentTypes(t *testing.T) {
	type entityShape struct {
		Id       int64  `json:"-"`
		UniqueId string `json:"uniqueId"`
		Title    string `json:"title"`
		Internal string `json:"-"`
	}
	type publicDto struct {
		UniqueId string `json:"uniqueId"`
		Title    string `json:"title"`
	}

	src := entityShape{Id: 42, UniqueId: "abc-123", Title: "hello", Internal: "should not leak"}

	out, err := ConvertViaJSON[publicDto](src)
	if err != nil {
		t.Fatalf("ConvertViaJSON error: %v", err)
	}
	if out.UniqueId != "abc-123" || out.Title != "hello" {
		t.Fatalf("unexpected conversion result: %+v", out)
	}
}

func TestConvertViaJSON_NilSrcReturnsZeroValue(t *testing.T) {
	type dto struct {
		Name string `json:"name"`
	}

	out, err := ConvertViaJSON[dto](nil)
	if err != nil {
		t.Fatalf("ConvertViaJSON error: %v", err)
	}
	if out != (dto{}) {
		t.Fatalf("expected zero value, got %+v", out)
	}
}
