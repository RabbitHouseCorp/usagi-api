package usagiapi

import (
	"github.com/RabbitHouseCorp/usagi-api/rest"
) 
 


// Endpoint: Dance
func Dance() string {
	return rest.Api("/api/dance").Url
}

// Endpoint: Feed
func Feed() string {
	return rest.Api("/api/feed").Url
}

// Endpoint: Hug
func Hug() string {
	return rest.Api("/api/hug").Url
}

// Endpoint: Kiss
func Kiss() string {
	return rest.Api("/api/kiss").Url
}

// Endpoint: Pat
func Pat() string {
	return rest.Api("/api/pat").Url
}

// Endpoint: Poke
func Poke() string {
	return rest.Api("/api/poke").Url
}

// Endpoint: Slap
func Slap() string {
	return rest.Api("/api/slap").Url
}

// Endpoint: Tickle
func Tickle() string {
	return rest.Api("/api/tickle").Url
}