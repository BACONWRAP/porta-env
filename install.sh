#!/usr/bin/env bash

# Install the dependencies
sh <(curl -L https://nixos.org/nix/install) --daemon

# Nix
if [ -e '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh' ]; then
  . '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'
fi
# End Nix
PATH="$HOME/.nix-profile/bin:/nix/var/nix/profiles/default/bin:$PATH"
nix-channel --add https://github.com/nix-community/home-manager/archive/master.tar.gz home-manager
nix-channel --update
nix-shell '<home-manager>' -A install
cd ~/.config/home-manager && { curl -O https://raw.githubusercontent.com/BACONWRAP/porta-env/main/home.nix ; cd -; }
sed -i.bak 's/home.username.*/home.username = "${(whoami)}";/g' ~/.config/home-manager/home.nix && rm ~/.config/home-manager/home.nix.bak
sed -i.bak 's/home.username.*/home.username = "${(echo $HOME)}";/g' ~/.config/nixpkgs/home.nix && rm ~/.config/nixpkgs/home.nix.bak
sed 
home-manager switch
