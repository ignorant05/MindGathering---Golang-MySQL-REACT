package helpers

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

var Logger slog.Logger

func HashPwd(pwd string) string {
	bytesPwd := []byte(pwd)
	hashedPwd, err := bcrypt.GenerateFromPassword(bytesPwd, bcrypt.DefaultCost)
	if err != nil {
		Logger.Error("Error whiel hashing password", "error", err)
	}
	return string(hashedPwd)
}

func ComparePwd(targetPwd string, pwd string) bool {
	bytesPwd := []byte(pwd)
	targetBytes := []byte(targetPwd)

	err := bcrypt.CompareHashAndPassword(targetBytes, bytesPwd)
	if err != nil {
		Logger.Error("Passwords don't match", "error", err)
		return false
	}

	return true
}
