package twitter_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wipeinc/wipeinc/twitter"
	"github.com/wow-sweetlie/anaconda"
)

var api *anaconda.TwitterApi

var testBase string

var client *twitter.Client

func init() {
	var consumerKey = ""
	var consumerSecret = ""
	var acessToken = ""
	var acessTokenSecret = ""

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api = anaconda.NewTwitterApi(acessToken, acessTokenSecret)
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	parsed, _ := url.Parse(server.URL)
	testBase = parsed.String()
	api.SetLogger(anaconda.BasicLogger)
	api.SetBaseUrl(testBase)

	client = &twitter.Client{
		API: api,
	}

	var endpointElems [][]string
	filepath.Walk("json", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			elems := strings.Split(path, string(os.PathSeparator))[1:]
			endpointElems = append(endpointElems, elems)
		}

		return nil
	})

	for _, elems := range endpointElems {
		endpoint := strings.Replace("/"+path.Join(elems...), "_id_", "?id=", -1)
		filename := filepath.Join(append([]string{"json"}, elems...)...)

		mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
			// if one filename is the prefix of another, the prefix will always match
			// check if there is a more specific filename that matches this request

			// create local variable to avoid closing over `filename`
			sourceFilename := filename

			r.ParseForm()
			form := strings.Replace(r.Form.Encode(), "=", "_", -1)
			form = strings.Replace(form, "&", "_", -1)
			specific := sourceFilename + "_" + form
			_, err := os.Stat(specific)
			if err == nil {
				sourceFilename = specific
			} else {
				if err != nil && !os.IsNotExist(err) {
					fmt.Fprintf(w, "error: %s", err)
					return
				}
			}

			f, err := os.Open(sourceFilename)
			if err != nil {
				// either the file does not exist
				// or something is seriously wrong with the testing environment
				fmt.Fprintf(w, "error: %s", err)
			}
			defer f.Close()

			io.Copy(w, f)
		})
	}
}

func TestGetFriendsIds(t *testing.T) {
	assert := assert.New(t)
	ids, err := client.GetFriendsIds()
	assert.Nil(err)
	assert.Equal(769, len(ids))
}

func TestNewClient(t *testing.T) {
	assert := assert.New(t)
	var accessTokenTest = "ACCESSTOKEN"
	var accessTokenSecretTest = "ACCESSTOKENSECRET"
	var client = twitter.NewClient(accessTokenTest, accessTokenSecretTest)
	assert.Equal(client.API.Credentials.Token, accessTokenTest)
	assert.Equal(client.API.Credentials.Secret, accessTokenSecretTest)
}

func TestGetFollowersIds(t *testing.T) {
	assert := assert.New(t)
	ids, err := client.GetFollowersIds(nil)
	assert.Nil(err)
	assert.Equal(2872, len(ids))
}
