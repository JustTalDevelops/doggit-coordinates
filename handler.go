package doggit_coordinates

import (
	"github.com/JustTalDevelops/doggit"
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/player/title"
	"github.com/go-gl/mathgl/mgl64"
	"strconv"
)

type DoggitHandler struct {
	doggit.NopHandler
}

func (d DoggitHandler) HandleMove(p *player.Player, _ *event.Context, pos mgl64.Vec3, _ float64, _ float64) {
	p.SendTitle(title.New().WithActionText("X: " + strconv.Itoa(int(pos.X())) + " Y: " + strconv.Itoa(int(pos.Y())) + " Z: " + strconv.Itoa(int(pos.Z()))))
}