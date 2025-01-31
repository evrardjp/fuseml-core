// Code generated by goa v3.3.1, DO NOT EDIT.
//
// runnable HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/fuseml/fuseml-core/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	runnable "github.com/fuseml/fuseml-core/gen/runnable"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "runnable" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListRunnablePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("runnable", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the runnable list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*runnable.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("runnable", "list", "*runnable.ListPayload", v)
		}
		values := req.URL.Query()
		if p.Kind != nil {
			values.Add("kind", *p.Kind)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the runnable
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//	- "NotFound" (type *goa.ServiceError): http.StatusNotFound
//	- error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("runnable", "list", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateRunnableResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("runnable", "list", err)
			}
			res := NewListRunnableOK(body)
			return res, nil
		case http.StatusNotFound:
			var (
				body ListNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("runnable", "list", err)
			}
			err = ValidateListNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("runnable", "list", err)
			}
			return nil, NewListNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("runnable", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildRegisterRequest instantiates a HTTP request object with method and path
// set to call the "runnable" service "register" endpoint
func (c *Client) BuildRegisterRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RegisterRunnablePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("runnable", "register", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRegisterRequest returns an encoder for requests sent to the runnable
// register server.
func EncodeRegisterRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*runnable.Runnable)
		if !ok {
			return goahttp.ErrInvalidType("runnable", "register", "*runnable.Runnable", v)
		}
		body := NewRegisterRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("runnable", "register", err)
		}
		return nil
	}
}

// DecodeRegisterResponse returns a decoder for responses returned by the
// runnable register endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRegisterResponse may return the following errors:
//	- "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeRegisterResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body RegisterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("runnable", "register", err)
			}
			err = ValidateRegisterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("runnable", "register", err)
			}
			res := NewRegisterRunnableCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body RegisterBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("runnable", "register", err)
			}
			err = ValidateRegisterBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("runnable", "register", err)
			}
			return nil, NewRegisterBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("runnable", "register", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "runnable" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		runnableNameOrID string
	)
	{
		p, ok := v.(*runnable.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("runnable", "get", "*runnable.GetPayload", v)
		}
		runnableNameOrID = p.RunnableNameOrID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetRunnablePath(runnableNameOrID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("runnable", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetResponse returns a decoder for responses returned by the runnable
// get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetResponse may return the following errors:
//	- "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//	- "NotFound" (type *goa.ServiceError): http.StatusNotFound
//	- error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("runnable", "get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("runnable", "get", err)
			}
			res := NewGetRunnableOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("runnable", "get", err)
			}
			err = ValidateGetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("runnable", "get", err)
			}
			return nil, NewGetBadRequest(&body)
		case http.StatusNotFound:
			var (
				body GetNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("runnable", "get", err)
			}
			err = ValidateGetNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("runnable", "get", err)
			}
			return nil, NewGetNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("runnable", "get", resp.StatusCode, string(body))
		}
	}
}

// unmarshalRunnableResponseToRunnableRunnable builds a value of type
// *runnable.Runnable from a value of type *RunnableResponse.
func unmarshalRunnableResponseToRunnableRunnable(v *RunnableResponse) *runnable.Runnable {
	res := &runnable.Runnable{
		ID:      v.ID,
		Name:    *v.Name,
		Kind:    *v.Kind,
		Created: v.Created,
	}
	res.Image = unmarshalRunnableImageResponseToRunnableRunnableImage(v.Image)
	res.Inputs = make([]*runnable.RunnableInput, len(v.Inputs))
	for i, val := range v.Inputs {
		res.Inputs[i] = unmarshalRunnableInputResponseToRunnableRunnableInput(val)
	}
	res.Outputs = make([]*runnable.RunnableOutput, len(v.Outputs))
	for i, val := range v.Outputs {
		res.Outputs[i] = unmarshalRunnableOutputResponseToRunnableRunnableOutput(val)
	}
	res.Labels = make([]string, len(v.Labels))
	for i, val := range v.Labels {
		res.Labels[i] = val
	}

	return res
}

// unmarshalRunnableImageResponseToRunnableRunnableImage builds a value of type
// *runnable.RunnableImage from a value of type *RunnableImageResponse.
func unmarshalRunnableImageResponseToRunnableRunnableImage(v *RunnableImageResponse) *runnable.RunnableImage {
	res := &runnable.RunnableImage{
		RegistryURL: v.RegistryURL,
		Repository:  v.Repository,
		Tag:         v.Tag,
	}

	return res
}

