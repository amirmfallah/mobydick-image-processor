package lib

import (
	"os"
)

var (
	JWT_SECRET              = ""
	MAX_UPLOAD_SIZE         = 1024 * 1024
	MAX_PRODUCT_IMAGE_WIDTH = 512
	MAX_SHOP_IMAGE_WIDTH    = 192
	S3_SECRET_KEY           = ""
	S3_ACCESS_KEY           = ""
	S3_ENDPOINT             = ""
	S3_BUCKET_NAME          = ""
	TYPE_PROFILE            = "profiles"
	TYPE_SHOP               = "shops"
	TYPE_PRODUCT            = "products"
)

func init() {
	JWT_SECRET = os.Getenv("JWT_SECRET")
	S3_SECRET_KEY = os.Getenv("S3_SECRET_KEY")
	S3_ACCESS_KEY = os.Getenv("S3_ACCESS_KEY")
	S3_ENDPOINT = os.Getenv("S3_ENDPOINT")
	S3_BUCKET_NAME = os.Getenv("S3_BUCKET_NAME")
}
