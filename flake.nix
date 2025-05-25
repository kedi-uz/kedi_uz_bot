{
    description = "Telegram bot @kedi_uz_bot";

    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
        flake-utils.url = "github:numtide/flake-utils";

    };
    outputs = {self, nixpkgs, flake-utils, ... }: 
        # Attributes for each system
        flake-utils.lib.eachDefaultSystem (
            system: let
                pkgs = import nixpkgs {inherit system; };
                in {
                    formatter = pkgs.alejandra;

                    # Dev Shells
                    devShells.default = import ./shell.nix { inherit pkgs; };
                }
        );
}