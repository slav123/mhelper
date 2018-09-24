package mhelper

import (
"crypto/md5"
"encoding/hex"
"regexp"
"strings"
	"time"
	"github.com/araddon/dateparse"

	"log"
)

// GetMD5Hash calc md5 hash from string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// if your img's are properly formed with doublequotes then use this, it's more efficient.
// var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)

// FindImages - look for images in text
func FindImages(html string) []string {
	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(html, 1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

// CleanImgSrc  clean ? form src
func CleanImgSrc(imgSrc string) string {
	pos := strings.LastIndex(imgSrc, "?")

	if pos >= 0 {
		return imgSrc[0:pos]
	}

	return imgSrc
}

// clean and normalize string from standard stuff
func CleanString(str string) string {

	result := strings.Replace(str, "\r\n", " ", -1)
	result = strings.Replace(result, "\n", " ", -1)
	result = strings.Replace(result, "\t", " ", -1)

	for strings.Contains(result, "  ") {
		result = strings.Replace(result, "  ", " ", -1)
	}
	result = strings.Replace(result, "&#xA;", "", -1)

//	result = strings.Join(strings.Fields(result), " ")


	result = strings.TrimSpace(result)
	return result
}

// cleanDate - remove common elements from date
func CleanDate(date string) string {

	result := CleanString(strings.TrimLeft(date, "Posted on"))
	result = strings.TrimLeft(result, "Published on")
	result = strings.TrimLeft(result, "Posted")
	result = strings.TrimLeft(result, "Published:")

	return CleanString(result)
}

// parse date
func ParseDate(createdString, reg string) time.Time {

	loc, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		panic(err.Error())
	}
	time.Local = loc

	// extra cleanup
	re := regexp.MustCompile("\\n")
	createdString = re.ReplaceAllString(createdString, " ")

	// regexp ?
	if reg != "" {
		re = regexp.MustCompile(reg)
		matches := re.FindAllString(createdString, -1)
		log.Print(matches)
		createdString = strings.Join(matches[:], " ")
	}

	t, err := dateparse.ParseStrict(createdString)
	if err != nil {
		log.Printf("unable to recognize date :%+q\n:", createdString)

		//		return time.Now()
	} else {
		return t
	}

	// basic: Mon Jan 2 15:04:05 -0700 MST 2006

	//2018-09-11T00:00:00+08:00

	patterns := []string{"2 January 2006", "Jan 02 2006", "Mon 2 Jan 2006", "Monday 2 January 2006", "2 January 2006",
	"Mon, 02 Jan 2006 03:04:00 MST", "Monday, 02 January 2006 3:04:05 PM", "02/01/2006", "Mon, 02 Jan 2006 15:05:05 -07:00", "2006-01-02T15:04:05+07:00"}

	for _,pattern := range patterns {
		t, err = time.Parse(pattern, createdString)
		if err != nil {
			//log.Printf("\nfailed with to decode %s with %s", createdString, pattern)
		} else {

			log.Printf("\nmatched decode %s with %s", createdString, pattern)
			return t
		}
	}


	return time.Now()

}