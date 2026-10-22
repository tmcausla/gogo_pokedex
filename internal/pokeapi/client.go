package pokeapi
import (
	"net/http"
	"github.com/tmcausla/gogo_pokedex/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache 	   pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	newCache := pokecache.NewCache(interval)
	return Client{
		httpClient: http.Client{},
		cache: newCache,
	}
}

