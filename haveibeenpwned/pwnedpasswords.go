package haveibeenpwned

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func PwnedPassword(password string) (count int, err error) {

	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)

	prefix := strings.ToUpper(hex.EncodeToString(bs)[:5])
	suffix := strings.ToUpper(hex.EncodeToString(bs)[5:])

	res, err := PwnedPasswordRange(prefix)
	if err != nil {
		return 0, err
	}

	count, err = FindSuffix(suffix, res)
	if err != nil {
		return 0, err
	}

	return count, nil

}

func FindSuffix(suffix string, passwordRange []PasswordRange) (int, error) {
	for _, p := range passwordRange {
		if p.Suffix == suffix {
			return p.Count, nil
		}
	}
	return 0, nil
}

type PasswordRange struct {
	Suffix string
	Count  int
}

func PwnedPasswordRange(prefix string) (passwordRange []PasswordRange, err error) {
	u, err := url.Parse("https://api.pwnedpasswords.com/range/")
	if err != nil {
		return nil, err
	}

	u.Path += prefix

	res, err := ApiCall(u.String())

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

	}
	defer res.Body.Close()

	// slice from new lines
	newLines := strings.Split(string(body), "\n")

	// split newLines by colon
	for _, line := range newLines {
		split := strings.Split(line, ":")
		if len(split) == 2 {
			count, err := strconv.Atoi(strings.TrimRight(split[1], "\r\n"))
			if err != nil {
				return nil, err
			}
			t := PasswordRange{split[0], count}

			passwordRange = append(passwordRange, t)
		}
	}

	return passwordRange, nil
}

func ApiCall(endpoint string) (*http.Response, error) {
	client := &http.Client{}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
