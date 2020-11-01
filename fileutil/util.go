package fileutil

import (
	"errors"
	"io/ioutil"
	"os"
)

// MkdirAll takes in a path, expands it if necessary, and looks at permissions of the path,
// ensuring we are not attempting to overwrite any existing permissions. Finally, creates
// the directory accordingly with standardized, permissions for our project.
func MkdirAll(dir string) error {
	exists, err := hasDir(dir)
	if err != nil {
		return err
	}
	if exists {
		info, err := os.Stat(dir)
		if err != nil {
			return err
		}
		if info.Mode().Perm() != 0700 {
			return errors.New("dir already exists with wrong permissions")
		}
	}
	return os.MkdirAll(dir, 0700)
}

// WriteFile is the static-analysis enforced method for writing binary data to a file
// in our project, enforcing a single entrypoint with standardized permissions.
func WriteFile(file string, data []byte) error {
	exists, err := hasFile(file)
	if err != nil {
		return err
	}
	if exists {
		info, err := os.Stat(file)
		if err != nil {
			return err
		}
		if info.Mode() != 0600 {
			return errors.New("file already exists with wrong permissions")
		}
	}
	return ioutil.WriteFile(file, data, 0600)
}

func hasDir(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false, nil
	}
	if info == nil {
		return false, err
	}
	return info.IsDir(), err
}

func hasFile(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return false, err
		}
		return false, nil
	}
	return info != nil && !info.IsDir(), nil
}
