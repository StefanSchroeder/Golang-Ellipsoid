#
# Run the example from hello_world with geod from the proj4 library.
# It is available in Cygwin and Debian (proj).
#
geod +ellps=WGS84 << EOT -I +units=m -f %.12f -F %.12f
37.619002N 122.374843W 33.942536N 118.408074W
EOT
