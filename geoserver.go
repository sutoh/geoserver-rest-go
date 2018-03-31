package geoserver

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

//GeoServer is the configuration Object
type GeoServer struct {
	WorkspaceName string `yaml:"workspace"`
	ServerURL     string `yaml:"geoserver_url"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	httpClient    *http.Client
	logger        *logrus.Logger
}

//LoadConfig load geoserver config from yaml file
func (g *GeoServer) LoadConfig(configFile string) *GeoServer {

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, g)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	g.SetLogger()
	return g
}

//SetLogger sets instance logger
func (g *GeoServer) SetLogger() *GeoServer {
	g.logger = GetLogger()
	return g
}

// GetGeoserverRequest creates a HTTP request with geoserver credintails and header
func (g *GeoServer) GetGeoserverRequest(
	targetURL string,
	method string,
	accept string,
	data io.Reader,
	contentType string) (request *http.Request, err error) {
	request, err = http.NewRequest(method, targetURL, data)
	if err != nil {
		g.logger.Fatal("Error Creating Geoserver Request\t", err)
		return
	}
	if data != nil {
		request.Header.Set(contentTypeHeader, contentType)
	}
	if accept != "" {
		request.Header.Set(acceptHeader, accept)
	}

	request.SetBasicAuth(g.Username, g.Password)
	return request, err
}
