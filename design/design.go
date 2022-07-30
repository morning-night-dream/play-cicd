//go:generate goa gen github.com/morning-night-dream/play-cicd/design --output ../
package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("todo", func() {
	Title("Todo Service")
	Description("Service for todo manage")
	Server("todo", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("todo", func() {
	Description("The todo service.")

	Method("ReadTodoList", func() {
		Result(ReadTodoListResponseBody)

		HTTP(func() {
			GET("/todos")
			Response(StatusOK)
		})
	})

	Method("CreateTodo", func() {
		Payload(CreateTodoRequestBody)

		Result(CreateTodoResponseBody)

		HTTP(func() {
			POST("/todos")
			Response(StatusOK)
		})
	})

	Method("UpdateTodo", func() {
		Payload(func() {
			Field(1, "id", func() {
				Pattern(FormatUUID)
			})
			Attribute("task", String, func() {
				Example("Golang coding")
			})
			Required("id", "task")
		})

		Result(UpdateTodoResponseBody)

		HTTP(func() {
			PUT("/todos/{id}")
			Response(StatusOK)
		})
	})

	Method("DeleteTodo", func() {
		Payload(func() {
			Field(1, "id", func() {
				Pattern(FormatUUID)
			})
			Required("id")
		})

		HTTP(func() {
			DELETE("/todos/{id}")
			Response(StatusOK)
		})
	})
})

var Todo = Type("Todo", func() {
	Attribute("id", String, func() {
		Format(FormatUUID)
		Example("ac65bc46-3854-455f-82b9-54b6ec733b53")
	})
	Attribute("task", String, func() {
		Example("Golang coding")
	})
	Attribute("createdAt", String, func() {
		Format(FormatDateTime)
		Example("2006-01-02T15:04:05Z")
	})
	Attribute("updatedAt", String, func() {
		Format(FormatDateTime)
		Example("2006-01-02T15:04:05Z")
	})
})

var ReadTodoListResponseBody = Type("ReadTodoListResponseBody", func() {
	Attribute("todos", ArrayOf(Todo), "todos")
})

var CreateTodoRequestBody = Type("CreateTodoRequestBody", func() {
	Attribute("task", String, func() {
		Example("Golang coding")
	})
})

var CreateTodoResponseBody = Type("CreateTodoResponseBody", func() {
	Attribute("task", Todo, func() {})
})

var UpdateTodoResponseBody = Type("UpdateTodoResponseBody", func() {
	Attribute("task", Todo, func() {})
})
