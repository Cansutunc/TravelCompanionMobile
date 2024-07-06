import 'package:equatable/equatable.dart';

class Location extends Equatable {
  final String country;
  final String province;
  final String town;
  final String suburb;
  final String road;
  final String amenity;

  const Location(
      this.country,
      this.province,
      this.town,
      this.suburb,
      this.road,
      this.amenity,
      );

  Location.fromJson(Map<String, dynamic> json)
      : country = json["country"],
        province = json["province"] ?? "",
        town = json["town"] ?? "",
        suburb = json["suburb"] ?? "",
        road = json["road"] ?? "",
        amenity = json["amenity"] ?? "";




  @override
  List<Object> get props =>
      [country, province, town,suburb,road,amenity];

  static const empty = Location("", "","","","","");
}
