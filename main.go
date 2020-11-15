package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dreeks/AudhumlaVersionImporter/audhumla"
	"github.com/dreeks/AudhumlaVersionImporter/mojang"
)

func main() {

	versionUrl := "https://launchermeta.mojang.com/v1/packages/1c31823bcb47514d958cff19be3d3960d4b8f4ee/1.16.3.json"
	client := &http.Client{Timeout: 10 * time.Second}

	r, err := client.Get(versionUrl)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	mVersion := mojang.Manifest{}
	json.NewDecoder(r.Body).Decode(&mVersion)

	aVersion := audhumla.Version{
		ID:        mVersion.ID,
		Name:      mVersion.ID,
		MainClass: mVersion.MainClass,
		Release:   mVersion.Time,
		Type:      "vanilla",
	}

	b, _ := json.MarshalIndent(aVersion, "", "  ")
	ioutil.WriteFile(mVersion.ID+"_audhumla.json", b, os.ModePerm)
}
