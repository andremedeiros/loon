_l() {
  case $1 in
    cd)
      root=$(~/src/github.com/andremedeiros/loon/loon path $2)
      cd "${root}"
      ;;
    *)
      "~/src/github.com/andremedeiros/loon/loon" "$@"
      ;;
  esac
}
