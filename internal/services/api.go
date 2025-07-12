// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package services

import (
	"dadjoke/internal/logger"
	"dadjoke/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetJoke(jokeType int) utils.JokeAPIResp {
	var resp *http.Response
	var err error
	var data []utils.JokeAPIResp

	if jokeType == 1 {
		resp, err = http.Get(utils.JOKE_API_GENERAL)
	} else {
		resp, err = http.Get(utils.JOKE_API_PROGRAMMING)
	}

	logger.CheckErr(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	logger.CheckErr(err)

	err = json.Unmarshal(body, &data)
	logger.CheckErr(err)

	return data[0]
}

func StartWithSave(jokeType int, filename string) {
	data := GetJoke(jokeType)
	output := fmt.Sprintf("%s\n- %s\n", data.Setup, data.Punchline)
	if filename != "" {
		var outputBytes []byte
		var err error

		if strings.Contains(filename, "json") {
			outputBytes, err = json.MarshalIndent(data, "", "   ")
			logger.CheckErr(err)
		} else {
			outputBytes = []byte(output)
		}

		err = os.WriteFile(filename, outputBytes, 0644)
		logger.CheckErr(err)

		fmt.Printf("Dad Joke saved to %s\n", filename)
	} else {
		fmt.Print(output)
	}
}
