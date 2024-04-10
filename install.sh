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
home-manager switch
