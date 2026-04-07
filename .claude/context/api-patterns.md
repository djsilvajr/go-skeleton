# API Response Pattern

## Responses 

All responses must follow:

- data

Example:

{
	"data": {
		"id": 1,
		"name": "Admin",
		"email": "admin@example.com",
		"role": "admin",
		"CreatedAt": "2026-04-06T10:26:43.536Z",
		"UpdatedAt": "2026-04-06T10:26:43.536Z",
		"DeletedAt": null
	}
}

## Delete

- Aways write Deletes as SoftDeletes unless I specify in task.


