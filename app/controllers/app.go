package controllers

import (
	"github.com/revel/revel"
	"math/rand"
	"time"
)

type App struct {
	*revel.Controller
}

func getRandomBandouImage() (result string) {
	imageNames := []string{
		"ban1.jpg",
		"ban2.jpg",
		"ban3.jpg",
		"matsu1.jpg",
		"matsu2.jpg",
		"matsu3.jpg",
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return imageNames[rand.Intn(len(imageNames))]
}

func (c App) Index() revel.Result {
	greeting := "hello"
	imageName := getRandomBandouImage()
	
	return c.Render(greeting, imageName)
}

func (c App) Hello(myName string) revel.Result {

	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	imageName := getRandomBandouImage()

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName, imageName)
}
