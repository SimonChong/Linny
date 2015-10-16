package creator

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/satori/go.uuid"
)

func Create(adDirName string) {

	_, err := os.Stat("configLinny.json")
	if os.IsNotExist(err) {
		ioutil.WriteFile("configLinny.json", []byte(configLinny(adDirName)), os.ModePerm)
	}
	os.Mkdir(adDirName, os.ModePerm)
	ioutil.WriteFile(adDirName+"/configAd.json", []byte(configAd()), os.ModePerm)
	ioutil.WriteFile(adDirName+"/header.frag", []byte(headerFrag), os.ModePerm)
	ioutil.WriteFile(adDirName+"/footer.frag", []byte(footerFrag), os.ModePerm)
	os.Mkdir(adDirName+"/assets", os.ModePerm)
	ioutil.WriteFile(adDirName+"/assets/index.html", []byte(index), os.ModePerm)
}

func configLinny(adDir string) string {
	return `{
    "ContentRoot": "./` + adDir + `"
}`
}

func configAd() string {

	return `{
	"Id" : "` + strings.Replace(uuid.NewV1().String(), "-", "", -1) + `",
	"Name": "Example Campaign",
	"HeaderFrag" : "header.frag",
	"FooterFrag" : "footer.frag"
}`
}

var headerFrag = `<html>`

var footerFrag = `</html>`

var index = `<div>Hello World</div>`
