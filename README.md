<h1 align="center">
  <br>
  <img width="150" height="150" src="gh-s-logo.png">
  <br>
</h1>

<h2 align="center">
  <a href="#" onclick="return false;">
    <img alt="PR" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat"/>
  </a>
  <a href="https://golang.org/">
    <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=flat&logo=go&logoColor=white"/>
  </a>
  <a href="https://github.com/gennaro-tedesco/gh-s/releases">
    <img alt="releases" src="https://img.shields.io/github/release/gennaro-tedesco/gh-s"/>
  </a>
</h2>

<h4 align="center">search github repositories interactively</h4>
<h3 align="center">
  <a href="#Installation">Installation</a> •
  <a href="#Usage">Usage</a> •
  <a href="#Feedback">Feedback</a>
</h3>

Search GitHub repositories interactively from the command line. Start the prompt and browse the results! The name of that repository 🤔? Written in rust, a list of awesome projects...

...well say no more:
```
gh s -l rust -d list
```

<img alt="example_image" src="https://user-images.githubusercontent.com/15387611/151635859-4a8a2200-b000-4e03-888a-2dc8ddcef009.png">

## Installation
```
gh extension install gennaro-tedesco/gh-s
```
This being a `gh` extension, you of course need [gh cli](https://github.com/cli/cli) as prerequisite.

## Usage
Get started!
```
gh s
```

![demo](https://user-images.githubusercontent.com/15387611/151630538-07574523-662a-4e74-b117-4afec38794ad.gif)

Without any argument (or with flags only) `gh s` starts a prompt to insert the search query; after the search a list of results is shown. Navigate the list to show details, stars counts, URL and more. If instead you want to do all in one line
```
gh s [search] [flag]
```
takes one of the following arguments or flags

| flags        | description                                      | multiple   | example
|:------------ |:------------------------------------------------ |:---------- |:--------
| -l, --lang   | narrow down the search to a specific language    | yes (OR)   | gh s prompt -l go -l lua
| -d, --desc   | search for keyword in the repository description | no         | gh s neovim -d plugin
| -u, --user   | restrict the search to a specific user           | no         | gh s lsp -u neovim
| -t, --topic  | narrow down the search to specific topics        | yes (AND)  | gh s nvim -t plugin -l lua
| -c, --colour | change colour of the prompt                      | no         | gh s nvim -c magenta
| -h, --help   | show the help page                               | no         | gh s -h
| -V, --version| print the current version                        | no         | gh s -V

The prompt accepts the following navigation commands:

| key           | description
|:------------- |:-----------------------------------
| arrow keys    | browse results list
| `/`           | toggle search in results list
| `enter (<CR>)`| open selected repository in web browser


### Execute commands
`gh-s` must be intended as a filter prompt returning the URL of the selection; as such, the best and most flexible way to execute commands with the results is to pipe it into and from `stdin/stdout`. Have a look at the Wiki for some common examples!

## Feedback
If you find this application useful consider awarding it a ⭐, it is a great way to give feedback! Otherwise, any additional suggestions or merge request is warmly welcome!

See also a companion extension to snap around your git workflow: [gh-f](https://github.com/gennaro-tedesco/gh-f).
