import 'package:casu_mobile/model/hobbie.dart';
import 'package:casu_mobile/model/route.dart';
import 'package:equatable/equatable.dart';

class RouteResponse {
  final List<Routes> routes;
  final bool hasError;
  final String error;

  RouteResponse(this.routes, this.hasError, this.error);

  RouteResponse.fromJson(Map<String, dynamic> json)
      : routes = (json["data"] as List)
      .map((i) => Routes.fromJson(i))
      .toList(),

        hasError = false,
        error = "";

  RouteResponse.withError(String errorValue)
      : routes = [],
        hasError = true,
        error = errorValue;
}