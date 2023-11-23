package liter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
)

const (
	DefaultApiServer     = "https://api.mojang.com/"
	DefaultAuthServer    = "https://authserver.mojang.com/"
	DefaultSessionServer = "https://sessionserver.mojang.com/session/minecraft/"
)

var (
	ErrPlayerNameInvalid = errors.New("Player name is invalid")
	ErrPlayerNotExists   = errors.New("Player not exists on the server")
)

type AuthClient struct {
	Client *http.Client

	ApiServer     string
	AuthServer    string
	SessionServer string
}

var DefaultAuthClient = &AuthClient{
	Client: &http.Client{},

	ApiServer:     DefaultApiServer,
	AuthServer:    DefaultAuthServer,
	SessionServer: DefaultSessionServer,
}

type HttpStatusError struct {
	Code int
}

var _ error = (*HttpStatusError)(nil)

func (e *HttpStatusError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, http.StatusText(e.Code))
}

func (cli *AuthClient) GetPlayerInfo(name string) (player PlayerInfo, err error) {
	const EndPoint = "/users/profiles/minecraft"
	if len(name) == 0 {
		err = ErrPlayerNameInvalid
		return
	}
	url := joinUrl(cli.ApiServer, EndPoint, name)
	var res *http.Response
	if res, err = cli.Client.Get(url); err != nil {
		return
	}
	if res.StatusCode == http.StatusNoContent {
		err = ErrPlayerNotExists
		return
	}
	if res.StatusCode != http.StatusOK {
		err = &HttpStatusError{res.StatusCode}
		return
	}
	var body []byte
	if body, err = readBody(res); err != nil {
		return
	}
	if err = json.Unmarshal(body, &player); err != nil {
		return
	}
	return
}

func (cli *AuthClient) GetPlayerProfile(id UUID) (profile *PlayerProfile, err error) {
	const EndPoint = "/profile"
	url := joinUrl(cli.SessionServer, EndPoint, id.String())
	var res *http.Response
	if res, err = cli.Client.Get(url); err != nil {
		return
	}
	if res.StatusCode == http.StatusNoContent {
		err = ErrPlayerNotExists
		return
	}
	if res.StatusCode != http.StatusOK {
		err = &HttpStatusError{res.StatusCode}
		return
	}
	var body []byte
	if body, err = readBody(res); err != nil {
		return
	}
	profile = new(PlayerProfile)
	if err = json.Unmarshal(body, profile); err != nil {
		return
	}
	return
}

func (cli *AuthClient) auth_request(point string, req any, res any) (err error) {
	url := joinUrl(cli.AuthServer, point)
	var buf []byte
	if buf, err = json.Marshal(req); err != nil {
		return
	}
	var res0 *http.Response
	if res0, err = cli.Client.Post(url, "application/json", bytes.NewReader(buf)); err != nil {
		return
	}

	var body []byte
	if body, err = readBody(res0); err != nil {
		return
	}

	failed := res0.StatusCode/100 != 2 /* if not 2xx */
	if failed {
		yggErr := new(YggError)
		if err = json.Unmarshal(body, yggErr); err != nil {
			return
		}
		return yggErr
	}
	if res0.StatusCode != http.StatusNoContent {
		if len(body) > 0 {
			if err = json.Unmarshal(body, res); err != nil {
				return
			}
		}
	}
	return
}

type YggError struct {
	Code         int    `json:"-"`
	ErrorShort   string `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	Cause        string `json:"cause,omitempty"`
}

var _ error = (*YggError)(nil)

func (e *YggError) Error() (s string) {
	s = fmt.Sprintf("HTTP %d %s: %s", e.Code, e.ErrorShort, e.ErrorMessage)
	if len(e.Cause) > 0 {
		s += " caused by " + e.Cause
	}
	return
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessData struct {
	ClientToken string `json:"clientToken"`
	AccessToken string `json:"accessToken"`
}

type Prop struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Props []Prop

type UserData struct {
	Username   string `json:"username"`
	Id         string `json:"id"`
	Properties Props  `json:"properties"`
}

type AuthData struct {
	LoginData
	ClientToken string `json:"clientToken"`
}

type AuthResponse struct {
	User        UserData `json:"user"`
	ClientToken string   `json:"clientToken"`
	AccessToken string   `json:"accessToken"`

	AvailableProfiles []PlayerInfo `json:"availableProfiles"`
	SelectedProfile   PlayerInfo   `json:"selectedProfile"`
}

func (cli *AuthClient) Auth(data AuthData) (res *AuthResponse, err error) {
	const EndPoint = "/authenticate"
	type agentT struct {
		Name    string `json:"name"`
		Version int    `json:"version"`
	}
	type payload struct {
		Agent agentT `json:"agent"`
		AuthData
		RequestUser bool `json:"requestUser"`
	}
	res = new(AuthResponse)
	if err = cli.auth_request(EndPoint, payload{
		Agent: agentT{
			Name:    "Minecraft",
			Version: 1,
		},
		AuthData:    data,
		RequestUser: true,
	}, res); err != nil {
		return
	}
	return
}

type RefreshResponse struct {
	User        UserData `json:"user"`
	ClientToken string   `json:"clientToken"`
	AccessToken string   `json:"accessToken"`

	SelectedProfile PlayerInfo `json:"selectedProfile"`
}

func (cli *AuthClient) Refresh(data AccessData) (res *RefreshResponse, err error) {
	const EndPoint = "/refresh"
	type payload struct {
		AccessData
		RequestUser bool `json:"requestUser"`
	}
	res = new(RefreshResponse)
	if err = cli.auth_request(EndPoint, payload{
		AccessData:  data,
		RequestUser: true,
	}, res); err != nil {
		return
	}
	return
}

func (cli *AuthClient) Validate(data AccessData) (err error) {
	const EndPoint = "/validate"
	if err = cli.auth_request(EndPoint, data, nil); err != nil {
		return
	}
	return
}

func (cli *AuthClient) Signout(data LoginData) (err error) {
	const EndPoint = "/signout"
	if err = cli.auth_request(EndPoint, data, nil); err != nil {
		return
	}
	return
}

func (cli *AuthClient) Invalidate(data AccessData) (err error) {
	const EndPoint = "/invalidate"
	if err = cli.auth_request(EndPoint, data, nil); err != nil {
		return
	}
	return
}

func joinUrl(base string, paths ...string) (url string) {
	var scheme string
	if i := strings.Index(base, "://"); i != -1 {
		scheme, base = base[:i], base[i+3:]
	}
	url = path.Join(base, path.Join(paths...))
	if len(scheme) > 0 {
		url = scheme + "://" + url
	}
	return
}

func readBody(res *http.Response) (body []byte, err error) {
	defer res.Body.Close()
	if res.StatusCode == http.StatusNoContent {
		return
	}
	if l := res.ContentLength; l >= 0 {
		body = make([]byte, l)
		if _, err = io.ReadFull(res.Body, body); err != nil {
			return
		}
	} else if body, err = io.ReadAll(res.Body); err != nil {
		return
	}
	return
}
