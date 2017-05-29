package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".eml") {

			// Check full email, and remove full email checks.
			rawEmailData, _ := ioutil.ReadFile(file.Name())
			rawEmail := string(rawEmailData)
			if !strings.HasPrefix(rawEmail, "ContentType +=+=+= text/plain; charset=ascii|||Content +=+=+= ") {
				fmt.Println("wrong email! " + file.Name())
				continue
			} else {
				rawEmail = strings.TrimPrefix(rawEmail, "ContentType +=+=+= text/plain; charset=ascii|||Content +=+=+= ")
			}

			if !strings.HasSuffix(rawEmail, "|||||") {
				fmt.Println("Corrupt! " + file.Name())
				continue
			} else {
				rawEmail = strings.TrimSuffix(rawEmail, "|||||")
			}

			ioutil.WriteFile(file.Name()+".enc", []byte(rawEmail), file.Mode())

			// Process email
			emailId := strings.TrimSuffix(file.Name(), ".eml")
			cmd := exec.Command("gpg", "--passphrase", os.Getenv("GPG_PW"), "-o", "dump/"+emailId+"_decrypted.txt", "-d", file.Name()+".enc")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Start()
			err := cmd.Wait()
			if err != nil {
				fmt.Println("error : ", err)
				fmt.Println(cmd.Output())
				os.Remove(file.Name() + ".enc")
				continue
			}

			os.Remove(file.Name() + ".enc")
			os.Remove(file.Name())

		}
	}
}
