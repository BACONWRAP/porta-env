#!/usr/bin/env bash

# Install the dependencies
sh <(curl -L https://nixos.org/nix/install) --daemon

# Nix
if [ -e '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh' ]; then
  . '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'
fi
# End Nix
exec PATH="$HOME/.nix-profile/bin:/nix/var/nix/profiles/default/bin:$PATH"
nix-channel --add https://github.com/nix-community/home-manager/archive/master.tar.gz home-manager
nix-channel --update
nix-shell '<home-manager>' -A install
home-manager switch
