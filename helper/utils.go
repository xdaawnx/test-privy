package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// get JWT data
func GetUserData(ctx echo.Context) (map[string]interface{}, error) {
	u := []byte(ctx.Request().Header.Get("u"))
	var m map[string]interface{}
	err := json.Unmarshal(u, &m)
	if err != nil {
		return nil, errors.New("error filter list")
	}

	return m, nil
}

// Md5 hash
func Md5(p []byte) string {
	hash := md5.New()
	hash.Write(p)
	return hex.EncodeToString(hash.Sum(nil))
}

// IsPasswordValid for validating password
func IsPasswordValid(hshd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hshd), []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}

// InArray is used to find data string in array
func InArray(val string, array []interface{}) (exists bool, index int) {
	exists = false
	index = -1
	for i, v := range array {
		if val == v {
			index = i
			exists = true
			return
		}
	}
	return
}

// GetUID is used to get the user ID based on the JWT via header
func GetUID(ctx echo.Context) (int, error) {
	uid := ctx.Request().Header.Get("uid")
	id, err := strconv.Atoi(uid)
	if err != nil {
		log.Println(err)
		return 0, errors.New("Can't Get User ID")
	}
	return id, nil
}

// GenerateRandomString for token which containing random string
func GenerateRandomString() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//RandStringBytesFixedLength string fixed length
func RandStringBytesFixedLength(n int) string {
	const specialChars = "!#$%&*,"
	chars := letterBytes + specialChars
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

//GenerateUsername username
func GenerateUsername(name string) string {
	hours, minutes, second := time.Now().Clock()
	year, _, day := time.Now().Date()
	t := time.Date(year, time.Now().Month(), day, hours,
		minutes, second, 0, time.UTC)
	nama := strings.ReplaceAll(name, " ", "")
	clientcode := strings.ToUpper(nama[0:4])
	yrday := t.YearDay()
	frmt := strconv.Itoa(yrday) + strconv.Itoa(t.Hour()) + strconv.Itoa(t.Minute()) + strconv.Itoa(t.Second())
	randomstring := RandStringBytesFixedLength(17)

	username := clientcode + "-" + string(frmt) + "-" + randomstring

	return username
}

//GeneratePassword password
func GeneratePassword() string {
	password := RandStringBytesFixedLength(8)

	return password
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func UploadPicture(f *multipart.FileHeader) (string, error) {
	var ext = []interface{}{
		".png", ".jpg", ".jpeg", ".bmp",
	}

	src, err := f.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	fileExtension := filepath.Ext(TempDownload + "" + f.Filename)
	istrue, _ := InArray(fileExtension, ext)

	if !istrue {
		return "", errors.New("wrong file extension, your file extension: " + fileExtension)
	}

	if f.Size > 200000 {
		return "", errors.New("file size is too big")
	}
	randName := GenerateRandomString() + GenDateNameFile() + fileExtension
	filepath := DirUploadedFile + randName

	dst, err := os.Create(filepath)

	if err != nil {
		return "", err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return randName, nil
}
