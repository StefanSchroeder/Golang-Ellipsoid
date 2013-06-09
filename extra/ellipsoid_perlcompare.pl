# Some tests.
#
#
use strict;
use warnings;
use Geo::Ellipsoid;

my @origin = ( 37.619002, -122.374843 );    # SFO
my @dest = ( 33.942536, -118.408074 );      # LAX

{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'WGS84', units=>'degrees');
	my ( $range, $bearing ) = $geo->to( @origin, @dest );
	print "1 dist = $range bear = $bearing\n";
}

{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees');
	my ( $range, $bearing ) = $geo->to( @origin, @dest );
	print "2 dist = $range bear = $bearing\n";
}

{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'radian', bearing=>0);
	my ( $range, $bearing ) = $geo->to( @origin, @dest );
	print "3 dist = $range bear = $bearing\n";
}

{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'radian', bearing=>1);
	my ( $range, $bearing ) = $geo->to( @origin, @dest );
	print "4 dist = $range bear = $bearing\n";
}

{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'kilometer', bearing=>1);
	my ( $range, $bearing ) = $geo->to( @origin, @dest );
	print "5 dist = $range bear = $bearing\n";
}

{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'foot', bearing=>1);
	my ( $range, $bearing ) = $geo->to( @origin, @dest );
	print "6 dist = $range bear = $bearing\n";
}
{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'foot', bearing=>1, longitude=>1);
	my ($lat,$lon) = $geo->at( @origin, 2000.0, 45.0 );
	print "7 lat  = $lat lon  = $lon\n";
}
{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'meter', bearing=>1, longitude=>1);
	my ($lat,$lon) = $geo->at( @origin, 2000, 45.0 );
	print "8 lat  = $lat lon  = $lon\n";
}
{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'meter', bearing=>1, longitude=>1);
	my ($lat,$lon) = $geo->at(90.0, 90.0, 1000.0, 90.0);
	print "9 lat  = $lat lon  = $lon\n";
}
{
	my $geo = Geo::Ellipsoid->new(ellipsoid=>'WGS84', units=>'degrees', distance_units => 'nm');
	@origin = (73.06,19.11); # mumbai
	@dest   = (4.89,52.37); # amsterdam

	my ($range,$bear) = $geo->to( @origin, @dest );
	print "10 dist = $range bear = $bear\n";
}

