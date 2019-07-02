/*
Copyright 2017 - The TXTDirect Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package txtdirect

import (
	"net/http"
	"sync"

	"github.com/mholt/caddy/caddy/caddymain"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}

var torOnce sync.Once

var allOptions = []string{"host", "path", "gometa", "www"}

func removeArrayFromArray(array, toBeRemoved []string) []string {
	tmp := make([]string, len(array))
	copy(tmp, array)
	for _, toRemove := range toBeRemoved {
		for i, option := range tmp {
			if option == toRemove {
				tmp[i] = tmp[len(tmp)-1]
				tmp = tmp[:len(tmp)-1]
				break
			}
		}
	}
	return tmp
}

// Redirect is middleware to redirect requests based on TXT records
type TXTDirect struct {
	Next   httpserver.Handler
	Config Config
}

func (rd TXTDirect) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	if err := Redirect(w, r, rd.Config); err != nil {
		if err.Error() == "option disabled" {
			return rd.Next.ServeHTTP(w, r)
		}
		return http.StatusInternalServerError, err
	}

	// Count total redirects if prometheus is enabled
	if w.Header().Get("Status-Code") == "301" || w.Header().Get("Status-Code") == "302" {
		if rd.Config.Prometheus.Enable {
			RequestsCount.WithLabelValues(r.Host).Add(1)
		}
	}

	return 0, nil
}
