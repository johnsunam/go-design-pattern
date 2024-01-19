package main

type Mobile interface {
	ChargeAppleMobile()
}

type apple struct {
	color string
}

func (a *apple) ChargeAppleMobile() {
	println("Apple color is", a.color)
}

type client struct{}

func (c *client) ChargeMobile(m Mobile) {
	m.ChargeAppleMobile()
}

// adapter implementation for android mobile

type AndroidAdapter struct {
	android *android
}

func (a *AndroidAdapter) ChargeAppleMobile() {
	a.android.ChargeAndroidMobile()
}

type android struct {
	size int
}

func (a *android) ChargeAndroidMobile() {
	println("Android size is", a.size)
}

func main() {
	app := &apple{
		color: "red",
	}
	clt := &client{}
	clt.ChargeMobile(app)

	// adding the new functionality to existing
	android := &android{
		size: 10,
	}

	androidAdapter := &AndroidAdapter{
		android: android,
	}
	clt.ChargeMobile(androidAdapter)

}
