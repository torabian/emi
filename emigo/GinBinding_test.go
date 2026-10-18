package emigo

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func multipartContext(t *testing.T, build func(w *multipart.Writer)) *gin.Context {
	t.Helper()
	gin.SetMode(gin.TestMode)
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	build(w)
	w.Close()
	req := httptest.NewRequest(http.MethodPost, "/", &buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = req
	return c
}

func addFile(t *testing.T, w *multipart.Writer, field, name, content string) {
	t.Helper()
	fw, err := w.CreateFormFile(field, name)
	if err != nil {
		t.Fatal(err)
	}
	io.WriteString(fw, content)
}

type heldFile struct {
	Name string
	Data string
}

func TestBindMultipartKeepsEveryFileOfARepeatedField(t *testing.T) {
	c := multipartContext(t, func(w *multipart.Writer) {
		w.WriteField("title", "hello")
		addFile(t, w, "attachments[]", "a.txt", "AAA")
		addFile(t, w, "attachments[]", "b.txt", "BBB")
	})

	var body struct {
		Title       string              `json:"title"`
		Attachments []BindMultipartFile `json:"attachments"`
	}
	if err := BindGinMultipartForm(c, &body); err != nil {
		t.Fatal(err)
	}
	if body.Title != "hello" || len(body.Attachments) != 2 {
		t.Fatalf("want title and 2 files, got %+v", body)
	}
	if body.Attachments[0].Filename != "a.txt" && body.Attachments[1].Filename != "a.txt" {
		t.Fatalf("a.txt lost: %+v", body.Attachments)
	}
}

func TestBindMultipartListSuffixWithOneFileIsStillAList(t *testing.T) {
	c := multipartContext(t, func(w *multipart.Writer) { addFile(t, w, "attachments[]", "a.txt", "AAA") })
	var body struct {
		Attachments []BindMultipartFile `json:"attachments"`
	}
	if err := BindGinMultipartForm(c, &body); err != nil {
		t.Fatal(err)
	}
	if len(body.Attachments) != 1 || string(body.Attachments[0].Data) != "AAA" {
		t.Fatalf("got %+v", body.Attachments)
	}
}

func TestBindMultipartConvertHookBypassesJson(t *testing.T) {
	old := ConvertMultipartFile
	defer func() { ConvertMultipartFile = old }()
	ConvertMultipartFile = func(fh *multipart.FileHeader) (any, error) {
		f, err := fh.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()
		b, _ := io.ReadAll(f)
		return &heldFile{Name: fh.Filename, Data: string(b)}, nil // pointer on purpose, field holds values
	}

	c := multipartContext(t, func(w *multipart.Writer) {
		addFile(t, w, "docs", "a.txt", "AAA")
		addFile(t, w, "docs", "b.txt", "BBB")
		addFile(t, w, "cover", "c.txt", "CCC")
		w.WriteField("note", "n")
	})
	var body struct {
		Docs  []heldFile `json:"docs"`
		Cover heldFile   `json:"cover"`
		Note  string     `json:"note"`
	}
	if err := BindGinMultipartForm(c, &body); err != nil {
		t.Fatal(err)
	}
	if len(body.Docs) != 2 || body.Cover.Data != "CCC" || body.Note != "n" {
		t.Fatalf("unexpected body: %+v", body)
	}
}

func TestBindMultipartConvertHookRejectsSeveralFilesForASingleField(t *testing.T) {
	old := ConvertMultipartFile
	defer func() { ConvertMultipartFile = old }()
	ConvertMultipartFile = func(fh *multipart.FileHeader) (any, error) { return heldFile{Name: fh.Filename}, nil }

	c := multipartContext(t, func(w *multipart.Writer) {
		addFile(t, w, "cover", "a.txt", "A")
		addFile(t, w, "cover", "b.txt", "B")
	})
	var body struct {
		Cover heldFile `json:"cover"`
	}
	if err := BindGinMultipartForm(c, &body); err == nil {
		t.Fatal("two files for a single-file field must be an error")
	}
}
