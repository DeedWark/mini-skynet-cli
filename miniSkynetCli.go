//CLI - SIA Skynet (upload/download)
//Kenji DURIEZ - [DeedWark] - 2020
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	skynet "github.com/NebulousLabs/go-skynet"
)

var fileU string          //file to upload
var fileDir string        //directory to upload
var fileD string          //file to download
var linkD string          //uploaded file link
var client = skynet.New() //INIT THE CLIENT

//
//UPLOAD A FILE (without encryption)
//
func uploadFile(fileU string) {
	skylink, err := client.UploadFile(fileU, skynet.DefaultUploadOptions)
	if err != nil {
		panic("Unable to upload: " + err.Error())
	}
	fmt.Printf("Upload successful, Skylink: %v\n", skylink)
}

//
//UPLOAD A DIRECTORY
//
func uploadDir(fileDir string) {
	url, err := client.UploadDirectory(fileDir, skynet.DefaultUploadOptions)
	if err != nil {
		panic("Unable to upload: " + err.Error())
	}
	fmt.Printf("Upload successful, url: %v\n", url)
}

//
//DOWNLOAD A FILE (without decryption)
//
func downloadFile(linkD string, fileD string) {

	cutLink := strings.Split(linkD, "/") //Remove "/""
	linkDL := cutLink[len(cutLink)-1]    //Get only ID

	err := client.DownloadFile(fileD, linkDL, skynet.DefaultDownloadOptions)
	if err != nil {
		panic("Something went wrong, please try again.\nError: " + err.Error())
	}
	fmt.Println("Dowload successful -> " + fileD)
}

//
//UPLOAD A FILE (with encryption)
//
func secureUpload(fileU string) {
	opts := skynet.DefaultUploadOptions

	//ASK USER SKYKEY
	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("Skykey: ")
	sc.Scan()
	skykey := sc.Text()

	opts.SkykeyName = skykey
	skylink, err := client.UploadFile(fileU, opts)
	if err != nil {
		panic("Unable to upload: " + err.Error())
	}
	fmt.Printf("Upload successful, skylink: %v\n", skylink)
}

//
//DOWNLOAD A FILE (with decryption)
//
func secureDownload(linkD string, fileD string) {
	opts := skynet.DefaultDownloadOptions

	//ASK USER THE KEY WORD
	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("Skykey: ")
	sc.Scan()
	skykey := sc.Text()

	opts.SkykeyName = skykey
	err := client.DownloadFile(fileD, linkD, opts)
	if err != nil {
		panic("Something went wrong, please try again.\nError: " + err.Error())
	}
	fmt.Println("Download successful")
}

func usage() {
	fmt.Println(`Tools to simplify SIA Skynet - File sharing plateform for devs.

USAGE:
-u, --upload		    upload a file
-a, --upload-directory  upload a directory (Must be one file minimum in the directory)
-d, --download		    download a file (default file "downloadedFileSkynet") - Add -o or --output
-s                      upload/download with a Skykey (if needed)
		   Upload   : You must type your skykey to encrypt
		   Download : A skykey is required to download an encrypted file`)
	ex := "Example\r\n" +
		"UPLOAD        : " + os.Args[0] + " -u file.png\r\n" +
		"              : " + os.Args[0] + " -a dir\r\n" +
		"DOWNLOAD      : " + os.Args[0] + " -d vAJjNMDWDTIhZISFiXesRcjgAMfL -o file.png\r\n" +
		"                " + os.Args[0] + " -d https://siasky.net/vAJjNMDWDTIhZISFiXesRcjgAMfL -o file.png\r\n" +
		"(EN|DE)CRYPTED: " + os.Args[0] + " -s -u file.png \r\n" +
		"                " + os.Args[0] + " -s -d vAJjNMDWDTIhZISFiXesRcjgAMfL -o file.png\r\n"
	fmt.Printf(ex)
	os.Exit(1)
}

func main() {

	//IF NO ARGS
	if len(os.Args[1:]) < 1 {
		usage()
	}

	//HELP
	if os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help" {
		usage()
	}

	//FLAGS
	flag.StringVar(&fileU, "u", "", "Upload a file in SIA Skynet\n")
	flag.StringVar(&fileU, "upload", "", "Upload a file in SIA Skynet\n")
	flag.StringVar(&fileDir, "a", "", "Upload a directory in SIA Skynet\n")
	flag.StringVar(&fileDir, "upload-directory", "", "Upload a directory in SIA Skynet\n")
	flag.StringVar(&linkD, "d", "", "Download a file from SIA Skynet\n")
	flag.StringVar(&linkD, "download", "", "Download a file from SIA Skynet\n")
	flag.StringVar(&fileD, "o", "downloadedFileSkynet", "Output file (ex: file.png)\n")
	flag.StringVar(&fileD, "output", "downloadedFileSkynet", "Output file (ex: file.png)\n")
	sec := flag.Bool("s", false, "Encrypted upload / Decrypted download (Skykey required if needed)")
	flag.Parse()

	switch os.Args[1] {
	case "upload":
		fileU = os.Args[2]
		uploadFile(fileU)
		//os.Exit(0)
	case "download":
		linkD = os.Args[2]
		fileD = "downloadedFileSIASkynet"
		if os.Args[3] != "" {
			fileD = os.Args[3]
		}
	}
		//os.Exit(0)

	if fileU != "" {
		if *sec == true {
			secureUpload(fileU)
		} else {
			uploadFile(fileU)
		}
	} else if linkD != "" {
		if *sec == true {
			secureDownload(linkD, fileD)
		} else {
			downloadFile(linkD, fileD)
		}
	} else if fileDir != "" {
		uploadDir(fileDir)
	} else {
		usage()
	}
}
