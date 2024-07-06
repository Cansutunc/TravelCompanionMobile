import 'package:equatable/equatable.dart';

class Hobbie extends Equatable {
  final String book_id;
  final String name;
  final String code;
  final String author;
  final String homepage;

  const Hobbie(
      this.book_id,
      this.name,
      this.code,
      this.author,
      this.homepage,
      );

  Hobbie.fromJson(Map<String, dynamic> json)
      : book_id = json["book_id"],
        name = json["name"] ?? "",
        code = json["code"] ?? "",
        author = json["author"] ?? "",
        homepage = json["homepage"] ?? "";

  Map<String, dynamic> toMap() {
    return {
      'name': name,
      'code': code,
      'author': author,
      'homepage': homepage,
    };
  }


  @override
  List<Object> get props =>
      [book_id, name, code,author,homepage];

  static const empty = Hobbie("", "","","","");
}
