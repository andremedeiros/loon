package service

type Redis struct{}

func (r *Redis) String() string {
	return "Redis"
}

func (r *Redis) Environ() []string {
	return []string{"REDIS_URL=redis://localhost:6379"}
}

func (r *Redis) Start() []string {
	return []string{
		"redis-server",
		"--dir /tmp",
		"--port 6379",
		"--pidfile /tmp/redis.pid",
	}
}

func (r *Redis) Stop() error {
	return nil
}
