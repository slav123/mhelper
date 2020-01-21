package mhelper

import (
	"testing"
	"time"
)

func TestGetMD5Hash(t *testing.T) {
	data := "http://www.gex.pl/test.html"
	result := GetMD5Hash(data)
	if result != "a3512f5a47ca6fec208b69c546c1022c" {
		t.Errorf("Failed to extract BaseURL %s", result)
	}
}

func TestCleanImgSrc(t *testing.T) {

	data := "/MEDIACCB/PUBLISHINGIMAGES/COMMUNITY/MOVEMENT%20MAY.JPG?RenditionID=7"
	result := CleanImgSrc(data)
	if result != "/MEDIACCB/PUBLISHINGIMAGES/COMMUNITY/MOVEMENT%20MAY.JPG" {
		t.Errorf("Failed to extract BaseURL %s", result)

	}
}

func TestFindImages(t *testing.T) {
	data := "<div><img width=\"620\" height=\"265\" src=\"http://cdn.sydneymedia.com.au/assets/20180329204648/Youth-in-the-City_school-hols-620x265.jpg\" class=\"attachment-thumbnail size-thumbnail wp-post-image\" alt=\"Youth in the City_school hols\"></div>Sydney’s young people will come together in a range of exhilarating events ranging from scooter and skateboard competitions to a basketball contest with the police force as part of Youth Week. Live music, performances and other entertainment will launch the &#8230; <a href=\"http://www.sydneymedia.com.au/fast-pace-for-youth-week/\">Continued</a>"

	result := FindImages(data)

	if result[0] != "http://cdn.sydneymedia.com.au/assets/20180329204648/Youth-in-the-City_school-hols-620x265.jpg" {
		t.Errorf("Failed to 1 FindImages %s", result)
	}

	data = "<div><img src=\"http://cdn.sydneymedia.com.au/assets/20180329204648/Youth-in-the-City_school-hols-620x265.jpg\" width=\"620\" height=\"265\" class=\"attachment-thumbnail size-thumbnail wp-post-image\" alt=\"Youth in the City_school hols\"></div>Sydney’s young people will come together in a range of exhilarating events ranging from scooter and skateboard competitions to a basketball contest with the police force as part of Youth Week. Live music, performances and other entertainment will launch the &#8230; <a href=\"http://www.sydneymedia.com.au/fast-pace-for-youth-week/\">Continued</a>"

	result = FindImages(data)

	if result[0] != "http://cdn.sydneymedia.com.au/assets/20180329204648/Youth-in-the-City_school-hols-620x265.jpg" {
		t.Errorf("Failed to 2 FindImages %s", result)
	}

}

func TestCleanDate(t *testing.T) {
	samples := []string{"Published on 1/1/2011", "Posted on: January"}
	results := []string{"1/1/2011", "January"}

	var result string

	for index, sample := range samples {
		result = CleanDate(sample)

		if results[index] != result {
			t.Errorf("Failed to clean date %s", result)
		}
	}

}

func TestParseDate(t *testing.T) {

	samples := []struct {
		src    string
		result string
		reg    string
	}{
		{"06 August 2018 9:30 am - 03 December 2018 3:00 pm", "06 Aug 2018 00:00:00 UTC", "^(\\d{2})\\s(\\w+)\\s(\\d{4})"},
		{"04 August 2018 2:00 pm", "04 Aug 2018 00:00:00 UTC", "^(\\d{2})\\s(\\w+)\\s(\\d{4})"},
		{"2 July 2014", "02 Jul 2014 00:00:00 UTC", ""},
		{"Mon, 17 Sep 2018 06:38:00 GMT", "17 Sep 2018 06:38:00 UTC", ""},
		{"Wednesday, 12 September 2018 8:30:34 AM", "12 Sep 2018 08:30:34 UTC", ""},
		{"Thursday 13 September 2018", "13 Sep 2018 00:00:00 UTC", ""},
		{"14/09/2018", "14 Sep 2018 00:00:00 UTC", ""},
		{"Thu, 13 Sep 2018 16:30:59 +10:00", "13 Sep 2018 16:30:59 AEST", ""},
		{"2018-09-11T00:00:00+08:00", "11 Sep 2018 00:00:00 AWST", ""},
		{"Wed 03 Oct", "03 Oct 2018 00:00:00 UTC", ""},
	}

	var rtime time.Time
	var result time.Time

	for _, sample := range samples {
		rtime, _ = time.Parse("02 Jan 2006 15:04:05 MST", sample.result)
		result = ParseDate(sample.src, sample.reg)
		if rtime != result {
			t.Errorf("Failed to parse %s got %s expected %s", sample.src, result, rtime)
		}
	}
}

func TestCleanString(t *testing.T) {
	samples := []struct {
		src    string
		result string
	}{
		{"		Boondall,  				                	24 Eton Avenue                        ", "Boondall, 24 Eton Avenue"},
	}

	var result string

	for _, sample := range samples {
		result = CleanString(sample.src)
		if result != sample.result {
			t.Errorf("Failed to clean '%s' : '%s", sample.src, result)
		}
	}

}
