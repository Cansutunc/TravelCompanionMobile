import 'package:casu_mobile/pages/register_screen/register_controller.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import '../../widgets/common_widgets.dart';
import 'bloc/register_blocs.dart';
import 'bloc/register_events.dart';
import 'bloc/register_states.dart';

class RegisterScreen extends StatefulWidget {
  const RegisterScreen({super.key});


  @override
  State<RegisterScreen> createState() => _RegisterScreenState();
}

class _RegisterScreenState extends State<RegisterScreen> {
  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => RegisterBlocs(),
      child: BlocBuilder<RegisterBlocs, RegisterStates>(
        builder: (context, state) {
          return Container(
            color: Colors.white,
            child: SafeArea(
              child: Scaffold(
                backgroundColor: Colors.white,
                appBar: buildAppBar("Register"),
                body: SingleChildScrollView(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      SizedBox(height: 20.h),
                      Center(
                          child: reusableText(
                              "Enter your details belw and free sign up")),
                      Container(
                        margin: EdgeInsets.only(top: 66.h),
                        padding: EdgeInsets.only(left: 25.w, right: 25.w),
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            reusableText("Username"),
                            buildTextField(
                                "Enter your user name", "name", "user",
                                (value) {
                              context
                                  .read<RegisterBlocs>()
                                  .add(UserNameEvent(value));
                            }),
                            reusableText("Email"),
                            buildTextField(
                                "Enter your email address", "email", "user",
                                (value) {
                              context
                                  .read<RegisterBlocs>()
                                  .add(EmailEvent(value));
                            }),
                            reusableText("Password"),
                            buildTextField(
                                "Enter your password", "password", "lock",
                                (value) {
                              context
                                  .read<RegisterBlocs>()
                                  .add(PasswordEvent(value));
                            }),
                            reusableText("Confirm Password"),
                            buildTextField("Enter your confirm password",
                                "password", "lock", (value) {
                              context
                                  .read<RegisterBlocs>()
                                  .add(RePasswordEvent(value));
                            }),
                          ],
                        ),
                      ),
                      Container(
                          margin: EdgeInsets.only(left: 25.w),
                          child: reusableText(
                              "Enter your details below and free sign up")),
                      buildLoginAndReqButton("Sign Up", "login", () {
                        RegisterController(context: context)
                            .handleEmailRegister();
                      }),
                    ],
                  ),
                ),
              ),
            ),
          );
        },
      ),
    );
  }
}
