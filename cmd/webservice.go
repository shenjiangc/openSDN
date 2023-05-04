package main

import (
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
)

func regionWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/regions").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"Region"}

	ws.Route(ws.GET("/").To(findAllRegions).
		// docs
		Doc("get all regions").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]Region{}).
		Returns(200, "OK", []Region{}))

	ws.Route(ws.GET("/{region-id}").To(findRegion).
		// docs
		Doc("get a region").
		Param(ws.PathParameter("region-id", "identifier of the region").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(Region{}). // on the response
		Returns(200, "OK", Region{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{region-id}").To(updateRegion).
		// docs
		Doc("update a region").
		Param(ws.PathParameter("region-id", "identifier of the region").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(Region{})) // from the request

	ws.Route(ws.PUT("").To(createRegion).
		// docs
		Doc("create a region").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(Region{})) // from the request

	ws.Route(ws.DELETE("/{region-id}").To(removeRegion).
		// docs
		Doc("delete a region").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("region-id", "identifier of the region").DataType("string")))

	return ws
}

// GET http://localhost:8080/regions
func findAllRegions(request *restful.Request, response *restful.Response) {
	list := []Region{}
	for _, each := range regions {
		list = append(list, each)
	}
	response.WriteEntity(list)
}

// GET http://localhost:8080/regions/1
func findRegion(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("region-id")
	region := regions[id]
	if len(region.ID) == 0 {
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(region)
	}
}

// PUT http://localhost:8080/regions/1
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
func updateRegion(request *restful.Request, response *restful.Response) {
	region := new(Region)
	err := request.ReadEntity(&region)
	if err == nil {
		regions[region.ID] = *region
		response.WriteEntity(region)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// PUT http://localhost:8080/regions/1
// <User><Id>1</Id><Name>Melissa</Name></User>
func createRegion(request *restful.Request, response *restful.Response) {
	region := Region{ID: request.PathParameter("region-id")}
	err := request.ReadEntity(&region)
	if err == nil {
		regions[region.ID] = region
		response.WriteHeaderAndEntity(http.StatusCreated, region)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// DELETE http://localhost:8080/regions/1
func removeRegion(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("region-id")
	delete(regions, id)
}
