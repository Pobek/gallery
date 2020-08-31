package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/pobek/gallery/server/api/controllers"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		log.Printf("Could not parse Jaeger env vars: %s", err.Error())
		return
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)
	testSpan := tracer.StartSpan("test")
	defer testSpan.Finish()

	app := controllers.App{}
	app.Init(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		tracer,
	)

	app.RunServer()
}
