package resources

import ()

var DefaultFields []Field = []Field{
	{
		Name:     "id",
		DataType: "uid",
		Constraints: []Constraint{
			{
				Value: "not null",
				Type:  "sql",
			},
		},
	}, {
		Name:     "created",
		DataType: "unix",
		Constraints: []Constraint{
			{
				Value: "not null",
				Type:  "sql",
			},
		},
	}, {
		Name:     "updated",
		DataType: "unix",
		Constraints: []Constraint{
			{
				Value: "not null",
				Type:  "sql",
			},
		},
	}, {
		Name:     "archived",
		DataType: "unix",
	},
}
