package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func LogrusMiddleware(log *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Add request ID
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
			c.Set("X-Request-ID", requestID)
		}

		if log.IsLevelEnabled(logrus.DebugLevel) {
			log.WithFields(logrus.Fields{
				"request_id": requestID,
				"method":     c.Method(),
				"path":       c.Path(),
				"headers":    c.GetReqHeaders(),
				"body":       string(c.Body()),
			}).Debug("Incoming request")
		}

		err := c.Next()

		fields := logrus.Fields{
			"request_id": requestID,
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     c.Response().StatusCode(),
			"latency":    time.Since(start).String(),
			"ip":         c.IP(),
			"user_agent": c.Get("User-Agent"),
		}

		if err != nil {
			log.WithFields(fields).WithError(err).Error("Request failed")
		} else {
			if log.IsLevelEnabled(logrus.DebugLevel) {
				fields["response_body"] = string(c.Response().Body())
				log.WithFields(fields).Debug("Request completed with response")
			} else {
				log.WithFields(fields).Info("Request completed")
			}
		}

		return err
	}
}
