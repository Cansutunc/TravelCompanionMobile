import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../common/constants.dart';
import '../../common/global.dart';
import '../../widgets/toast_info.dart';
import 'bloc/signin_blocs.dart';

class SignInController {
  final BuildContext context;

  const SignInController({required this.context});

  Future<void> resetPassword() async {
    final state = context.read<SignInBloc>().state;
    String emailAddress = state.email;
    try {
      await FirebaseAuth.instance
          .sendPasswordResetEmail(email: emailAddress);
      toastInfo(msg: "Send reset password link");
    }on FirebaseAuthException catch (e) {
    }catch(e){
      print(e);
    }
  }

  Future<void> handleSignIn(String type) async {
    if (!context.mounted) return;

    try {
      if (type == "email") {
        // BlocProvider.of<SignInBloc>(context).state
        final state = context.read<SignInBloc>().state;
        String emailAddress = state.email;
        String password = state.password;

        if (emailAddress.isEmpty) {
          toastInfo(msg: "You need to fill email address");
        }

        if (password.isEmpty) {
          toastInfo(msg: "You need to fill password");
        }

        try {
          final credential = await FirebaseAuth.instance
              .signInWithEmailAndPassword(
                  email: emailAddress, password: password);
          if (credential.user == null) {
            toastInfo(msg: "You don't exist");
          }

          if (!credential.user!.emailVerified) {
            toastInfo(msg: "You need to verify your email account");
          }

          var user = credential.user;

          if (user != null) {
            if (!context.mounted) return;
            Global.storageService
                .setString(AppConstants.STORAGE_USER_TOKEN_KEY, user.uid);
            if (user.displayName!.isNotEmpty) {
              Global.storageService.setString(
                  AppConstants.STORAGE_USER_PROFILE, user.displayName ?? "");
            }
            //
            Navigator.of(context)
                .pushNamedAndRemoveUntil("/main", (route) => false);
          } else {
            toastInfo(msg: "Currently you are not a user of this app");
          }
        } on FirebaseAuthException catch (e) {
          if (e.code == 'user-not-found') {
            toastInfo(msg: "No use found for that email");
          } else if (e.code == 'wrong-password') {
            toastInfo(msg: "Wrong password provided for that user");
          } else if (e.code == 'invalid-email') {
            toastInfo(msg: "Your email address format is wrong");
          } else if (e.code == 'invalid-credential') {
            toastInfo(msg: "Invalid Credential");
          } else {
            toastInfo(msg: e.code);
          }
        }
      }
    } catch (e) {
      print("Error");
    }
  }
}
