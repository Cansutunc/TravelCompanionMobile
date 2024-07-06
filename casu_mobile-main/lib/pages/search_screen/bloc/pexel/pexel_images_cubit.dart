import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:casu_mobile/repositories/movie_repository.dart';

import '../../../../model/pexel_image.dart';
part 'pexel_images_state.dart';

class PexelImagesCubit extends Cubit<PexelImagesState> {
  PexelImagesCubit({required this.repository})
      : super(const PexelImagesState.loading());

  final MovieRepository repository;

  Future<void> getPexelImage(String query) async {
    try {
      final movieResponse = await repository.getPexelImages(query);
      emit(PexelImagesState.success(movieResponse.images));
    } on Exception {
      emit(const PexelImagesState.failure());
    }
  }
}