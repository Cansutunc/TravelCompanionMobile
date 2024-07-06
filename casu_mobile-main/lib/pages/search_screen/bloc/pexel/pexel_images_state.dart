part of 'pexel_images_cubit.dart';

enum ListStatus { loading, success, failure }

class PexelImagesState extends Equatable {
  const PexelImagesState._({
    this.status = ListStatus.loading,
    this.images = const <PexelImage>[],
  });

  const PexelImagesState.loading() : this._();

  const PexelImagesState.success(List<PexelImage> images)
      : this._(status: ListStatus.success, images: images);

  const PexelImagesState.failure() : this._(status: ListStatus.failure);

  final ListStatus status;
  final List<PexelImage> images;

  @override
  List<Object> get props => [status, images];
}