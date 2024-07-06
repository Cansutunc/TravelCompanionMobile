import 'package:equatable/equatable.dart';

class PexelImage extends Equatable {
  final int id;
  final String title;
  final String imgUrl;

  const PexelImage(this.id, this.title,this.imgUrl);

  PexelImage.fromJson(Map<String, dynamic> json)
      : id = json["id"],
        title = json["alt"] ?? "",
        imgUrl = json["src"]["original"] ?? "";


  @override
  List<Object> get props =>
      [id, imgUrl, title];

  static const empty = PexelImage(0,"", "");
}
