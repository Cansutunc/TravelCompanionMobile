import 'dart:math';

import 'package:bloc/bloc.dart';
import 'package:casu_mobile/common/constants.dart';
import 'package:casu_mobile/model/user_location.dart';
import 'package:equatable/equatable.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:geolocator/geolocator.dart';
import 'package:meta/meta.dart';

import '../../../../common/global.dart';
import '../../../../common/location/location.dart';
import '../../../../model/location.dart';
import '../../../../repositories/movie_repository.dart';


part 'location_state.dart';


class LocationCubit extends Cubit<LocationState> {
  LocationCubit({required this.repository}) : super(const LocationState.loading()){
    getLocation();
  }
  final MovieRepository repository;

  Future<void> getLocation() async {
    try {
      Position? _currentPosition = await LocationHandler.getCurrentPosition();
      final locationResponse = await repository.getLocation(_currentPosition!);

      Global.storageService.setString(AppConstants.LOCATION, locationResponse.location.province);
      String userId = Global.storageService.getUserId();
      String userName = Global.storageService.getUsername();
      await repository.addUserLocation(UserLocation("", locationResponse.location.province, userId, locationResponse.location.country,userName));
      emit(LocationState.success(locationResponse.location));
    } on Exception {
      emit(const LocationState.failure());
    }
  }

}


