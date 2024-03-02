{ pkgs, ... }:

{
  env.GREET = "Welcome to capwaspoc";

  packages = [
    pkgs.cocoapods
    pkgs.npm-check-updates
  ];

  languages.go.enable = true;
  languages.javascript.enable = true;

  enterShell = ''
    echo $GREET
    echo "node version:"
    node --version
    go version
  '';
}
