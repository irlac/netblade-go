package main

import (
	"fmt"
	"os"
	"bufio"
	"net/url"
	"regexp"
	"github.com/anaskhan96/soup"
)

var save_dest string

func main() {
	fmt.Println("Save destination is ", save_dest)
	choice := SiteChoice()
	var p string
	if choice == "i" {
		p = "insta"
		Retrieve("camden_richter", p)
	} else if choice == "v" {
		p = "vsco"
		Retrieve("richtercamden", p)
	} else {
		p = "both"
	}
	
}

func SiteChoice() string {
	fmt.Print("[i]nstagram, [v]sco, or [b]oth? ")
	reader := bufio.NewReader(os.Stdin)
	platform, _ := reader.ReadString('\n')
	platform = string(platform[0])
	if platform == "i" {
		return "i"
	} else if platform == "v" {
		return "v"
	} else {
		return "b"
	}
}

func Retrieve(uname, pform string) () {
	save_dest = "./" + uname + "/"
	fmt.Printf("Username: %s\nPlatform: %s\n", uname, pform)
	if pform == "vsco" {
		base := "https://vsco.co/" + uname + "/images/"
		res, _ := soup.Get(base)
		body := soup.HTMLParse(res)
		preval := body.FindAll("script")
		re := regexp.MustCompile("responsiveUrl\":\"(.+?)\"")
		for _, val := range preval {
			matches := re.FindAllStringSubmatch(val.Text(), -1)
			if len(matches) > 0 {
				fmt.Println("Found " , len(matches) , " urls.")
				fmt.Print("Downloading")
			}
			for _, ret := range matches {
				for x := range ret {
					parse, _ := url.Parse(ret[x])
					if parse != nil {
						fmt.Print(".")
						parse_string := parse.String()
						parse_slice := []rune(parse_string)
						var name string
						if len(parse_slice) > 60 {
							name = string(parse_slice[65:])
						}
						out_url := "https://" + parse_string
						out, _ := soup.Get(out_url)
						Call(out, name)
					}
				}
			}
		}
	} else if pform == "insta" {
		base := "https://www.instagram.com/" + uname
		res, _ := soup.Get(base)
		body := soup.HTMLParse(res)
		preval := body.FindAll("script")
		re := regexp.MustCompile("shortcode\":\"(.+?)\"")
		for _, val := range preval {
			matches := re.FindAllStringSubmatch(val.Text(), -1)
			if len(matches) > 0 {
				fmt.Println("Found " , len(matches) , " urls.")
				fmt.Print("Downloading")
			}
			for _, ret := range matches {
				for x := range ret {
					parse, _ := url.Parse(ret[x])
					if parse != nil {
						fmt.Print(".")
						fmt.Print(parse)
						parse_string := parse.String()
						parse_slice := []rune(parse_string)
						var name string
						if len(parse_slice) > 60 {
							name = string(parse_slice[65:])
						}
						out_url := "https://instagram.com/p/" + parse_string + "/media?size=l"
						out, _ := soup.Get(out_url)
						Call(out, name)
					}
				}
			}
		}
	}
}

func Call(response, name string) {
	if _, err := os.Stat(save_dest); os.IsNotExist(err) {
    	os.Mkdir(save_dest, os.ModeDir)
	}
	if _, err := os.Stat(save_dest + name); os.IsNotExist(err) {
    	file, _ := os.Create(save_dest + name)
		file.WriteString(response)
	}
}