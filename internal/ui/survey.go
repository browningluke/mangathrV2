package ui

import "github.com/AlecAivazis/survey/v2"

func Checkboxes(label string, opts []string) []string {
	survey.MultiSelectQuestionTemplate = `
{{- define "option"}}
    {{- if eq .SelectedIndex .CurrentIndex }}{{color .Config.Icons.SelectFocus.Format }}{{ .Config.Icons.SelectFocus.Text }}{{color "reset"}}{{else}} {{end}}
    {{- if index .Checked .CurrentOpt.Index }}{{color .Config.Icons.MarkedOption.Format }} {{ .Config.Icons.MarkedOption.Text }} {{else}}{{color .Config.Icons.UnmarkedOption.Format }} {{ .Config.Icons.UnmarkedOption.Text }} {{end}}
    {{- color "reset"}}
    {{- " "}}{{- .CurrentOpt.Value}}
{{end}}
{{- if .ShowHelp }}{{- color .Config.Icons.Help.Format }}{{ .Config.Icons.Help.Text }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color .Config.Icons.Question.Format }}{{ .Config.Icons.Question.Text }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }}{{ .FilterMessage }}{{color "reset"}}
{{- if .ShowAnswer}}{{color "cyan"}} {{"✔"}}{{color "reset"}}{{"\n"}}
{{- else }}
	{{- "  "}}{{- color "cyan"}}[Use arrows to move, space to select, <right> to all, <left> to none, type to filter{{- if and .Help (not .ShowHelp)}}, {{ .Config.HelpInput }} for more help{{end}}]{{color "reset"}}
  {{- "\n"}}
  {{- range $ix, $option := .PageEntries}}
    {{- template "option" $.IterateOption $ix $option}}
  {{- end}}
{{- end}}`

	var res []string
	prompt := &survey.MultiSelect{
		Message:  label,
		Options:  opts,
		PageSize: 10,
	}
	err := survey.AskOne(prompt, &res, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Text = ""
		icons.Question.Format = "yellow+hb"
	}), survey.WithKeepFilter(true))
	if err != nil {
		panic(err)
	}

	return res
}

func SingleCheckboxes(label string, opts []string) string {
	var res string
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	err := survey.AskOne(prompt, &res, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Text = ""
		icons.Question.Format = "yellow+hb"
	}))
	if err != nil {
		panic(err)
	}

	return res
}

func ConfirmPrompt(label string) bool {
	var res bool

	prompt := &survey.Confirm{
		Message: label,
	}

	err := survey.AskOne(prompt, &res)
	if err != nil {
		panic(err)
	}

	return res
}

func InputPrompt(label string) string {
	res := ""
	prompt := &survey.Input{
		Message: label,
	}
	err := survey.AskOne(prompt, &res)
	if err != nil {
		panic(err)
	}

	return res
}
