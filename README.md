# Mini Skynet CLI

This tools allow you to upload/download file from SIA Skynet shared plateform easily
(Made during a Bootcamp project [Theme: "Choosing a file hosting platform and develop a CLI"])
## Installation

```bash
go build -o skynetcli miniSkynetCli.go
```
## Usage

- **Upload**
```bash
skynetcli upload file.png
```
>*or*
``` bash
skynetcli -u file.png
skynetcli --upload file.png
```
- **Download**
```bash
skynetcli download skylink #vAJjNMDWDTIhZISFiXesRcjgAMfL
```
> *or*
``` bash
skynetcli -d vAJjNMDWDTIhZISFiXesRcjgAMfL
skynetcli --upload https://siasky.net/vAJjNMDWDTIhZISFiXesRcjgAMfL
```
- **Help**
```bash
skynetcli -h #--help / help
```
## Me
[LinkedIn](https://fr.linkedin.com/in/kenji-duriez-9b93bb141)
