import 'package:casu_mobile/model/pexel_image.dart';


class PexelImageResponse {
  final List<PexelImage> images;

  PexelImageResponse(this.images);

  PexelImageResponse.fromJson(Map<String, dynamic> json)
      : images = (json["photos"] as List)
            .map((i) => PexelImage.fromJson(i))
            .toList();

  PexelImageResponse.withError(String errorValue)
      : images = [];
}
