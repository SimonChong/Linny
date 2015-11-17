package creator

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/satori/go.uuid"
	"github.com/simonchong/linny/creator/resources"
)

//go:generate wgf -i=../resources/creator/header.frag -o=./resources/headerFrag.go -p=resources -c=HeaderFrag
//go:generate wgf -i=../resources/creator/footer.frag -o=./resources/footerFrag.go -p=resources -c=FooterFrag
//go:generate wgf -i=../resources/creator/configLinny.json -o=./resources/configLinnyJSON.go -p=resources -c=ConfigLinnyJSON
//go:generate wgf -i=../resources/creator/configAd.json -o=./resources/configAdJSON.go -p=resources -c=ConfigAdJSON
//go:generate wgf -i=../resources/creator/index.html -o=./resources/indexHTML.go -p=resources -c=IndexHTML

func Create(adDirName string) {

	_, err := os.Stat("configLinny.json")
	if os.IsNotExist(err) {
		ioutil.WriteFile("configLinny.json", []byte(configLinny(adDirName)), os.ModePerm)
	}
	os.Mkdir(adDirName, os.ModePerm)
	ioutil.WriteFile(adDirName+"/configAd.json", []byte(configAd()), os.ModePerm)
	ioutil.WriteFile(adDirName+"/header.frag", []byte(resources.HeaderFrag), os.ModePerm)
	ioutil.WriteFile(adDirName+"/footer.frag", []byte(resources.FooterFrag), os.ModePerm)
	os.Mkdir(adDirName+"/assets", os.ModePerm)
	ioutil.WriteFile(adDirName+"/assets/index.html", []byte(resources.IndexHTML), os.ModePerm)
}

func configLinny(adDir string) string {

	return strings.Replace(resources.ConfigLinnyJSON, "{{DIR}}", adDir, -1)

}

func configAd() string {

	uuid := strings.Replace(uuid.NewV1().String(), "-", "", -1)
	return strings.Replace(resources.ConfigAdJSON, "{{ID}}", uuid, -1)

}
