import 'package:equatable/equatable.dart';

class UserLocation extends Equatable {
  final String country_id;
  final String province;
  final String user_id;
  final String country;
  final String user_name;

  const UserLocation(this.country_id, this.province, this.user_id, this.country,this.user_name);

  UserLocation.fromJson(Map<String, dynamic> json)
      : country_id = json["country_id"] ??"",
        province = json["province"] ??"",
        user_id = json["user_id"] ?? "",
        user_name = json["user_name"] ?? "",
        country = json["country"] ?? "";


  Map<String, dynamic> toMap() {
    return {
      'country_id': country_id,
      'province': province,
      'user_id': user_id,
      'country': country,
      'user_name': user_name,
    };
  }


  @override
  List<Object> get props =>
      [country_id, province, user_id, country,user_name];

  static const empty = UserLocation("", "", "", "","");
}
