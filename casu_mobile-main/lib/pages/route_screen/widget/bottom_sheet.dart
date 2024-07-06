import 'package:casu_mobile/model/route.dart';
import 'package:casu_mobile/pages/route_screen/bloc/route_bloc/route_cubit.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';


import '../../../common/colors.dart';
import '../../../common/global.dart';

class CustomRouteModelBottomSheet extends StatelessWidget {

  CustomRouteModelBottomSheet({super.key,required this.context2,required this.userId,required this.username});
  TextEditingController titleController = TextEditingController();
  TextEditingController descriptionController = TextEditingController();
  TextEditingController maxPersonController = TextEditingController();


  final BuildContext context2;
  final String userId;
  final String username;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(12.0),
      child: Column(
        children: [
          Text("Add Activity"),
          SizedBox(height: 10),
          Form(
              child: Column(
            children: [
              Padding(
                padding: EdgeInsets.all(8.0),
                child: TextField(
                  controller: titleController,
                  keyboardType: TextInputType.multiline,
                  decoration: InputDecoration(
                    hintText: "hintText",
                    border: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    enabledBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    disabledBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    focusedBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    hintStyle: TextStyle(
                      color: Colors.grey.withOpacity(0.5),
                    ),
                  ),
                  style: TextStyle(
                      color: AppColors.primaryText,
                      fontFamily: "Avenir",
                      fontWeight: FontWeight.normal,
                      fontSize: 14.sp),
                  autocorrect: false,

                ),
              ),
              Padding(
                padding: EdgeInsets.all(8.0),
                child: TextField(
                  keyboardType: TextInputType.multiline,
                  controller: descriptionController,
                  decoration: InputDecoration(
                    hintText: "hintText",
                    border: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    enabledBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    disabledBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    focusedBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    hintStyle: TextStyle(
                      color: Colors.grey.withOpacity(0.5),
                    ),
                  ),
                  style: TextStyle(
                      color: AppColors.primaryText,
                      fontFamily: "Avenir",
                      fontWeight: FontWeight.normal,
                      fontSize: 14.sp),
                  autocorrect: false,

                ),
              ),
              Padding(
                padding: EdgeInsets.all(8.0),
                child: TextField(
                  keyboardType: TextInputType.multiline,
                  controller: maxPersonController,
                  decoration: InputDecoration(
                    hintText: "hintText",
                    border: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    enabledBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    disabledBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    focusedBorder: const OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black)),
                    hintStyle: TextStyle(
                      color: Colors.grey.withOpacity(0.5),
                    ),
                  ),
                  style: TextStyle(
                      color: AppColors.primaryText,
                      fontFamily: "Avenir",
                      fontWeight: FontWeight.normal,
                      fontSize: 14.sp),
                  autocorrect: false,

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
                  ElevatedButton(onPressed: () {
                    context2.read<RouteCubit>().addRoute(Routes("route_id", titleController.value.text, descriptionController.value.text,userId,username, int.parse(maxPersonController.value.text), []));
                    Navigator.pop(context);

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
