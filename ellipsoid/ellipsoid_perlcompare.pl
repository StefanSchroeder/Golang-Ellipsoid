use Geo::Ellipsoid;

@origin = ( 37.619002, -122.374843 );    # SFO
@dest = ( 33.942536, -118.408074 );      # LAX

{
$geo = Geo::Ellipsoid->new(ellipsoid=>'WGS84', units=>'degrees');
( $range, $bearing ) = $geo->to( @origin, @dest );
print "1 dist = $range bear = $bearing\n";
}

{
$geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees');
( $range, $bearing ) = $geo->to( @origin, @dest );
print "2 dist = $range bear = $bearing\n";
}

{
$geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'radian', bearing=>0);
( $range, $bearing ) = $geo->to( @origin, @dest );
print "3 dist = $range bear = $bearing\n";
}

{
$geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'radian', bearing=>1);
( $range, $bearing ) = $geo->to( @origin, @dest );
print "4 dist = $range bear = $bearing\n";
}

{
$geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'kilometer', bearing=>1);
( $range, $bearing ) = $geo->to( @origin, @dest );
print "5 dist = $range bear = $bearing\n";
}

{
$geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'foot', bearing=>1);
( $range, $bearing ) = $geo->to( @origin, @dest );
print "6 dist = $range bear = $bearing\n";
}
{
$geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'foot', bearing=>1, longitude=>1);
  ($lat,$lon) = $geo->at( @origin, 2000.0, 45.0 );
  print "7 lat  = $lat lon  = $lon\n";
}
{
	$geo = Geo::Ellipsoid->new(ellipsoid=>'AIRY', units=>'degrees', distance_units => 'meter', bearing=>1, longitude=>1);
	($lat,$lon) = $geo->at( @origin, 2000, 45.0 );
	print "8 lat  = $lat lon  = $lon\n";
}
{
	$geo = Geo::Ellipsoid->new(ellipsoid=>'WGS84', units=>'degrees', distance_units => 'nm');
	@origin = (73.06,19.11); # mumbai
	@dest   = (4.89,52.37); # amsterdam

	$range,$bear) = $geo->to( @origin, @dest );
	print "10 dist = $range bear = $bearing\n";
}
