package service

import (
	"crypto/sha256"
	"deathlog-tracker/domain/entity"
	"fmt"
)

func Hash(player entity.Player) string {
	h := sha256.New()
	fmt.Fprintf(h, "%s%d", player.Name, player.Timestamp)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
