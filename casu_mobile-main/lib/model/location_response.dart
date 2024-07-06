import 'location.dart';
import 'video.dart';

class LocationResponse {
  final Location location;
  final String error;

  LocationResponse(this.location, this.error);

  LocationResponse.fromJson(Map<String, dynamic> json)
      : location = Location.fromJson(json["address"]),
        error = "";

  LocationResponse.withError(String errorValue)
      : location = Location.empty,
        error = errorValue;
}
