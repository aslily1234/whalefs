package api

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

//TODO: delete
func DeleteFile(host string, port int, vid int, fid uint64, filename string) error {
	url := fmt.Sprintf("http://%s:%d/%d/%d/%s", host, port, vid, fid, filename)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusAccepted {
		return nil
	}else {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}
}