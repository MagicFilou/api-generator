package docs

import (
	"api-builder/templates"
)

const Root templates.Template = `
get:
  summary: '{{ .Name.Spaced }}: Get all'
  description: This endpoint retrieves and returns all {{ .Name.PluralUnderscored }}.
  tags:
    - {{ .Name.CamelLower }}
  responses:
    "200":
      description: Success
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                $ref: "#/components/schemas/Status" 
              length:
                $ref: "#/components/schemas/Length"
              data:
                type: array
                items:
                  allOf:
                    - $ref: "#/components/schemas/SharedModel"
                    - $ref: "../schemas/{{ .Name.CamelLower }}.yaml"
    "204":
      description: Empty
      $ref: "#/components/responses/204Empty"
    "400":
      description: Request malformed
      $ref: "#/components/responses/400BadRequest"
    "401":
      description: Unauthorized
      $ref: "#/components/responses/401Unauthorized"
    "500":
      description: Server error
      $ref: "#/components/responses/500ServerError"

post:
  summary: '{{ .Name.Spaced }}: Create'
  description: This endpoint accepts a POST body and creates one {{ .Name.SingularUnderscored }}, returning the created object.
  tags:
    - {{ .Name.CamelLower }}
  requestBody:
    description: An {{ .Name.SingularUnderscored }}
    required: true
    content:
      application/json:
        schema:
          $ref: "../schemas/{{ .Name.CamelLower }}.yaml"
  responses:
    "200":
      description: Success
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                $ref: "#/components/schemas/Status"
              data:
                  allOf:
                    - $ref: "#/components/schemas/SharedModel"
                    - $ref: "../schemas/{{ .Name.CamelLower }}.yaml"
    "400":
      description: Request malformed
      $ref: "#/components/responses/400BadRequest"
    "401":
      description: Unauthorized
      $ref: "#/components/responses/401Unauthorized"
    "500":
      description: Server error
      $ref: "#/components/responses/500ServerError"

patch:
  summary: '{{ .Name.Spaced }}: Update'
  description: This endpoint accepts a POST body and patches any relevant fields, returning the updated {{ .Name.SingularUnderscored }}.
  tags:
    - {{ .Name.CamelLower }}
  parameters:
    - in: body
      name: id
      description: A single ID of a {{ .Name.SingularUnderscored }}
      required: true
      schema:
        type: string
  requestBody:
    description: An {{ .Name.SingularUnderscored }}
    required: true
    content:
      application/json:
        schema:
          $ref: "../schemas/{{ .Name.CamelLower }}_with_id.yaml"
  responses:
    "200":
      description: Success
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                $ref: "#/components/schemas/Status"
              data:
                $ref: "../schemas/{{ .Name.CamelLower }}.yaml"
                allOf:
                  - $ref: "#/components/schemas/SharedModel"
                  - $ref: "../schemas/{{ .Name.CamelLower }}.yaml"
    "400":
      description: Request malformed
      $ref: "#/components/responses/400BadRequest"
    "401":
      description: Unauthorized
      $ref: "#/components/responses/401Unauthorized"
    "422":
      description: Missing ID
      $ref: "#/components/responses/400BadRequest"
    "500":
      description: Server error
      $ref: "#/components/responses/500ServerError"

`

const ByID templates.Template = `

delete:
  summary: '{{ .Name.Spaced }}: Delete one'
  description: This endpoint deletes a single {{ .Name.SingularUnderscored }}, referenced by its ID.
  tags:
    - {{ .Name.CamelLower }}
  parameters:
    - in: query
      name: id
      description: A single ID
      required: true
      schema:
        type: string
  responses:
    "200":
      description: Success
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                $ref: "#/components/schemas/Status" 
    "204":
      description: Empty
      $ref: "#/components/responses/204Empty"
    "400":
      description: Request malformed
      $ref: "#/components/responses/400BadRequest"
    "401":
      description: Unauthorized
      $ref: "#/components/responses/401Unauthorized"
    "500":
      description: Server error
      $ref: "#/components/responses/500ServerError"
`
