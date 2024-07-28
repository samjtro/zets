# big idea

templates are essentially just map[string]*Template
- when you compose templates together, you populate this template

{{template "name"}} you are doing a map[string]*Template lookup for name

one of them is marked as the current template

set current template: `t.Lookup("name") // == nil, if DNE`
execute current template: `t.Execute(os.Stdout, data)`

template inheritance; go doesn't need it, because refs are dynamically resolved

base.html:
```
{{template "name"}}
{{define "name"}}
    placeholder
{{end}}
```
index.html
```
{{define "name"}}
    placeholder2
{{end}}
```

depending on where this template is executed from, the content will be changed

# syntax keywords

1. {{define ...}}
2. {{template ...}}
3. {{block ...}} {{end}}
4. {{if `<var>`} {{else if `<var>`}} {{else}} {{end}}
5. {{with `<var>`}} {{end}}
6. {{range `<var>`}} {{end}}

... but first:
# variables

every template has access to one var, `.` notation, passed into `Execute`/`ExecuteTemplate()`

e.g. dot is `struct{`
```
    Params: Params{
        Author: Author{
            FirstName: "John",
            LastName: "Doe",
        },
    },
}
```
`{{.Params.Author.LastName}}`
e.g. dot is `map[string]interface{}{`
```
    "Params": map[string]interface{}{
        "Author": map[string]interface{}{
            "FirstName": "John",
            "LastName": "Doe",
        },
    },
}
```

## vars must be passed explicitly down to templates

```
{{define "footer"}}
    {{.Params.Author.FirstName}}
{{end}}
{{template "footer"}} //fails
{{template "footer" .}} //passes

{{define "footer"}}
    {{.FirstName}}
{{end}}
{{template "footer" .Params.Author}}
```

## vars can be assigned, and reassigned

```
{{define "greet"}}
    {{$firstname := .Params.User.FirstName}}
    <p>Hello, my name is {{$firstname}}</p>
    {{$firstname = "Joe"}}
    <p>Actually, it's {{$firstname}}</p>
{{end}}
```

var scope extends to {{end}} of current block

... now back to:
# syntax

## 4. {{if <var>}} {{else if <var>}} {{else}} {{end}}

```
{{if .User.IsAuthorized}}
    <p>Welcome</p>
{{else}}
    <p>Not welcome...</p>
{{end}}
```

eval = truthy/falsy
    - falsy = `false`, `0`, `nil`, empty array/slice/map/string
    - anything else

empty slice = `false`

### dot, dollar

dot refers to current template var, dollar also does so

dot changes depending on context, aka within various blocks
whereas, $ remains the top level variable

## 5. {{with <var>}} {{end}}

behaves like an {{if}}, if `<var>` is falsy, the block is not executed
otherwise, dot is set to var - anything outside of dot, use $

## 5. {{range <var>}} {{else}} {{end}}

loops over `<var>.<var>` - must be array, slice, channel

if len == 0, {{else}} is executed

dot is changed to each successive element of the array within the range

anything outside of dot should use $

basic:
```
{{range .Users}}
    <p>User: {{.}}</p>
    <p>FirstName: {{.FirstName}}</p>
    <p>Today is {{$.Today}}</p>
{{end}}
```

template var:
```
{"Users": [
    {
        "FirstName": "John",
        ...
    }
    {
        ...
        ...
    }
],
"Today": "Thursday"
}
```

if slice:
`{{range $i, $user := .Users}}`...

if map:
`{{range $key, $val := .User}}`...

if comparable with ==, it will be sorted, otherwise not

# funcs, methods, pipes

templates have some default funcs

- and `{{if and <cond> <cond> ...}} {{end}}`
- or `{{if or <cond> <cond> ...}} {{end}}`
- not `{{if not <cond>}} {{end}}`
- len `{{len <var>}}` // `<var>` is one of string | slice | array | map | channel
- index `{{index <slice> <num> ...}}` | `{{index <map> <key> ...}}`

## can call user defined funcs

install by calling `.Funcs(map[string]inteface{})`

```
template.
    New("t").
    Funcs(map[string]interface{}{
        "greet": func(name string) string {
            return "hello " + name
        }
    }).
    Parse(`{{greet "bob"}}`)
```

can take any number of args, must return one result

can optionally return an error, if non-nil halts execution

## templates can call methods

## templates can't directly call vars that are funcs

## results of one function can be piped to another

```
...
    Parse(`{{upper . | greet "hello"}}`)
t.Execute(os.Stdout, "bob")
```

# more

templates can be parsed separately, then combined
- just merging two map[string]*Template's

if two templates combine, their functions (map[string]interface{}'s) will also get merged together

`html/template` automatically secures html/css/hs/url input
- even upcoming {{break}} and {{continue}} will respect this

functions must be defined before Parse(), but then they can be overwritten
