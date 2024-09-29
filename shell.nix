with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    git
    go-task
    
    # go development
    go
    go-outline
    gopls
    gopkgs
    go-tools
  ];
  shellHook = ''
    set -a
    set +a
  '';
}