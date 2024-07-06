import 'package:casu_mobile/model/hobbie.dart';
import 'package:equatable/equatable.dart';

class HobbieResponse {
  final List<Hobbie> hobbies;
  final bool hasError;
  final String error;

  HobbieResponse(this.hobbies, this.hasError, this.error);

  HobbieResponse.fromJson(Map<String, dynamic> json)
      : hobbies = (json["data"] as List)
      .map((i) => Hobbie.fromJson(i))
      .toList(),

        hasError = false,
        error = "";

  HobbieResponse.withError(String errorValue)
      : hobbies = [],
        hasError = true,
        error = errorValue;
}