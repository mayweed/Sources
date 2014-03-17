#!/bin/bash

# ameliorer le harcode de l'interface, verifier que usage est root
# recuperer les messages d'erreur dans des logs
ifconfig eth0 down
ifconfig ethO up
wpa_supplicant -B -Dwext -i eth0 -c /etc/wpa_supplicant.conf
dhcpcd eth0
