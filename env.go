package main

type (
	Firebase struct {
		ProjectID     string `env:"PROJECT_ID,required"`
		Configuration string `env:"CONFIGURATION_FILE,required"`
	}
	Environment struct {
		Database string `env:"DATABASE_FILE,required"`
		Secret   string `env:"SECRET,required"`
		Firebase `env:",prefix=FIREBASE_"`
	}
)
