package testdata

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
)

func UseOsMkdirAllAndWriteFile() {
	randPath, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	p := filepath.Join(os.TempDir(), fmt.Sprintf("/%d", randPath))
	_ = os.MkdirAll(p, os.ModePerm) // want "os and ioutil dir and file writing functions are not permissions-safe, use fileutil"
	someFile := filepath.Join(p, "some.txt")
	_ = ioutil.WriteFile(someFile, []byte("hello"), os.ModePerm) // want "os and ioutil dir and file writing functions are not permissions-safe, use fileutil"
}
