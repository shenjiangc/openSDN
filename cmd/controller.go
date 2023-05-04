package main

import (
	"fmt"
	"log"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	//init command
	initCmd()

	//init configure
	initConfig()

	//init logger
	err := InitLogger(logFile, logLevel, logMaxSize, logMaxBackups, logMaxAge)
	if err != nil {
		fmt.Println(err)
	}
	defer SyncLog() //flush log

	//init cql
	err = initCql()
	if err != nil {
		vlog.Fatalf("init CQL failed (%v)", err)
	}
	defer releaseCql()

	//log example
	vlog.Infof("start opensdn......")
	zlog.Info("abc", zap.String("t1", "t1value"), zap.Int("t2", 2))

	validate = validator.New()
	validateVariable()

	//start web serve
	startRestful()
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:   "opensdn service",
			Version: "1.0.0",
		},
	}
}

func startRestful() {
	restful.Add(regionWebService())

	//swagger
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.Add(restfulspec.NewOpenAPIService(config))

	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("./swagger-ui/dist"))))

	log.Printf("Get the API using http://localhost:8080/apidocs.json")
	log.Printf("Open Swagger UI using http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