// unmarshalRunnableInputResponseToRunnableRunnableInput builds a value of type
// *runnable.RunnableInput from a value of type *RunnableInputResponse.
func unmarshalRunnableInputResponseToRunnableRunnableInput(v *RunnableInputResponse) *runnable.RunnableInput {
	res := &runnable.RunnableInput{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Runnable != nil {
		res.Runnable = unmarshalRunnableRefResponseToRunnableRunnableRef(v.Runnable)
	}
	if v.Parameter != nil {
		res.Parameter = unmarshalInputParameterResponseToRunnableInputParameter(v.Parameter)
	}

	return res
}

// unmarshalRunnableRefResponseToRunnableRunnableRef builds a value of type
// *runnable.RunnableRef from a value of type *RunnableRefResponse.
func unmarshalRunnableRefResponseToRunnableRunnableRef(v *RunnableRefResponse) *runnable.RunnableRef {
	if v == nil {
		return nil
	}
	res := &runnable.RunnableRef{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Labels != nil {
		res.Labels = make([]string, len(v.Labels))
		for i, val := range v.Labels {
			res.Labels[i] = val
		}
	}

	return res
}

// unmarshalInputParameterResponseToRunnableInputParameter builds a value of
// type *runnable.InputParameter from a value of type *InputParameterResponse.
func unmarshalInputParameterResponseToRunnableInputParameter(v *InputParameterResponse) *runnable.InputParameter {
	if v == nil {
		return nil
	}
	res := &runnable.InputParameter{
		Datatype: v.Datatype,
		Optional: v.Optional,
		Default:  v.Default,
	}

	return res
}

// unmarshalRunnableOutputResponseToRunnableRunnableOutput builds a value of
// type *runnable.RunnableOutput from a value of type *RunnableOutputResponse.
func unmarshalRunnableOutputResponseToRunnableRunnableOutput(v *RunnableOutputResponse) *runnable.RunnableOutput {
	res := &runnable.RunnableOutput{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Runnable != nil {
		res.Runnable = unmarshalRunnableRefResponseToRunnableRunnableRef(v.Runnable)
	}
	if v.Metadata != nil {
		res.Metadata = unmarshalInputParameterResponseToRunnableInputParameter(v.Metadata)
	}

	return res
}

// marshalRunnableRunnableImageToRunnableImageRequestBody builds a value of
// type *RunnableImageRequestBody from a value of type *runnable.RunnableImage.
func marshalRunnableRunnableImageToRunnableImageRequestBody(v *runnable.RunnableImage) *RunnableImageRequestBody {
	res := &RunnableImageRequestBody{
		RegistryURL: v.RegistryURL,
		Repository:  v.Repository,
		Tag:         v.Tag,
	}

	return res
}

// marshalRunnableRunnableInputToRunnableInputRequestBody builds a value of
// type *RunnableInputRequestBody from a value of type *runnable.RunnableInput.
func marshalRunnableRunnableInputToRunnableInputRequestBody(v *runnable.RunnableInput) *RunnableInputRequestBody {
	res := &RunnableInputRequestBody{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Runnable != nil {
		res.Runnable = marshalRunnableRunnableRefToRunnableRefRequestBody(v.Runnable)
	}
	if v.Parameter != nil {
		res.Parameter = marshalRunnableInputParameterToInputParameterRequestBody(v.Parameter)
	}

	return res
}

// marshalRunnableRunnableRefToRunnableRefRequestBody builds a value of type
// *RunnableRefRequestBody from a value of type *runnable.RunnableRef.
func marshalRunnableRunnableRefToRunnableRefRequestBody(v *runnable.RunnableRef) *RunnableRefRequestBody {
	if v == nil {
		return nil
	}
	res := &RunnableRefRequestBody{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Labels != nil {
		res.Labels = make([]string, len(v.Labels))
		for i, val := range v.Labels {
			res.Labels[i] = val
		}
	}

	return res
}

// marshalRunnableInputParameterToInputParameterRequestBody builds a value of
// type *InputParameterRequestBody from a value of type
// *runnable.InputParameter.
func marshalRunnableInputParameterToInputParameterRequestBody(v *runnable.InputParameter) *InputParameterRequestBody {
	if v == nil {
		return nil
	}
	res := &InputParameterRequestBody{
		Datatype: v.Datatype,
		Optional: v.Optional,
		Default:  v.Default,
	}

	return res
}

// marshalRunnableRunnableOutputToRunnableOutputRequestBody builds a value of
// type *RunnableOutputRequestBody from a value of type
// *runnable.RunnableOutput.
func marshalRunnableRunnableOutputToRunnableOutputRequestBody(v *runnable.RunnableOutput) *RunnableOutputRequestBody {
	res := &RunnableOutputRequestBody{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Runnable != nil {
		res.Runnable = marshalRunnableRunnableRefToRunnableRefRequestBody(v.Runnable)
	}
	if v.Metadata != nil {
		res.Metadata = marshalRunnableInputParameterToInputParameterRequestBody(v.Metadata)
	}

	return res
}

// marshalRunnableImageRequestBodyToRunnableRunnableImage builds a value of
// type *runnable.RunnableImage from a value of type *RunnableImageRequestBody.
func marshalRunnableImageRequestBodyToRunnableRunnableImage(v *RunnableImageRequestBody) *runnable.RunnableImage {
	res := &runnable.RunnableImage{
		RegistryURL: v.RegistryURL,
		Repository:  v.Repository,
		Tag:         v.Tag,
	}

	return res
}

// marshalRunnableInputRequestBodyToRunnableRunnableInput builds a value of
// type *runnable.RunnableInput from a value of type *RunnableInputRequestBody.
func marshalRunnableInputRequestBodyToRunnableRunnableInput(v *RunnableInputRequestBody) *runnable.RunnableInput {
	res := &runnable.RunnableInput{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Runnable != nil {
		res.Runnable = marshalRunnableRefRequestBodyToRunnableRunnableRef(v.Runnable)
	}
	if v.Parameter != nil {
		res.Parameter = marshalInputParameterRequestBodyToRunnableInputParameter(v.Parameter)
	}

	return res
}

// marshalRunnableRefRequestBodyToRunnableRunnableRef builds a value of type
// *runnable.RunnableRef from a value of type *RunnableRefRequestBody.
func marshalRunnableRefRequestBodyToRunnableRunnableRef(v *RunnableRefRequestBody) *runnable.RunnableRef {
	if v == nil {
		return nil
	}
	res := &runnable.RunnableRef{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Labels != nil {
		res.Labels = make([]string, len(v.Labels))
		for i, val := range v.Labels {
			res.Labels[i] = val
		}
	}

	return res
}

// marshalInputParameterRequestBodyToRunnableInputParameter builds a value of
// type *runnable.InputParameter from a value of type
// *InputParameterRequestBody.
func marshalInputParameterRequestBodyToRunnableInputParameter(v *InputParameterRequestBody) *runnable.InputParameter {
	if v == nil {
		return nil
	}
	res := &runnable.InputParameter{
		Datatype: v.Datatype,
		Optional: v.Optional,
		Default:  v.Default,
	}

	return res
}

// marshalRunnableOutputRequestBodyToRunnableRunnableOutput builds a value of
// type *runnable.RunnableOutput from a value of type
// *RunnableOutputRequestBody.
func marshalRunnableOutputRequestBodyToRunnableRunnableOutput(v *RunnableOutputRequestBody) *runnable.RunnableOutput {
	res := &runnable.RunnableOutput{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Runnable != nil {
		res.Runnable = marshalRunnableRefRequestBodyToRunnableRunnableRef(v.Runnable)
	}
	if v.Metadata != nil {
		res.Metadata = marshalInputParameterRequestBodyToRunnableInputParameter(v.Metadata)
	}

	return res
}

// unmarshalRunnableImageResponseBodyToRunnableRunnableImage builds a value of
// type *runnable.RunnableImage from a value of type *RunnableImageResponseBody.
func unmarshalRunnableImageResponseBodyToRunnableRunnableImage(v *RunnableImageResponseBody) *runnable.RunnableImage {
	res := &runnable.RunnableImage{
		RegistryURL: v.RegistryURL,
		Repository:  v.Repository,
		Tag:         v.Tag,
	}

	return res
}

// unmarshalRunnableInputResponseBodyToRunnableRunnableInput builds a value of
// type *runnable.RunnableInput from a value of type *RunnableInputResponseBody.
func unmarshalRunnableInputResponseBodyToRunnableRunnableInput(v *RunnableInputResponseBody) *runnable.RunnableInput {
	res := &runnable.RunnableInput{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Runnable != nil {
		res.Runnable = unmarshalRunnableRefResponseBodyToRunnableRunnableRef(v.Runnable)
	}
	if v.Parameter != nil {
		res.Parameter = unmarshalInputParameterResponseBodyToRunnableInputParameter(v.Parameter)
	}

	return res
}

// unmarshalRunnableRefResponseBodyToRunnableRunnableRef builds a value of type
// *runnable.RunnableRef from a value of type *RunnableRefResponseBody.
func unmarshalRunnableRefResponseBodyToRunnableRunnableRef(v *RunnableRefResponseBody) *runnable.RunnableRef {
	if v == nil {
		return nil
	}
	res := &runnable.RunnableRef{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Labels != nil {
		res.Labels = make([]string, len(v.Labels))
		for i, val := range v.Labels {
			res.Labels[i] = val
		}
	}

	return res
}

// unmarshalInputParameterResponseBodyToRunnableInputParameter builds a value
// of type *runnable.InputParameter from a value of type
// *InputParameterResponseBody.
func unmarshalInputParameterResponseBodyToRunnableInputParameter(v *InputParameterResponseBody) *runnable.InputParameter {
	if v == nil {
		return nil
	}
	res := &runnable.InputParameter{
		Datatype: v.Datatype,
		Optional: v.Optional,
		Default:  v.Default,
	}

	return res
}

// unmarshalRunnableOutputResponseBodyToRunnableRunnableOutput builds a value
// of type *runnable.RunnableOutput from a value of type
// *RunnableOutputResponseBody.
func unmarshalRunnableOutputResponseBodyToRunnableRunnableOutput(v *RunnableOutputResponseBody) *runnable.RunnableOutput {
	res := &runnable.RunnableOutput{
		Name: v.Name,
		Kind: v.Kind,
	}
	if v.Runnable != nil {
		res.Runnable = unmarshalRunnableRefResponseBodyToRunnableRunnableRef(v.Runnable)
	}
	if v.Metadata != nil {
		res.Metadata = unmarshalInputParameterResponseBodyToRunnableInputParameter(v.Metadata)
	}

	return res
}
