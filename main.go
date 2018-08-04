package main

import (
	"bufio"
	"fmt"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type keys struct {
	Token  string `yaml:"Token"`
	Key    string `yaml:"Key"`
	Secret string `yaml:"Secret"`
}

type path struct {
	Folder string `yaml:"Folder"`
}

func main() {
	k, err := getKeys("./secret.yaml")
	if err != nil {
		panic(err)
	}

	config := dropbox.Config{
		Token:    k.Token,
		LogLevel: dropbox.LogOff,
	}

	// TODO: use oauth2 instead of access token(may be needed to use fmt.Scan to send auth code)
	//conf := &oauth2.Config{
	//	ClientID:     k.Token,
	//	ClientSecret: k.Secret,
	//	//Endpoint:     dropbox.OAuthEndpoint(domain),
	//}
	//
	//ctx := context.Background()
	//token, err = conf.Exchange(ctx, code)
	//
	//config := dropbox.Config{
	//	Token:    token.,
	//	LogLevel: dropbox.LogInfo,
	//}

	//dbx := users.New(config)
	//_, err = dbx.GetCurrentAccount()
	c := files.New(config)
	if err != nil {
		panic(err)
	}

	p, err := getFolderPath("./path.yaml")
	if err != nil {
		panic(err)
	}

	arg := files.NewListFolderArg(p.Folder)
	res, err := c.ListFolder(arg)
	if err != nil {
		panic(err)
	}

	fw, err := os.Create("./undownloaded.log")
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	w := bufio.NewWriter(fw)
	defer w.Flush()

	for _, entry := range res.Entries {
		switch f := entry.(type) {
		case *files.FileMetadata:
			fmt.Println("download: ", f.Name)

			arg := files.NewDownloadArg(p.Folder + "/" + f.Name)
			res, contents, err := c.Download(arg)
			if err != nil {
				panic(err)
			}
			// TODO: mkdir if the download folder is not found
			fw, err := os.Create("contents/" + res.Name)
			if err != nil {
				panic(err)
			}
			io.Copy(fw, contents)

			contents.Close()
			fw.Close()
		case *files.FolderMetadata:
			w.Write([]byte(f.Name))
		}
	}
}

func getKeys(path string) (*keys, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	k := &keys{}
	r := bufio.NewReader(f)
	if err != nil {
		return nil, err
	}

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(l, &k)
		if err != nil {
			return nil, err
		}
	}
	return k, nil
}

func getFolderPath(filepath string) (*path, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := &path{}
	r := bufio.NewReader(f)
	if err != nil {
		return nil, err
	}

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(l, &p)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}
