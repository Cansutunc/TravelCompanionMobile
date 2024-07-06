import 'package:casu_mobile/pages/hobbie_screen/bloc/hobbie_cubit.dart';
import 'package:casu_mobile/repositories/movie_repository.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../common/global.dart';
import 'bottom_sheet.dart';


class HobbieView extends StatelessWidget {
  const HobbieView({super.key, required MovieRepository movieRepository});

  void _addModal(BuildContext context2) {
    showModalBottomSheet(
      context: context2,
      builder: (context) {
        return SingleChildScrollView(
          child: Container(
            padding: EdgeInsets.only(
              bottom: MediaQuery.of(context).viewInsets.bottom,
            ),
            child: CustomModelBottomSheet(context2: context2),
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    String userId = Global.storageService.getUserId();
    return BlocBuilder<HobbieCubit, HobbieState>(
      buildWhen: (previous, current) {
        if (previous.hobbies.length != current.hobbies.length) {
          return true;
        }
        return false;
      },
      builder: (context, state) {
        return Scaffold(
          appBar: AppBar(
            title: Text(
              "Your Hobbies",
              style: TextStyle(color: Colors.black),
            ),
          ),
          floatingActionButton: FloatingActionButton(
            child: Icon(Icons.add),
            onPressed: () {
              _addModal(context);
            },
          ),
          body: ListView.builder(
            itemCount: state.hobbies.length,
            itemBuilder: (context, index) {
              if (userId == state.hobbies[index].code){
                return ListTile(
                  leading: CircleAvatar(radius:15,child: Text((index+1).toString()),),
                  title: Text(state.hobbies[index].name),
                  trailing: IconButton(
                    onPressed: () {
                      context
                          .read<HobbieCubit>()
                          .removeHobbies(state.hobbies[index]);
                    },
                    icon: Icon(Icons.remove),
                  ),
                );
              }

            },
          ),
        );
      },
    );
  }
}
