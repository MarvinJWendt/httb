package service

import (
	"net/http"
	"strings"

	"github.com/marvinjwendt/httb/internal/pkg/api"
)

func (s Service) GetPing(w http.ResponseWriter, r *http.Request, params api.GetPingParams) {
	if ok := s.Validate(w, &params); !ok {
		return
	}

	sendFormattedResponse(w, r, "pong", "message")
}

// getRealIP extracts the real IP address from the request, respecting reverse proxy headers
func getRealIP(r *http.Request) string {
	// Check X-Forwarded-For header (most common)
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ips := strings.Split(forwardedFor, ",")
		if len(ips) > 0 {
			ip := strings.TrimSpace(ips[0])
			if ip != "" {
				return ip
			}
		}
	}

	// Check X-Real-IP header
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	// Check X-Client-IP header
	if clientIP := r.Header.Get("X-Client-IP"); clientIP != "" {
		return clientIP
	}

	// Check CF-Connecting-IP header (Cloudflare)
	if cfIP := r.Header.Get("CF-Connecting-IP"); cfIP != "" {
		return cfIP
	}

	// Check True-Client-IP header (Akamai and Cloudflare)
	if trueClientIP := r.Header.Get("True-Client-IP"); trueClientIP != "" {
		return trueClientIP
	}

	// Fall back to RemoteAddr
	return r.RemoteAddr
}

func (s Service) GetIp(w http.ResponseWriter, r *http.Request, params api.GetIpParams) {
	if ok := s.Validate(w, &params); !ok {
		return
	}

	ip := getRealIP(r)
	sendFormattedResponse(w, r, ip, "ip")
}
