import 'package:casu_mobile/repositories/movie_repository.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../common/global.dart';
import '../../../model/hobbie.dart';

part 'hobbie_state.dart';

class HobbieCubit extends Cubit<HobbieState> {
  HobbieCubit({required this.repository}) : super(HobbieState());
  String userId = Global.storageService.getUserId();
  final MovieRepository repository;

  Future<void> fetchHobbies() async {
    try {
      final hobbieResponse = await repository.getHobbies();
      hobbieResponse.hobbies.removeWhere((element) => element.code != userId );
      emit(HobbieState(hobbies: hobbieResponse.hobbies));
    } on Exception {
      // emit(const HobbieState.failure());
    }
  }

  Future<void> removeHobbies(Hobbie hobbie) async {
    try {
      await repository.removeHobbies(hobbie.book_id);
      final state = this.state;
      emit(HobbieState(hobbies: List.from(state.hobbies)..remove(hobbie)));
    } on Exception {
      // emit(const HobbieState.failure());
    }
  }

  Future<void> addHobbie(Hobbie hobbie) async {
    try {
      bool response = await repository.addHobbie(hobbie);
      final state = this.state;
      if (response) {
        emit(HobbieState(hobbies: List.from(state.hobbies)..add(hobbie)));
      }
    } on Exception {
      // emit(const HobbieState.failure());
    }
  }
}
