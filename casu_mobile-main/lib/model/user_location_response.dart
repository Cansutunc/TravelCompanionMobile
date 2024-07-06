import 'package:casu_mobile/model/pexel_image.dart';
import 'package:casu_mobile/model/user_location.dart';


class UserLocationResponse {
  final List<UserLocation> userLocations;
  final bool hasError;
  final String error;

  UserLocationResponse(this.userLocations, this.hasError, this.error);

  UserLocationResponse.fromJson(Map<String, dynamic> json)
      : userLocations = (json["data"] as List)
      .map((i) => UserLocation.fromJson(i))
      .toList(),

        hasError = false,
        error = "";

  UserLocationResponse.withError(String errorValue)
      : userLocations = [],
        hasError = true,
        error = errorValue;
}