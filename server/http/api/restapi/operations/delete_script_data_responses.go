// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteScriptDataNoContentCode is the HTTP code returned for type DeleteScriptDataNoContent
const DeleteScriptDataNoContentCode int = 204

/*DeleteScriptDataNoContent delete script data no content

swagger:response deleteScriptDataNoContent
*/
type DeleteScriptDataNoContent struct {
}

// NewDeleteScriptDataNoContent creates DeleteScriptDataNoContent with default headers values
func NewDeleteScriptDataNoContent() *DeleteScriptDataNoContent {

	return &DeleteScriptDataNoContent{}
}

// WriteResponse to the client
func (o *DeleteScriptDataNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}