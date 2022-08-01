package docs

import (
	"api-builder/templates"
)

const OpenAPISpec templates.Template = `
openapi: 3.1.0
info:
  title: Wult internal REST API
  description: |
    Full API specification for Wult's internal API.
  version: 0.0.1
tags:
{{ range .TagDescriptions }}
  {{ . }}
{{ end }}
x-tagGroups:
  - name: Resources
    tags:
    {{ range .Tags }}
      {{ . }}
    {{ end }}
components:

  responses:
  
    '204Empty':
      description: Empty. The requested resource was not found, the request was otherwise well formed.
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                $ref: "#/components/schemas/Status"
              length:
                $ref: "#/components/schemas/Length"

    '400BadRequest':

      description: Request malformed. Failure due to errors in the request.
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                $ref: "#/components/schemas/Status"
              error:
                $ref: "#/components/schemas/Error"

    '401Unauthorized':

      description: Unauthorized. The requester does not have access to the resource.
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                $ref: "#/components/schemas/Status"
              error:
                $ref: "#/components/schemas/Error"

    '500ServerError':

      description: Server error. This is almost always still a frontend issue.
      content:
        application/json:
          schema:
            type: object
            properties:
              status:
                $ref: "#/components/schemas/Status"
              error:
                $ref: "#/components/schemas/Error"

  schemas:

    Error:
      type: string

    Status:
      type: string

    Length:
      type: integer
      example: 1

    Count:
      type: integer
      example: 0
    
    id:
      type: string
      example: 760c23a4-4f9b-415c-97b0-960a23d02d24

    unix:
      type: integer
      example: 1656533128

    SharedModel:
      type: object
      properties:
        id:
          $ref: "#/components/schemas/id"
        created:
          $ref: "#/components/schemas/unix"
        updated:
          $ref: "#/components/schemas/unix"
        archived:
          $ref: "#/components/schemas/unix"

paths:
    {{ range .Paths }}
      {{ . }}
    {{ end }}

  /serviceconfigurations/nested:
    $ref: "./paths/serviceConfigurationNested.yaml"
`
