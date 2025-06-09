{
  description = "A Nix-flake-based Go 1.24 + OCAML development environment";

  inputs = {
    # List of platform identifiers, e.g. "x86_64-linux" etc.
    systems.url = "github:nix-systems/default";

    # Snapshot of nixpkgs, pinned by a FlakeHub wildcard.
    nixpkgs.url = "nixpkgs/nixos-unstable";

    flake-utils.url = "github:numtide/flake-utils";
  };

  # ──────────────────────────────────────────────────────────
  # outputs : receives materialized inputs and *returns* an attr‑set
  # ──────────────────────────────────────────────────────────
  outputs =
    {
      self,
      nixpkgs,
      systems,
      flake-utils,
    }:
    let
      lib = nixpkgs.lib; # Nixpkgs standard library
      eachSystem = lib.genAttrs (import systems);
    in
    {
      devShells = eachSystem (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          python = pkgs.python312.withPackages (
            ps: with ps; [
              numpy
              pandas
              matplotlib
              jupyterlab
              ipykernel
              scikit-learn
              xgboost
            ]
          );
        in
        {
          default = pkgs.mkShell {
            dontDetectOcamlConflicts = true;

            # packages placed on $PATH
            packages = with pkgs; [
              # --- Python toolchain ---
              python
              sqlite
              dotnetCorePackages.sdk_9_0
            ];
          };
        }
      );
    };
}
