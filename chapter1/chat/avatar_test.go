package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)

	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}

	// set a value
	testURL := "http://url-to-gravatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testURL, nil)

	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present")
	} else {
		if url != testURL {
			t.Error("AuthAvatar.GetAvatarURL should return correct URL")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}

	url, err := gravatar.GetAvatarURL(user)
	if err != nil {
		t.Error("gravatar.GetAvatarURL should not return an error")
	}

	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("Gravatar.GetAvatarURL wrongly returned %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	// make a test avatar file
	filename := path.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var fileAvatar FileSystemAvatar
	// seed user with wrong ID
	user := &chatUser{uniqueID: "123"}
	url, err := fileAvatar.GetAvatarURL(user)
	// ensure error returned is ErrNoAvatarURL
	if err != ErrNoAvatarURL {
		t.Error("fileAvatar.GetAvatarURL should not return an error")
	}

	user.uniqueID = "abc"
	url, err = fileAvatar.GetAvatarURL(user)
	if url != "/avatars/abc.jpg" {
		t.Errorf("fileAvatar.GetAvatarURL wrongly returned %s", url)
	}
}
