// 📦 Usagi API
//
//
// 🔗 Github Package: https://github.com/RabbitHouseCorp/usagi-api
package usagiapi

import "github.com/RabbitHouseCorp/usagi-api/v1/usagirest"

// Endpoint: Dance
func Dance() string {
	return usagirest.Api("/api/dance").Url
}

// Endpoint: Feed
func Feed() string {
	return usagirest.Api("/api/feed").Url
}

// Endpoint: Hug
func Hug() string {
	return usagirest.Api("/api/hug").Url
}

// Endpoint: Kiss
func Kiss() string {
	return usagirest.Api("/api/kiss").Url
}

// Endpoint: Pat
func Pat() string {
	return usagirest.Api("/api/pat").Url
}

// Endpoint: Poke
func Poke() string {
	return usagirest.Api("/api/poke").Url
}

// Endpoint: Slap
func Slap() string {
	return usagirest.Api("/api/slap").Url
}

// Endpoint: Tickle
func Tickle() string {
	return usagirest.Api("/api/tickle").Url
}