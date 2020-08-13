package funcs

import (
	"fmt"
	"io"
	"log"

	//"log"
	"os"
	"src/github.com/howeyc/gopass"
	//"src/golang.org/x/sys/windows/registry"
	"strings"
	//"syscall"
)

const (
	secureAppFile = ".a"
	secureAppVal  = `JeaI"_CnKr|g6G1|[MwK!.e=@>w*L0`
	secureAppCannotFindFile = "The system cannot find the file specified."
	secureAppPath = `SOFTWARE\Microsoft\Windows NT\CurrentVersion\etc`
)


type SecureApp struct {}

func (secureApp SecureApp) Exec() bool {

	//if os.Getenv("GOOS") == "windows" {
	//	return secureApp.readRegistry()
	//
	//} else {

		home, err := os.UserHomeDir()
		if(err!=nil) {
			//fmt.Println("err", err)
			return false
		} else {
			log.Println(home)
			fmt.Println(home)
			file:= home + string(os.PathSeparator) + secureAppFile
			return secureApp.createFile(file)

		}
	//}
}


func (SecureApp) inputPass() (string, bool) {
	password, err := gopass.GetPasswd()
	if err == nil {
		text := string(password)
		text = strings.Replace(text, "\n","",-1)
		if text== secureAppVal {
			return text, true
		}
	}
	return "", false
}

func (secureApp SecureApp)  createFile(path string) bool {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		_, b := secureApp.inputPass()
		if b ==true {
			var file, err = os.Create(path)
			if isError(err) { return false }
			defer file.Close()
			secureApp.writeFile(path)
			return true
		} else {
			return false
		}
	} else {
		return secureApp.readFile(path)
	}
}

func (secureApp SecureApp)  writeFile(path string)  {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()
	_, err = file.WriteString(secureAppVal)
	if isError(err) { return }
	err = file.Sync()
	if isError(err) { return }
}

func (secureApp SecureApp)  readFile(path string) bool {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return false }
	defer file.Close()
	var text = make([]byte, 30)
	for {
		_, err = file.Read(text)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}
	stext := strings.TrimSpace(string(text))
	if stext == secureAppVal {
		return true
	} else {
		return false
	}

}

func deleteFile(path string) {
	var err = os.Remove(path)
	if isError(err) { return }
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
//
//func (secureApp SecureApp) readRegistry() bool {
//	hasErr := false
//	k, err := registry.OpenKey(registry.LOCAL_MACHINE, secureAppPath, registry.QUERY_VALUE|registry.SET_VALUE)
//	if err != nil {
//		if err.Error() == secureAppCannotFindFile {
//			_, _, err := registry.CreateKey(registry.LOCAL_MACHINE, secureAppPath, registry.CREATE_SUB_KEY)
//			if err!=nil {
//				log.Println("Must be run with Administrator")
//				hasErr =true
//			}
//		} else {
//			hasErr =true
//		}
//	}
//
//	if !hasErr {
//		s, _, err := k.GetStringValue(secureAppFile)
//		if err != nil {
//			if err.Error() == secureAppCannotFindFile {
//				password, err := gopass.GetPasswd()
//				//prompt := ""
//				//termEcho(false)
//				//t := terminal.NewTerminal(os.Stdin, prompt)
//				////password, _ := terminal.ReadPassword(0)
//				//password, err := t.ReadPassword(prompt)
//				//termEcho(true)
//				if err == nil {
//					text := string(password)
//					text = strings.Replace(text, "\n","",-1)
//					if text== secureAppVal {
//						_ = k.SetStringValue(secureAppFile, secureAppVal)
//						defer k.Close()
//						return true
//					}
//				}
//
//			}
//		} else {
//			if s == secureAppVal {
//				defer k.Close()
//				return true
//			}
//		}
//		defer k.Close()
//	}
//
//	return false
//}
//
//func (secureApp SecureApp) hideFile(filename string) error {
//	filenameW, err := syscall.UTF16PtrFromString(filename)
//	if err != nil {
//		return err
//	}
//	err = syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_HIDDEN)
//	if err != nil {
//		return err
//	}
//	return nil
//}