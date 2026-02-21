package main

import (
	"CleanArchitecture/internal/domain/model"
	"CleanArchitecture/internal/handler"
	postgres "CleanArchitecture/internal/infrastructure/postgress"
	redis "CleanArchitecture/internal/infrastructure/redis"
	"CleanArchitecture/internal/router"
	"CleanArchitecture/internal/usecase/auth"
	"CleanArchitecture/internal/utils/jwt"
	"log"

	_ "CleanArchitecture/docs"
)

func main() {
	db := postgres.NewPostgresDB()        //   بالا اوردن  پوست گرس
	redisClient := redis.NewRedisClient() //  بالا اوردن ردیس

	err := jwt.InitJWTKeys(
		"internal/utils/jwt/keys/private.pem",
		"internal/utils/jwt/keys/public.pem",
	)
	if err != nil {
		log.Fatal("cannot init keys: ", err)
	}

	userRepo := postgres.NewUserPostgresRepository(db)    //  کار های دیتا  بیسی با کاربر
	otpRepo := redis.NewOtpRedisRepository(redisClient)   // ذخیره و گرفتن او تی پی
	ratelimiter := redis.NewRateLimiterRedis(redisClient) //  برای ریت لیمیت

	authUsecase := auth.NewAuthUsecase(otpRepo, userRepo, ratelimiter)
	authHandler := handler.NewAuthHandler(authUsecase)

	r := router.SetupRoutes(authHandler)

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
