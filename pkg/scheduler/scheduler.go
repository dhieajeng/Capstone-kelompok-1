package scheduler

import (
	"github.com/google/uuid"
	"log"

	"github.com/bloomingbug/depublic/configs"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type scheduler struct {
	rdb *redis.Pool
	cfg configs.NamespaceConfig
}

func (s *scheduler) SendOTP(email, otp string) {
	var enqueuer = work.NewEnqueuer(s.cfg.Namespace, s.rdb)

	_, err := enqueuer.Enqueue("send_otp", work.Q{"email_address": email, "otp_code": otp})
	if err != nil {
		log.Fatal(err)
	}
}

func (s *scheduler) SendToken(email, link string, token uuid.UUID) {
	var enqueuer = work.NewEnqueuer(s.cfg.Namespace, s.rdb)

	_, err := enqueuer.Enqueue("send_token", work.Q{"email_address": email, "token": token, "link": link})
	if err != nil {
		log.Fatal(err)
	}
}

type Scheduler interface {
	SendOTP(email, otp string)
	SendToken(email, link string, token uuid.UUID)
}

func NewScheduler(rdb *redis.Pool, cfg configs.NamespaceConfig) Scheduler {
	return &scheduler{
		rdb: rdb,
		cfg: cfg,
	}
}
