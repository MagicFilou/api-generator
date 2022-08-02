package resources

var DefaultFields []Field = []Field{
	{
		Name:     "id",
		DataType: "serial",
	}, {
		Name:     "created",
		DataType: "unix",
		Constraints: []Constraint{
			{
				Value: "not null",
				Type:  "sql",
			},
		},
	},
}
