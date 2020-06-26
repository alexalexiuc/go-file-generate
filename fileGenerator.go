package fileGenerator

import (
	"math/rand"
	"os"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateFile(fullPath string, fileSize int64) error {
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
	}()

	fi, err := f.Stat()
	if err != nil {
		return err
	}
	currentSize := fi.Size()
	iterations := 0
	for currentSize < fileSize {
		iterations++
		availableSizeRemained := fileSize - currentSize
		randStringSize := 50000
		if 50000 > availableSizeRemained {
			randStringSize = int(availableSizeRemained)
		}
		value := randStringRunes(randStringSize)
		_, err := f.Write([]byte(value))
		if err != nil {
			return err
		}
		fi, err := f.Stat()
		if err != nil {
			return err
		}
		currentSize = fi.Size()
	}
	return nil
}
