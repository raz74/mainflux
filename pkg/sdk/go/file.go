package sdk

import (
	"bytes"
	"fmt"
	"github.com/mainflux/mainflux/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func (sdk mfSDK) UploadFile(file File, fileName, token string) (string, errors.SDKError) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormField("name")
	if err != nil {
	}
	_, err = io.Copy(fw, strings.NewReader(file.Name))
	if err != nil {
		return "", errors.NewSDKError(err)
	}

	fw, err = writer.CreateFormField("version")
	if err != nil {
	}
	_, err = io.Copy(fw, strings.NewReader(file.Version))
	if err != nil {
		return "", errors.NewSDKError(err)
	}

	fw, err = writer.CreateFormFile("file", "file")
	if err != nil {
	}
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, f)
	if err != nil {
		return "", errors.NewSDKError(err)
	}

	err = writer.Close()
	if err != nil {
		return "", nil
	}

	url := "http://localhost:8000/create-file"

	writer.FormDataContentType()
	headers, _, sdkerr := sdk.processRequest(http.MethodPost, url, token, writer.FormDataContentType(), body.Bytes(), http.StatusCreated)
	if sdkerr != nil {
		return "", sdkerr
	}

	id := strings.TrimPrefix(headers.Get("Location"), fmt.Sprintf("/%s/", usersEndpoint))
	return id, nil
}
