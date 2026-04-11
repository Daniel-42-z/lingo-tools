{
  description = "Development environment for lingo-tools";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [
            go
            gopls
          ];

          shellHook = ''
            export PATH="$PWD/build:$PATH"

            # ---------------------------------------------------------
            # 1. AUTO-BUILD THE BINARY
            # ---------------------------------------------------------
            # The reason completion was failing/missing flags is likely 
            # because the binary in build/ was compiled BEFORE the Cobra 
            # refactor, so it didn't understand the 'completion' command 
            # or the new flags! We build it automatically here to ensure
            # it's always up to date with the latest code.
            if [[ -f "main.go" ]]; then
              echo "Building lingo-tools..."
              go build -o build/lingo-tools main.go
            fi

            # ---------------------------------------------------------
            # 2. GENERATE COMPLETIONS
            # ---------------------------------------------------------
            if [[ -f "./build/lingo-tools" ]]; then
              # Zsh completion via FPATH
              mkdir -p .zsh-completions
              ./build/lingo-tools completion zsh > .zsh-completions/_lingo-tools 2>/dev/null || true
              export FPATH="$PWD/.zsh-completions:$FPATH"

              # Bash completion
              source <(./build/lingo-tools completion bash) 2>/dev/null || true
            fi

            # ---------------------------------------------------------
            # 3. ZSH CACHE INVALIDATION
            # ---------------------------------------------------------
            # Oh My Zsh caches completions in ~/.zcompdump. If we add a new 
            # FPATH directory, it ignores it until the cache is cleared.
            # We silently delete the cache here so Zsh rebuilds it on startup.
            if [[ -n "$ZSH_VERSION" || "$SHELL" == *"zsh"* ]]; then
               rm -f ~/.zcompdump* 2>/dev/null || true
            fi

            echo "Lingo-tools dev shell active. 'build/' added to PATH."
            
            # Hide the bash warning if they actually used -c zsh
            parent_cmd=$(ps -o args= -p $PPID 2>/dev/null || echo "")
            if [[ "$parent_cmd" != *"zsh"* && "$SHELL" != *"zsh"* ]]; then
               echo "Note: You are currently in BASH. Start with 'nix develop -c zsh' for Zsh."
            fi
          '';
        };
      }
    );
}
