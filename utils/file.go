package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"path/filepath"
)

func FileExt(fname string) string {
	return filepath.Ext(fname)
}

func FileMD5(f io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	md5Hash := hash.Sum(nil)
	md5String := hex.EncodeToString(md5Hash)

	return md5String, nil
}
