import 'package:casu_mobile/model/hobbie.dart';
import 'package:casu_mobile/pages/hobbie_screen/bloc/hobbie_cubit.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';


import '../../../common/global.dart';

class CustomModelBottomSheet extends StatelessWidget {

  CustomModelBottomSheet({super.key,required this.context2});
  TextEditingController nameController = TextEditingController();

  String userId = Global.storageService.getUserId();
  String userName = Global.storageService.getUsername();

  final BuildContext context2;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(12.0),
      child: Column(
        children: [
          Text("Add Hobbie"),
          SizedBox(height: 10),
          Form(
              child: Column(
            children: [
              Padding(
                padding: EdgeInsets.all(8.0),
                child: TextField(
                  controller: nameController,
                  autofocus: true,
                  decoration: InputDecoration(
                      label: Text("Title"), border: OutlineInputBorder()),
                ),
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  TextButton(
                      onPressed: () {
                        Navigator.pop(context);
                      },
                      child: Text("Cancel")),
                  ElevatedButton(focusNode: FocusNode(),onPressed: () {
                    if(nameController.value.text.isNotEmpty){
                      context2.read<HobbieCubit>().addHobbie(Hobbie(userId, nameController.value.text, userId, userName, "homepage"));
                      Navigator.pop(context);
                    }

                  }, child: Text("Save"))
                ],
              )
            ],
          ))
        ],
      ),
    );
  }
}
