rm -rf $HOME/.relayer
rly config init
rly config add-chains ibc/chains

ibcplug="bottom dad own clap trade dice action ripple crater obvious mountain cement penalty first doctor advice keen reform coyote fun leader verb december dry"
ibconp="kick oyster love green crowd culture course zoo wood female shed forward track small urban flame imitate analyst combine protect beauty blur pear donate"

rly keys restore plugchain_520-1 ibcplug "$ibcplug"
rly keys restore onp_521-1 ibconp "$ibconp"
rly config add-paths ibc/paths
rly start ibc-plug