package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"path"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")
	file, header, err := r.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	filename := path.Join("avatars", userID+path.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, "Successful")
}
