package templategeneration

const ProjectMarkdownTemplate = `
# {{.Title}}

| Package |
| - | {{range $Package := .AlphabeticalRows}}
[{{$Package.PackageName}}]({{$Package.PackageName}}/{{$Package.OutputFile}})|{{end}}

Integration diagram:

![alt text]({{.RootLevelIntegrationDiagram.AppName}}.svg)
`

const PackageMarkdownTemplate = `
[Back](../README.md)
# Package {{.PackageName}}

## Sequence Diagrams
| AppName | Endpoint |
| - | - | {{range $Diagram := .SequenceDiagrams}}
| {{$Diagram.AppName}} | [{{$Diagram.EndpointName}}]({{$Diagram.OutputFileName__}}.md) |{{end}}
`

const EmbededSvgTemplate = `
[Back](README.md)

![alt text]({{.OutputFileName__}}.svg)

`

// [Back](../README.md)
// # Package {{.PackageName}}

// ## Data Model Diagrams

// ![alt text]({{.PackageName}}_datamodel.svg)

// ## Sequence Diagrams
// | AppName | Endpoint |
// | - | - | {{range $Diagram := .SequenceDiagrams}}
// | {{$Diagram.AppName}} | [{{$Diagram.OutputFileName__}}]({{$Diagram.OutputFileName__}}.md) |{{end}}