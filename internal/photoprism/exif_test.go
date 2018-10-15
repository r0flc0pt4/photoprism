package photoprism

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMediaFile_GetExifData(t *testing.T) {
	conf := NewTestConfig()

	conf.InitializeTestData(t)

	image1, err := NewMediaFile(conf.ImportPath + "/iphone/IMG_6788.JPG")

	assert.Nil(t, err)

	info, err := image1.GetExifData()

	assert.Empty(t, err)

	assert.IsType(t, &ExifData{}, info)

	assert.Equal(t, "iPhone SE", info.CameraModel)

	image2, err := NewMediaFile(conf.ImportPath + "/raw/IMG_1435.CR2")

	assert.Nil(t, err)

	info, err = image2.GetExifData()

	assert.Empty(t, err)

	assert.IsType(t, &ExifData{}, info)

	assert.Equal(t, "Canon EOS M10", info.CameraModel)
}
