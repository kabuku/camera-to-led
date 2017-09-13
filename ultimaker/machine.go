package ultimaker

import (
	"github.com/go-openapi/strfmt"
	apiclient "github.com/kabuku/camera-to-led/ultimaker/api/client"
	"github.com/kabuku/camera-to-led/ultimaker/api/client/printer"
	"github.com/kabuku/camera-to-led/ultimaker/api/models"
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
	"log"
	"math"
)

type Ultimaker3 struct {
	config       *Ultimaker3MachineConfig
	client       *apiclient.Ultimaker
	FirstColor   *color.RGBA
	CurrentColor *color.RGBA
}

type Ultimaker3MachineConfig struct {
	Host     string
	BasePath string `mapstructure:"base_path"`
	Auth     *Ultimaker3AuthConfig
}

type Ultimaker3AuthConfig struct {
	Id  string
	Key string
}

func New(config *Ultimaker3MachineConfig) (*Ultimaker3, error) {

	basePath := "/api/v1"
	if config.BasePath != "" {
		basePath = config.BasePath
	}

	// create the transport
	transport := apiclient.NewDigestAuthRuntime(config.Host, basePath, nil, config.Auth.Id, config.Auth.Key)

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	u := &Ultimaker3{
		config,
		client,
		nil,
		nil,
	}

	return u, nil
}

func (u *Ultimaker3) Client() *apiclient.Ultimaker {
	return u.client
}

func (u *Ultimaker3) SetLedColor(rbga color.RGBA) error {
	log.Println(rbga)

	if u.FirstColor != nil && !isDefferentColor(u.FirstColor, &rbga) {
		log.Println("it is same fisrt color")
		return nil
	}

	if u.CurrentColor != nil && !isDefferentColor(u.CurrentColor, &rbga) {
		log.Println("it is same current color")
		return nil
	}
	u.CurrentColor = &rbga

	c := colorful.Color{float64(rbga.R) * colorful.Delta, float64(rbga.G) * colorful.Delta, float64(rbga.B) * colorful.Delta}

	h, s, _ := c.Hsv()

	p := printer.NewPutPrinterLedParams()
	p.SetHue(&models.Led{
		h, float64(s * 100), 100,
	})

	_, err := u.client.Printer.PutPrinterLed(p)

	return err

}

func isDefferentColor(rbga1, rbga2 *color.RGBA) bool {

	delta := math.Abs(colorDiff(rbga1, rbga2))
	log.Println(delta)
	return delta > 0.3

}

func colorDiff(rbga1, rbga2 *color.RGBA) float64 {
	c1 := colorful.Color{float64(rbga1.R) * colorful.Delta, float64(rbga1.G) * colorful.Delta, float64(rbga1.B) * colorful.Delta}
	c2 := colorful.Color{float64(rbga2.R) * colorful.Delta, float64(rbga2.G) * colorful.Delta, float64(rbga2.B) * colorful.Delta}

	return c1.DistanceLab(c2)

}
