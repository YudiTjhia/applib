package ftp

import (
	bytes2 "bytes"
	"encoding/json"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"strconv"
)

type FtpClientConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func (config FtpClientConfig) ToAddr() string {
	return config.Host + ":" + strconv.Itoa(config.Port)
}

func (config FtpClientConfig) ReadConfigFile(configFile string) error {
	bytes, err := ioutil.ReadFile(configFile)
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return err
	}
	return nil
}

type FtpClient struct {
	conn *ftp.ServerConn
}

func CreateFtpClient(config FtpClientConfig) (*FtpClient, error) {
	ftpClient := FtpClient{}
	conn, err := ftp.Dial(config.ToAddr())
	if err != nil {
		return nil, err
	}
	ftpClient.conn = conn

	err = conn.Login(config.User, config.Pass)
	if err != nil {
		return nil, err
	}
	return &ftpClient, nil
}

func (client *FtpClient) Close() {

	client.conn.Logout()

	client.conn.Quit()

}

func (client FtpClient) ReadFile(filePath string) ([]byte, error) {
	resp, err := client.conn.Retr(filePath)
	if err != nil {
		return []byte{}, err
	}

	respBytes, err := ioutil.ReadAll(resp)
	resp.Close()
	if err != nil {
		return []byte{}, err
	}
	return respBytes, nil
}

func (client FtpClient) WriteFile(bytes []byte, dstFilePath string) error {
	buff := bytes2.NewBuffer(bytes)
	err := client.conn.Stor(dstFilePath, buff)
	if err != nil {
		return err
	}
	return nil
}

func (client FtpClient) RemoveFile(dstFilePath string) error {
	err := client.conn.Delete(dstFilePath)
	if err != nil {
		return err
	}
	return nil
}

func (client FtpClient) Mkdir(dstFilePath string) error {
	list, err := client.conn.List(".")
	found := false
	if err == nil {
		for _, d := range list {
			if d.Name == dstFilePath {
				found = true
				break
			}
		}
	}
	if !found {
		err := client.conn.MakeDir(dstFilePath)
		if err != nil {
			return err
		}
	}

	return nil
}
