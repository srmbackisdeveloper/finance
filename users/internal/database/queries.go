package database

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (s *service) Health() string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	sqlDB, err := s.db.DB()
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to get db connection: %v", err))
	}

	err = sqlDB.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return "postgres is healthy"
}
