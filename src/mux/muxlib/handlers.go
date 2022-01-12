package muxlib

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"

	"image-processor-dokkaan.ir/lib"
)

func initLog(r *http.Request) {
	fmt.Printf("%s-%s-%v-", r.Method, r.URL.Path, r.URL.Query())
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	HttpError404(w, "Requested resource doesn't exist. Please check your path.")
}

func ProductImageProcessHandler(w http.ResponseWriter, r *http.Request) {
	HttpOptionsResponseHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	initLog(r)

	bearerUser, err := lib.GetBearerUser(r.Header)
	if err != nil {
		fmt.Println("ProductImageProcessHandler-GetBearerUser error:", err)
		HttpHandleTokenErrors(w, err)
		return
	}
	fmt.Println(bearerUser.UserId)

	r.Body = http.MaxBytesReader(w, r.Body, int64(lib.MAX_UPLOAD_SIZE))
	img, _, err := image.Decode(r.Body)
	if err != nil {
		fmt.Println("ProductImageProcessHandler-Decode error:")
		HttpError400(w, lib.ErrBadRequest.Error())
		return
	}

	output := lib.SquareFitCrop(img)
	output = lib.ImageResize(output, uint(lib.MAX_PRODUCT_IMAGE_WIDTH), uint(lib.MAX_PRODUCT_IMAGE_WIDTH))

	var reader io.Reader
	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, output, nil)
	if err != nil {
		fmt.Println("ProductImageProcessHandler-Encode error:")
		HttpError500(w)
		return
	}
	reader = buffer

	uploadedFile, err := lib.UploadImage(bearerUser.UserId, lib.TYPE_PRODUCT, reader)
	jsonBytes, err := uploadedFile.ToJson()
	if err != nil {
		fmt.Println("ProductImageProcessHandler-json.Marshal error:", err)
		HttpError500(w)
		return
	}

	HttpSuccessResponse(w, http.StatusCreated, jsonBytes)
}
