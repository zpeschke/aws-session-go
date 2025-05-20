package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os/user"
)

func updateConfig(credFile, profile, accessKey, secretKey, sessionToken, expiration string) {
	if credFile == "" {
		usr, err := user.Current()
		if err != nil {
			fmt.Println("Error getting current user: ", err)
			return
		}
		credFile = usr.HomeDir + "/.aws/credentials"
	}
	cfg, err := ini.Load(credFile)
	if err != nil {
		fmt.Println("Error loading credentials file: ", err)
		return
	}

	if !cfg.HasSection(profile) {
		cfg.NewSection(profile)
	}

	cfg.Section(profile).Key("aws_access_key_id").SetValue(accessKey)
	cfg.Section(profile).Key("aws_secret_access_key").SetValue(secretKey)
	cfg.Section(profile).Key("aws_session_token").SetValue(sessionToken)
	cfg.Section(profile).Key("aws_session_expiration").SetValue(expiration)

	err = cfg.SaveTo(credFile)
	if err != nil {
		fmt.Println("Error saving to credentials file: ", err)
		return
	}

	fmt.Println("Updated credentials file with new session token. Profile: ", profile)
}
