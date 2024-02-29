package once

import (
	"image"
	"strings"
	"sync"
)

var mu sync.RWMutex // guards icons
var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func loadIcon(s string) image.Image {
	var img image.RGBA
	switch strings.TrimSpace(s) {
	case "spades.png":
	case "hearts.png":
	case "diamonds.png":
	case "clubs.png":
	}
	return &img
}

// Icon Concurrency-safe.
func Icon(name string) image.Image {
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()

	// acquire an exclusive lock
	mu.Lock()
	if icons == nil { // NOTE: must recheck for nil
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}
