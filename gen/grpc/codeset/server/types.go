// Code generated by goa v3.3.1, DO NOT EDIT.
//
// codeset gRPC server types
//
// Command:
// $ goa gen github.com/fuseml/fuseml-core/design

package server

import (
	codeset "github.com/fuseml/fuseml-core/gen/codeset"
	codesetpb "github.com/fuseml/fuseml-core/gen/grpc/codeset/pb"
	goa "goa.design/goa/v3/pkg"
)

// NewListPayload builds the payload of the "list" endpoint of the "codeset"
// service from the gRPC request type.
func NewListPayload(message *codesetpb.ListRequest) *codeset.ListPayload {
	v := &codeset.ListPayload{}
	if message.Project != "" {
		v.Project = &message.Project
	}
	if message.Label != "" {
		v.Label = &message.Label
	}
	return v
}

// NewListResponse builds the gRPC response type from the result of the "list"
// endpoint of the "codeset" service.
func NewListResponse(result []*codeset.Codeset) *codesetpb.ListResponse {
	message := &codesetpb.ListResponse{}
	message.Field = make([]*codesetpb.Codeset2, len(result))
	for i, val := range result {
		message.Field[i] = &codesetpb.Codeset2{
			Name:    val.Name,
			Project: val.Project,
		}
		if val.Description != nil {
			message.Field[i].Description = *val.Description
		}
		if val.Labels != nil {
			message.Field[i].Labels = make([]string, len(val.Labels))
			for j, val := range val.Labels {
				message.Field[i].Labels[j] = val
			}
		}
	}
	return message
}

// NewRegisterPayload builds the payload of the "register" endpoint of the
// "codeset" service from the gRPC request type.
func NewRegisterPayload(message *codesetpb.RegisterRequest) *codeset.RegisterPayload {
	v := &codeset.RegisterPayload{
		Location: message.Location,
	}
	if message.Codeset != nil {
		v.Codeset = protobufCodesetpbCodeset2ToCodesetCodeset(message.Codeset)
	}
	return v
}

// NewRegisterResponse builds the gRPC response type from the result of the
// "register" endpoint of the "codeset" service.
func NewRegisterResponse(result *codeset.Codeset) *codesetpb.RegisterResponse {
	message := &codesetpb.RegisterResponse{
		Name:    result.Name,
		Project: result.Project,
	}
	if result.Description != nil {
		message.Description = *result.Description
	}
	if result.Labels != nil {
		message.Labels = make([]string, len(result.Labels))
		for i, val := range result.Labels {
			message.Labels[i] = val
		}
	}
	return message
}

// NewGetPayload builds the payload of the "get" endpoint of the "codeset"
// service from the gRPC request type.
func NewGetPayload(message *codesetpb.GetRequest) *codeset.GetPayload {
	v := &codeset.GetPayload{
		Project: message.Project,
		Name:    message.Name,
	}
	return v
}

// NewGetResponse builds the gRPC response type from the result of the "get"
// endpoint of the "codeset" service.
func NewGetResponse(result *codeset.Codeset) *codesetpb.GetResponse {
	message := &codesetpb.GetResponse{
		Name:    result.Name,
		Project: result.Project,
	}
	if result.Description != nil {
		message.Description = *result.Description
	}
	if result.Labels != nil {
		message.Labels = make([]string, len(result.Labels))
		for i, val := range result.Labels {
			message.Labels[i] = val
		}
	}
	return message
}

// ValidateRegisterRequest runs the validations defined on RegisterRequest.
func ValidateRegisterRequest(message *codesetpb.RegisterRequest) (err error) {
	if message.Codeset == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("codeset", "message"))
	}
	return
}

// ValidateCodeset2 runs the validations defined on Codeset2.
func ValidateCodeset2(message *codesetpb.Codeset2) (err error) {

	return
}

// protobufCodesetpbCodeset2ToCodesetCodeset builds a value of type
// *codeset.Codeset from a value of type *codesetpb.Codeset2.
func protobufCodesetpbCodeset2ToCodesetCodeset(v *codesetpb.Codeset2) *codeset.Codeset {
	res := &codeset.Codeset{
		Name:    v.Name,
		Project: v.Project,
	}
	if v.Description != "" {
		res.Description = &v.Description
	}
	if v.Labels != nil {
		res.Labels = make([]string, len(v.Labels))
		for i, val := range v.Labels {
			res.Labels[i] = val
		}
	}

	return res
}

// svcCodesetCodesetToCodesetpbCodeset2 builds a value of type
// *codesetpb.Codeset2 from a value of type *codeset.Codeset.
func svcCodesetCodesetToCodesetpbCodeset2(v *codeset.Codeset) *codesetpb.Codeset2 {
	res := &codesetpb.Codeset2{
		Name:    v.Name,
		Project: v.Project,
	}
	if v.Description != nil {
		res.Description = *v.Description
	}
	if v.Labels != nil {
		res.Labels = make([]string, len(v.Labels))
		for i, val := range v.Labels {
			res.Labels[i] = val
		}
	}

	return res
}
