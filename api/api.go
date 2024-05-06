package api

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"strings"
)

var DefaultClient = &Client{}

func Run(parameters Parameters) (string, error) {
	return DefaultClient.Run(parameters)
}

type Client struct {
	HTTPClient http.Client
	URL        string
}

type Parameters struct {
	File   fs.File
	URL    string
	Text   string
	Input  string
	Output string
	Force  bool
	DieOn  string
	Time   string
	Md     map[string]string
}

func (a *Client) Run(parameters Parameters) (string, error) {
	buffer := &bytes.Buffer{}
	data := multipart.NewWriter(buffer)

	if parameters.File == nil {
		data.CreateFormFile("file", "")
	} else {
		w, _ := data.CreateFormFile("file", "index.bs")
		if _, err := io.Copy(w, parameters.File); err != nil {
			return "", err
		}
	}

	w, _ := data.CreateFormField("url")
	w.Write([]byte(parameters.URL))

	w, _ = data.CreateFormField("text")
	w.Write([]byte(parameters.Text))

	w, _ = data.CreateFormField("input")
	if parameters.Input == "" {
		w.Write([]byte("spec"))
	} else {
		w.Write([]byte(parameters.Input))
	}

	w, _ = data.CreateFormField("output")
	if parameters.Output == "" {
		w.Write([]byte("html"))
	} else {
		w.Write([]byte(parameters.Output))
	}

	if parameters.Force {
		w, _ = data.CreateFormField("force")
		w.Write([]byte("1"))
	}

	w, _ = data.CreateFormField("die-on")
	if parameters.DieOn == "" {
		w.Write([]byte("fatal"))
	} else {
		w.Write([]byte(parameters.DieOn))
	}

	w, _ = data.CreateFormField("time")
	w.Write([]byte(parameters.Time))

	w, _ = data.CreateFormField("action")
	w.Write([]byte("Process"))

	for key, value := range parameters.Md {
		w, _ = data.CreateFormField("md-" + key)
		w.Write([]byte(value))
	}

	data.Close()

	req, _ := http.NewRequest("POST", a.URL, buffer)
	req.Header.Set("Content-Type", data.FormDataContentType())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "curl/0.0.0")

	res, err := a.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("unexpected status code: " + res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	text := string(bytes)

	if strings.Contains(text, `.api.csswg.org`) {
		panic("Run: got the form page instead of rendered HTML")
	}

	return text, nil
}
