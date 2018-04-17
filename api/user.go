package api

import (
	"nirvana-cms-auth/controllers/users"

	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/operators/validator"
)

var User = definition.Descriptor{
	Path:        "/user",
	Description: "user api",
	Children: []definition.Descriptor{
		{
			Path:        "/login",
			Description: "login api",
			Definitions: []definition.Definition{
				{
					Method:   definition.Create,
					Function: users.Login,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Form,
							Name:        "account",
							Description: "user account",
							Operators:   []definition.Operator{validator.String("min=5"), validator.String("max=16")},
						},
						{
							Source:      definition.Form,
							Name:        "password",
							Description: "password",
							Operators:   []definition.Operator{validator.String("min=6"), validator.String("max=20")},
						},
					},
					Results: []definition.Result{
						{
							Destination: definition.Meta,
							Description: "session id",
						},
						{
							Destination: definition.Error,
							Description: "error info",
						},
					},
				},
			},
		},
	},
}
