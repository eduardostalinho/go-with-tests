package osexec

import (
	"bytes"
	"io/ioutil"
	"io"
	"encoding/xml"
	"os/exec"
	"strings"
)

type Payload struct {
	Message string `xml:"message"`
}

func GetData(r io.Reader) string {
	var payload Payload
	xml.NewDecoder(r).Decode(&payload)
	return strings.ToUpper(payload.Message)
}

func getXMLFromCommand() io.Reader {
	cmd := exec.Command("cat", "msg.xml")
	out, _ := cmd.StdoutPipe()

	cmd.Start()
	data, _ := ioutil.ReadAll(out)
	cmd.Wait()

	return bytes.NewReader(data)
}