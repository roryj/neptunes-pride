package lib

import (
	"net/http"
	"mime/multipart"
	"bytes"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const (
	neptuneBaseEnpdoint = "https://np.ironhelmet.com"
	loginEndpoint = "arequest/login"
	getShipsEndpoint = "grequest/intel_data"
	cookieHeaderKey = "Set-Cookie"
)

type NeptuneClientInterface interface {
	Login()
	GetShips()
}

type NeptuneClient struct {
	GameConfig
	apiKey	string
	authCookies	string
}

func NewNeptuneClient(apiKey string, config GameConfig) *NeptuneClient {
	return &NeptuneClient{
		apiKey: apiKey,
		authCookies: "",
		GameConfig: config,
	}
}

func (np *NeptuneClient) GetShips() GetIntelDataResponse {
	orderkeys := map[string]string{"type": "intel_data", "game_number": np.GameId}

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

	fmt.Printf("Full request: %v\n",  req)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	//fmt.Printf("Result: %v\n", res)

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	var intelResponse GetIntelDataResponse

	json.Unmarshal(bodyBytes, &intelResponse)

	fmt.Printf("Parsed body: %v\n", intelResponse)

	return intelResponse
}

func (np *NeptuneClient) Login(username, password string) {

	loginKeys := map[string]string{"type": "login", "alias": username, "password": password}

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
		fw, err := w.CreateFormField(key);
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
