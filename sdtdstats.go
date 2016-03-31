package sdtd_stats

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type SdtdStats struct {
	AdminUser  string
	AdminToken string
	Host       string
	Port       int
}

func (s *SdtdStats) SetConnection(host string, port int) {
	s.Host = host
	s.Port = port
}

func (s *SdtdStats) SetLogin(user string, token string) {
	s.AdminUser = user
	s.AdminToken = token
}

func (s *SdtdStats) makeCall(endPoint string) ([]byte, error) {
	return s.makeCallWithParams(endPoint, "")
}

func (s *SdtdStats) makeCallWithParams(endPoint string, params string) ([]byte, error) {
	var body []byte
	var resp *http.Response
	var err error

	resp, err = http.Get(fmt.Sprintf("http://%v:%v/api/%v?adminuser=%v&admintoken=%v&%v", s.Host, s.Port, endPoint, s.AdminUser, s.AdminToken, params))
	if err != nil {
		return body, err
	}

	body, ierr := ioutil.ReadAll(resp.Body)
	if ierr != nil {
		return body, ierr
	}

	return body, nil
}
