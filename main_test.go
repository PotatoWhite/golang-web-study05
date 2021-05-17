package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPage(t *testing.T)  {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestDecoHandler(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf)

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	reader := bufio.NewReader(buf)
	line, _, err := reader.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER 1] Started")

	line, _, err = reader.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER 2] Param")

	line, _, err = reader.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER 2] Completed")

	line, _, err = reader.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER 1] Completed")
}