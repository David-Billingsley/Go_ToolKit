package toolkit

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type Tools struct {
	MaxFileSize      int
	AllowedFileTypes []string
}

// RandomString returns a random string with the length found in variable n
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

// UploadedFile is a struct used to save information about an uploaded file
type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int64
}

// UploadFiles retrives files from a web based front end and allow the files to be saved
func (t *Tools) UploadFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	var uploadedFiles []*UploadedFile

	if t.MaxFileSize == 0 {
		t.MaxFileSize = 1024 * 1024 * 1024
	}

	err := r.ParseMultipartForm(int64(t.MaxFileSize))
	if err != nil {
		return nil, errors.New("the uploaded file is too big")
	}

	for _, fHeaders := range r.MultipartForm.File {
		for _, hdr := range fHeaders {
			uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadedFile, error) {
				var uploadedFile UploadedFile
				infile, err := hdr.Open()
				if err != nil {
					return nil, err
				}
				defer infile.Close()

				buff := make([]byte, 512)
				_, err = infile.Read(buff)
				if err != nil {
					return nil, err
				}

				// check to see if the file type is permitted
				allowed := false
				fileType := http.DetectContentType(buff)

				if len(t.AllowedFileTypes) > 0 {
					for _, x := range t.AllowedFileTypes {
						if strings.EqualFold(fileType, x) {
							allowed = true
						}
					}
				} else {
					allowed = true
				}

				if !allowed {
					return nil, errors.New("the uploaded file type is not permitted")
				}

				_, err = infile.Seek(0, 0)
				if err != nil {
					return nil, err
				}

				if renameFile {
					uploadedFile.NewFileName = fmt.Sprintf("%s%s", t.RandomString(25), filepath.Ext(hdr.Filename))
				} else {
					uploadedFile.NewFileName = hdr.Filename
				}

				var outfile *os.File
				defer outfile.Close()

				if outfile, err = os.Create(filepath.Join(uploadDir, uploadedFile.NewFileName)); err != nil {
					return nil, err
				} else {
					fileSize, err := io.Copy(outfile, infile)
					if err != nil {
						return nil, err
					}
					uploadedFile.FileSize = fileSize
				}

				uploadedFiles = append(uploadedFiles, &uploadedFile)

				return uploadedFiles, nil
			}(uploadedFiles)
			if err != nil {
				return uploadedFiles, err
			}
		}
	}
	return uploadedFiles, nil
}

// CreateDirIfNotExist creates a directory, and all necessary parents, if it does not exist
func (t *Tools) CreateDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteDir deletes a directory & files
func (t *Tools) DeleteDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// CopyDir copies a directory & Files
func (t *Tools) CopyDir(path string, orgpath string) error {
	destfile, err := os.Open(path)
	if err != nil {
		return err
	}

	topath, err := os.Open(orgpath)
	if err != nil {
		return err
	}
	defer topath.Close()

	defer destfile.Close()

	_, err = io.Copy(destfile, topath)
	if err != nil {
		return err
	}
	return nil
}

// fixes the json files to remove the items in the front of it to allow the system to read it correctly
func (t *Tools) FixJson(body string, arrayKey string) (arrayContent string) {
	fileContent := string(body)

	// Removes the beging part of the json API return
	resultsarrayKey := arrayKey

	if strings.Contains(fileContent, resultsarrayKey) {

		startIndex := strings.Index(fileContent, resultsarrayKey) + len(resultsarrayKey)

		arrayContent := strings.TrimSpace(fileContent[startIndex:])
		// removes the end } as the first part is removed in arrayKey
		arrayContent = strings.TrimSuffix(arrayContent, "}")
		return arrayContent
	}
	return string(body)
}

// converts epoch time to current time
func (t *Tools) EpochConverMil(epochTime int64) time.Time {
	convertedtime := time.UnixMilli(epochTime)
	return convertedtime
}

// Parses string date into a date format.  Returns either the date or an error
func (t *Tools) DateStrParse(dateStr string) (time.Time, error) {
	formats := []string{"1/2/2006", "1-2-2006"}

	var parsedDate time.Time
	var err error

	for _, format := range formats {
		parsedDate, err = time.Parse(format, dateStr)
		if err == nil {
			break
		}
	}

	return parsedDate, err
}
