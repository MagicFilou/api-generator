package name

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	SINGULAR                 string = "<SINGULAR>"
	SINGULARLOWER            string = "<SINGULARLOWER>"
	SINGULARUNDERSCORED      string = "<SINGULARUNDERSCORED>"
	SINGULARUNDERSCOREDLOWER string = "<SINGULARUNDERSCOREDLOWER>"
	CAMEL                    string = "<CAMEL>"
	CAMELLOWER               string = "<CAMELLOWER>"
	PLURAL                   string = "<PLURAL>"
	PLURALLOWER              string = "<PLURALLOWER>"
	PLURALUNDERSCORED        string = "<PLURALUNDERSCORED>"
	PLURALUNDERSCOREDLOWER   string = "<PLURALUNDERSCOREDLOWER>"
	SHORT                    string = "<SHORT>"
)

type Name struct {
	Singular            string `yaml:"singular,omitempty"`    // eventlog
	SingularUnderscored string `yaml:"singular_underscored"`  // event_log
	CamelLower          string `yaml:"camel_lower,omitempty"` // eventLog
	CamelUpper          string `yaml:"camel_upper,omitempty"` // EventLog
	Plural              string `yaml:"plural,omitempty"`      // eventlogs
	PluralUnderscored   string `yaml:"plural_underscored"`    // event_logs
	Spaced              string `yaml:"spaced"`                // Event Log
	Short               string `yaml:"short"`                 // e
}

func (n *Name) Build() error {

	if n.SingularUnderscored == "" || n.PluralUnderscored == "" {
		return fmt.Errorf("missing SingularUnderscored and/or PluralUnderscored")
	}

	if n.Short == "" {

		fmt.Println("SHORT NAME NOT SET")
		fmt.Println("consider assigning a value manually to avoid conflicts")
		fmt.Println("setting short name to first character of PluralUnderscored")
	}

	n.Singular = strings.ToLower(strings.ReplaceAll(n.SingularUnderscored, "_", ""))

	n.CamelLower = strcase.ToLowerCamel(n.SingularUnderscored)

	n.CamelUpper = strcase.ToCamel(n.SingularUnderscored)

	n.Plural = strings.ToLower(strings.ReplaceAll(n.PluralUnderscored, "_", ""))

	n.Spaced = cases.Title(language.Und, cases.NoLower).String(strcase.ToDelimited(n.CamelUpper, ' '))

	return nil
}

func (n Name) ToPreload() string {

	return ".Preload(\"" + strcase.ToCamel(n.PluralUnderscored) + "\")"
}
