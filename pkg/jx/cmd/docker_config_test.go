package cmd

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type DockerConfigTest struct {
	Auths map[string]*AuthTest `json:"auths,omitempty"`
}

type AuthTest struct {
	Auth  string `json:"auth,omitempty"`
	Email string `json:"email,omitempty"`
}

func TestManipulateDockerConfig(t *testing.T) {

	testFile, err := ioutil.ReadFile("docker_config_test.json.json")
	assert.NoError(t, err)

	config := &DockerConfigTest{}
	err = json.Unmarshal(testFile, config)
	/*assert.NoError(t, err)

	assert.Equal(t, 1, len(config.Auths))*/

	/*for k, v := range config.Auths {
		assert.Equal(t, "https://index.docker.io/v1/", k)
		assert.Equal(t, "foo@gmail.com", v.Email)
		assert.Equal(t, "BASE64_VALUE_OF_USERNAME:PASSWORD", v.Auth)
	}*/
	newConfigData := &AuthTest{}
	newConfigData.Auth = b64.StdEncoding.EncodeToString([]byte("something" + ":" + "something"))
	newConfigData.Email = "something"
	if config.Auths == nil {
		config.Auths = map[string]*AuthTest{}
	}
	config.Auths["something"] = newConfigData

}
