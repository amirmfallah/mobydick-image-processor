package lib

type BearerUser struct {
	UserId string "json:'userId'"
}

type UploadedFile struct {
	Location string "json:'location'"
	Key      string "json:'key'"
}
