package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

const (
	neptuneBaseEnpdoint = "https://np.ironhelmet.com"
	loginEndpoint       = "arequest/login"
	getShipsEndpoint    = "grequest/intel_data"
	cookieHeaderKey     = "Set-Cookie"
)

type NeptuneClientInterface interface {
	Login()
	GetShips()
}

type NeptuneClient struct {
	userConfig  UserConfig
	gameID      string
	apiKey      string
	authCookies string
}

func NewNeptuneClient(gameID string, apiKey string, config UserConfig) *NeptuneClient {
	return &NeptuneClient{
		gameID:      gameID,
		apiKey:      apiKey,
		authCookies: "",
		userConfig:  config,
	}
}

func (np *NeptuneClient) GetShips() GetIntelDataResponse {
	orderkeys := map[string]string{"type": "intel_data", "game_number": np.gameID}

	formData, formDataContentType, err := createRequestFormData(orderkeys)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", neptuneBaseEnpdoint, getShipsEndpoint), &formData)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", formDataContentType)
	req.Header.Set("Cookie", np.authCookies)

	// fmt.Printf("Full request: %v\n", req)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	//fmt.Printf("Result: %v\n", res)

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	var intelResponse GetIntelDataResponse

	_ = json.Unmarshal(bodyBytes, &intelResponse)

	fmt.Printf("Parsed body: %v\n", intelResponse)

	return intelResponse
}

func (np *NeptuneClient) Login() {

	loginKeys := map[string]string{"type": "login", "alias": np.userConfig.Username, "password": np.userConfig.Password}

	formData, formDataContentType, err := createRequestFormData(loginKeys)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", neptuneBaseEnpdoint, loginEndpoint), &formData)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", formDataContentType)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	np.authCookies = res.Header.Get(cookieHeaderKey)
}

func createRequestFormData(formData map[string]string) (bytes.Buffer, string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	defer w.Close()
	for key, value := range formData {
		fw, err := w.CreateFormField(key)
		if err != nil {
			return b, "", err
		}

		if _, err = fw.Write([]byte(value)); err != nil {
			return b, "", err
		}
	}

	formDataContentType := w.FormDataContentType()

	return b, formDataContentType, nil
}
