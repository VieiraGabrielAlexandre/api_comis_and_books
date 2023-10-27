package documentation

import (
	"net/http"
)

// @title           Comics API
// @version         1.0
// @description    	This comics API is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host
// @BasePath  /api/

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/// ErrorResponse represents the error response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// Create godoc
// @Summary Create a new comic
// @Description Create a new comic entry
// @Accept json
// @Produce json
// @Param comic body Comics true "Comic object"
// @Success 201 {string} string "Comic created successfully"
// @Router /comics [post]
func Create(w http.ResponseWriter, r *http.Request) {
	// Your existing code here
}

// @Summary Get a comic by ID
// @Description Get a comic by its ID
// @Produce json
// @Param id path int true "Comic ID"
// @Success 200 {object} Comics "Comic object"
// @Failure 500 {object} ErrorResponse "Error response"
// @Router /comics/{id} [get]
func GetByID(w http.ResponseWriter, r *http.Request) {
	// Your existing code here
}

// @Summary Update a comic by ID
// @Description Update a comic by its ID
// @Accept json
// @Produce json
// @Param id path int true "Comic ID"
// @Param comic body Comics true "Updated comic object"
// @Success 200 {string} string "Comic updated successfully"
// @Router /comics/{id} [put]
func Update(w http.ResponseWriter, r *http.Request) {
	// Your existing code here
}

// @Summary Get comics with pagination
// @Description Get a list of comics with pagination
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Results per page"
// @Param orderBy query string false "Order by"
// @Param filter query string false "Filter"
// @Param column query string false "Column"
// @Param order query string false "Order"
// @Success 200 {object} Comics "Comics response"
// @Failure 500 {object} ErrorResponse "Error response"
// @Router /comics [get]
func GetAllPaginate(w http.ResponseWriter, r *http.Request) {
	// Your existing code here
}

// @Summary Delete a comic by ID
// @Description Delete a comic by its ID
// @Produce json
// @Param id path int true "Comic ID"
// @Success 200 {string} string "Comic deleted successfully"
// @Router /comics/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	// Your existing code here
}

type Comics struct {
	Title     string `json:"title"`
	Volume    int    `json:"volume"`
	Publisher string `json:"publisher"`
	Editor    string `json:"editor"`
}
