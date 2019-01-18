package global

import "github.com/go-redis/redis"

func InitRedis() {
	r := Cfg.Section("redis")
	host := r.Key("HOST").MustString("127.0.0.1")
	port := r.Key("PORT").MustString("6379")
	pass := r.Key("PASS").MustString("")
	db := r.Key("DB").MustInt(0)
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       db,
	})
	err := RedisConn.Ping().Err()
	if err != nil {
		panic("redis连接失败: " + err.Error())
	}
}
