package elasticstack

import (
	"fmt"
	"path"

	"f3s.tech/hey-utils/fileutil"
)

var (
	elasticsearch = &app{
		name:      "elasticsearch",
		appFolder: "elasticsearch-8.2.0",
		imageUrl:  "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.2.0-windows-x86_64.zip",
	}

	kibana = &app{
		name:      "kibana",
		appFolder: "kibana-8.2.0",
		imageUrl:  "https://artifacts.elastic.co/downloads/kibana/kibana-8.2.0-windows-x86_64.zip",
	}
)

type Installer struct {
	InstallFolder string
}

func (i *Installer) InstallAll() error {
	fmt.Println("installing elasticstack version 8.2.0 ...")
	if err := i.InstallElasticsearch(); err != nil {
		return err
	}
	if err := i.InstallKibana(); err != nil {
		fmt.Println("elasticstack version 8.2.0 not fully installed")
		return err
	} else {
		fmt.Println("elasticstack version 8.2.0 installed")
		return nil
	}

}

func (i *Installer) InstallElasticsearch() error {
	err := elasticsearch.download(i)
	if err != nil {
		return err
	}
	return elasticsearch.unzip(i)
}

func (i *Installer) InstallKibana() error {
	if err := kibana.download(i); err != nil {
		return err
	}
	return kibana.unzip(i)
}

func (i *Installer) getDownloadPath() string {
	return path.Join(i.InstallFolder, "downloads")
}

func (i *Installer) getAppsPath() string {
	return path.Join(i.InstallFolder, "apps")
}

type app struct {
	name      string
	appFolder string
	imageUrl  string
}

func (a *app) download(installer *Installer) error {

	if !fileutil.Exists(a.getDownloadImagePath(installer)) {
		fmt.Println("downloding: ", a.name)
		return fileutil.Download(a.imageUrl, installer.getDownloadPath())
	}
	return nil
}

func (a *app) unzip(installer *Installer) error {
	if !fileutil.Exists(a.getAppFolder(installer)) {
		if fileutil.Exists(a.getDownloadImagePath(installer)) {
			fileutil.Unzip(a.getDownloadImagePath(installer), installer.getAppsPath())
		}
	}
	return nil
}

func (a *app) getAppFolder(installer *Installer) string {
	return path.Join(installer.getAppsPath(), a.appFolder)
}

func (a *app) getDownloadImagePath(installer *Installer) string {
	return path.Join(installer.getDownloadPath(), a.getImageName())
}

func (a *app) getImageName() string {
	return path.Base(a.imageUrl)
}
