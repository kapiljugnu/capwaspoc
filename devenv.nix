{ pkgs, ... }:

{
  env.GREET = "Welcome to capwaspoc";

  packages = with pkgs; [
    cocoapods
    fswatch
    npm-check-updates
    templ
  ];

  dotenv.enable = true;

  languages.go.enable = true;
  languages.javascript.enable = true;

  enterShell = ''
    echo $GREET
    echo "node version:"
    node --version
    go version
  '';
}
