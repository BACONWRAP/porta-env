## how to use

```bash
bash <(curl https://raw.githubusercontent.com/BACONWRAP/porta-env/main/install.sh)
```

### What it does

- Installs nix
- Installs home-manager
- Places the configuration in ~/.config/nixpkgs/home.nix
- Modifies the home.nix config to use this machine's env vars
- Installs the packages in the home.nix config

### What it still needs

- A way to add more packages to the home.nix config
- A way to manage any configurations for those packages
- A way to back all this up and store it in a remote repo
