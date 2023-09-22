package helpers

// All used by the application to render the templates.
var All = map[string]any{
	// "form":    form.Form,
	// "formFor": tags.FormFor,

	"hasPrefix":    hasPrefix,
	"urlWithParam": urlWithParams,
	"urlHas":       urlHas,

	"currentPathWithoutFields": currentPathWithoutFields,
	"currentPathWithout":       currentPathWithout,
	"urlMatches":               urlMatches,
	"setActiveClass":           setActiveClass,
}
