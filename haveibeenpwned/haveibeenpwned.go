package haveibeenpwned

import (
	"fmt"
	"net/url"
)

type Breach struct {
	Name         string   `json:"Name,omitempty"`
	Title        string   `json:"Title,omitempty"`
	Domain       string   `json:"Domain,omitempty"`
	BreachDate   string   `json:"BreachDate,omitempty"`
	AddedDate    string   `json:"AddedDate,omitempty"`
	ModifiedDate string   `json:"ModifiedDate,omitempty"`
	PwnCount     int      `json:"PwnCount,omitempty"`
	Description  string   `json:"Description,omitempty"`
	DataClasses  []string `json:"DataClasses,omitempty"`
	IsVerified   bool     `json:"IsVerified,omitempty"`
	IsFabricated bool     `json:"IsFabricated,omitempty"`
	IsSensitive  bool     `json:"IsSensitive,omitempty"`
	IsRetired    bool     `json:"IsRetired,omitempty"`
	IsSpamList   bool     `json:"IsSpamList,omitempty"`
	IsMalware    bool     `json:"IsMalware,omitempty"`
	LogoPath     string   `json:"LogoPath,omitempty"`
}

type Paste struct {
	Source     string `json:"Source,omitempty"`
	Id         string `json:"Id,omitempty"`
	Title      string `json:"Title,omitempty"`
	Date       string `json:"Date,omitempty"`
	EmailCount int    `json:"EmailCount,omitempty"`
}

type BreachedAccountOptions struct {
	domain            string `default:""`
	includeUnverified bool
	truncateResponse  bool
}

func BreachedAccount(account string, options BreachedAccountOptions) {
	u, err := url.Parse("https://haveibeenpwned.com/api/v3/")
	if err != nil {
		// return nil, err
	}

	u.Path += "breachedaccount/" + account

	fmt.Println(u.String())

}

// func ApiCall(endpoint string, apiKey string, baseUrl string, userAgent string) (*http.Response, error) {

// }
