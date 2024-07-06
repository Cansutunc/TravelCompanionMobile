import 'package:equatable/equatable.dart';

class Routes extends Equatable {
  final String city_id;
  final String title;
  final String description;
  final String createdUserId;
  final String createdUsername;
  final int maxPerson;
  final List<dynamic> allPersons;

  const Routes(
      this.city_id,
      this.title,
      this.description,
      this.createdUserId,
      this.createdUsername,
      this.maxPerson,
      this.allPersons,
      );

  Routes.fromJson(Map<String, dynamic> json)
      : city_id = json["city_id"],
        title = json["title"] ?? "",
        description = json["description"] ?? "",
        createdUserId = json["created_user_id"] ?? "",
        createdUsername = json["created_username"] ?? "",
        maxPerson = json["max_person"] ?? 0,
        allPersons = json["all_person"] ?? [];

  Map<String, dynamic> toMap() {
    return {
      'title': title,
      'description': description,
      'created_user_id': createdUserId,
      'created_username': createdUsername,
      'max_person': maxPerson,
    };
  }


  @override
  List<Object> get props =>
      [city_id, title, description,createdUserId,createdUsername,maxPerson,allPersons];

  static const empty = Routes("", "","","","",0,[]);
}
