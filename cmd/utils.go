package cmd

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func executeCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating StdoutPipe for Cmd", err)
	}

	defer stdOut.Close()

	scanner := bufio.NewScanner(stdOut)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("Error creating StderrPipe for Cmd", err)
	}

	defer stdErr.Close()

	stdErrScanner := bufio.NewScanner(stdErr)
	go func() {
		for stdErrScanner.Scan() {

			txt := stdErrScanner.Text()

			if !strings.Contains(txt, "no buildable Go source files in") {
				fmt.Printf("%s\n", txt)
			}
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal("Error starting Cmd", err)
	}

	err = cmd.Wait()
	if err != nil {
		if err.Error() != "exit status 1" {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
}

// Given a regular expression, search a file and replace the first instance of the match
func replaceInFile(file string, regex regexp.Regexp, replacement string) {
	input, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	inputString := string(input)
	outputString := regex.ReplaceAllString(inputString, replacement)
	err = os.WriteFile(file, []byte(outputString), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
// Untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func Untar(dst string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}
