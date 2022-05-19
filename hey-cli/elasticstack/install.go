package elasticstack

import (
	"path"

	"f3s.tech/hey-utils/fileutil"
)

var (
	elasticsearch = &app{
		name:     "elasticsearch",
		imageUrl: "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.2.0-windows-x86_64.zip",
	}

	kibana = &app{
		name:     "kibana",
		imageUrl: "https://artifacts.elastic.co/downloads/kibana/kibana-8.2.0-windows-x86_64.zip",
	}
)

type Installer struct {
	InstallFolder string
}

func (i *Installer) InstallAll() error {
	if err := i.InstallElasticsearch(); err != nil {
		return err
	}
	return i.InstallKibana()

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
	name     string
	imageUrl string
}

func (a *app) download(installer *Installer) error {

	if !fileutil.Exists(a.getDownloadImagePath(installer)) {
		return fileutil.Download(a.imageUrl, installer.getDownloadPath())
	}
	return nil
}

func (a *app) unzip(installer *Installer) error {
	if fileutil.Exists(a.getDownloadImagePath(installer)) {
		fileutil.Unzip(a.getDownloadImagePath(installer), installer.getAppsPath())
	}
	return nil
}

func (a *app) getDownloadImagePath(installer *Installer) string {
	return path.Join(installer.getDownloadPath(), a.getImageName())
}

func (a *app) getImageName() string {
	return path.Base(a.imageUrl)
}
