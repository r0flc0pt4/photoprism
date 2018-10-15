package photoprism

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"os"
	"testing"
)

const testDataPath = "testdata"
const testDataUrl = "https://www.dropbox.com/s/na9p9wwt98l7m5b/import.zip?dl=1"
const testDataHash = "ed3bdb2fe86ea662bc863b63e219b47b8d9a74024757007f7979887d"
const testConfigFile = "../../configs/photoprism.yml"

var darktableCli = "/usr/bin/darktable-cli"
var testDataZip = GetExpandedFilename(testDataPath + "/import.zip")
var assetsPath = GetExpandedFilename("../../assets")
var thumbnailsPath = GetExpandedFilename(testDataPath + "/thumbnails")
var originalsPath = GetExpandedFilename(testDataPath + "/originals")
var importPath = GetExpandedFilename(testDataPath + "/import")
var exportPath = GetExpandedFilename(testDataPath + "/export")
var databaseDriver = "mysql"
var databaseDsn = "photoprism:photoprism@tcp(database:3306)/photoprism?parseTime=true"

func (c *Config) RemoveTestData(t *testing.T) {
	os.RemoveAll(c.ImportPath)
	os.RemoveAll(c.ExportPath)
	os.RemoveAll(c.OriginalsPath)
	os.RemoveAll(c.ThumbnailsPath)
}

func (c *Config) DownloadTestData(t *testing.T) {
	if fileExists(testDataZip) {
		hash := fileHash(testDataZip)

		if hash != testDataHash {
			os.Remove(testDataZip)
			t.Logf("Removed outdated test data zip file (fingerprint %s)\n", hash)
		}
	}

	if !fileExists(testDataZip) {
		fmt.Printf("Downloading latest test data zip file from %s\n", testDataUrl)

		if err := downloadFile(testDataZip, testDataUrl); err != nil {
			fmt.Printf("Download failed: %s\n", err.Error())
		}
	}
}

func (c *Config) UnzipTestData(t *testing.T) {
	if _, err := unzip(testDataZip, testDataPath); err != nil {
		t.Logf("Could not unzip test data: %s\n", err.Error())
	}
}

func (c *Config) InitializeTestData(t *testing.T) {
	t.Log("Initializing test data")

	c.RemoveTestData(t)

	c.DownloadTestData(t)

	c.UnzipTestData(t)
}

func NewTestConfig() *Config {
	return &Config{
		Debug:          false,
		AssetsPath:     assetsPath,
		ThumbnailsPath: thumbnailsPath,
		OriginalsPath:  originalsPath,
		ImportPath:     importPath,
		ExportPath:     exportPath,
		DarktableCli:   darktableCli,
		DatabaseDriver: databaseDriver,
		DatabaseDsn:    databaseDsn,
	}
}

func getTestCliContext() *cli.Context {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Bool("debug", false, "doc")
	globalSet.String("config-file", testConfigFile, "doc")
	globalSet.String("assets-path", assetsPath, "doc")
	globalSet.String("originals-path", originalsPath, "doc")
	globalSet.String("darktable-cli", darktableCli, "doc")

	c := cli.NewContext(nil, globalSet, nil)

	c.Set("config-file", testConfigFile)
	c.Set("assets-path", assetsPath)
	c.Set("originals-path", originalsPath)
	c.Set("darktable-cli", darktableCli)

	return c
}

func TestNewConfig(t *testing.T) {
	context := getTestCliContext()

	assert.True(t, context.IsSet("assets-path"))
	assert.False(t, context.Bool("debug"))

	c := NewConfig(context)

	assert.IsType(t, &Config{}, c)

	assert.Equal(t, assetsPath, c.AssetsPath)
	assert.False(t, c.Debug)
}

func TestConfig_SetValuesFromFile(t *testing.T) {
	c := NewConfig(getTestCliContext())

	c.SetValuesFromFile(GetExpandedFilename(testConfigFile))

	assert.Equal(t, "/var/photoprism", c.AssetsPath)
	assert.Equal(t, "/var/photoprism/thumbnails", c.ThumbnailsPath)
	assert.Equal(t, "/var/photoprism/photos/originals", c.OriginalsPath)
	assert.Equal(t, "/var/photoprism/photos/import", c.ImportPath)
	assert.Equal(t, "/var/photoprism/photos/export", c.ExportPath)
	assert.Equal(t, databaseDriver, c.DatabaseDriver)
	assert.Equal(t, databaseDsn, c.DatabaseDsn)
}

func TestConfig_ConnectToDatabase(t *testing.T) {
	c := NewTestConfig()

	c.ConnectToDatabase()

	db := c.GetDb()

	assert.IsType(t, &gorm.DB{}, db)
}
