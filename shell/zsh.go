package shell

// Zsh is the zsh shell integration.
var Zsh = Shell(`# Put the line below in ~/.zshrc:
#
#   eval "$(jump shell zsh)"
#
# The following lines are autogenerated:

__jump_chpwd() {
  jump chdir
}

typeset -gaU chpwd_functions
chpwd_functions+=__jump_chpwd

{{.Bind}}() {
  local dir="$(jump cd $@)"
  test -d "$dir" && cd "$dir"
}

jump_completion() {
  reply=($(jump hint $@))
}

compctl -U -K jump_completion {{.Bind}}
`)
