package usagiapi

import (
	"github.com/RabbitHouseCorp/usagi-api/rest"
) 
 


// Endpoint: Dance
func Dance() string {
	return rest.Api("/api/dance").Dance
}

// Endpoint: Feed
func Feed() string {
	return rest.Api("/api/feed").Feed
}

// Endpoint: Hug
func Hug() string {
	return rest.Api("/api/hug").Hug
}

// Endpoint: Kiss
func Kiss() string {
	return rest.Api("/api/kiss").Kiss
}

// Endpoint: Pat
func Pat() string {
	return rest.Api("/api/pat").Pat
}

// Endpoint: Poke
func Poke() string {
	return rest.Api("/api/poke").Poke
}

// Endpoint: Slap
func Slap() string {
	return rest.Api("/api/slap").Slap
}

// Endpoint: Tickle
func Tickle() string {
	return rest.Api("/api/tickle").Tickle
}