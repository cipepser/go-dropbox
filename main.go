package main

import (
	"bufio"
	"fmt"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type keys struct {
	Token  string `yaml:"Token"`
	Secret string `yaml:"Secret"`
}

func main() {
	k, err := getKeys("./secret.yaml")
	if err != nil {
		panic(err)
	}

	config := dropbox.Config{
		Token: k.Token,
	}
	fmt.Println(config)
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
