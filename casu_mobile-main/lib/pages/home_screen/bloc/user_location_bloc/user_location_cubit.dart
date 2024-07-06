import 'package:casu_mobile/model/hobbie.dart';
import 'package:casu_mobile/model/user_location.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../common/global.dart';
import '../../../../repositories/movie_repository.dart';


part 'user_location_state.dart';


class UserLocationCubit extends Cubit<UserLocationState> {
  UserLocationCubit({required this.repository}) : super(const UserLocationState.loading()){
    getUserLocation();
  }
  final MovieRepository repository;
  String username = Global.storageService.getUsername();

  Future<void> getUserLocation() async {
    try {
      final res = await repository.getUserLocations();

      res.userLocations.removeWhere((item) => item.user_name == username );
      final resHobbies = await repository.getHobbies();
      emit(UserLocationState.success(res.userLocations,resHobbies.hobbies));
    } on Exception {
      emit(const UserLocationState.failure());
    }
  }

}


