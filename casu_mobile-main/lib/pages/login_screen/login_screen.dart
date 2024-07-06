import 'package:casu_mobile/pages/login_screen/sign_in_controller.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import '../../common/colors.dart';
import '../../widgets/common_widgets.dart';
import 'bloc/signin_blocs.dart';
import 'bloc/signin_events.dart';
import 'bloc/signin_states.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => SignInBloc(),
      child: BlocBuilder<SignInBloc, SignInState>(
        builder: (context, state) =>
            Container(
              color: Colors.white,
              child: SafeArea(
                child: Scaffold(
                  backgroundColor: Colors.white,
                  appBar: buildAppBar("LogIn"),
                  body: SingleChildScrollView(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        buildThirdPartyLogin(context),
                        Center(
                            child:
                            reusableText("Or use your email account to login")),
                        Container(
                          margin: EdgeInsets.only(top: 66.h),
                          padding: EdgeInsets.only(left: 25.w, right: 25.w),
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              reusableText("Email"),
                              SizedBox(
                                height: 5.h,
                              ),
                              buildTextField(
                                  "Enter your email address", "email", "user",
                                      (value) {
                                    context.read<SignInBloc>().add(
                                        EmailEvent(value));
                                  }),
                              reusableText("Password"),
                              buildTextField(
                                  "Enter your email address", "password",
                                  "lock",
                                      (value) {
                                    context
                                        .read<SignInBloc>()
                                        .add(PasswordEvent(value));
                                  }),
                              Container(
                                margin: EdgeInsets.only(left: 10.w),
                                width: 260.w,
                                height: 44.h,
                                child: GestureDetector(
                                  onTap: () {
                                    SignInController(context: context).resetPassword();
                                  },
                                  child: Text(
                                    "Forgot password",
                                    style: TextStyle(
                                        color: AppColors.primaryText,
                                        // decoration: TextDecoration.underline,
                                        decorationColor: AppColors.primaryText,
                                        fontSize: 12.sp),
                                  ),
                                ),
                              ),
                              buildLoginAndReqButton("Log In", "login", () {
                                SignInController(context: context)
                                    .handleSignIn("email");
                              }),
                              buildLoginAndReqButton(
                                  "Register", "register", () {
                                Navigator.of(context).pushNamed("/register");
                              }),
                            ],
                          ),
                        )
                      ],
                    ),
                  ),
                ),
              ),
            ),
      ),
    );
  }
}
