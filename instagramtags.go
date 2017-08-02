package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
	"bufio"
)

type pictureTagStruct struct {
	Results []struct {
		TaggingID interface{} `json:"tagging_id"`
		Image string `json:"image"`
		Tags []struct {
			Confidence float64 `json:"confidence"`
			Tag string `json:"tag"`
		} `json:"tags"`
	} `json:"results"`
}

func main() {
client := &http.Client{}
api_key := ""
api_secret := ""
image_url := "http://docs.imagga.com/static/images/docs/sample/japan-605234_1280.jpg" // change this url to get hashtags for different images

//post image to an API that finds associated words - imagaa (can register a free account with limited requests):
req, _ := http.NewRequest("GET", "https://api.imagga.com/v1/tagging?url="+image_url, nil)
req.SetBasicAuth(api_key, api_secret)

resp, err := client.Do(req)

if err != nil {
fmt.Println("Error when sending request to the server")
return
}

//get the response from the API
defer resp.Body.Close()
resp_body, _ := ioutil.ReadAll(resp.Body)

var pictureTag pictureTagStruct
json.Unmarshal(resp_body, &pictureTag)

// Ask user to give the image a caption
reader := bufio.NewReader(os.Stdin)
fmt.Print("What caption do you want to give your photo: ")
myPhotoComment, _ := reader.ReadString('\n')

// format the list of tags returned from API, only return 20 tags
fmt.Println(resp.Status)
count := 0
var lsOfTags [] string
for _, tag := range pictureTag.Results[0].Tags{
	if count < 20 {
		word := strings.Replace(tag.Tag, " ", "", -1)
		lsOfTags = append(lsOfTags, word)
		count += 1
	} else {
		break
	}
}

// Turn tags into hashtag list
finalHashTagList := ""
for _, each := range lsOfTags {
	hashtag := "#" + each
	finalHashTagList = finalHashTagList + " " + hashtag
}

// print the hashtag list at the end of the image caption given by the user
fmt.Println(myPhotoComment + finalHashTagList)
}